# Frontend Routing Architecture

## Overview

El sistema de navegación de Nexora fue diseñado utilizando Vue Router como capa responsable de gestionar la navegación dentro de la aplicación.

La arquitectura de rutas fue organizada siguiendo los mismos principios utilizados en el resto del frontend:

- Separación de responsabilidades.
- Organización modular.
- Protección de navegación.
- Bajo acoplamiento.
- Escalabilidad por dominio.

El router no contiene lógica de negocio.

Su responsabilidad es determinar:

- Qué vistas existen.
- Qué rutas son públicas.
- Qué rutas requieren autenticación.
- Qué rutas requieren contexto específico.

La seguridad definitiva permanece en backend.

---

# Routing Principles

Durante el diseño del sistema de rutas se establecieron los siguientes principios:

## Routes Represent Application Navigation

Las rutas representan caminos dentro de la aplicación, no procesos de negocio.

Ejemplo:

Correcto:

```

/erp/invoices

```

Representa una sección de la aplicación.

Incorrecto:

```

/invoice/pay

```

Porque el pago es una operación de negocio que pertenece al backend.

---

## Router Does Not Replace Backend Authorization

El frontend puede impedir navegación innecesaria, pero nunca reemplaza la autorización real.

Flujo correcto:

```

Frontend Guard

↓

Backend Middleware

↓

Permission Validation

↓

Service Authorization

```

La protección frontend mejora UX.

La seguridad backend protege el sistema.

---

# Route Organization

Las rutas fueron organizadas por módulos funcionales.

Conceptualmente:

```

router/

├── index.ts

├── auth.routes.ts

├── erp.routes.ts

├── store.routes.ts

└── marketplace.routes.ts

```

Cada módulo mantiene sus propias rutas evitando concentrar toda la navegación en un único archivo.

---

# Public Routes

Las rutas públicas son aquellas disponibles sin sesión autenticada.

Ejemplos:

```

/login

/register

/recovery

```

Estas rutas no requieren:

- JWT.
- Tenant Context.
- Permisos.

---

# Protected Routes

Las rutas privadas requieren una sesión válida.

Flujo:

```

Navigation Attempt

↓

Check Token

↓

Load Authentication State

↓

Validate Context

↓

Allow Access

```

Si no existe una sesión válida:

```

Redirect → Login

```

---

# Authentication Guard

El guard de autenticación verifica que exista una sesión activa.

Responsabilidades:

- Confirmar existencia de usuario autenticado.
- Validar estado inicial de aplicación.
- Evitar acceso a módulos privados.

No realiza:

- Validaciones de negocio.
- Consultas directas a base de datos.
- Reglas complejas.

---

# Permission Based Navigation

Además de autenticación, ciertas rutas requieren permisos específicos.

Ejemplo conceptual:

```

ERP Invoice Management

requires:

invoice.read

```

El guard consulta el contexto autenticado:

```

TenantContext

↓

Permissions

↓

Route Access

```

Si el permiso no existe:

```

Redirect / Forbidden

```

---

# Tenant Based Navigation

Debido al modelo multi-tenant, la navegación puede cambiar según el tipo de usuario.

Ejemplo:

Usuario empresa:

```

ERP Dashboard

Invoices

Inventory

Purchases

```

Usuario cliente:

```

Store

Orders

Profile

```

El router utiliza la información obtenida desde backend para construir una experiencia acorde al contexto.

---

# Application Initialization

Antes de permitir navegación completa, la aplicación debe reconstruir el estado inicial.

Flujo:

```

Application Start

↓

Bootstrap Auth

↓

Check Stored Token

↓

Request /auth/me

↓

Load Tenant Context

↓

Initialize Router

↓

Render Application

```

Esto evita mostrar interfaces incorrectas durante la carga inicial.

---

# Layout Routing

La aplicación utiliza layouts para separar experiencias completas.

Ejemplo:

```

PublicLayout

```

Login

```

ERPLayout

```

Dashboard
Invoices
Inventory

```

StoreLayout

```

Products
Cart
Orders

```

```

Los layouts contienen estructura general.

Las vistas contienen contenido específico.

---

# Nested Routes

La estructura permite utilizar rutas anidadas.

Ejemplo conceptual:

```

/erp

```

/dashboard

/invoices

/products

/inventory

```

```

Esto permite mantener:

- Navegación consistente.
- Layout compartido.
- Organización modular.

---

# Route Metadata

Las rutas pueden contener información adicional mediante metadata.

Ejemplo:

```ts
{
    path: "/invoices",
    meta: {
        requiresAuth: true,
        permission: "invoice.read"
    }
}
```

Esto permite que los guards sean genéricos y reutilizables.

---

# Navigation Flow

El flujo completo de navegación:

```
User Click

↓

Vue Router

↓

Route Guard

↓

Authentication Check

↓

Permission Check

↓

Layout Selection

↓

View Rendering
```

Cada etapa tiene una responsabilidad concreta.

---

# Error Navigation Handling

El sistema contempla escenarios donde una navegación no puede completarse.

Ejemplos:

## Usuario no autenticado

```
Redirect Login
```

## Usuario autenticado sin permiso

```
Redirect Forbidden
```

## Ruta inexistente

```
404 Page
```

---

# Relationship With Backend

La relación entre navegación frontend y seguridad backend es:

```
Frontend Router

↓

User Experience Layer


Backend Middleware

↓

Security Layer


Backend Service

↓

Business Rules
```

Cada capa protege una responsabilidad diferente.

---

# Scalability

La arquitectura de routing permite agregar nuevos módulos fácilmente.

Nuevo módulo:

```
New Domain

↓

New Route File

↓

New Views

↓

New Layout if required

↓

New Permissions
```

Sin modificar rutas existentes.

---

# Architectural Evaluation

El sistema de routing de Nexora fue diseñado como una capa independiente de navegación y experiencia de usuario.

Las principales decisiones fueron:

- Separación de rutas por módulos.
- Uso de guards reutilizables.
- Integración con Tenant Context.
- Protección basada en permisos.
- Uso de layouts compartidos.
- Separación entre navegación y seguridad.

Este diseño permite que la aplicación pueda incorporar nuevos dominios manteniendo una navegación organizada, consistente y alineada con la arquitectura general del sistema.

```

```
