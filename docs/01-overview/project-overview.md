# Nexora — Project Overview

## 1. Introducción

Nexora es una plataforma empresarial orientada a la gestión integral de operaciones comerciales, diseñada bajo un modelo SaaS multi-tenant que permite administrar diferentes contextos de negocio dentro de una misma infraestructura.

El sistema combina capacidades de ERP, Store y Marketplace, permitiendo manejar procesos relacionados con empresas, clientes, productos, inventario, ventas, compras, facturación y pagos desde una arquitectura centralizada y escalable.

El objetivo principal del proyecto fue construir una plataforma moderna donde los diferentes dominios de negocio pudieran evolucionar independientemente, manteniendo una separación clara de responsabilidades, seguridad robusta y una estructura preparada para crecimiento futuro.

---

# 2. Contexto del Sistema

Nexora nace como una plataforma que busca unificar diferentes flujos comerciales dentro de un único ecosistema.

El sistema contempla dos grandes escenarios de uso:

## Modelo empresarial (ERP)

Orientado a organizaciones que necesitan administrar:

- Productos.
- Clientes.
- Ventas.
- Compras.
- Inventario.
- Facturación.
- Pagos.
- Métricas operativas.

## Modelo usuario final (Store / Marketplace)

Orientado a clientes individuales que pueden:

- Explorar productos.
- Gestionar carrito.
- Realizar compras.
- Generar órdenes.
- Interactuar con el ecosistema comercial.

Ambos modelos conviven dentro de la misma plataforma utilizando un sistema de identidad y contexto multi-tenant.

---

# 3. Objetivos Técnicos

Desde el inicio del desarrollo se establecieron objetivos arquitectónicos orientados a construir un sistema mantenible y extensible.

Los principales objetivos fueron:

- Diseñar una arquitectura desacoplada.
- Separar responsabilidades entre capas.
- Mantener una clara división entre infraestructura y lógica de negocio.
- Implementar seguridad basada en autenticación y autorización.
- Soportar múltiples empresas dentro del mismo sistema.
- Permitir crecimiento modular sin afectar funcionalidades existentes.
- Construir una base preparada para evolución hacia una plataforma productiva.

---

# 4. Arquitectura General

Nexora está compuesto por dos aplicaciones principales:

```
Frontend
(Vue 3)

        |
        |
        v

Backend API
(Go)

        |
        |
        v

Database
(PostgreSQL)
```

Cada capa posee responsabilidades independientes.

La comunicación entre frontend y backend se realiza mediante una API REST.

El backend actúa como fuente única de verdad respecto al dominio, reglas de negocio y seguridad, mientras que el frontend administra la presentación, navegación y experiencia del usuario.

---

# 5. Principios Arquitectónicos

El diseño general del sistema se basa en los siguientes principios:

## Separación de responsabilidades

Cada componente tiene una función definida dentro del sistema.

Ejemplo:

- Middleware → seguridad transversal.
- Controller → comunicación HTTP.
- Service → lógica de negocio.
- Repository → persistencia.
- Componentes Vue → presentación.

---

## Bajo acoplamiento

Los módulos fueron diseñados para minimizar dependencias directas.

Cada dominio mantiene sus propias responsabilidades y se comunica mediante contratos definidos.

Esto permite evolucionar partes del sistema sin generar efectos colaterales.

---

## Modularización por dominio

La organización del sistema sigue los conceptos del dominio del negocio.

Los principales módulos incluyen:

- Auth.
- Users.
- Companies.
- Clients.
- Products.
- Inventory.
- Invoice.
- Payments.
- Purchases.
- Dashboard.
- Store.
- Marketplace.

Cada módulo representa una responsabilidad funcional concreta.

---

## Seguridad como parte de la arquitectura

La seguridad no fue agregada como una funcionalidad secundaria.

Forma parte del flujo principal del sistema mediante:

- JWT.
- Middleware de autenticación.
- Control de permisos.
- Tenant Context.
- Aislamiento multi-tenant.

---

# 6. Stack Tecnológico

## Backend

| Tecnología  | Uso                            |
| ----------- | ------------------------------ |
| Go          | Lenguaje principal del backend |
| net/http    | Servidor HTTP                  |
| Gorilla Mux | Routing                        |
| PostgreSQL  | Persistencia                   |
| JWT         | Autenticación                  |
| Docker      | Entorno de ejecución           |

---

## Frontend

| Tecnología   | Uso                 |
| ------------ | ------------------- |
| Vue 3        | Framework principal |
| Vite         | Build tooling       |
| TypeScript   | Tipado estático     |
| Pinia        | Gestión de estado   |
| Vue Router   | Navegación          |
| Tailwind CSS | Sistema de estilos  |

---

## Infraestructura

| Tecnología       | Uso                        |
| ---------------- | -------------------------- |
| Docker           | Contenedores               |
| Render           | Deployment                 |
| Supabase Storage | Almacenamiento de recursos |
| Git/GitHub       | Control de versiones       |

---

# 7. Módulos Principales

## Authentication & Authorization

Responsable de:

- Login.
- Generación de JWT.
- Gestión de permisos.
- Construcción del contexto del usuario.

---

## ERP

Módulo empresarial encargado de:

- Gestión de compañías.
- Clientes.
- Productos.
- Inventario.
- Ventas.
- Compras.
- Facturación.

---

## Invoice

Dominio central del sistema.

Gestiona:

- Creación de facturas.
- Estados.
- Items.
- Pagos.
- Flujo de emisión.

La entidad Invoice funciona como una pieza central alrededor de la cual interactúan múltiples procesos del negocio.

---

## Inventory

Responsable de:

- Control de stock.
- Reservas.
- Finalización de movimientos.
- Consistencia entre ventas y disponibilidad.

---

## Store / Marketplace

Representan los flujos comerciales orientados al usuario final.

Incluyen:

- Catálogo.
- Carrito.
- Checkout.
- Generación de órdenes.

---

## Dashboard

Implementa una capa de composición orientada a obtener una visión global del negocio.

No representa una entidad propia, sino un servicio encargado de consolidar información transversal.

---

# 8. Características Principales

Nexora implementa:

## Arquitectura Multi-Tenant

Permite múltiples organizaciones dentro de una misma aplicación manteniendo aislamiento lógico mediante:

- Company Context.
- Tenant Context.
- Validaciones de acceso.
- Filtros por contexto.

---

## Sistema RBAC

El acceso está basado en permisos dinámicos.

El sistema diferencia:

- Identidad del usuario.
- Roles.
- Permisos específicos.

Esto permite una autorización flexible y escalable.

---

## Workflows de Negocio

Los procesos importantes no fueron implementados como simples operaciones CRUD.

Ejemplos:

- Submit Invoice.
- Pay Invoice.
- Checkout.
- Reserva de inventario.

Estos workflows encapsulan reglas reales del negocio.

---

## Manejo Centralizado de Errores

El sistema posee mecanismos globales para:

- Errores funcionales.
- Errores técnicos.
- Respuestas consistentes.
- Recuperación ante fallos.

---

## Observabilidad

Se incorporaron mecanismos para facilitar diagnóstico:

- Request ID.
- Logs estructurados.
- Middleware de observabilidad.
- Seguimiento de errores.

---

# 9. Filosofía del Proyecto

La construcción de Nexora no estuvo orientada únicamente a entregar funcionalidades.

El objetivo principal fue desarrollar una base tecnológica que pudiera evolucionar.

Cada decisión arquitectónica buscó equilibrar:

- simplicidad actual;
- capacidad de crecimiento;
- facilidad de mantenimiento;
- claridad del dominio.

La arquitectura fue diseñada pensando en que nuevos módulos, procesos y reglas de negocio puedan incorporarse sin necesidad de reconstruir el sistema existente.

---

# 10. Evolución Futura

La arquitectura actual permite una evolución progresiva hacia escenarios más avanzados.

Posibles líneas futuras:

- Mayor separación por dominios.
- Arquitectura hexagonal.
- Servicios independientes.
- Observabilidad avanzada.
- Automatización CI/CD.
- Escalamiento horizontal.
- Nuevos módulos empresariales.

La base actual fue construida buscando que estas evoluciones sean incrementales y no requieran cambios estructurales completos.

---

# Conclusión

Nexora representa una plataforma Full Stack diseñada bajo principios modernos de ingeniería de software.

Su arquitectura combina un backend orientado a dominio, un frontend modular y una infraestructura preparada para crecimiento.

Más allá de sus funcionalidades actuales, el principal valor técnico del proyecto reside en las decisiones de diseño que permiten mantenerlo, extenderlo y evolucionarlo de manera sostenible.
