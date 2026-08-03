# Invoice Lifecycle — Nexora

## Introducción

La entidad `Invoice` representa el núcleo del dominio comercial de Nexora.

Más que una simple tabla de facturación, funciona como un documento de negocio que controla múltiples procesos relacionados:

- Ventas.
- Compras.
- Pagos.
- Inventario.
- Store.
- Marketplace.

Debido a esta importancia, Invoice posee un ciclo de vida definido mediante estados, donde cada transición representa una operación de negocio válida.

El objetivo principal del diseño fue evitar modificaciones arbitrarias y garantizar que cada operación respete el estado actual del documento.

---

# Concepto de State Machine

El ciclo de vida de una invoice fue diseñado como una máquina de estados.

Esto significa que:

- Una invoice no puede cambiar libremente entre estados.
- Cada transición debe cumplir reglas determinadas.
- Las operaciones disponibles dependen del estado actual.

Conceptualmente:

```text
Estado actual

↓

Validación de reglas

↓

Transición permitida

↓

Nuevo estado
```

---

# Estados Principales

La invoice posee diferentes estados según su etapa dentro del flujo comercial.

## Draft

Representa una factura en construcción.

Características:

- Puede modificarse.
- Puede agregar o eliminar items.
- Puede utilizarse como base para futuros procesos.

Es el estado inicial de muchos flujos.

---

## Pending

Representa una invoice lista para continuar el proceso comercial.

En este estado:

- Los datos principales ya fueron validados.
- Los items están definidos.
- Puede continuar hacia pago u otras operaciones.

---

## Paid

Representa una invoice cuyo monto total fue cubierto.

Condiciones:

- El total pagado debe coincidir con el monto de la invoice.
- Los pagos asociados quedan registrados.

Una invoice pagada no debería volver a estados anteriores.

---

## Completed

Representa la finalización completa del flujo.

Puede implicar:

- Confirmación comercial.
- Actualización definitiva de inventario.
- Cierre del proceso.

---

# Transiciones Permitidas

El flujo general puede representarse:

```text
Draft

↓

Pending

↓

Paid

↓

Completed
```

Cada transición posee validaciones propias.

---

# Reglas de Estado

## Draft → Pending

Para enviar una invoice debe cumplirse:

- La invoice existe.
- El usuario tiene permisos.
- Posee contexto válido.
- Tiene cliente asociado cuando corresponde.
- Posee items válidos.
- Los cálculos internos son correctos.

Si alguna validación falla, la transición no ocurre.

---

## Pending → Paid

Para realizar un pago:

- La invoice debe encontrarse en estado permitido.
- El monto ingresado debe ser válido.
- No puede superar el total pendiente.
- El usuario debe poseer autorización.

Proceso:

```text
Registrar pago

↓

Actualizar paid_amount

↓

Comparar total

↓

Cambiar estado si corresponde
```

---

## Paid → Completed

Representa la finalización del flujo.

Puede incluir:

- Confirmación de inventario.
- Cierre operativo.
- Registro final.

---

# Invoice como Aggregate Root

Dentro del modelo de dominio, Invoice funciona como Aggregate Root.

Esto significa que controla la consistencia de entidades dependientes.

Principalmente:

```text
Invoice

├── Invoice Items

└── Payments
```

Las entidades relacionadas no deberían modificarse ignorando las reglas del agregado principal.

---

# Relación con Invoice Items

Los items representan líneas internas de una invoice.

Su ciclo de vida depende del documento principal.

Ejemplos:

- Una invoice cerrada no debería permitir modificaciones de items.
- Una invoice enviada mantiene sus items asociados.
- Los cálculos deben permanecer consistentes.

---

# Relación con Payments

Los pagos forman parte del flujo financiero de una invoice.

Reglas principales:

- No se puede pagar una invoice inexistente.
- No se puede superar el monto total.
- No se puede pagar una invoice en estado incorrecto.

El servicio de pago controla estas invariantes.

---

# Relación con Inventario

En determinados flujos, la invoice interactúa con inventario.

Ejemplo:

```text
Checkout

↓

Invoice creada

↓

Submit Invoice

↓

Pago confirmado

↓

Confirmar stock reservado
```

La actualización de inventario debe respetar la consistencia del proceso completo.

---

# Flujos Principales

## Flujo ERP

```text
Crear Invoice

↓

Agregar Items

↓

Submit Invoice

↓

Procesar pago

↓

Completar operación
```

---

## Flujo Store

```text
Carrito

↓

Checkout

↓

Crear Invoice

↓

Validar stock

↓

Procesar operación comercial
```

---

## Flujo Marketplace

```text
Productos externos

↓

Creación Invoice

↓

Validaciones

↓

Procesamiento comercial
```

---

# Consistencia Transaccional

Las operaciones críticas sobre invoice utilizan transacciones.

Ejemplo: Pago

Dentro de una única operación:

```text
Crear Payment

+

Actualizar paid_amount

+

Actualizar estado Invoice

+

Actualizar Inventario
```

Si una operación falla:

```text
Rollback

↓

Estado anterior restaurado
```

Esto evita información parcial.

---

# Validaciones de Seguridad

Antes de modificar una invoice se valida:

- Usuario autenticado.
- Permisos correspondientes.
- Tenant Context.
- Ownership.
- Estado permitido.

El acceso nunca depende únicamente del ID enviado por cliente.

---

# Separación de Responsabilidades

La implementación evita concentrar todo dentro de Invoice.

La distribución es:

## Controller

Responsable de:

- HTTP.
- Entrada.
- Respuesta.

---

## Service

Responsable de:

- Reglas de negocio.
- Estados.
- Workflow.

---

## Repository

Responsable de:

- Persistencia.
- Consultas.
- Escritura de datos.

---

# Beneficios del Diseño

El modelo de lifecycle permite:

## Mayor consistencia

Las operaciones inválidas son bloqueadas.

## Mejor trazabilidad

Cada transición representa una acción concreta.

## Escalabilidad

Nuevos estados o workflows pueden agregarse sin romper la estructura existente.

## Seguridad

El acceso depende del contexto y no solamente de identificadores.

---

# Evolución Futura

El modelo actual permite evolucionar hacia estados adicionales.

Ejemplos:

```text
Cancelled

Refunded

Expired

Archived
```

También permite incorporar nuevos procesos:

- Facturación electrónica.
- Contabilidad.
- Logística.
- Auditoría financiera.

---

# Evaluación Arquitectónica

El ciclo de vida de Invoice es una de las piezas centrales de Nexora.

La utilización de una máquina de estados permite transformar una entidad tradicional en un proceso de negocio controlado.

La separación entre estados, workflows, servicios y persistencia permite mantener reglas claras y evita inconsistencias a medida que el sistema crece.

Invoice no funciona únicamente como un registro almacenado, sino como el eje alrededor del cual se coordinan múltiples dominios del sistema.
