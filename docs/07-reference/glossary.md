# Glossary — Nexora

## Introducción

Este documento contiene la definición de los principales conceptos técnicos y arquitectónicos utilizados dentro del proyecto Nexora.

El objetivo es mantener un lenguaje común entre desarrollo, documentación y evolución futura del sistema.

Los términos están definidos según su aplicación concreta dentro de la arquitectura del proyecto.

---

# Aggregate Root

Entidad principal dentro de un dominio que controla la consistencia y ciclo de vida de otras entidades relacionadas.

En Nexora, Invoice funciona como uno de los principales Aggregate Roots.

Ejemplo:

```
Invoice

├── Invoice Items

└── Invoice Payments

```

Las entidades relacionadas no deben modificar reglas críticas fuera del agregado principal.

---

# API Layer

Capa encargada de encapsular la comunicación HTTP entre frontend y backend.

Responsabilidades:

- endpoints.
- requests.
- responses.
- serialización.
- configuración HTTP.

En frontend permite evitar que los componentes interactúen directamente con Axios.

---

# Application Service

Servicio encargado de coordinar operaciones del sistema.

Normalmente representa casos de uso completos.

Ejemplos:

- Submit Invoice.
- Pay Invoice.
- Checkout.

No representa simplemente operaciones CRUD.

---

# Authentication

Proceso encargado de validar la identidad del usuario.

Responde:

> ¿Quién está realizando esta petición?

En Nexora se realiza mediante JWT.

Flujo:

```
JWT

↓

Auth Middleware

↓

Tenant Context

```

---

# Authorization

Proceso encargado de determinar si un usuario tiene permiso para ejecutar una acción.

Responde:

> ¿Este usuario puede realizar esta operación?

En Nexora se implementa mediante permisos y middleware de autorización.

---

# Controller

Capa responsable de manejar la entrada y salida HTTP.

Responsabilidades:

- recibir requests.
- validar formato básico.
- llamar servicios.
- devolver responses.

No contiene lógica de negocio.

---

# DTO (Data Transfer Object)

Objeto utilizado para transportar información entre capas.

Se utiliza para:

- requests.
- responses.
- comunicación entre módulos.

Permite evitar exponer directamente modelos internos.

---

# Domain Module

Módulo que representa una responsabilidad de negocio específica.

Ejemplos:

- Invoice.
- Product.
- Client.
- Inventory.

Cada módulo contiene sus propias responsabilidades y evita depender innecesariamente de otros.

---

# Entity

Objeto que representa una entidad del dominio.

Posee:

- identidad.
- estado.
- reglas asociadas.

Ejemplos:

```
Invoice

Product

Client

Company

```

---

# JWT

JSON Web Token utilizado para representar la identidad autenticada.

En Nexora contiene información necesaria para construir el contexto del usuario.

Ejemplo:

```
user_id

company_id

client_id

permissions

roles

```

---

# Middleware

Componente que intercepta requests antes o después de llegar al controlador.

En Nexora se utiliza para:

- autenticación.
- autorización.
- manejo de errores.
- recuperación de panics.
- observabilidad.

---

# Multi-Tenancy

Arquitectura que permite que múltiples organizaciones utilicen el mismo sistema manteniendo aislamiento de información.

En Nexora se implementa mediante:

```
company_id

+

Tenant Context

+

Repository Filtering

```

---

# Tenant Context

Objeto que representa la identidad completa de una petición.

Contiene información como:

- usuario.
- empresa.
- cliente.
- permisos.
- roles.

Permite mantener el aislamiento multi-tenant durante todo el flujo.

---

# RBAC

Role Based Access Control.

Modelo de autorización basado en roles.

Nexora utiliza RBAC combinado con permisos.

Modelo:

```
User

↓

Role

↓

Permission

```

---

# Permission

Unidad mínima utilizada para autorizar acciones dentro del sistema.

Ejemplo:

```
invoice.pay

product.create

inventory.update

```

Los endpoints validan permisos, no roles directamente.

---

# Repository

Capa encargada del acceso a datos.

Responsabilidades:

- consultas SQL.
- persistencia.
- lectura de datos.

No contiene reglas de negocio.

---

# Service

Capa donde vive la lógica de negocio.

Responsabilidades:

- reglas del dominio.
- validaciones.
- coordinación entre repositorios.
- ejecución de casos de uso.

---

# Shared Module

Módulo transversal utilizado por diferentes dominios.

Ejemplos:

- respuestas HTTP.
- errores.
- contexto.
- validadores.

Su objetivo es evitar duplicación manteniendo bajo acoplamiento.

---

# State Machine

Modelo utilizado para controlar estados y transiciones válidas.

En Nexora Invoice posee un ciclo de vida definido.

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

Cada transición posee reglas propias.

---

# Workflow

Proceso de negocio compuesto por múltiples pasos.

No representa una operación CRUD simple.

Ejemplos:

## Pay Invoice

```
Validate Invoice

↓

Validate Amount

↓

Register Payment

↓

Update Status

```

---

# Composition Service

Servicio encargado de combinar información de múltiples dominios para una necesidad específica.

Ejemplo:

Dashboard Overview.

Puede consultar:

- ventas.
- pagos.
- inventario.
- métricas.

No representa una entidad propia del dominio.

---

# Clean Architecture

Patrón arquitectónico basado en separación de responsabilidades.

Busca que:

- dominio no dependa de infraestructura.
- reglas de negocio sean independientes.
- módulos sean reemplazables.

---

# Layered Architecture

Arquitectura organizada por capas.

En Nexora:

```
Controller

↓

Service

↓

Repository

↓

Database

```

Cada capa tiene responsabilidades definidas.

---

# Dependency Injection

Patrón utilizado para entregar dependencias a los componentes desde fuera.

Permite:

- menor acoplamiento.
- mejor testabilidad.
- reemplazo de implementaciones.

---

# State

Estado actual de una entidad dentro de su ciclo de vida.

Ejemplo:

```
Invoice.status

```

El estado determina operaciones permitidas.

---

# Soft Delete

Estrategia donde un registro no se elimina físicamente.

En lugar de borrar información:

```
deleted_at = timestamp

```

permite mantener historial.

---

# Observability

Capacidad de entender qué ocurre dentro del sistema.

Incluye:

- logs.
- request_id.
- seguimiento de errores.
- métricas.

---

# Request ID

Identificador único asociado a una petición.

Permite reconstruir el recorrido de una operación dentro del sistema.

Útil para:

- debugging.
- auditoría.
- análisis de errores.

---

# PWA

Progressive Web Application.

Modelo de aplicación web que permite características similares a una aplicación instalada.

Nexora frontend fue preparado para soportar evolución hacia este modelo.

---

# Frontend Composition API

Modelo de organización lógica utilizado en Vue 3.

Permite encapsular comportamiento reutilizable mediante composables.

Ejemplo:

```
useAuth()

useProducts()

useInvoices()

```

---

# Store

Estado global administrado por una librería de gestión de estado.

En Nexora frontend se utiliza Pinia.

Contiene información compartida como:

- usuario autenticado.
- tenant context.
- permisos.

---

# Layout

Estructura visual reutilizable del frontend.

Ejemplos:

```
PublicLayout

ERPLayout

StoreLayout

```

Permite separar navegación y contenido.

---

# Resumen

El vocabulario técnico de Nexora refleja las decisiones arquitectónicas principales:

- dominio separado.
- seguridad centralizada.
- multi-tenancy.
- workflows de negocio.
- bajo acoplamiento.
- escalabilidad.

Mantener un lenguaje común permite que el sistema continúe evolucionando sin perder consistencia arquitectónica.
