# CRM-System-Sales

Scalable multi-tenant ERP & commerce backend built with Go following clean architecture and modular domain-driven design principles.

The project focuses on secure business workflows, invoice management, inventory control, authentication, and financial operations for B2B/B2C environments.

---

# Features

* Multi-tenant architecture
* JWT authentication & authorization
* Role & permission system
* Company & client management
* Product management
* Invoice draft lifecycle
* Invoice item workflows
* Inventory reservation system
* Payment-ready invoice engine
* Secure access scopes
* Transaction-safe operations
* Soft delete & audit fields
* Centralized error handling
* Request tracing & metrics
* RESTful API design

---

# Architecture

The project follows a modular clean architecture approach.

Each domain is separated into independent modules:

* controller
* service
* repository
* dto
* model
* access
* middleware

The application is designed to scale into larger ERP and commerce workflows while maintaining low coupling between domains.

---

# Modules

## Auth

Authentication, JWT tokens, roles and permissions.

## Client

Client management and multi-tenant ownership validation.

## Company

Company administration and tenant isolation.

## Product

Product catalog, stock management and company ownership.

## Invoice

Invoice lifecycle, draft workflows, financial states and payments.

## Invoice Item

Collaborative invoice item management with inventory reservation.

## Inventory

Reserved stock handling and inventory consistency.

## Core

Shared infrastructure:

* middleware
* validators
* observability
* tracing
* response handlers
* access scopes

---

# Technologies

* Go
* PostgreSQL
* Gorilla Mux
* JWT
* Docker
* REST APIs

Architecture & Concepts:

* Clean Architecture
* Modular Monolith
* Multi-Tenancy
* RBAC Authorization
* Transaction Management
* Observability

---

# API Example

## Create Draft Invoice

POST /invoices

Request:

```json
{
  "buyer_client_id": 8,
  "seller_company_id": 2
}
```

Response:

```json
{
  "status": "success",
  "code": 200,
  "data": {
    "id": 15,
    "status_invoice": "draft"
  }
}
```

---

# Observability

The system includes:

* centralized logging
* request tracing
* execution timing
* middleware metrics
* structured error handling

Example trace:

```text
[REQ] method=POST path=/invoices status=200 duration=12ms
→ SERVICE CreateInvoiceDraft
→ REPOSITORY CreateInvoice
```

---

# Security

* JWT authentication
* Role-based permissions
* Multi-tenant access scopes
* Ownership validation
* Transaction-safe financial operations
* Draft state validation
* Soft delete strategy

---

# Run Locally

```bash
git clone <repository>

docker-compose up

go run cmd/api/main.go
```

---

# Future Improvements

* payment gateways
* invoice reconciliation
* analytics dashboards
* event-driven workflows
* notifications
* WebSocket real-time updates
* distributed tracing
* CI/CD pipelines
* Kubernetes deployment
