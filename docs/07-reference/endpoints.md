# API Endpoints Reference — Nexora

## Introducción

Nexora expone una API REST organizada por módulos funcionales del dominio.

La organización de endpoints sigue los mismos principios aplicados en la arquitectura general:

- separación por dominios.
- rutas públicas y privadas diferenciadas.
- controllers enfocados en transporte HTTP.
- servicios encargados de lógica de negocio.
- autorización mediante permisos.
- contexto multi-tenant para aislamiento de datos.

La API representa casos de uso del negocio, no solamente operaciones CRUD.

---

# Convenciones Generales

## Base URL

Ejemplo:

```
/api
```

---

## Formato General

Las rutas siguen la estructura:

```
/api/{modulo}/{recurso}/{accion}
```

Ejemplo:

```
/api/invoices/pay
```

---

# Métodos HTTP

## GET

Utilizado para consultas.

Ejemplos:

```
GET /api/products

GET /api/invoices/{id}

```

---

## POST

Utilizado para creación o ejecución de procesos.

Ejemplos:

```
POST /api/invoices

POST /api/invoices/pay

```

---

## PUT

Utilizado para actualización de recursos existentes.

Ejemplo:

```
PUT /api/products/{id}

```

---

## DELETE

Utilizado para eliminación lógica o física según corresponda.

Ejemplo:

```
DELETE /api/products/{id}

```

---

# Autenticación

Los endpoints privados requieren:

```
Authorization: Bearer {JWT}

```

Flujo:

```
Request

↓

JWT Validation

↓

Tenant Context

↓

Permission Middleware

↓

Controller

```

---

# Módulo Authentication

Responsabilidad:

Gestionar identidad y sesión del usuario.

---

## Login

```
POST /api/auth/login

```

Permite autenticar un usuario y generar un JWT.

---

## Usuario autenticado

```
GET /api/auth/me

```

Devuelve el contexto actual del usuario.

Incluye:

- usuario.
- empresa asociada.
- cliente asociado.
- roles.
- permisos.

---

# Módulo Users

Responsabilidad:

Gestión de usuarios del sistema.

Ejemplos:

```
GET /api/users

POST /api/users

PUT /api/users/{id}

```

Operaciones relacionadas:

- creación de usuarios.
- actualización.
- consulta.
- asignación de roles.

---

# Módulo Companies

Responsabilidad:

Gestión de empresas dentro del modelo multi-tenant.

Ejemplos:

```
GET /api/companies

POST /api/companies

PUT /api/companies/{id}

```

Las operaciones respetan el contexto empresarial correspondiente.

---

# Módulo Clients

Responsabilidad:

Gestión de clientes y relaciones comerciales.

Ejemplos:

```
GET /api/clients

POST /api/clients

PUT /api/clients/{id}

```

Los clientes representan actores comerciales asociados a operaciones del sistema.

---

# Módulo Products

Responsabilidad:

Gestión del catálogo de productos y servicios.

Ejemplos:

```
GET /api/products

POST /api/products

PUT /api/products/{id}

DELETE /api/products/{id}

```

Incluye:

- productos.
- servicios.
- categorías.
- tipos de producto.

---

# Módulo Inventory

Responsabilidad:

Control de disponibilidad y movimientos de stock.

Ejemplos:

```
GET /api/inventory

POST /api/inventory/update

```

El inventario mantiene sus propias reglas y no debe ser modificado directamente desde otros dominios.

---

# Módulo Invoice

Responsabilidad:

Gestionar el ciclo comercial principal del sistema.

Invoice representa uno de los Aggregate Roots principales.

---

## Crear Invoice

```
POST /api/invoices

```

Crea una nueva factura.

---

## Consultar Invoice

```
GET /api/invoices/{id}

```

Obtiene información de una factura específica.

---

## Listar Invoices

```
GET /api/invoices

```

Consulta facturas disponibles según contexto y permisos.

---

## Submit Invoice

```
POST /api/invoices/{id}/submit

```

Proceso de negocio encargado de validar y enviar una Invoice.

Incluye:

- validación de estado.
- validación de datos.
- transición de State Machine.

---

## Pay Invoice

```
POST /api/invoices/{id}/pay

```

Proceso encargado de registrar pagos.

Incluye:

- validación de estado.
- control de monto.
- actualización del flujo financiero.

---

# Módulo Invoice Payments

Responsabilidad:

Gestionar pagos asociados a facturas.

Ejemplos:

```
GET /api/invoice-payments

POST /api/invoice-payments

```

Los pagos respetan las reglas definidas por el ciclo de vida de Invoice.

---

# Módulo Store

Responsabilidad:

Gestionar operaciones del entorno comercial personal.

Incluye:

- carrito.
- compras personales.
- checkout.

---

## Cart

Ejemplos:

```
GET /api/store/cart

POST /api/store/cart/items

```

---

## Checkout

```
POST /api/store/checkout

```

Representa un workflow completo:

```
Cart

↓

Validation

↓

Stock

↓

Invoice

↓

Confirmation

```

---

# Módulo Marketplace

Responsabilidad:

Gestionar operaciones comerciales sobre productos publicados.

Incluye:

- catálogo.
- disponibilidad.
- compra.

---

# Módulo Dashboard

Responsabilidad:

Composición de información para vistas generales.

No representa una entidad de dominio propia.

Es un Composition Service.

Ejemplo:

```
GET /api/dashboard/overview

```

Devuelve:

- métricas.
- estados.
- indicadores principales.

---

# Endpoints Públicos

Los endpoints públicos son aquellos que no requieren autenticación.

Ejemplos:

- información pública.
- catálogo limitado.
- recursos necesarios para navegación inicial.

---

# Endpoints Privados

Los endpoints privados requieren:

1. JWT válido.

2. Tenant Context.

3. Permisos suficientes.

4. Validaciones de negocio.

---

# Reglas de Diseño

## Los endpoints no contienen lógica de negocio

Incorrecto:

```
Controller

↓

Validaciones complejas

↓

Cambios de estado

```

Correcto:

```
Controller

↓

Service

↓

Repository

```

---

## Los procesos complejos tienen endpoints propios

Ejemplo:

Incorrecto:

```
PUT /invoice

```

para representar:

- enviar factura.
- pagar.
- completar proceso.

Correcto:

```
POST /invoice/{id}/submit

POST /invoice/{id}/pay

```

---

# Manejo de Respuestas

Todas las respuestas siguen una estructura consistente.

Éxito:

```
{
 data,
 message
}

```

Error:

```
{
 code,
 message
}

```

Los errores son centralizados mediante middleware.

---

# Seguridad

La API garantiza:

- autenticación mediante JWT.
- autorización mediante permisos.
- aislamiento multi-tenant.
- validación de acceso por entidad.
- ocultamiento de información sensible.

---

# Evolución de la API

Para agregar nuevas funcionalidades:

- crear nuevos endpoints.
- mantener contratos existentes.
- evitar cambios incompatibles.
- versionar cuando sea necesario.

La API fue diseñada para crecer agregando nuevos módulos sin modificar dominios existentes.

---

# Resumen Arquitectónico

La API de Nexora está organizada como una capa de entrada al dominio.

Su responsabilidad es:

```
Recibir Request

↓

Validar Seguridad

↓

Ejecutar Caso de Uso

↓

Devolver Response

```

Los endpoints representan acciones reales del negocio y no simples operaciones sobre tablas.

Esta decisión permite mantener una API escalable, mantenible y alineada con la arquitectura interna del sistema.
