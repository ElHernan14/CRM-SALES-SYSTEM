# Permissions Reference — Nexora

## Introducción

Nexora utiliza un modelo de autorización basado en permisos mediante RBAC (Role Based Access Control).

Aunque el sistema contempla roles, la autorización real de acceso a funcionalidades se realiza mediante permisos específicos.

La decisión de utilizar permisos como unidad principal de autorización permite:

- mayor granularidad.
- menor acoplamiento entre usuarios y funcionalidades.
- escalabilidad del modelo de seguridad.
- incorporación de nuevos módulos sin modificar reglas existentes.

El flujo general es:

```
Usuario

↓

Roles

↓

Permisos

↓

JWT Claims

↓

Tenant Context

↓

Permission Middleware

↓

Endpoint

```

---

# Modelo de Seguridad

El modelo está compuesto por:

```
User

↓

User Role

↓

Role Permission

↓

Permission

```

Cada usuario puede tener uno o múltiples roles.

Cada rol agrupa permisos.

Los permisos determinan qué acciones puede ejecutar un usuario.

---

# Roles

Los roles representan agrupaciones funcionales.

Ejemplo conceptual:

```
ADMIN

SALES_MANAGER

EMPLOYEE

CUSTOMER

```

Sin embargo, los roles no son utilizados directamente para proteger endpoints.

Su responsabilidad principal es administrar conjuntos de permisos.

---

# Permissions

Los permisos representan acciones concretas dentro del sistema.

Ejemplos:

```
invoice.read

invoice.create

invoice.update

invoice.submit

invoice.pay

product.read

product.create

inventory.update

```

Cada endpoint puede requerir uno o más permisos específicos.

---

# Autenticación vs Autorización

Nexora separa ambos conceptos.

## Autenticación

Responde:

"¿Quién es este usuario?"

Proceso:

```
JWT

↓

Auth Middleware

↓

Validación Token

↓

Tenant Context

```

---

## Autorización

Responde:

"¿Puede este usuario realizar esta acción?"

Proceso:

```
Tenant Context

↓

Permissions

↓

Permission Middleware

↓

Controller

```

---

# JWT Claims

El token JWT contiene información necesaria para construir el contexto del usuario.

Ejemplo conceptual:

```
{
 user_id,
 email,
 company_id,
 client_id,
 roles,
 permissions
}

```

Los permisos dentro del JWT evitan consultas adicionales a la base de datos en cada request.

Esto mejora rendimiento y reduce carga sobre la infraestructura.

El tiempo de expiración del token debe mantenerse controlado para reducir riesgos asociados.

---

# Tenant Context

Luego de validar el JWT se construye el Tenant Context.

Este objeto representa la identidad completa de la request.

Incluye información como:

```
User ID

Company ID

Client ID

Roles

Permissions

```

Todas las capas posteriores utilizan este contexto cuando necesitan conocer:

- usuario actual.
- empresa asociada.
- permisos disponibles.
- alcance de acceso.

---

# Permission Middleware

El middleware de permisos se ejecuta antes de ingresar al controller.

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

Su responsabilidad es:

- obtener permisos del contexto.
- comparar contra permisos requeridos.
- permitir o rechazar acceso.

El controller no vuelve a realizar esta validación.

---

# Permisos por Dominio

Los permisos siguen la organización modular del sistema.

Ejemplo:

## Invoice

```
invoice.read

invoice.create

invoice.update

invoice.submit

invoice.pay

```

---

## Products

```
product.read

product.create

product.update

product.delete

```

---

## Inventory

```
inventory.read

inventory.update

```

---

## Clients

```
client.read

client.create

client.update

```

---

# Administración Global

Existe un concepto administrativo especial.

Los usuarios administrativos pueden tener acceso completo mediante validación de rol.

Este comportamiento se utiliza únicamente en escenarios específicos y no reemplaza el modelo general basado en permisos.

La arquitectura mantiene los permisos como mecanismo principal de autorización.

---

# Reglas de Diseño

## No validar permisos dentro del Controller

Incorrecto:

```
if user.Permission == "invoice.pay"

```

El controller solamente ejecuta el caso de uso autorizado.

---

## No confiar en información enviada por Frontend

El frontend puede ocultar funcionalidades visualmente.

Pero la seguridad real siempre pertenece al backend.

---

## No utilizar únicamente roles

Los roles son útiles para agrupación.

Los permisos son necesarios para autorización granular.

---

# Agregar un Nuevo Permiso

Cuando se incorpora una nueva funcionalidad:

1. Definir el permiso.

Ejemplo:

```
invoice.cancel

```

2. Registrar el permiso en base de datos.

3. Asociarlo al rol correspondiente.

4. Proteger el endpoint mediante middleware.

5. Actualizar documentación.

---

# Filosofía del Sistema

El modelo de permisos de Nexora busca mantener un equilibrio entre:

- seguridad.
- flexibilidad.
- mantenibilidad.
- escalabilidad.

La autorización no depende de lógica dispersa dentro del código, sino de un flujo centralizado:

```
JWT

↓

Tenant Context

↓

Permissions

↓

Middleware

↓

Endpoint autorizado

```

Este diseño permite agregar nuevos módulos y funcionalidades manteniendo una política de acceso consistente.
