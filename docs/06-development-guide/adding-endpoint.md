# Adding a New Endpoint — Nexora Development Guide

## Introducción

Nexora fue diseñado siguiendo una separación clara entre responsabilidades. Agregar un nuevo endpoint no consiste únicamente en crear una nueva ruta HTTP, sino en incorporar un nuevo caso de uso respetando la arquitectura existente.

La creación de nuevos endpoints debe mantener los siguientes principios:

- Controllers delgados.
- Lógica de negocio dentro de Services.
- Acceso a datos únicamente mediante Repositories.
- Validaciones compartidas mediante módulos comunes.
- Errores centralizados.
- Uso del Tenant Context cuando corresponda.
- Separación entre operaciones CRUD y procesos de negocio.

El objetivo es que cada nuevo endpoint pueda incorporarse sin aumentar el acoplamiento del sistema.

---

# Flujo Arquitectónico de un Endpoint

El ciclo completo de una request sigue la siguiente estructura:

```
Request

↓

Router

↓

Middlewares

↓

Controller

↓

Service

↓

Repository

↓

Database

```

Cada capa posee una responsabilidad específica.

---

# 1. Definir el Caso de Uso

Antes de crear un endpoint se debe determinar qué representa dentro del dominio.

La primera pregunta no debe ser:

"¿Qué tabla necesito modificar?"

Sino:

"¿Qué proceso de negocio estoy representando?"

Ejemplos:

Incorrecto:

```
PUT /invoice
```

para modificar cualquier comportamiento relacionado con una factura.

Correcto:

```
POST /invoice/{id}/submit

POST /invoice/{id}/pay

POST /checkout

```

porque representan acciones reales del dominio.

---

# CRUD vs Business Workflow

Antes de implementar un endpoint debe definirse si corresponde a:

## Operación CRUD

Características:

- operación simple sobre una entidad.
- no contiene reglas complejas.
- no involucra múltiples dominios.

Ejemplo:

```
GET /products

POST /clients

DELETE /categories/{id}

```

---

## Proceso de negocio

Características:

- contiene múltiples validaciones.
- modifica varias entidades.
- representa una acción real del sistema.

Ejemplos:

```
POST /invoice/{id}/submit

POST /invoice/{id}/pay

POST /store/checkout

```

Estos procesos deben implementarse como servicios específicos.

---

# 2. Crear la Ruta

Las rutas deben organizarse por módulo funcional.

Ejemplo conceptual:

```
routes/

├── invoice.routes.go
├── product.routes.go
├── client.routes.go

```

No se deben mezclar rutas pertenecientes a distintos dominios.

Ejemplo:

Correcto:

```
/api/invoices
/api/products
/api/clients

```

Incorrecto:

```
/api/business/actions
```

con múltiples responsabilidades mezcladas.

---

# 3. Implementar el Controller

El controller representa la capa HTTP.

Sus responsabilidades son:

- recibir la request.
- validar estructura básica.
- obtener parámetros.
- llamar al servicio correspondiente.
- devolver la respuesta.

El controller NO debe:

- consultar la base de datos.
- implementar reglas de negocio.
- validar estados complejos.
- decidir permisos.

Ejemplo conceptual:

```
Controller

↓

Parse Request

↓

Service.Execute()

↓

Response

```

---

# 4. Crear DTOs

Las entradas y salidas HTTP deben utilizar estructuras propias.

No se recomienda exponer directamente modelos internos de base de datos.

Ejemplo:

Modelo interno:

```
Invoice

- id
- company_id
- internal_status
- created_at

```

DTO público:

```
InvoiceResponse

- id
- status
- total
- items

```

Esto permite evolucionar la base de datos sin romper contratos externos.

---

# 5. Implementar el Service

El Service representa el caso de uso.

Aquí debe vivir:

- lógica de negocio.
- validaciones funcionales.
- coordinación entre módulos.
- manejo de workflows.

Ejemplo:

```
PayInvoiceService

```

Responsabilidades:

- validar estado de invoice.
- verificar monto.
- registrar pago.
- actualizar estado.
- mantener consistencia transaccional.

---

# 6. Uso de Repositories

Los servicios nunca acceden directamente a PostgreSQL.

El acceso se realiza mediante interfaces.

Ejemplo:

```
Service

↓

InvoiceRepository Interface

↓

Postgres Implementation

```

Los repositories únicamente conocen:

- persistencia.
- consultas.
- operaciones de datos.

No conocen reglas del dominio.

---

# 7. Seguridad del Endpoint

Cada endpoint debe definir si requiere:

## Autenticación

La request necesita:

```
JWT válido

↓

Auth Middleware

↓

Tenant Context

```

---

## Autorización

Cuando corresponde:

```
Permission Middleware

↓

Validación Permission

↓

Controller

```

El controller no vuelve a verificar permisos.

La autorización ya fue resuelta previamente.

---

# 8. Uso del Tenant Context

Los endpoints que trabajan con información empresarial deben utilizar el contexto del usuario autenticado.

Ejemplo:

```
Tenant Context

{
 user_id,
 company_id,
 client_id,
 permissions
}

```

Los repositories deben utilizar esta información para garantizar aislamiento.

Nunca debe recibirse un `company_id` confiando únicamente en el request.

---

# 9. Manejo de Errores

Los endpoints no construyen respuestas de error manualmente.

Los errores deben seguir el flujo:

```
Service

↓

Error de negocio

↓

Error Handler

↓

Middleware

↓

Response estándar

```

Ejemplo:

```
invoice no encontrada

↓

AppError

↓

404 Response

```

Errores técnicos:

```
Database error

↓

Internal error response

+

Log completo

```

---

# 10. Respuestas HTTP

Todas las respuestas deben mantener consistencia.

Ejemplo conceptual:

Success:

```
{
 data: {},
 message: "",
 status: 200
}

```

Error:

```
{
 error: true,
 message: "",
 status: 400
}

```

La estructura debe mantenerse para todos los módulos.

---

# 11. Procesos Transaccionales

Si el endpoint modifica múltiples entidades relacionadas debe evaluarse el uso de transacciones.

Ejemplos:

- pagos.
- checkout.
- actualización de stock.
- generación de invoice.

La regla general:

Si una operación debe completarse completamente o fallar completamente, debe ejecutarse dentro de una transacción.

---

# 12. Checklist para Crear un Endpoint

Antes de considerar terminado un endpoint:

## Arquitectura

- [ ] ¿Pertenece realmente a un módulo existente?
- [ ] ¿Representa CRUD o Workflow?
- [ ] ¿La lógica está ubicada en la capa correcta?

## Seguridad

- [ ] ¿Necesita autenticación?
- [ ] ¿Necesita permisos?
- [ ] ¿Respeta Tenant Context?

## Código

- [ ] Controller delgado.
- [ ] Service independiente.
- [ ] Repository desacoplado.
- [ ] DTOs definidos.
- [ ] Errores centralizados.

## Calidad

- [ ] Respuestas consistentes.
- [ ] Logs adecuados.
- [ ] Validaciones completas.
- [ ] Transacciones si corresponde.

---

# Filosofía General

Agregar un endpoint en Nexora no significa agregar una ruta.

Significa incorporar una nueva capacidad al dominio manteniendo las reglas arquitectónicas existentes.

La prioridad siempre es:

```
Dominio

↓

Caso de Uso

↓

Service

↓

Persistencia

↓

HTTP

```

El endpoint es únicamente la puerta de entrada al sistema, no el lugar donde vive la lógica principal.

Mantener esta filosofía permite que Nexora continúe creciendo sin perder mantenibilidad ni consistencia arquitectónica.
