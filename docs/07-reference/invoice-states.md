# Invoice States Reference — Nexora

## Introducción

Invoice representa una de las entidades centrales del dominio de Nexora.

Su ciclo de vida está controlado mediante una máquina de estados (State Machine), donde cada estado representa una etapa concreta del proceso comercial.

Las operaciones relacionadas con:

- pagos.
- envío de factura.
- modificaciones.
- inventario.
- checkout.
- procesos comerciales.

dependen directamente del estado actual de la Invoice.

El estado de la Invoice funciona como fuente principal de verdad para determinar qué operaciones están permitidas.

---

# Modelo General

El flujo conceptual es:

```
Draft

↓

Pending Payment

↓

Paid

↓

Completed

```

No todas las operaciones pueden ejecutarse en cualquier estado.

Cada transición posee reglas específicas.

---

# Estados Principales

## Draft

### Descripción

Estado inicial de una Invoice.

Representa una factura creada pero todavía no confirmada dentro del flujo comercial.

Puede contener:

- datos básicos.
- items.
- cantidades.
- información del cliente.

---

### Operaciones permitidas

En este estado se permite:

- agregar items.
- modificar información.
- eliminar items.
- validar datos antes del envío.

---

### Restricciones

No permite:

- realizar pagos.
- completar procesos contables.
- finalizar operaciones comerciales.

---

# Pending Payment

## Descripción

Representa una Invoice confirmada y lista para recibir pagos.

En este estado la factura ya fue validada y comienza el proceso financiero.

---

## Operaciones permitidas

Se permite:

- registrar pagos.
- consultar información.
- continuar flujo financiero.

---

## Restricciones

No permite:

- modificar información crítica.
- alterar items.
- cambiar cantidades.

La Invoice debe mantener consistencia con el monto esperado.

---

# Paid

## Descripción

Representa una Invoice cuyo monto total fue cubierto.

El pago acumulado alcanza el monto requerido.

---

## Operaciones permitidas

Se permite:

- consultar información histórica.
- continuar procesos posteriores.

---

## Restricciones

No permite:

- registrar nuevos pagos.
- modificar valores comerciales.

---

# Completed

## Descripción

Estado final del proceso comercial.

Representa una Invoice completamente procesada.

---

## Operaciones permitidas

Únicamente:

- consulta.
- auditoría.
- reportes.

---

# Transiciones Permitidas

La máquina de estados acepta únicamente transiciones válidas.

Ejemplo:

```
Draft

↓

Pending Payment

↓

Paid

↓

Completed

```

No son válidas:

```
Draft

↓

Paid

```

o

```
Completed

↓

Draft

```

---

# Validación de Transiciones

Antes de ejecutar una operación de negocio se valida:

1. Estado actual de Invoice.

2. Acción solicitada.

3. Transición permitida.

4. Reglas adicionales del dominio.

Ejemplo:

Un pago solamente puede ejecutarse si:

```
Invoice.status == Pending Payment

```

---

# Invoice Payment Rules

Los pagos poseen reglas adicionales.

Antes de registrar un pago se valida:

## Invoice pagable

La Invoice debe encontrarse en un estado permitido.

---

## Monto válido

El monto ingresado no puede superar el saldo pendiente.

Ejemplo:

```
Invoice Total: 100

Paid Amount: 80

Remaining: 20

```

Un pago de:

```
30

```

debe ser rechazado.

---

# Invoice Items Rules

Los items pertenecen al ciclo de vida de la Invoice.

Su modificación depende del estado principal.

Ejemplo:

Permitido:

```
Draft

+

Agregar Item

```

No permitido:

```
Paid

+

Modificar Item

```

Esto evita inconsistencias comerciales.

---

# Inventario y Estado

El estado de Invoice también determina cuándo interactuar con inventario.

Los procesos como:

- checkout.
- reserva de stock.
- confirmación comercial.

dependen del workflow definido.

El inventario no debe actualizarse fuera del proceso autorizado.

---

# Seguridad y Acceso

Antes de ejecutar cualquier operación sobre Invoice se realizan controles previos:

```
JWT Authentication

↓

Tenant Context

↓

Permission Validation

↓

Invoice Access Validation

↓

Business Rules

↓

Operation

```

Esto garantiza que:

- el usuario está autenticado.
- posee permisos.
- pertenece al contexto correcto.
- la Invoice pertenece al alcance permitido.

---

# Reglas Fundamentales

## Una Invoice controla su ciclo de vida

Las entidades relacionadas no deben modificar su estado independientemente.

---

## Los servicios de negocio controlan transiciones

Las operaciones complejas no son simples updates.

Ejemplo:

Incorrecto:

```
UPDATE invoice SET status='Paid'

```

Correcto:

```
PayInvoice()

↓

Validate State

↓

Validate Amount

↓

Register Payment

↓

Update Status

```

---

## El estado protege la consistencia del dominio

La State Machine evita:

- pagos inválidos.
- modificaciones fuera de tiempo.
- datos inconsistentes.
- operaciones fuera del flujo comercial.

---

# Resumen Visual

```
                +---------+
                |  Draft  |
                +---------+
                     |
                     |
                     v
        +-----------------------+
        | Pending Payment       |
        +-----------------------+
                     |
                     |
                     v
             +-------------+
             |    Paid     |
             +-------------+
                     |
                     |
                     v
          +----------------+
          |   Completed    |
          +----------------+

```

---

# Filosofía del Diseño

Invoice no fue diseñada como una tabla con un campo status.

Fue diseñada como una entidad con comportamiento y reglas de negocio.

El estado representa una etapa real del proceso empresarial y funciona como mecanismo de protección de consistencia dentro del dominio.

Esta decisión permite que nuevas funcionalidades puedan incorporarse respetando el flujo existente sin romper reglas comerciales.
