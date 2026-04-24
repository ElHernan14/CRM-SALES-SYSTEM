# CRM-SALES-SYSTEM
BACKEND OF CRM SALES SYSTEM
# 🧠 CRM SaaS Multi-Tenant (Go + PostgreSQL)

## 📌 Descripción

Sistema SaaS multi-tenant de gestión de ventas y marketplace B2B/B2C.
Permite que empresas publiquen productos y usuarios compren tanto como individuos como organizaciones.

---

# 🏗️ Arquitectura

## 🔹 Stack

* **Backend:** Go (net/http + gorilla/mux)
* **DB:** PostgreSQL
* **Auth:** JWT + bcrypt
* **Arquitectura:** Clean-ish modular

---

## 🔹 Capas

```text
controller (handler)
    ↓
service (lógica de negocio + seguridad real)
    ↓
repository (acceso a datos)
```

---

## 🔹 Estructura de carpetas (recomendada)

```text
internal/
├── modules/
│   ├── auth/
│   ├── users/
│   ├── clients/
│   ├── companies/
│   ├── products/
│   ├── invoices/
│   └── invoice_items/
│
├── middleware/
│   ├── auth.middleware.go
│   ├── permission.middleware.go
│   └── context.go
│
├── core/                # (opcional pero recomendado)
│   └── tenant.go
│
├── utils/
│   ├── jwt.go
│   └── permissions.go
│
├── database/
│   └── connection.go
│
└── routes/
    └── routes.go
```

---

# 🔐 Autenticación

## 🔹 JWT Claims

```go
type Claims struct {
    UserID      int
    Email       string
    CompanyID   *int
    Roles       []string
    Permissions []string
    jwt.RegisteredClaims
}
```

---

## 🔹 Flujo de Login

```text
Login → validar password (bcrypt)
     → traer roles + permisos desde DB
     → generar JWT
```

---

## 🔹 Generación de Token

* Incluye:

  * user_id
  * email
  * company_id
  * roles
  * permissions

---

# 🧠 Contexto de Usuario (Multi-Tenant)

## 🔹 TenantContext

```go
type TenantContext struct {
    UserID      int
    Email       string
    CompanyID   *int
    Roles       []string
    Permissions []string
}
```

---

## 🔹 Middleware de Auth

Responsabilidades:

* Validar JWT
* Parsear claims
* Construir `TenantContext`
* Inyectarlo en `context.Context`

---

## 🔹 Acceso al contexto

```go
tenant := GetTenant(ctx)
```

---

# 🔑 Permisos

## 🔹 Formato

```text
entity:action
```

Ejemplos:

```text
invoice:view
invoice:read
invoice:create
invoice:pay
```

---

## 🔹 Tipos de permisos

| Tipo                    | Uso               |
| ----------------------- | ----------------- |
| `:view`                 | Frontend (UI)     |
| `:read`                 | Backend           |
| `:create/update/delete` | CRUD              |
| acciones especiales     | (`pay`, `cancel`) |

---

## 🔹 Regla clave

```text
Permisos → permiten ejecutar acción
Service → decide acceso real a los datos
```

---

# 👥 Roles

```text
super_admin
company_user
individual_user
```

---

## 🔹 Descripción

| Rol             | Descripción             |
| --------------- | ----------------------- |
| super_admin     | acceso global           |
| company_user    | opera dentro de empresa |
| individual_user | usuario B2C             |

---

# 🏢 Multi-Tenant

## 🔹 Tipos

### 🧩 B2B (empresa)

* `company_id != null`
* acceso limitado a su empresa

### 🧩 B2C (individual)

* `company_id = null`
* acceso solo a sus propios datos

---

## 🔹 Regla clave

```text
❌ NO filtrar en middleware
✔ SIEMPRE filtrar en service
```

---

# 🔒 Seguridad

## ❌ Nunca confiar en el request

```json
{
  "company_id": 999 ❌
}
```

---

## ✔ Siempre usar contexto

```go
tenant.CompanyID
tenant.UserID
```

---

# 🧩 Ejemplo de Endpoint

```http
GET /clients/{id}/invoices
```

---

## 🔹 Middleware

```go
RequirePermission("invoice:read")
```

---

## 🔹 Service (control real)

```go
if super_admin → acceso total

if company_user → validar company_id

if individual_user → validar user_id
```

---

# 🗄️ Base de Datos (simplificada)

## Tablas principales

```text
users
clients (user_id, company_id nullable)
companies
roles
permissions
role_permissions
user_rol
products
invoices
invoice_items
```

---

# 🔄 Flujo completo

```text
Request
  ↓
AuthMiddleware
  ↓
TenantContext
  ↓
RequirePermission
  ↓
Controller
  ↓
Service (🔥 lógica real)
  ↓
Repository
  ↓
Database
```

---

# 🧠 Decisiones de diseño

## ✔ Permisos simples

* Evitar sobreingeniería (`:own`, `:all`, etc.)
* Resolver alcance en services

## ✔ Multi-tenant en service

* Más flexible
* Más seguro

## ✔ JWT con contexto suficiente

* evita queries innecesarias

---

# 🚀 Estado del sistema

✔ Auth completa (JWT + bcrypt)
✔ Permisos dinámicos desde DB
✔ Middleware desacoplado
✔ Multi-tenant (B2B + B2C)
✔ Arquitectura modular
✔ Base lista para escalar

---

# 🔮 Próximos pasos

* Refresh tokens
* Auditoría (logs de acciones)
* Soft delete
* Cache
* Rate limiting
* Frontend (React/Vue)

---

# 🧠 Conclusión

Este proyecto ya implementa:

* autenticación real
* autorización basada en permisos
* multi-tenancy sólido
* separación de responsabilidades

👉 Base lista para un SaaS real.

---
