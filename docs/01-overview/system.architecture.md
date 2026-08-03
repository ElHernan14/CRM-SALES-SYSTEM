# System Architecture

## Overview

Nexora fue diseñado como una plataforma empresarial Full Stack orientada a la gestión comercial, integrando funcionalidades ERP, Store y Marketplace bajo una arquitectura modular y escalable.

Desde el inicio del desarrollo se priorizó la construcción de una base arquitectónica preparada para evolucionar, evitando una implementación centrada únicamente en funcionalidades aisladas.

La arquitectura general fue diseñada siguiendo principios de:

- Separation of Concerns
- Layered Architecture
- Domain-Oriented Design
- Clean Architecture principles
- Modularización por responsabilidades
- Bajo acoplamiento entre componentes

El objetivo principal fue construir un sistema donde cada capa y módulo tuviera una responsabilidad claramente definida, permitiendo agregar nuevas funcionalidades sin comprometer la estabilidad del sistema existente.

---

# High-Level Architecture

La arquitectura general de Nexora está compuesta por tres grandes bloques:

```
                Client Applications
                       |
                       |
                 Vue Frontend
                       |
                       |
                  REST API
                       |
                       |
                 Go Backend
                       |
                       |
                PostgreSQL Database
```

Cada bloque posee responsabilidades independientes:

## Frontend

Responsable de:

- Presentación visual.
- Interacción con usuarios.
- Gestión de estado de interfaz.
- Navegación.
- Consumo de casos de uso expuestos por la API.

El frontend no contiene reglas de negocio críticas.

Su responsabilidad es representar el estado del sistema y permitir la interacción del usuario.

---

## Backend

El backend funciona como núcleo del sistema.

Responsable de:

- Autenticación.
- Autorización.
- Validación de reglas de negocio.
- Ejecución de workflows.
- Control de estados.
- Persistencia.
- Integridad de datos.

El backend es la única fuente de verdad del dominio.

---

## Database

La base de datos es responsable únicamente del almacenamiento y persistencia.

No contiene reglas de negocio complejas.

Las decisiones de negocio permanecen dentro de los servicios del backend.

---

# Backend Architecture Overview

El backend fue organizado utilizando una combinación de arquitectura por capas y separación por dominios.

El flujo general de ejecución es:

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
```

Cada capa posee una responsabilidad específica.

---

# Request Lifecycle

Una request atraviesa diferentes etapas antes de completar su ejecución.

Flujo completo:

```
Incoming Request

↓

Public / Private Route

↓

Authentication Middleware

↓

Authorization Middleware

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

## Public Routes

Las rutas públicas no requieren autenticación.

Ejemplos:

- Login.
- Registro.
- Endpoints públicos.

Estas rutas ingresan directamente hacia la capa Controller.

---

## Private Routes

Las rutas privadas requieren:

- JWT válido.
- Usuario autenticado.
- Contexto de tenant.
- Permisos suficientes.

Antes de llegar al Controller atraviesan la capa de seguridad.

---

# Middleware Layer

Los middlewares representan la primera capa de procesamiento de una request.

Sus responsabilidades principales son:

## Authentication Middleware

Responsable de:

- Validar JWT.
- Extraer claims.
- Identificar usuario.
- Construir Tenant Context.

No conoce reglas de negocio.

Su única responsabilidad es establecer identidad.

---

## Authorization Middleware

Responsable de:

- Validar permisos requeridos.
- Comparar permisos del usuario autenticado.
- Permitir o bloquear acceso al endpoint.

La autorización está separada de la autenticación para mantener responsabilidades independientes.

---

## Recovery Middleware

Captura errores inesperados y panics globales.

Su objetivo es evitar que un fallo interno derribe el servidor completo.

---

## Observability Middleware

Responsable de:

- Request ID.
- Seguimiento de request.
- Información temporal.
- Registro del flujo.

Permite reconstruir operaciones durante debugging.

---

# Controller Layer

Los controllers representan la capa de entrada HTTP.

Sus responsabilidades son:

- Recibir requests.
- Parsear parámetros.
- Validar estructura de entrada.
- Convertir datos HTTP a modelos internos.
- Ejecutar servicios correspondientes.
- Construir respuestas HTTP.

Los controllers no contienen lógica de negocio.

Ejemplo conceptual:

```
HTTP Request

↓

Controller

↓

Service

↓

HTTP Response
```

Toda decisión de negocio pertenece al Service Layer.

---

# Service Layer

La capa Service representa el núcleo de aplicación.

Aquí viven:

- Casos de uso.
- Reglas de negocio.
- Validaciones complejas.
- Transiciones de estado.
- Orquestación entre dominios.

Los servicios representan acciones reales del negocio.

Ejemplos:

- Submit Invoice.
- Pay Invoice.
- Checkout.
- Inventory Reservation.
- Dashboard Overview.

Estos procesos no son simples operaciones CRUD.

Representan workflows completos.

---

# Repository Layer

Los repositories son responsables exclusivamente de persistencia.

Sus responsabilidades:

- Consultar información.
- Insertar registros.
- Actualizar datos.
- Ejecutar operaciones SQL.

Los repositories no deciden:

- Estados válidos.
- Permisos.
- Reglas comerciales.
- Flujos de negocio.

La capa Service mantiene el control sobre esas decisiones.

---

# Domain Organization

Los módulos fueron organizados alrededor de conceptos del negocio.

Principales dominios:

```
Auth

User

Company

Client

Product

Inventory

Invoice

Payments

Purchases

Dashboard

Store

Marketplace
```

Cada dominio mantiene sus propias responsabilidades:

- Modelos.
- Servicios.
- Repositories.
- Controllers.
- DTOs.

Esto evita mezclar lógica entre módulos.

---

# Shared Components

Existen componentes compartidos utilizados transversalmente:

Ejemplos:

- Tenant Context.
- Error Handling.
- Response Handling.
- Transaction Management.
- Access Validation.
- Utilities.

Estos componentes no representan dominios de negocio.

Funcionan como infraestructura común.

---

# Dependency Injection

La aplicación utiliza inyección de dependencias para construir los módulos.

El flujo conceptual es:

```
Repository

↓

Service

↓

Controller

↓

Router
```

Esto permite:

- Reemplazar implementaciones.
- Reducir acoplamiento.
- Facilitar testing.
- Mantener responsabilidades claras.

---

# Security Architecture

La seguridad está diseñada alrededor de:

- JWT Authentication.
- RBAC Authorization.
- Tenant Isolation.

El flujo es:

```
JWT

↓

Claims

↓

Tenant Context

↓

Permissions

↓

Endpoint Access
```

Los permisos determinan qué operaciones puede realizar cada usuario.

---

# Multi-Tenant Architecture

Nexora utiliza un modelo multi-tenant basado en base de datos compartida.

El aislamiento se realiza mediante:

```
company_id
```

Cada request autenticada posee un Tenant Context que determina:

- Usuario actual.
- Empresa asociada.
- Cliente asociado.
- Permisos disponibles.

Los repositories utilizan este contexto para asegurar separación de información.

---

# Frontend Architecture

El frontend mantiene una arquitectura alineada con el backend.

Principios utilizados:

- Vue 3 Composition API.
- Pinia State Management.
- API Layer desacoplada.
- Router modular.
- Componentización.
- Organización por dominios.

Flujo:

```
View

↓

Composable

↓

API Layer

↓

Axios

↓

Backend
```

---

# Database Architecture

La persistencia utiliza PostgreSQL.

El modelo está basado en entidades centrales:

- Company.
- User.
- Client.
- Product.
- Invoice.

Invoice funciona como una entidad central del dominio comercial.

A partir de ella se relacionan procesos como:

- Items.
- Payments.
- Stock.
- Estados comerciales.

---

# Architectural Goals Achieved

La arquitectura permitió conseguir:

## Escalabilidad

Agregar nuevos módulos sin modificar grandes partes del sistema.

---

## Mantenibilidad

Cada responsabilidad posee una ubicación definida.

---

## Seguridad

Control centralizado de autenticación, autorización y aislamiento tenant.

---

## Evolución

La estructura permite incorporar nuevos dominios y workflows.

---

# Final Architecture Summary

La arquitectura de Nexora fue diseñada para mantener un equilibrio entre simplicidad y preparación futura.

No busca únicamente resolver los requerimientos actuales, sino establecer una base profesional donde:

- El backend mantiene la lógica de negocio.
- El frontend administra presentación e interacción.
- Los dominios están desacoplados.
- La seguridad es transversal.
- La persistencia está aislada.
- Los workflows complejos poseen servicios propios.
- Los módulos pueden evolucionar independientemente.

Esta organización permite que el sistema pueda crecer desde un MVP avanzado hacia una plataforma empresarial manteniendo claridad arquitectónica.
