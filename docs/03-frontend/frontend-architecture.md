# Frontend Architecture

## Overview

El frontend de Nexora fue diseñado como una aplicación web moderna orientada a módulos de negocio, priorizando la separación de responsabilidades, escalabilidad y mantenibilidad a largo plazo.

El objetivo principal no fue únicamente construir una interfaz funcional, sino desarrollar una arquitectura capaz de evolucionar junto al backend manteniendo una estructura clara, predecible y desacoplada.

La arquitectura frontend fue concebida para reflejar los principios utilizados en backend, manteniendo una separación clara entre:

- Presentación.
- Estado global.
- Comunicación HTTP.
- Navegación.
- Configuración.
- Lógica reutilizable.
- Componentes visuales.

El frontend funciona como consumidor de casos de uso expuestos por la API, manteniendo al backend como única fuente de verdad respecto a reglas de negocio y autorización.

---

# Architectural Principles

Durante el diseño del frontend se priorizaron los siguientes principios:

## Separation of Responsibilities

Cada capa posee una responsabilidad concreta:

- Components → representación visual.
- Views → composición de pantallas completas.
- Composables → lógica reutilizable.
- Stores → estado global compartido.
- API Layer → comunicación con backend.
- Router → navegación y protección de rutas.
- Config → variables y configuración externa.

El objetivo fue evitar componentes monolíticos con responsabilidades mezcladas.

---

## Backend as Single Source of Truth

El frontend no replica reglas de negocio críticas.

Las validaciones principales permanecen en backend:

- permisos.
- estados de entidades.
- reglas comerciales.
- acceso multi-tenant.
- consistencia de datos.

El frontend solamente implementa validaciones preventivas para mejorar la experiencia del usuario.

Ejemplo:

Una invoice puede mostrar u ocultar acciones según permisos disponibles, pero la autorización real siempre es validada nuevamente por backend.

---

# Technology Stack

El frontend fue construido utilizando:

- Vue 3.
- Composition API.
- Vite.
- TypeScript.
- Pinia.
- Axios.
- Tailwind CSS.

La selección tecnológica permitió construir una aplicación moderna, ligera y preparada para crecer.

---

# Application Architecture

La arquitectura general puede representarse como:

```

User Interaction

```

    ↓

```

Vue Components

```

    ↓

```

Views

```

    ↓

```

Composables

```

    ↓

```

Stores / State

```

    ↓

```

API Layer

```

    ↓

```

Axios Client

```

    ↓

```

Backend API

```

Cada capa tiene una responsabilidad independiente.

Los componentes no conocen detalles de HTTP.

Los servicios API no conocen reglas visuales.

Los stores no contienen lógica de presentación.

---

# Project Organization

La estructura principal del proyecto fue organizada por responsabilidad:

```

src/

├── api/
├── assets/
├── components/
├── composables/
├── config/
├── layouts/
├── router/
├── stores/
├── types/
├── utils/
├── views/
└── App.vue

```

Esta organización evita una estructura basada únicamente en tipos de archivo y favorece una separación conceptual del sistema.

---

# Configuration Layer

La carpeta `config` centraliza información dependiente del entorno.

Ejemplo:

```

API URL

Storage URL

Environment Variables

```

Ningún componente accede directamente a valores de entorno.

Toda configuración pasa por una capa intermedia.

Esto permite ejecutar la aplicación en distintos escenarios:

- Desarrollo local.
- Docker.
- Render.
- Producción.

Sin modificar código fuente.

---

# Component Architecture

Los componentes fueron diseñados siguiendo una filosofía de composición.

Se evita construir componentes gigantes con múltiples responsabilidades.

La estructura busca separar:

```

Componentes reutilizables

-

Componentes específicos de dominio

```

Ejemplo:

Un componente genérico de tabla puede utilizarse en múltiples módulos.

Mientras que un componente específico de invoice mantiene únicamente lógica relacionada con facturación.

---

# Views Layer

Las vistas representan pantallas completas de la aplicación.

Su responsabilidad principal es:

- Componer componentes.
- Coordinar información necesaria.
- Utilizar composables.
- Definir la estructura visual de una página.

Las vistas no contienen lógica compleja de negocio.

---

# Composition API

La lógica reutilizable fue desarrollada mediante Composition API.

Se evitó utilizar patrones como mixins debido a sus problemas de escalabilidad:

- Dependencias ocultas.
- Difícil seguimiento.
- Colisiones de nombres.

Los composables permiten encapsular comportamiento reutilizable.

Ejemplos:

```

useAuth()

useInvoices()

useProducts()

useCompany()

```

Cada composable mantiene una responsabilidad concreta.

---

# Domain-Oriented Frontend

Aunque la estructura física está organizada por responsabilidades técnicas, conceptualmente el frontend sigue los dominios principales del negocio.

Ejemplos:

- Authentication.
- ERP.
- Store.
- Marketplace.
- Products.
- Inventory.
- Invoice.
- Dashboard.

Esto mantiene una correspondencia directa con la arquitectura backend.

---

# State Management

El estado global utiliza Pinia.

La decisión principal fue mantener el estado global limitado únicamente a información realmente compartida.

Ejemplos:

- Usuario autenticado.
- Tenant Context.
- Permisos.
- Sesión.
- Estados globales necesarios.

No se utilizan stores globales para información temporal de una única pantalla.

---

# HTTP Communication

Toda comunicación con backend pasa por una única capa API.

Flujo:

```

Component

↓

Composable

↓

API Module

↓

Axios Instance

↓

Backend

```

Esto evita que componentes visuales conozcan detalles HTTP.

Beneficios:

- Menor acoplamiento.
- Mayor facilidad de mantenimiento.
- Cambios centralizados.
- Código más testeable.

---

# Authentication Architecture

La autenticación sigue el mismo modelo conceptual del backend.

Flujo general:

```

Login

↓

JWT

↓

Persist Token

↓

Request /auth/me

↓

Tenant Context

↓

Pinia Store

↓

Router Access

```

El frontend no interpreta directamente la estructura interna del JWT.

Consume el contexto generado por backend.

---

# Multi-Tenant Frontend

El frontend fue preparado para trabajar con múltiples tipos de usuarios:

- Usuarios empresa.
- Clientes individuales.
- Usuarios con distintos permisos.

La interfaz se adapta según el contexto recibido:

- Empresa asociada.
- Permisos.
- Roles.
- Módulos disponibles.

Sin duplicar aplicaciones.

---

# Routing Architecture

La navegación fue organizada mediante Vue Router.

Existen diferentes tipos de rutas:

- Públicas.
- Privadas.
- Protegidas por permisos.

Los guards tienen responsabilidad únicamente de navegación.

No reemplazan la seguridad backend.

---

# Layout Architecture

La aplicación utiliza layouts reutilizables.

Ejemplos:

```

PublicLayout

ERPLayout

StoreLayout

MarketplaceLayout

```

Cada layout define:

- Navegación.
- Estructura general.
- Elementos compartidos.

Las vistas solamente representan contenido.

---

# Error Handling

El manejo de errores fue diseñado para mantener consistencia con backend.

El frontend diferencia:

## Business Errors

Ejemplo:

- Invoice no disponible.
- Permiso insuficiente.
- Operación inválida.

## Technical Errors

Ejemplo:

- Servidor caído.
- Error inesperado.
- Problemas de red.

Los errores técnicos muestran mensajes seguros al usuario mientras mantienen información detallada en logs del sistema.

---

# User Experience Considerations

Durante el desarrollo se contemplaron estados adicionales:

## Loading States

Cada operación asíncrona posee estados visibles:

- Carga inicial.
- Procesamiento.
- Bloqueo preventivo.

## Empty States

Se diferencia entre:

- No existen datos.
- No existen resultados por filtros.

## Feedback

Las acciones críticas muestran información clara al usuario.

Ejemplo:

- Confirmaciones.
- Errores.
- Estados exitosos.

---

# Scalability Approach

La arquitectura permite incorporar nuevos módulos siguiendo un patrón definido:

Nuevo módulo implica:

```

Nueva vista

-

Nuevos componentes

-

Nuevo composable

-

Nuevo API module

-

Nuevas rutas

-

Nuevo estado si corresponde

```

Sin modificar módulos existentes.

---

# Architectural Evaluation

Desde una perspectiva técnica, el frontend de Nexora presenta una arquitectura modular, desacoplada y preparada para evolución.

Las principales decisiones arquitectónicas fueron:

- Separación estricta de responsabilidades.
- Comunicación HTTP centralizada.
- Estado global controlado.
- Organización orientada a dominios.
- Integración directa con el modelo multi-tenant del backend.
- Componentización reutilizable.
- Configuración independiente del entorno.

El resultado es una aplicación frontend empresarial donde la interfaz funciona como una capa independiente de presentación, manteniendo coherencia arquitectónica con el backend y permitiendo crecimiento progresivo sin degradar mantenibilidad.

```

```
