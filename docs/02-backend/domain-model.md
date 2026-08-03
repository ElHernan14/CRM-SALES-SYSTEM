# Domain Model — Nexora

## Introducción

El modelo de dominio de Nexora fue diseñado tomando como punto central los procesos comerciales relacionados con la gestión de facturación, ventas, compras, inventario y relaciones entre usuarios, clientes y empresas.

La construcción del dominio no comenzó desde una representación puramente técnica de tablas, sino desde la identificación de las entidades principales que representan conceptos reales del negocio.

El núcleo principal identificado fue la entidad `Invoice`, debido a que representa el eje alrededor del cual giran múltiples procesos:

- creación de documentos comerciales.
- gestión de estados.
- pagos.
- control de items.
- reserva y actualización de inventario.
- flujos de Store y Marketplace.
- operaciones administrativas del ERP.

A partir de este núcleo se definieron los distintos dominios y responsabilidades que componen la aplicación.

La intención principal fue evitar construir un sistema basado únicamente en operaciones CRUD sobre tablas, y en cambio modelar procesos reales mediante servicios de dominio y workflows encapsulados.

---

# Principios utilizados

Durante el diseño del modelo de dominio se aplicaron principalmente los siguientes principios:

- Separación por dominios funcionales.
- Responsabilidad única por módulo.
- Encapsulamiento de reglas de negocio.
- Aggregate Roots definidos.
- Servicios de dominio para procesos complejos.
- Evitar lógica distribuida entre múltiples capas.
- Mantener invariantes del negocio centralizadas.

La arquitectura busca que cada módulo tenga un lenguaje y responsabilidad clara, evitando que una entidad conozca detalles internos innecesarios de otros dominios.

---

# Dominio principal: Invoice

## Invoice como núcleo del sistema

`Invoice` representa el agregado principal del sistema comercial.

Una invoice no es simplemente un registro almacenado en base de datos, sino una entidad con ciclo de vida propio y reglas asociadas.

Desde esta entidad nacen múltiples procesos:

- creación de documentos.
- incorporación de productos.
- validaciones comerciales.
- envío de factura.
- pago.
- actualización de inventario.
- seguimiento de estados.

La invoice actúa como punto de coordinación entre distintos dominios sin absorber responsabilidades que pertenecen a otros módulos.

---

# Aggregate Roots principales

Los principales Aggregate Roots identificados son:

## Invoice

Representa el agregado comercial principal.

Responsabilidades:

- controlar el ciclo de vida de la operación comercial.
- mantener coherencia de sus estados.
- validar transiciones permitidas.
- coordinar procesos relacionados.

Entidades asociadas:

- Invoice Items.
- Invoice Payments.

---

## Client

`Client` representa la raíz del subdominio de relaciones comerciales.

Responsabilidades:

- representar compradores.
- mantener información comercial.
- vincular operaciones realizadas por usuarios clientes.

El cliente posee ciclo de vida propio y no depende directamente de una invoice específica.

---

## Company

`Company` representa el tenant empresarial dentro del sistema.

Responsabilidades:

- definir el contexto organizacional.
- agrupar usuarios empresariales.
- aislar información mediante multi-tenancy.

Una compañía existe independientemente de las operaciones comerciales realizadas.

---

## Product

`Product` representa el catálogo comercial.

Responsabilidades:

- definir productos disponibles.
- almacenar información comercial.
- relacionarse con inventario.
- participar en operaciones de venta y compra.

Su ciclo de vida es independiente de invoices específicas.

---

# Entidades dependientes y ciclo de vida

No todas las entidades poseen el mismo nivel de autonomía.

## Invoice Item

Los items pertenecen directamente al agregado Invoice.

Su ciclo de vida está condicionado por la factura:

Ejemplo:

- una invoice en estado draft puede modificarse.
- una invoice enviada no debería permitir cambios arbitrarios.
- una invoice cancelada no puede continuar procesos comerciales.

El item no tiene sentido fuera del contexto de la invoice.

---

## Invoice Payment

Los pagos están asociados al ciclo de vida financiero de una invoice.

Reglas principales:

- no se puede pagar una invoice inexistente.
- no se puede pagar una invoice en estados inválidos.
- el monto acumulado no puede superar el total permitido.
- una invoice pasa a estado pagado cuando corresponde.

---

# Invariantes del dominio

Las invariantes representan reglas que siempre deben mantenerse verdaderas dentro del sistema.

## Invoice

Principales invariantes:

- Una invoice debe pertenecer a un contexto válido.
- Una invoice debe respetar su máquina de estados.
- Una transición debe ser válida antes de ejecutarse.
- Los montos financieros deben permanecer consistentes.

---

## Invoice Items

Invariantes:

- Un item pertenece a una única invoice.
- Los valores calculados deben mantener consistencia con la factura.
- No puede modificarse fuera de estados permitidos.

---

## Payments

Invariantes:

- El monto pagado nunca puede superar el total.
- Una invoice completamente pagada debe reflejar correctamente su estado.
- Cada pago debe estar asociado a una invoice válida.

---

# State Machine de Invoice

Uno de los conceptos centrales del dominio es la máquina de estados.

La invoice controla su comportamiento mediante estados definidos.

Ejemplo conceptual:

```

Draft

↓

Pending

↓

Paid

```

Cada transición representa una acción de negocio, no una simple modificación de columna.

Ejemplos:

- Submit Invoice.
- Pay Invoice.
- Cancel Invoice.

Cada operación valida previamente:

- estado actual.
- permisos.
- ownership.
- reglas comerciales.
- consistencia de datos.

---

# Separación entre CRUD y procesos de negocio

Una decisión importante del diseño fue diferenciar operaciones simples de procesos reales.

## Operaciones CRUD

Ejemplos:

- crear producto.
- actualizar cliente.
- consultar información.

Características:

- afectan una entidad específica.
- tienen pocas reglas.
- no representan procesos completos.

---

## Procesos de negocio

Ejemplos:

- Submit Invoice.
- Pay Invoice.
- Store Checkout.
- Marketplace Checkout.
- Dashboard Overview.
- Inventory Reservation.

Características:

- involucran múltiples entidades.
- contienen reglas de negocio.
- coordinan distintos dominios.
- requieren consistencia transaccional.

Estos procesos fueron encapsulados mediante servicios específicos.

---

# Domain Services y Workflows

Cuando una operación no pertenece exclusivamente a una entidad, se utiliza un servicio de dominio o workflow.

Ejemplo:

## SubmitInvoice

Responsabilidades:

- validar invoice.
- validar cliente.
- verificar ownership.
- validar items.
- controlar transición de estado.
- ejecutar transacción.

---

## PayInvoice

Responsabilidades:

- registrar pago.
- actualizar estado financiero.
- finalizar reservas.
- coordinar inventario.

---

## Store Checkout

Responsabilidades:

- validar carrito.
- generar invoice.
- reservar stock.
- completar flujo comercial.

---

Estos servicios representan casos de uso reales del negocio y no simples operaciones técnicas.

---

# Organización de dominios

Los principales dominios definidos son:

```

Auth

User

Company

Client

Product

Inventory

Invoice

Payment

Purchases

Dashboard

Store

Marketplace

```

Cada dominio mantiene sus propias responsabilidades:

- modelos.
- repositorios.
- servicios.
- reglas específicas.

---

# Módulos compartidos

No todo componente representa un dominio.

Existen módulos transversales que colaboran con múltiples dominios.

Ejemplos:

## Tenant Context

Permite mantener identidad y contexto durante una request.

---

## Invoice Access

Centraliza validaciones relacionadas con:

- ownership.
- permisos.
- relación usuario/invoice.

---

## Transaction Utilities

Proveen herramientas comunes para operaciones atómicas.

---

## Error Handling

Centraliza convenciones de errores.

---

# Límites entre dominios

Un dominio termina cuando sus reglas dejan de ser autónomas.

Ejemplos:

## Invoice

Responsabilidad:

- representar documento comercial.
- manejar estados.
- controlar operación de facturación.

No debería conocer:

- cómo se almacena stock.
- cómo funciona inventario internamente.

---

## Inventory

Responsabilidad:

- disponibilidad.
- reservas.
- actualización de stock.

No debería conocer:

- reglas completas de facturación.

---

## Payment

Responsabilidad:

- gestión financiera relacionada al pago.

No debería conocer:

- detalles internos del catálogo.

---

Esta separación permite que cada dominio evolucione independientemente.

---

# Decisiones arquitectónicas relevantes

Durante la construcción del modelo de dominio se tomaron las siguientes decisiones:

## Invoice como entidad central

Se eligió una entidad comercial principal desde la cual derivar procesos relacionados.

Esto permitió construir un modelo coherente alrededor del ciclo de vida de una operación comercial.

---

## Servicios para lógica transversal

Cuando una regla involucraba múltiples dominios, no se forzó dentro de una entidad.

Se creó un servicio específico encargado de coordinar el proceso.

---

## Repositorios sin reglas de negocio

Los repositorios fueron mantenidos como componentes de persistencia.

Responsabilidades:

- obtener información.
- guardar información.
- ejecutar consultas.

No contienen decisiones comerciales.

---

## Bajo acoplamiento entre módulos

Los módulos colaboran mediante contratos definidos.

No dependen de implementaciones internas de otros dominios.

Esto permite incorporar nuevos módulos sin modificar grandes partes del sistema.

---

# Evolución futura del dominio

El modelo actual permite extender Nexora hacia nuevos dominios.

Ejemplos:

## Logística

Podría evolucionar desde Inventory hacia un dominio independiente.

Agregar:

- envíos.
- seguimiento.
- distribución.

---

## Contabilidad

Podría construirse sobre los eventos comerciales existentes.

Agregar:

- balances.
- movimientos contables.
- reportes financieros.

---

## Proveedores

Podría evolucionar desde Client/Product hacia un dominio comercial completo.

Agregar:

- proveedores.
- órdenes de compra.
- contratos.

---

# Evaluación arquitectónica

El modelo de dominio de Nexora fue construido priorizando coherencia, separación de responsabilidades y evolución futura.

La arquitectura evita tratar las entidades como simples tablas y establece una separación clara entre:

- entidades del negocio.
- procesos comerciales.
- infraestructura.
- persistencia.

La utilización de Aggregate Roots, servicios de dominio y workflows permite mantener reglas complejas encapsuladas y facilita que nuevos módulos puedan incorporarse sin degradar la arquitectura existente.

El resultado es un modelo preparado para evolucionar desde un MVP avanzado hacia una plataforma empresarial con múltiples dominios de negocio.

```

```
