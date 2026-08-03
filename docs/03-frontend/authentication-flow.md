# Frontend Authentication Flow

## Overview

El flujo de autenticación del frontend de Nexora fue diseñado para integrarse con el modelo de seguridad implementado en backend.

La responsabilidad principal del frontend es gestionar:

- Inicio de sesión.
- Persistencia del token.
- Recuperación del contexto del usuario.
- Gestión del estado autenticado.
- Protección de navegación.

La autenticación real y autorización definitiva permanecen en backend.

El frontend no decide si un usuario puede ejecutar una operación de negocio, únicamente utiliza la información recibida para adaptar la interfaz.

---

# Authentication Principles

Durante el diseño del flujo se siguieron los siguientes principios:

## Backend as Security Authority

El backend es la única fuente de verdad respecto a:

- Identidad del usuario.
- Empresa asociada.
- Cliente asociado.
- Permisos.
- Roles.
- Acceso a recursos.

El frontend nunca genera permisos ni modifica reglas de acceso.

---

## Centralized Authentication State

Toda la información relacionada con sesión se mantiene centralizada.

El objetivo fue evitar:

- Estados duplicados.
- Información inconsistente entre componentes.
- Lecturas repetidas del token.

La aplicación trabaja sobre un único contexto autenticado.

---

# Authentication Flow Overview

El flujo completo puede representarse de la siguiente manera:

```

User Login

↓

Login View

↓

Auth Composable

↓

Auth API

↓

Backend /auth/login

↓

JWT Response

↓

Persist Token

↓

Request /auth/me

↓

Tenant Context

↓

Pinia Store

↓

Router Navigation

↓

Protected Application

```

Cada etapa tiene una responsabilidad específica.

---

# Login Process

El proceso inicia cuando el usuario ingresa sus credenciales.

Flujo:

```

User

↓

Login Component

↓

useAuth()

↓

auth.api.ts

↓

POST /auth/login

↓

Backend Validation

```

El frontend solamente envía las credenciales recibidas.

No realiza validaciones de seguridad propias.

---

# JWT Handling

Cuando el backend valida correctamente las credenciales devuelve un JWT.

El frontend realiza:

- Recepción del token.
- Persistencia local.
- Configuración del cliente HTTP.
- Recuperación del contexto.

El token representa una sesión autenticada.

Sin embargo, el frontend no interpreta directamente la lógica interna del token.

---

# Token Persistence

El token es almacenado para mantener la sesión entre navegaciones y recargas.

El proceso general:

```

JWT

↓

Storage

↓

Axios Authorization Header

↓

Authenticated Requests

```

La persistencia permite reconstruir el estado de la aplicación al iniciar nuevamente.

---

# Authentication Bootstrap

Uno de los puntos importantes del diseño fue separar:

Login inicial

de

Restauración de sesión existente.

Al cargar la aplicación se ejecuta un proceso de bootstrap.

Flujo:

```

Application Start

↓

Check Stored Token

↓

Set Auth State

↓

Request /auth/me

↓

Load User Context

↓

Initialize Application

```

Esto evita depender únicamente del login inicial.

---

# /auth/me Flow

Después de obtener el token, el frontend solicita la información completa del usuario autenticado.

Ejemplo conceptual:

```

GET /auth/me

```

La respuesta contiene el contexto necesario para operar:

- Usuario.
- Email.
- Company.
- Client.
- Roles.
- Permissions.
- Tenant Information.

Este endpoint funciona como fuente del contexto actual de sesión.

---

# Tenant Context Integration

El frontend utiliza el concepto de Tenant Context proveniente del backend.

Este contexto permite conocer:

```

¿Quién es el usuario?

¿A qué empresa pertenece?

¿Qué permisos posee?

¿Qué módulos puede utilizar?

```

El frontend utiliza esta información para:

- Mostrar módulos disponibles.
- Configurar navegación.
- Controlar experiencia visual.

---

# Authentication Store

La información autenticada se almacena mediante Pinia.

Responsabilidades principales:

- Mantener usuario actual.
- Mantener token.
- Guardar Tenant Context.
- Exponer estado autenticado.
- Controlar loading inicial.

El store no contiene reglas de negocio.

Solamente administra estado de sesión.

---

# Authorization in Frontend

Aunque existe control visual basado en permisos, la autorización real continúa siendo responsabilidad del backend.

El frontend utiliza permisos para:

- Ocultar opciones no disponibles.
- Evitar navegación innecesaria.
- Mejorar experiencia.

Ejemplo:

Un usuario sin permiso de crear invoices puede no visualizar el botón correspondiente.

Sin embargo:

```

Frontend Permission Check

≠

Backend Authorization

```

El backend siempre valida nuevamente.

---

# Route Protection

Las rutas privadas utilizan guards.

Flujo:

```

Navigation Attempt

↓

Check Authentication

↓

Check Context

↓

Check Permission

↓

Allow / Redirect

```

El router únicamente protege la navegación.

---

# Authentication States

La aplicación contempla diferentes estados:

## Initial Loading

La aplicación todavía está verificando sesión.

Ejemplo:

```

Checking authentication...

```

---

## Authenticated

Existe:

- Token válido.
- Usuario cargado.
- Contexto disponible.

---

## Unauthenticated

No existe una sesión válida.

Acción:

- Redirección a login.

---

## Authentication Error

La sesión existe pero no puede restaurarse.

Ejemplo:

- Token expirado.
- Usuario eliminado.
- Error backend.

Acción:

- Limpieza de sesión.
- Redirección segura.

---

# Logout Flow

El cierre de sesión elimina completamente el contexto autenticado.

Flujo:

```

Logout

↓

Remove Token

↓

Clear Store

↓

Reset Context

↓

Redirect Login

```

No quedan datos residuales de sesión.

---

# Security Considerations

Durante el diseño se contemplaron las siguientes decisiones:

## No Business Logic in Frontend

El frontend no decide:

- Estados válidos de entidades.
- Acceso definitivo.
- Reglas comerciales.

---

## No Trust Client State

Aunque el frontend almacene permisos, estos no representan una autorización real.

Toda operación sensible vuelve a validarse en backend.

---

## Expiration Handling

Cuando el token deja de ser válido:

- Se limpia sesión.
- Se informa al usuario.
- Se solicita nueva autenticación.

---

# Relationship With Backend Security

El flujo completo entre frontend y backend queda:

```

Frontend

↓

JWT

↓

Backend Auth Middleware

↓

Claims

↓

Tenant Context

↓

Permission Middleware

↓

Controller

↓

Service

```

El frontend inicia la sesión, pero backend controla la seguridad.

---

# Architectural Evaluation

El sistema de autenticación del frontend fue diseñado siguiendo principios de desacoplamiento y seguridad empresarial.

Las principales decisiones fueron:

- Centralizar estado autenticado.
- Consumir contexto generado por backend.
- Separar autenticación de autorización.
- Evitar lógica de negocio en cliente.
- Utilizar guards únicamente como capa de experiencia.
- Mantener una única fuente de verdad.

Esta arquitectura permite que nuevos módulos frontend puedan incorporarse utilizando el mismo modelo de seguridad sin duplicar lógica ni crear flujos paralelos.

```

```
