# Request Lifecycle

## Overview

El ciclo de vida de una request en Nexora fue diseñado siguiendo una separación estricta de responsabilidades entre capas.

Cada petición atraviesa una serie de etapas claramente definidas, donde cada componente tiene una función específica dentro del procesamiento.

El objetivo principal es garantizar:

- Seguridad.
- Trazabilidad.
- Validación correcta.
- Separación de responsabilidades.
- Consistencia de respuestas.
- Bajo acoplamiento entre capas.

El flujo completo de una request es:

```

Request

↓

Router

↓

Middleware Layer

↓

Controller

↓

Service

↓

Repository

↓

Database

↓

Response

```

---

# Request Entry Point

Toda petición HTTP ingresa inicialmente por la capa de routing.

El router es responsable de:

- Resolver el endpoint solicitado.
- Determinar si la ruta es pública o privada.
- Asociar los middlewares correspondientes.
- Delegar la ejecución al controller adecuado.

El router no contiene lógica de negocio.

Su responsabilidad es únicamente definir el camino que seguirá la petición.

Ejemplo conceptual:

```

POST /api/invoices/pay

↓

Invoice Controller

↓

Pay Invoice Service

```

---

# Middleware Layer

Antes de alcanzar cualquier controller, las requests privadas atraviesan la capa de middleware.

Esta capa implementa funcionalidades transversales que deben ejecutarse antes del caso de uso.

Los principales middlewares son:

- Recovery Middleware.
- Authentication Middleware.
- Permission Middleware.
- Error Middleware.
- Observability Middleware.

---

# Recovery Middleware

El primer objetivo es garantizar estabilidad del servidor.

El middleware de recuperación captura errores inesperados mediante mecanismos de recuperación de panic.

Su responsabilidad es:

- Evitar caída completa del servidor.
- Registrar el error.
- Devolver una respuesta controlada.

Flujo:

```

Panic

↓

Recovery Middleware

↓

Log interno

↓

HTTP 500

```

---

# Authentication Middleware

La autenticación es la primera validación de seguridad.

Su responsabilidad es determinar quién está realizando la petición.

El flujo es:

```

Request

↓

JWT Token

↓

Validate Token

↓

Extract Claims

↓

Build Tenant Context

```

El JWT contiene información necesaria para identificar el contexto del usuario:

- user_id.
- company_id.
- client_id.
- roles.
- permissions.
- email.

Una vez validado, esta información es transformada en un Tenant Context asociado a la request.

---

# Tenant Context Construction

El Tenant Context representa la identidad completa dentro del sistema.

Contiene información necesaria para que las siguientes capas puedan operar correctamente.

Ejemplo conceptual:

```

TenantContext

{
UserID,
CompanyID,
ClientID,
Roles,
Permissions
}

```

Este contexto permite:

- Identificar al usuario.
- Aislar datos por empresa.
- Aplicar reglas de autorización.
- Mantener consistencia durante toda la ejecución.

El contexto acompaña a la request durante todo el flujo.

---

# Permission Middleware

Luego de autenticar al usuario se ejecuta la autorización.

La autenticación responde:

> ¿Quién eres?

La autorización responde:

> ¿Puedes realizar esta operación?

El middleware de permisos:

1. Obtiene los permisos desde Tenant Context.
2. Compara contra los permisos requeridos por el endpoint.
3. Permite o bloquea la ejecución.

Flujo:

```

Tenant Context

↓

Permissions

↓

Endpoint Requirement

↓

Allow / Deny

```

Si el permiso no existe:

```

HTTP 403 Forbidden

```

---

# Controller Layer

Una vez superados los middlewares, la request llega al controller.

El controller representa la capa HTTP del sistema.

Sus responsabilidades son:

- Leer parámetros.
- Parsear JSON.
- Validar estructura de entrada.
- Invocar servicios.
- Construir respuesta HTTP.

El controller no debe:

- Consultar base de datos.
- Ejecutar reglas de negocio.
- Manejar estados complejos.

Ejemplo:

```

HTTP Request

↓

Controller

↓

Service.Call()

```

---

# Service Layer

La capa service representa el núcleo de aplicación.

Es la responsable de ejecutar casos de uso completos.

Aquí viven:

- Reglas de negocio.
- Validaciones funcionales.
- Cambios de estado.
- Orquestación entre dominios.
- Coordinación transaccional.

Ejemplo:

Un pago de invoice no es simplemente:

```

UPDATE invoice

```

El service debe:

1. Validar existencia.
2. Verificar estado actual.
3. Validar monto.
4. Crear payment.
5. Actualizar invoice.
6. Coordinar inventario.
7. Confirmar operación.

El service es el dueño del caso de uso.

---

# Repository Layer

Cuando el service necesita información persistente delega la operación al repository correspondiente.

Los repositories son responsables únicamente de:

- Ejecutar consultas SQL.
- Obtener datos.
- Persistir modificaciones.
- Manejar transacciones cuando corresponde.

No contienen reglas de negocio.

Ejemplo:

```

Service

↓

InvoiceRepository

↓

PostgreSQL

```

---

# Database Layer

La base de datos representa la última capa del flujo.

Su responsabilidad es almacenar información y garantizar integridad mediante:

- Relaciones.
- Constraints.
- Transacciones.
- Consistencia estructural.

La lógica del negocio permanece fuera de esta capa.

---

# Transaction Flow

Las operaciones que modifican múltiples entidades utilizan transacciones SQL.

Ejemplo:

```

Begin Transaction

↓

Update Invoice

↓

Create Payment

↓

Update Inventory

↓

Commit

```

Si ocurre un error:

```

Rollback

↓

Estado anterior restaurado

```

Esto evita datos parciales o inconsistentes.

---

# Response Flow

Luego de completar el caso de uso, la respuesta vuelve siguiendo el camino inverso.

Flujo:

```

Database

↓

Repository

↓

Service

↓

Controller

↓

Middleware

↓

HTTP Response

```

El controller transforma el resultado en una respuesta HTTP consistente.

---

# Error Handling During Lifecycle

Los errores siguen un flujo centralizado.

Existen dos tipos principales:

## Business Errors

Errores esperados del dominio.

Ejemplos:

- Invoice inexistente.
- Estado inválido.
- Permiso insuficiente.
- Operación no permitida.

Estos errores contienen:

- Código HTTP.
- Mensaje funcional.

---

## Technical Errors

Errores inesperados.

Ejemplos:

- Fallo de conexión.
- Error SQL.
- Problemas internos.

Estos errores:

- Se registran internamente.
- Devuelven una respuesta genérica al usuario.

---

# Observability During Lifecycle

Cada request posee información de seguimiento.

Se utiliza:

- Request ID.
- Método HTTP.
- Path.
- Timestamp.
- Resultado final.

Esto permite reconstruir el recorrido completo de una operación.

Ejemplo:

```

Request ID: abc123

POST /api/invoices/pay

↓

Auth Middleware OK

↓

Permission OK

↓

Controller

↓

Service

↓

Repository

↓

Success 200

```

---

# Complete Request Flow

El flujo final de Nexora puede resumirse como:

```

HTTP Request

↓

Router

↓

Recovery Middleware

↓

Authentication Middleware

↓

Tenant Context

↓

Permission Middleware

↓

Controller

↓

Service

↓

Repository

↓

Database

↓

Repository

↓

Service

↓

Controller

↓

HTTP Response

```

---

# Architectural Objective

El diseño del ciclo de vida de una request busca que cada etapa tenga una responsabilidad única.

Esto permite:

- Mantener seguridad centralizada.
- Evitar duplicación de lógica.
- Facilitar debugging.
- Agregar nuevos módulos.
- Mantener consistencia en toda la aplicación.

El resultado es un flujo predecible, escalable y alineado con una arquitectura backend empresarial.

```

```
