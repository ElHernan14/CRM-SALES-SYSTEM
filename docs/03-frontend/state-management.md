# Frontend State Management

## Overview

El manejo de estado del frontend de Nexora fue diseñado utilizando Pinia como solución centralizada para administrar información compartida entre múltiples partes de la aplicación.

La decisión principal no fue almacenar todo el estado globalmente, sino definir claramente qué información realmente necesita persistencia transversal y cuál debe permanecer localizada dentro de cada módulo.

El objetivo fue evitar un frontend con estado excesivamente acoplado, donde cualquier componente pueda modificar cualquier información de la aplicación.

La arquitectura busca mantener:

- Estado global controlado.
- Responsabilidades claras.
- Bajo acoplamiento entre módulos.
- Flujo de datos predecible.

---

# State Management Principles

Durante el diseño del manejo de estado se siguieron los siguientes principios:

## Global State Only When Necessary

No toda información debe pertenecer al estado global.

Únicamente se centralizan datos que:

- Son utilizados por múltiples módulos.
- Representan contexto de aplicación.
- Necesitan mantenerse durante la sesión.
- Tienen impacto transversal.

Ejemplos:

- Usuario autenticado.
- Tenant Context.
- Permisos.
- Sesión actual.

---

## Local State For Local Concerns

La información propia de una pantalla o componente permanece dentro del mismo contexto.

Ejemplos:

- Filtros temporales.
- Estados de formularios.
- Modales abiertos.
- Datos utilizados únicamente por una vista.

Esto evita transformar el store global en un punto de dependencia innecesario.

---

# Pinia Architecture

Pinia fue seleccionado como mecanismo principal de estado debido a:

- Integración natural con Vue 3.
- Compatibilidad con Composition API.
- Tipado sencillo.
- Menor complejidad frente a soluciones más pesadas.

La estructura general:

```

Component

↓

Composable

↓

Pinia Store

↓

Application State

```

Los componentes no acceden directamente a lógica compleja del estado.

---

# Store Responsibilities

Cada store debe tener una responsabilidad concreta.

La filosofía aplicada fue:

```

One Store

=

One Context

```

No se construyeron stores gigantes que administren toda la aplicación.

Ejemplo:

Correcto:

```

authStore

productStore

invoiceStore

```

Incorrecto:

```

applicationStore
|
├── auth
├── products
├── invoices
├── dashboard
├── users

```

La segunda aproximación aumenta el acoplamiento y dificulta la evolución.

---

# Authentication Store

El store más importante del sistema es el relacionado con autenticación.

Su responsabilidad es administrar la sesión actual.

Mantiene información como:

```

Token

User

Tenant Context

Permissions

Roles

Authentication State

```

No contiene:

- Validaciones de negocio.
- Reglas de permisos.
- Procesos de backend.

Su función es representar el estado actual de la sesión.

---

# Tenant Context State

Debido al modelo multi-tenant del sistema, el frontend mantiene información del contexto actual.

Conceptualmente:

```

TenantContext

{

userId

email

companyId

clientId

roles

permissions

}

```

Este contexto permite que la aplicación conozca qué experiencia debe mostrar al usuario.

Ejemplos:

Usuario empresa:

```

ERP Dashboard

Invoices

Inventory

Purchases

```

Usuario cliente:

```

Store

Orders

Profile

```

---

# Data Flow

El flujo general del estado es:

```

Backend Response

↓

API Layer

↓

Composable

↓

Pinia Store

↓

Components

```

Los componentes consumen información ya procesada.

No conocen detalles de autenticación, HTTP o persistencia.

---

# Composables and Stores Relationship

Una decisión importante fue diferenciar claramente:

## Stores

Responsables de:

- Estado compartido.
- Persistencia durante sesión.
- Información transversal.

## Composables

Responsables de:

- Lógica reutilizable.
- Orquestación de acciones.
- Comunicación entre capas.

Ejemplo:

```

Component

↓

useAuth()

↓

authStore

↓

auth.api.ts

```

El composable actúa como capa intermedia.

---

# Avoiding Business Logic In Stores

Los stores no contienen reglas complejas de negocio.

Ejemplo incorrecto:

```

authStore.payInvoice()

```

Porque el pago pertenece al dominio invoice.

Ejemplo correcto:

```

invoiceService.pay()

↓

Backend

↓

Update Store State

```

El frontend mantiene estado, no dominio.

---

# State Synchronization

Cuando una operación modifica información del sistema, el flujo esperado es:

```

User Action

↓

Composable

↓

API Request

↓

Backend Processing

↓

Response

↓

Update Store / Local State

```

El frontend no asume que una operación fue exitosa hasta recibir confirmación del backend.

---

# Loading and Async State

El manejo de operaciones asíncronas fue contemplado como parte del estado.

Ejemplos:

- Cargando sesión inicial.
- Consultando información.
- Ejecutando acciones críticas.

Estados comunes:

```

idle

loading

success

error

```

Esto permite interfaces más predecibles.

---

# Error State Handling

Los errores no se almacenan globalmente sin necesidad.

Cada módulo administra sus errores funcionales.

Ejemplo:

Invoice:

```

invoiceError

```

Products:

```

productError

```

Mientras que errores globales de sesión sí pertenecen al contexto autenticado.

---

# Persistence Strategy

La persistencia del estado fue limitada.

Principalmente se mantiene:

- Token de autenticación.
- Información necesaria para reconstruir sesión.

No se persisten datos de negocio completos.

Ejemplo:

No se guarda:

```

Lista completa de productos

Todas las invoices

Dashboard completo

```

Estos datos siempre provienen del backend.

---

# Scalability Considerations

La estrategia de estado permite incorporar nuevos módulos sin modificar la arquitectura existente.

Agregar un nuevo dominio implica:

```

Nuevo módulo

↓

Nuevo Store si requiere estado global

↓

Nuevo Composable

↓

Nuevas APIs

↓

Nuevas Views

```

No se necesita alterar stores existentes.

---

# Architectural Evaluation

El sistema de gestión de estado de Nexora fue diseñado priorizando simplicidad y escalabilidad.

Las decisiones principales fueron:

- Uso de Pinia como estado global controlado.
- Separación entre estado y lógica reutilizable.
- Evitar stores monolíticos.
- Mantener backend como fuente de verdad.
- Utilizar Tenant Context como información transversal.
- Reducir dependencias entre módulos.

Este enfoque permite que el frontend pueda crecer incorporando nuevos dominios sin transformar el estado global en un punto centralizado de complejidad.

```

```
