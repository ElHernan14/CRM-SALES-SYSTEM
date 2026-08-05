# NEXORA - BACKEND

Scalable multi-tenant SaaS ERP & commerce backend built with Go following clean architecture and modular domain-driven design principles.

The platform is designed for secure business operations across B2B and B2C environments, including invoice workflows, inventory management, authentication, financial operations, and tenant-isolated company management.

The system focuses on scalability, security, maintainability, and enterprise-grade backend architecture.

---

# Features

- Multi-tenant SaaS architecture
- JWT authentication & authorization
- Role & permission system (RBAC)
- Company & client management
- Product catalog & inventory management
- Invoice draft lifecycle workflows
- Collaborative invoice item handling
- Inventory reservation system
- Payment-ready invoice engine
- Secure ownership & access scopes
- Transaction-safe operations
- Soft delete & audit fields
- Centralized error handling
- Request tracing & observability
- RESTful API architecture

---

# Architecture

The project follows a modular clean architecture approach designed for scalability and low coupling between business domains.

Each module is internally organized into independent layers:

- controller
- service
- repository
- dto
- model
- access
- middleware

The backend currently follows a modular monolith architecture with workflow orchestration patterns for complex financial operations.

---

# Modules

## Auth

Authentication, JWT tokens, roles and permission management.

## Client

Client management with ownership validation and tenant isolation.

## Company

Company administration and multi-tenant business structure.

## Product

Product catalog, stock management and company ownership validation.

## Invoice

Invoice lifecycle management, draft workflows, financial states and payment processing.

## Invoice Item

Collaborative invoice item handling with stock reservation logic.

## Inventory

Inventory consistency and reserved stock operations.

## Core

Shared infrastructure and cross-cutting concerns:

- middleware
- validators
- observability
- tracing
- response handlers
- access scopes

---

# Technologies

### Backend

- Go
- PostgreSQL
- Gorilla Mux
- JWT
- REST APIs

### Architecture & Concepts

- Clean Architecture
- Modular Monolith
- Workflow Orchestration
- Multi-Tenancy
- RBAC Authorization
- Transaction Management
- Observability

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

### Observability

## The system includes:

- centralized logging
- request tracing
- execution timing
- middleware metrics
- structured error handling

# Example trace:

```[REQ] method=POST path=/invoices status=200 duration=12ms
→ SERVICE CreateInvoiceDraft
→ REPOSITORY CreateInvoice
```

---

### Security

- JWT authentication
- Role-based permissions
- Multi-tenant access scopes
- Ownership validation
- Transaction-safe financial operations
- Draft state validation
- Soft delete strategy

### Project Vision

## CRM-System-Sales is designed as a scalable backend foundation for future ERP, commerce, and SaaS ecosystems.

# The architecture aims to support:

- enterprise workflows
- financial operations
- inventory systems
- analytics
- workflow automation
- real-time business management

# while maintaining a clean, modular, and extensible codebase.
