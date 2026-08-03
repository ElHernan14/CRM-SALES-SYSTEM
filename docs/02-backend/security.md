# Backend Security Architecture

## Overview

La seguridad de Nexora fue diseñada como una responsabilidad transversal de la aplicación, no como una funcionalidad aislada dentro de un único módulo.

El objetivo principal fue construir un flujo donde cada request tenga una identidad validada, un contexto asociado y permisos explícitos antes de acceder a cualquier caso de uso del sistema.

La arquitectura de seguridad se basa principalmente en:

- Autenticación mediante JWT.
- Autorización basada en permisos (RBAC).
- Construcción de Tenant Context.
- Middleware de seguridad.
- Validación de acceso por dominio.
- Aislamiento multi-tenant.

La responsabilidad de seguridad está distribuida correctamente entre infraestructura, middleware y lógica de dominio, evitando que los controllers o servicios tengan que resolver problemas de autenticación o autorización.

---

# Authentication Flow

La autenticación ocurre al inicio del ciclo de vida de una request.

El flujo general es:

```

Request

↓

Auth Middleware

↓

JWT Validation

↓

Claims Extraction

↓

Tenant Context Creation

↓

Next Middleware / Controller

```

El middleware de autenticación es responsable de:

- Obtener el token JWT recibido.
- Validar firma y expiración.
- Extraer los claims necesarios.
- Construir el contexto asociado al usuario autenticado.

Si el token no es válido, la request es rechazada antes de ingresar al dominio de negocio.

---

# JWT Design

El JWT contiene únicamente información necesaria para reconstruir la identidad y contexto operativo del usuario.

Los principales claims utilizados son:

```text
user_id

email

company_id

client_id

roles

permissions
```

## user_id

Identifica al usuario autenticado.

Permite relacionar la request actual con la entidad usuario dentro del sistema.

---

## company_id

Representa la empresa asociada al usuario.

Es un dato fundamental para el modelo multi-tenant, ya que permite determinar el contexto organizacional sobre el cual se ejecutarán las operaciones.

---

## client_id

Permite identificar usuarios asociados al flujo de clientes finales del sistema.

Esto permite soportar usuarios que no pertenecen necesariamente a una empresa.

---

## permissions

Contiene los permisos disponibles para el usuario autenticado.

Estos permisos son utilizados posteriormente por el middleware de autorización.

La decisión de incluirlos dentro del token evita consultas adicionales constantes a base de datos en cada request.

Como consecuencia, el tiempo de expiración del token debe mantenerse controlado para evitar inconsistencias ante cambios de permisos.

---

# Authorization Model

La autorización está implementada mediante un sistema RBAC basado en permisos.

El flujo es:

```
Request

↓

Permission Middleware

↓

Tenant Context

↓

Permissions Validation

↓

Controller Access
```

El middleware recibe los permisos requeridos por cada endpoint y valida si el usuario autenticado posee autorización suficiente.

Ejemplo conceptual:

```
POST /invoice/pay

Required Permission:

invoice.pay
```

Si el permiso existe dentro del Tenant Context:

```
Allow Request
```

En caso contrario:

```
403 Forbidden
```

---

# Why Authentication and Authorization Are Separate

La autenticación y autorización representan responsabilidades diferentes.

## Authentication

Responde:

> ¿Quién es el usuario?

Responsabilidades:

- Validar JWT.
- Identificar usuario.
- Crear contexto.

---

## Authorization

Responde:

> ¿Este usuario puede realizar esta acción?

Responsabilidades:

- Revisar permisos.
- Validar acceso al endpoint.
- Restringir operaciones.

Separar ambas capas permite mantener un sistema más claro y extensible.

La autenticación ocurre de manera global, mientras que la autorización depende del endpoint específico.

---

# Tenant Context

El Tenant Context es una de las piezas centrales de la arquitectura.

Su responsabilidad es mantener durante toda la request la información necesaria para identificar:

- Usuario actual.
- Empresa asociada.
- Cliente asociado.
- Roles.
- Permisos.

Conceptualmente:

```
JWT

↓

Claims

↓

Tenant Context

↓

Application Flow
```

El contexto evita que cada capa tenga que volver a consultar información de identidad.

---

# Security Responsibility by Layer

## Middleware

Responsable de:

- Autenticación.
- Autorización.
- Validación inicial.
- Construcción del contexto.

No conoce reglas de negocio.

---

## Controller

Responsable de:

- Recibir request.
- Validar estructura de entrada.
- Ejecutar caso de uso correspondiente.

No valida permisos nuevamente porque esa responsabilidad ya fue resuelta anteriormente.

---

## Service

Responsable de:

- Ejecutar reglas de negocio.
- Validar invariantes.
- Aplicar restricciones del dominio.

Puede realizar validaciones adicionales relacionadas al negocio.

---

## Repository

Responsable únicamente de persistencia.

No decide:

- permisos.
- ownership.
- reglas comerciales.

Su función es manipular información siguiendo las condiciones recibidas.

---

# Resource Ownership Validation

Además de permisos generales, ciertos recursos requieren validaciones adicionales.

Ejemplo:

Un usuario puede tener permiso para consultar invoices, pero no necesariamente cualquier invoice del sistema.

Para estos casos existen validaciones de ownership mediante:

- Tenant Context.
- Company ID.
- Client ID.
- Estados de entidad.

Esto evita accesos cruzados entre usuarios o empresas.

---

# Security and Multi-Tenancy

La seguridad está directamente relacionada con el aislamiento multi-tenant.

El flujo completo garantiza:

```
User Identity

↓

Tenant Context

↓

Repository Filtering

↓

Database Query
```

Cada operación sobre datos sensibles utiliza información obtenida desde el contexto actual.

Esto evita que un usuario pueda acceder accidentalmente a información perteneciente a otro tenant.

---

# Security Design Benefits

La arquitectura implementada permite:

- Agregar nuevos permisos sin modificar lógica existente.
- Incorporar nuevos módulos con autorización independiente.
- Mantener controllers simples.
- Evitar duplicación de validaciones.
- Separar seguridad de negocio.
- Mantener preparado el sistema para crecimiento futuro.

---

# Future Improvements

Para una versión productiva de mayor escala podrían incorporarse:

- Refresh tokens.
- Rotación de JWT.
- Revocación de sesiones.
- Auditoría de accesos.
- Políticas dinámicas de autorización.
- Integración con proveedores externos de identidad.
- Gestión avanzada de roles.

---

# Final Considerations

La seguridad de Nexora fue diseñada siguiendo una separación clara de responsabilidades.

La autenticación establece identidad.

La autorización determina capacidades.

El Tenant Context mantiene consistencia durante toda la operación.

Esta separación permite que la aplicación mantenga un modelo seguro, escalable y preparado para evolucionar junto con nuevos módulos y dominios de negocio.

```

```
