# Architecture Decisions

## Overview

Durante el desarrollo de Nexora se tomaron una serie de decisiones arquitectónicas orientadas a construir una plataforma escalable, mantenible y preparada para evolucionar.

Estas decisiones no estuvieron enfocadas únicamente en resolver los requerimientos actuales del sistema, sino en establecer una base técnica capaz de soportar crecimiento funcional, incorporación de nuevos dominios y evolución hacia una solución empresarial.

Este documento resume las principales decisiones de diseño tomadas durante la construcción del sistema y los motivos detrás de cada una.

---

# Decision 01 — Arquitectura Modular por Dominios

## Context

Nexora comenzó como una aplicación con múltiples áreas funcionales:

- ERP.
- Store.
- Marketplace.
- Gestión de usuarios.
- Inventario.
- Facturación.
- Pagos.

A medida que el sistema evolucionaba, mantener toda la lógica agrupada por funcionalidad técnica generaba riesgo de alto acoplamiento.

---

## Decision

Se decidió organizar el backend mediante módulos orientados a dominios de negocio.

Ejemplo:

```
invoice

product

inventory

client

company

payment

dashboard
```

Cada módulo mantiene sus propias responsabilidades:

- Controllers.
- Services.
- Repositories.
- Models.
- DTOs.

---

## Rationale

La separación por dominios permite:

- Reducir acoplamiento.
- Mantener responsabilidades claras.
- Facilitar incorporación de nuevos módulos.
- Evitar lógica distribuida entre diferentes partes del sistema.

La arquitectura se organiza alrededor del negocio y no únicamente alrededor de detalles técnicos.

---

# Decision 02 — Layered Architecture

## Context

El sistema necesitaba una separación clara entre entrada HTTP, lógica de aplicación y persistencia.

---

## Decision

Se implementó una arquitectura por capas:

```
Controller

↓

Service

↓

Repository

↓

Database
```

Con middleware transversal antes del Controller.

---

## Rationale

Cada capa posee una responsabilidad única.

### Controller

Responsable de:

- HTTP.
- Validación estructural.
- Request parsing.
- Response formatting.

No contiene reglas de negocio.

---

### Service

Responsable de:

- Casos de uso.
- Reglas de negocio.
- Validaciones.
- Workflows.

Es el dueño del comportamiento del sistema.

---

### Repository

Responsable de:

- Persistencia.
- Consultas.
- Operaciones SQL.

No conoce reglas de negocio.

---

## Result

La separación permite modificar una capa sin afectar las demás.

Ejemplo:

Cambiar PostgreSQL o modificar un endpoint no requiere alterar toda la aplicación.

---

# Decision 03 — Uso de Go como Backend Principal

## Context

El sistema requería:

- Buen rendimiento.
- Manejo eficiente de concurrencia.
- Simplicidad operacional.
- Bajo consumo de recursos.

---

## Decision

Se eligió Go como lenguaje principal del backend.

---

## Rationale

Go aporta:

- Concurrencia nativa mediante goroutines.
- Tipado fuerte.
- Interfaces simples.
- Bajo overhead.
- Excelente rendimiento para APIs.

Además, su filosofía favorece diseños con componentes pequeños y responsabilidades claras.

---

## Trade-off

Al no utilizar frameworks más pesados, varias piezas debieron construirse manualmente:

- Middleware.
- Dependency Injection.
- Manejo de errores.
- Routing.
- Arquitectura interna.

Esto aumenta el control, pero requiere mayor diseño inicial.

---

# Decision 04 — Uso de PostgreSQL sin ORM

## Context

El sistema posee operaciones donde la precisión de las consultas y el control transaccional son importantes.

Ejemplos:

- Facturación.
- Inventario.
- Dashboard.
- Reportes.

---

## Decision

Se decidió utilizar PostgreSQL directamente mediante consultas SQL.

---

## Rationale

Ventajas:

- Control total sobre queries.
- Mejor comprensión del modelo de datos.
- Optimización específica.
- Manejo explícito de transacciones.

---

## Trade-off

La ausencia de ORM implica:

- Mayor cantidad de código manual.
- Mayor responsabilidad del desarrollador.
- Mayor dificultad inicial.

Sin embargo, para un sistema ERP con operaciones críticas, el control explícito resulta beneficioso.

---

# Decision 05 — JWT Authentication con Tenant Context

## Context

Nexora necesitaba soportar múltiples tipos de usuarios:

- Usuarios empresariales.
- Clientes individuales.
- Usuarios con diferentes permisos.

---

## Decision

Se implementó autenticación basada en JWT.

El token contiene información necesaria para construir el contexto de ejecución:

```
user_id

email

company_id

client_id

roles

permissions
```

---

## Rationale

Esto permite:

- Identificar rápidamente al usuario.
- Evitar consultas adicionales en cada request.
- Construir Tenant Context desde el inicio del flujo.

---

## Trade-off

Mantener permisos dentro del token requiere controlar correctamente:

- Tiempo de expiración.
- Renovación.
- Invalidación.

Por esta razón los tokens poseen una duración limitada.

---

# Decision 06 — Separación Authentication / Authorization

## Context

Autenticar un usuario y permitir una acción son responsabilidades diferentes.

---

## Decision

Se separaron ambos procesos:

```
Authentication Middleware

↓

Authorization Middleware
```

---

## Rationale

Authentication responde:

> ¿Quién es el usuario?

Authorization responde:

> ¿Puede realizar esta operación?

Separarlos permite:

- Mayor claridad.
- Mayor reutilización.
- Mejor control por endpoint.

---

# Decision 07 — RBAC basado en Permisos

## Context

El sistema necesitaba controlar accesos de forma flexible.

---

## Decision

Se implementó un modelo basado en permisos en lugar de únicamente roles.

---

## Rationale

Los roles representan agrupaciones.

Los permisos representan capacidades reales.

Ejemplo:

```
invoice.create

invoice.update

invoice.pay

inventory.manage
```

Esto permite mayor granularidad y evita sistemas rígidos.

---

# Decision 08 — Multi-Tenant con Base Compartida

## Context

Nexora debía permitir usuarios pertenecientes a diferentes organizaciones.

---

## Decision

Se utilizó un modelo:

```
Shared Database

+

company_id isolation
```

---

## Rationale

Ventajas:

- Menor complejidad operacional.
- Modelo más simple.
- Fácil evolución inicial.
- Compartición controlada de estructura.

---

## Isolation Strategy

El aislamiento se garantiza mediante:

- Tenant Context.
- Validaciones de acceso.
- Filtros por company_id.
- Servicios compartidos de autorización.

---

# Decision 09 — Servicios Workflow para Procesos Complejos

## Context

Algunas operaciones no representan simples CRUD.

Ejemplos:

- Checkout.
- Submit Invoice.
- Pay Invoice.
- Inventory Reservation.

---

## Decision

Crear servicios especializados para procesos completos.

---

## Rationale

Estos servicios encapsulan:

- Reglas.
- Validaciones.
- Cambios múltiples de entidades.
- Transacciones.

Ejemplo:

```
Pay Invoice

↓

Update Payment

↓

Validate Amount

↓

Update Invoice State

↓

Finalize Stock
```

---

# Decision 10 — Manejo Centralizado de Errores

## Context

El sistema necesitaba respuestas consistentes.

---

## Decision

Se creó una estrategia centralizada mediante middleware.

---

## Rationale

Permite:

- Unificar respuestas HTTP.
- Separar errores técnicos de negocio.
- Evitar duplicación.

Flujo:

```
Service Error

↓

Error Handler

↓

Middleware

↓

HTTP Response
```

---

# Decision 11 — Observabilidad desde el Inicio

## Context

Durante desarrollo era necesario poder rastrear problemas complejos.

---

## Decision

Implementar observabilidad desde la arquitectura inicial.

---

## Components

- Request ID.
- Logging centralizado.
- Error tracking.
- Recovery middleware.

---

## Rationale

Permite reconstruir el recorrido de una operación cuando múltiples usuarios interactúan simultáneamente.

---

# Decision 12 — Frontend Desacoplado del Backend

## Context

El frontend debía evolucionar independientemente.

---

## Decision

Separar completamente:

- Presentación.
- Estado.
- Comunicación HTTP.

---

## Architecture

```
Vue Component

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

## Rationale

Esto permite:

- Cambiar endpoints sin modificar UI.
- Reutilizar lógica.
- Mantener componentes limpios.

---

# Decision 13 — State Management Controlado

## Context

No todo dato debe vivir globalmente.

---

## Decision

Utilizar Pinia únicamente para estado compartido real.

Ejemplo:

- Usuario.
- Tenant Context.
- Permisos.

---

## Rationale

Evita crear un estado global gigante y difícil de mantener.

---

# Decision 14 — Configuración Externalizada

## Context

El sistema debe ejecutarse en diferentes ambientes.

---

## Decision

Separar configuración mediante variables de entorno.

Ambientes:

- Desarrollo.
- Docker.
- Producción.

---

## Rationale

Permite cambiar infraestructura sin modificar código.

---

# Decision 15 — Preparación para Evolución Futura

## Context

Aunque Nexora actualmente funciona como una aplicación modular monolítica, debía existir una base preparada para crecer.

---

## Decision

Mantener módulos desacoplados y contratos claros.

---

## Future Evolution

Posibles evoluciones:

- Arquitectura hexagonal.
- Event-driven architecture.
- Servicios independientes.
- Observabilidad avanzada.
- Cache distribuido.
- Mensajería asíncrona.

---

# Final Summary

Las decisiones arquitectónicas tomadas en Nexora estuvieron orientadas a construir un sistema donde la complejidad estuviera controlada mediante separación de responsabilidades.

Los principales principios aplicados fueron:

- Modularidad.
- Bajo acoplamiento.
- Alta cohesión.
- Separación de responsabilidades.
- Seguridad transversal.
- Procesos de negocio explícitos.
- Escalabilidad progresiva.

Estas decisiones permiten que Nexora pueda evolucionar desde una aplicación empresarial modular hacia una plataforma de mayor escala manteniendo claridad técnica y estabilidad.
