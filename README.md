# Nexora

## Plataforma SaaS Multi-Tenant ERP + Commerce

Nexora es una plataforma Full Stack diseñada bajo una arquitectura SaaS multi-tenant, orientada a la gestión empresarial, operaciones comerciales y flujos de compra.

El proyecto combina un backend desarrollado en Go con un frontend moderno basado en Vue 3, implementando principios de arquitectura limpia, separación de responsabilidades y diseño orientado a dominios.

La plataforma fue construida pensando en escalabilidad, mantenibilidad y evolución hacia un producto empresarial real.

---

# Características Principales

## Arquitectura Multi-Tenant

Nexora permite que múltiples organizaciones utilicen la misma plataforma manteniendo aislamiento lógico entre sus datos.

Cada operación se ejecuta dentro de un contexto determinado:

- usuario autenticado;
- empresa asociada;
- roles;
- permisos;
- alcance de información.

---

## ERP Empresarial

El módulo ERP permite administrar operaciones internas de una organización.

Incluye:

- gestión de productos;
- categorías;
- clientes;
- inventario;
- ventas;
- compras;
- facturación.

---

## Store Commerce

La plataforma incorpora un flujo comercial orientado al usuario final.

Incluye:

- catálogo;
- carrito;
- checkout;
- compras personales;
- seguimiento de operaciones.

---

## Gestión de Facturación

Las invoices representan operaciones comerciales completas y poseen un ciclo de vida controlado mediante estados.

El sistema implementa:

- validación de transiciones;
- workflows de negocio;
- reglas de consistencia;
- separación entre operaciones organizacionales y personales.

---

# Arquitectura General

```text
                    Nexora Platform


        ┌─────────────────────────┐
        │        Frontend         │
        │       Vue 3 + Vite      │
        └────────────┬────────────┘
                     │
                     │ HTTP API
                     │
        ┌────────────▼────────────┐
        │        Backend          │
        │        Go API           │
        └────────────┬────────────┘
                     │
                     │
        ┌────────────▼────────────┐
        │      PostgreSQL         │
        │     Supabase Storage    │
        └─────────────────────────┘
```

````

---

# Stack Tecnológico

## Backend

- Go
- net/http
- PostgreSQL
- JWT Authentication
- Clean Architecture
- REST API
- Docker

---

## Frontend

- Vue 3
- Vite
- TypeScript
- Pinia
- Vue Router
- Tailwind CSS

---

## Infraestructura

- Docker
- Docker Compose
- Render
- Supabase Storage
- PostgreSQL

---

# Arquitectura Backend

El backend fue organizado siguiendo una separación por capas:

```text
Handler

↓

Service

↓

Repository

↓

Database
```

Cada capa posee responsabilidades definidas:

- Handler: comunicación HTTP.
- Service: reglas de negocio.
- Repository: persistencia.
- Database: almacenamiento.

---

# Arquitectura Frontend

El frontend utiliza una organización modular basada en dominios.

Principales capas:

```text
Views

↓

Components

↓

Composables

↓

Stores

↓

API Layer

↓

Backend
```

Los componentes visuales no contienen lógica de negocio ni comunicación directa con la API.

---

# Seguridad

La plataforma implementa:

- autenticación mediante JWT;
- autorización basada en permisos;
- control multi-tenant;
- validación backend de operaciones;
- separación entre usuarios y organizaciones.

---

# Modelo de Dominio

Los principales dominios del sistema son:

- Identity & Access.
- Multi-Tenant.
- ERP.
- Store.
- Product Catalog.
- Inventory.
- Invoice.
- Customer.
- Payment.

---

# Documentación Técnica

Toda la documentación arquitectónica se encuentra dentro de:

```text
/docs
```

Incluye:

- arquitectura general;
- backend;
- frontend;
- base de datos;
- deployment;
- workflows;
- diagramas.

---

# Ejecución Local

## Backend

```bash
git clone <repository>

cd backend

go mod download

go run main.go
```

---

## Frontend

```bash
cd frontend

npm install

npm run dev
```

---

## Docker Backend

```bash
docker compose up --build
```

---

# Objetivos del Proyecto

Nexora fue desarrollado como una plataforma orientada a demostrar:

- diseño de sistemas SaaS;
- arquitectura Full Stack moderna;
- modelado de dominios complejos;
- implementación de reglas de negocio reales;
- separación profesional entre capas.

---

# Estado del Proyecto

Actualmente Nexora cuenta con:

✅ Backend funcional
✅ Frontend funcional
✅ Arquitectura multi-tenant
✅ Sistema de permisos
✅ Gestión de productos
✅ Gestión comercial
✅ Invoice lifecycle
✅ Dockerización backend
✅ Deploy productivo

---

# Próximas Evoluciones

Posibles extensiones:

- sistema avanzado de reportes;
- analytics;
- notificaciones;
- automatizaciones;
- integraciones externas;
- aplicación móvil.

---

# Autor

Proyecto desarrollado como portfolio profesional Full Stack.

Tecnologías principales:

Go · Vue · PostgreSQL · Docker · SaaS Architecture

```

```
````
