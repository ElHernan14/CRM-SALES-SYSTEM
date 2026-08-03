# Business Workflows — Nexora

## Introducción

Dentro de Nexora existen operaciones que no representan simples modificaciones de datos, sino procesos completos del negocio.

Estos procesos involucran múltiples entidades, validaciones, reglas comerciales y coordinación entre distintos dominios.

Por esta razón, fueron diseñados como **workflows de negocio independientes**, separados de las operaciones CRUD tradicionales.

El objetivo principal fue evitar que la lógica compleja terminara distribuida entre controllers, repositories o múltiples servicios sin una responsabilidad clara.

La arquitectura utiliza servicios encapsulados que representan casos de uso completos.

Conceptualmente:

```

Request

↓

Controller

↓

Application Service

↓

Business Workflow

↓

Domain Services

↓

Repositories

↓

Database

```

Cada workflow representa una operación real del negocio y es responsable de coordinar las acciones necesarias para completarla.

---

# Diferencia entre CRUD y Workflow

## Operaciones CRUD

Las operaciones CRUD representan acciones simples sobre una entidad.

Ejemplos:

- Crear producto.
- Actualizar cliente.
- Consultar invoice.
- Eliminar registro.

Características:

- Trabajan principalmente sobre una entidad.
- Tienen pocas reglas.
- No representan un proceso empresarial completo.
- No coordinan múltiples dominios.

---

## Business Workflows

Los workflows representan procesos reales.

Ejemplos:

- Enviar una factura.
- Completar un pago.
- Realizar checkout.
- Actualizar inventario.
- Generar una vista empresarial consolidada.

Características:

- Involucran múltiples entidades.
- Ejecutan validaciones complejas.
- Coordinan distintos módulos.
- Mantienen consistencia transaccional.
- Representan acciones del negocio.

---

# Arquitectura de Workflows

Los workflows fueron diseñados como una capa intermedia entre los servicios de dominio y los casos de uso expuestos por la API.

Su responsabilidad principal es:

- coordinar procesos complejos.
- ordenar la ejecución de múltiples pasos.
- controlar transacciones.
- mantener consistencia entre dominios.

No reemplazan a los servicios de dominio.

Los servicios de dominio mantienen reglas específicas.

Los workflows coordinan esas reglas.

---

# Submit Invoice Workflow

## Objetivo

El workflow `SubmitInvoice` representa el proceso mediante el cual una invoice deja de estar en estado inicial y pasa a un estado comercial válido para continuar su ciclo de vida.

No es simplemente actualizar una columna de estado.

Representa una transición del negocio.

---

# Flujo conceptual

```

Submit Request

↓

Validación de Tenant

↓

Carga Invoice

↓

Validación Ownership

↓

Validación Cliente

↓

Validación Items

↓

Validación Totales

↓

Transacción SQL

↓

Cambio de estado Invoice

↓

Commit

```

---

# Responsabilidades

El workflow realiza:

## Obtención del contexto

Obtiene información del tenant actual:

- usuario.
- empresa.
- permisos.
- cliente asociado.

Esto garantiza que la operación pertenece al contexto correcto.

---

## Validación de existencia

Comprueba que:

- la invoice existe.
- pertenece al contexto permitido.
- sus entidades relacionadas son válidas.

---

## Validación de acceso

Antes de modificar información comercial se verifica:

- ownership.
- permisos.
- relación usuario-entidad.

---

## Validación de estado

La invoice debe encontrarse en un estado permitido.

Ejemplo:

Una invoice ya enviada no debería volver a enviarse nuevamente.

---

## Validación de items

Comprueba que:

- existen productos asociados.
- los datos comerciales son correctos.
- los valores calculados mantienen consistencia.

---

## Transacción

Una vez validadas las precondiciones comienza la operación atómica.

Dentro de la transacción:

- se actualiza el estado.
- se persisten cambios necesarios.

Si ocurre un error:

```

Rollback automático

```

evitando estados parciales.

---

# Pay Invoice Workflow

## Objetivo

El workflow `PayInvoice` representa la finalización del proceso financiero de una invoice.

No solamente registra un pago.

Coordina:

- actualización financiera.
- creación del pago.
- transición de estado.
- actualización de inventario.

---

# Flujo conceptual

```

Pay Request

↓

Validar Invoice

↓

Validar Estado

↓

Registrar Payment

↓

Actualizar Paid Amount

↓

Comparar Total

↓

Actualizar Estado Invoice

↓

Finalizar Reserva Inventario

↓

Commit

```

---

# Responsabilidades

## Registro del pago

El workflow crea la entidad de pago correspondiente.

Mantiene:

- monto.
- relación con invoice.
- fecha.
- información necesaria.

---

## Actualización financiera

Actualiza:

```

paid_amount

```

de la invoice.

Luego verifica si:

```

paid_amount == total_amount

```

En ese caso:

```

Invoice State = Paid

```

---

## Integración con Inventory

Cuando la operación financiera termina correctamente, se notifica al dominio de inventario.

El workflow utiliza:

```

FinalizeReservedStock()

```

para confirmar la salida definitiva del stock.

---

# Store Checkout Workflow

## Objetivo

Representa el proceso de compra desde el módulo Store.

No es simplemente crear una invoice.

Implica una operación comercial completa.

---

# Flujo conceptual

```

Cart

↓

Validar productos

↓

Validar disponibilidad

↓

Reservar stock

↓

Crear Invoice

↓

Crear Invoice Items

↓

Preparar Checkout

↓

Confirmar operación

```

---

# Responsabilidades

El checkout coordina:

- carrito.
- productos.
- inventario.
- invoice.
- cliente.

Cada módulo mantiene su propia responsabilidad.

El workflow únicamente coordina el proceso.

---

# Marketplace Checkout Workflow

El Marketplace utiliza una lógica similar al Store, pero orientada a operaciones comerciales entre distintos actores.

Responsabilidades principales:

- validar productos publicados.
- verificar condiciones comerciales.
- generar operación comercial.
- mantener aislamiento entre tenants.

---

# Inventory Workflow

## Objetivo

El módulo Inventory nació como una separación del dominio Product.

La decisión arquitectónica fue separar la responsabilidad de disponibilidad física del catálogo comercial.

Product representa:

- qué se vende.

Inventory representa:

- cuánto existe disponible.

---

# Responsabilidades

Inventory maneja:

- reservas.
- liberaciones.
- confirmaciones.
- actualización de cantidades.

Ejemplos:

```

ReserveStock()

ReleaseStock()

FinalizeReservedStock()

```

---

# Cart Ensure Workflow

## Objetivo

Mantener la existencia de un carrito válido para operaciones Store.

Responsabilidades:

- verificar carrito activo.
- crear draft si corresponde.
- mantener consistencia previa al checkout.

---

# Dashboard Overview Workflow

## Objetivo

Dashboard Overview es un caso especial.

No representa una escritura de negocio, sino un proceso de composición de información.

Es un **Composition Service**.

---

# Características

El workflow:

- recibe contexto tenant.
- obtiene información empresarial.
- consolida múltiples fuentes.
- devuelve un snapshot del estado actual.

---

# Dominios involucrados

Ejemplo:

- ventas.
- compras.
- inventario.
- métricas comerciales.

---

# Consistencia

Como es una operación de lectura:

No necesita transacciones de escritura.

La consistencia se garantiza mediante:

- mismo tenant context.
- mismo company_id.
- consultas coherentes.

---

# Comunicación entre workflows y dominios

Una decisión importante fue evitar que todos los servicios se conozcan entre sí.

La regla general:

```

Domain Service

↓

Workflow

↓

Otro Domain Service

```

Los servicios no llaman arbitrariamente otros servicios.

Cuando una operación cruza límites de dominio, se crea un workflow explícito.

---

# Transacciones y consistencia

Los workflows que modifican múltiples entidades utilizan transacciones SQL.

Principios aplicados:

## Atomicidad

Ocurre todo:

```

Commit

```

o nada:

```

Rollback

```

---

## Validación previa

Antes de iniciar una transacción:

- se validan entidades.
- se verifican estados.
- se comprueban permisos.

---

## Transaction Context

Cuando es necesario, los repositorios reciben el contexto transaccional.

Esto permite que todas las operaciones formen parte de la misma unidad atómica.

---

# Beneficios arquitectónicos

La utilización de workflows permitió:

## Evitar controllers complejos

Los controllers solamente traducen HTTP.

No contienen procesos de negocio.

---

## Evitar services gigantes

Los servicios mantienen responsabilidades específicas.

---

## Mantener dominios independientes

Cada módulo mantiene sus propias reglas.

---

## Facilitar evolución

Agregar nuevos procesos implica crear nuevos workflows sin modificar los existentes.

---

# Evaluación arquitectónica

Los workflows representan una de las decisiones más importantes del diseño de Nexora.

Permiten modelar operaciones reales del negocio evitando una arquitectura basada únicamente en CRUD.

La separación entre:

- entidades.
- servicios de dominio.
- workflows.
- infraestructura.

permite mantener un sistema donde procesos complejos pueden evolucionar sin perder claridad ni consistencia.

Este enfoque acerca la arquitectura a modelos utilizados en sistemas empresariales donde los casos de uso tienen prioridad sobre la estructura técnica de almacenamiento.

```

```
