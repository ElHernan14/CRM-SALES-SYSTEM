# Backend Architecture

## Overview

El backend de Nexora fue diseñado bajo una arquitectura modular orientada a dominios, combinando principios de Layered Architecture, Clean Architecture y separación de responsabilidades.

El objetivo principal fue construir una API empresarial mantenible y escalable, donde cada capa tenga una responsabilidad claramente definida y donde la lógica de negocio permanezca desacoplada de detalles externos como HTTP, persistencia o infraestructura.

La arquitectura busca evitar un crecimiento desordenado del sistema, permitiendo incorporar nuevos módulos y procesos de negocio sin afectar funcionalidades existentes.

---

# Architectural Principles

El backend fue construido siguiendo los siguientes principios:

## Separation of Responsibilities

Cada capa posee una responsabilidad específica.

Las reglas principales son:

- Middleware gestiona seguridad y aspectos transversales.
- Controllers gestionan transporte HTTP.
- Services contienen casos de uso y reglas de negocio.
- Repositories gestionan persistencia.
- Database almacena información.

Ninguna capa invade responsabilidades de otra.

---

## Domain-Oriented Organization

La estructura del backend está organizada alrededor de dominios funcionales y no únicamente alrededor de tipos técnicos.

Los módulos principales representan conceptos reales del negocio:

- Authentication.
- Users.
- Companies.
- Clients.
- Products.
- Inventory.
- Invoices.
- Payments.
- Dashboard.

Cada módulo mantiene su propia lógica y contratos.

---

## Dependency Direction

La dirección de dependencia sigue un flujo controlado:

Delivery Layer

↓

Application Layer

↓

Domain Logic

↓

Infrastructure Layer

Las capas superiores conocen las inferiores mediante contratos definidos.

Las capas inferiores no dependen de detalles específicos de las superiores.

---

# High Level Architecture

El flujo general del backend es:

HTTP Request

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

Cada etapa tiene una responsabilidad concreta dentro del ciclo de vida de la petición.

---

# Layer Responsibilities

## Router Layer

### Responsibility

Define cómo una petición entra al sistema.

Responsabilidades:

- Registro de endpoints.
- Separación entre rutas públicas y privadas.
- Asociación de middleware.
- Organización por módulos.

El router no contiene lógica de negocio.

Su única responsabilidad es dirigir la petición hacia el caso de uso correspondiente.

---

# Middleware Layer

## Responsibility

Los middlewares implementan comportamiento transversal que debe ejecutarse antes de llegar al caso de uso.

Principales responsabilidades:

- Autenticación.
- Autorización.
- Construcción del Tenant Context.
- Manejo de errores.
- Recuperación de panics.
- Observabilidad.
- Trazabilidad.

Ejemplos:

Auth Middleware

Permission Middleware

Error Middleware

Recovery Middleware

Observability Middleware

---

# Controller Layer

## Responsibility

El controller representa la capa de entrada HTTP.

Su responsabilidad es traducir una petición HTTP en una ejecución del sistema.

Realiza:

- Lectura de parámetros.
- Parseo de JSON.
- Validación estructural.
- Invocación del service correspondiente.
- Construcción de respuesta HTTP.

No contiene:

- Reglas de negocio.
- Consultas directas a base de datos.
- Validaciones complejas del dominio.

---

# Service Layer

## Responsibility

La capa de servicios representa el núcleo de aplicación.

Es responsable de ejecutar casos de uso completos.

Contiene:

- Reglas de negocio.
- Validaciones funcionales.
- Estados permitidos.
- Orquestación de procesos.
- Coordinación entre dominios.
- Uso de repositorios.

Ejemplo:

Una operación de pago no representa solamente actualizar una tabla.

El service:

- Verifica estado de invoice.
- Valida monto.
- Registra payment.
- Actualiza estado.
- Coordina inventario.

El service es el dueño del caso de uso.

---

# Repository Layer

## Responsibility

Los repositories representan la capa de persistencia.

Responsabilidades:

- Consultas SQL.
- Inserciones.
- Actualizaciones.
- Eliminaciones.
- Mapeo entre datos y modelos.

Los repositories:

- No conocen reglas de negocio.
- No deciden estados.
- No validan workflows.

Solamente ejecutan operaciones solicitadas por los servicios.

---

# Dependency Injection

La aplicación utiliza una construcción manual de dependencias.

El flujo de inicialización es:

Database Connection

↓

Repositories

↓

Services

↓

Controllers

↓

Routes

Esto permite:

- Dependencias explícitas.
- Fácil testing.
- Menor acoplamiento.
- Mayor control arquitectónico.

---

# Module Structure

Cada dominio sigue una estructura similar:

module/

├── controller/

├── service/

├── repository/

├── model/

├── dto/

└── routes/

Esto permite que cada módulo tenga una responsabilidad clara y pueda evolucionar de forma independiente.

---

# Shared Components

Existen componentes compartidos utilizados transversalmente:

Ejemplos:

- Tenant Context.
- Error Handling.
- Response Handler.
- Transaction Utilities.
- Access Validators.
- Authentication Utilities.

Estos componentes no representan dominios de negocio, sino capacidades reutilizables.

---

# Business Workflows

Cuando una operación involucra múltiples dominios, no se fuerza dentro de un único módulo.

Se utilizan servicios encapsulados de workflow.

Ejemplos:

- Submit Invoice.
- Pay Invoice.
- Store Checkout.
- Marketplace Checkout.
- Dashboard Overview.

Estos servicios actúan como orquestadores de procesos complejos.

---

# Transaction Management

Las operaciones críticas utilizan transacciones SQL.

El objetivo es garantizar:

- Atomicidad.
- Consistencia.
- Rollback automático.
- Evitar estados parciales.

Ejemplo conceptual:

Begin Transaction

Update Invoice

Create Payment

Update Inventory

Commit

or

Rollback

---

# Error Management

El manejo de errores está centralizado.

La arquitectura diferencia:

## Business Errors

Errores esperados del dominio.

Ejemplos:

- Invoice no encontrada.
- Estado inválido.
- Monto incorrecto.

## Technical Errors

Errores internos.

Ejemplos:

- Fallo de base de datos.
- Error inesperado.
- Problema de infraestructura.

Los controllers no manejan errores manualmente.

La responsabilidad pertenece al sistema centralizado de errores.

---

# Scalability Considerations

La arquitectura fue diseñada considerando crecimiento futuro.

Principales decisiones:

- Modularización por dominio.
- Servicios desacoplados.
- Interfaces entre capas.
- Repositorios independientes.
- Configuración centralizada.
- Componentes shared reutilizables.

Agregar nuevas funcionalidades implica extender módulos existentes o crear nuevos dominios sin modificar la arquitectura base.

---

# Architectural Benefits

La arquitectura actual permite:

## Maintainability

Cada responsabilidad tiene un lugar definido.

## Scalability

Nuevos dominios pueden agregarse sin reorganizar todo el sistema.

## Testability

Las dependencias pueden reemplazarse gracias al desacoplamiento.

## Security

La autenticación y autorización están centralizadas.

## Evolution

El sistema puede evolucionar hacia arquitecturas más complejas en el futuro.

---

# Final Architecture Summary

El backend de Nexora fue construido siguiendo una arquitectura donde:

- HTTP no conoce negocio.
- Controllers no conocen persistencia.
- Services controlan casos de uso.
- Repositories conocen únicamente datos.
- Middleware controla aspectos transversales.
- Los dominios permanecen desacoplados.

Esta separación permite que el sistema mantenga claridad estructural incluso al aumentar la cantidad de módulos, usuarios y procesos de negocio.
