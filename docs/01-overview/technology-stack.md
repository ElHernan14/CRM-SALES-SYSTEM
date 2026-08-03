# Technology Stack

## Overview

Nexora está construido utilizando un stack tecnológico orientado a aplicaciones empresariales modernas, priorizando escalabilidad, mantenibilidad, rendimiento y separación clara de responsabilidades.

La selección tecnológica no fue realizada únicamente por popularidad de herramientas, sino considerando la arquitectura planteada, la necesidad de crecimiento futuro y la capacidad de mantener un sistema modular con múltiples dominios de negocio.

El stack completo se divide principalmente en:

- Backend API
- Frontend SPA
- Database Layer
- Infrastructure
- Development Tooling

---

# Backend Stack

## Go

### Role

Go es el lenguaje principal utilizado para la construcción del backend de Nexora.

Fue seleccionado como base del sistema debido a sus características:

- Alto rendimiento.
- Modelo de concurrencia eficiente.
- Bajo consumo de recursos.
- Tipado fuerte.
- Simplicidad operacional.
- Excelente soporte para servicios backend.

Dentro del proyecto se utiliza para implementar:

- API REST.
- Lógica de negocio.
- Servicios de dominio.
- Middleware.
- Seguridad.
- Persistencia.
- Workflows transaccionales.

---

## Go net/http

### Role

El servidor HTTP fue construido utilizando el paquete estándar `net/http`.

La decisión permitió mantener un control explícito sobre:

- Ciclo de vida de requests.
- Middleware.
- Context handling.
- Manejo de errores.
- Inyección de dependencias.

Esto evita agregar abstracciones innecesarias y permite comprender completamente el flujo interno de cada petición.

---

## Gorilla Mux

### Role

Utilizado como router principal del backend.

Responsabilidades:

- Definición de endpoints.
- Organización de rutas públicas y privadas.
- Parametrización de URLs.
- Integración con middleware.

La estructura de rutas sigue una organización basada en módulos de negocio.

Ejemplo conceptual:

/api/auth

/api/invoices

/api/products

/api/inventory

---

# Database Stack

## PostgreSQL

### Role

PostgreSQL es la base de datos principal del sistema.

Fue seleccionado debido a:

- Robustez.
- Soporte transaccional.
- Integridad relacional.
- Capacidad para consultas complejas.
- Compatibilidad con sistemas empresariales.

Dentro de Nexora maneja:

- Usuarios.
- Empresas.
- Clientes.
- Productos.
- Facturación.
- Pagos.
- Inventario.
- Relaciones comerciales.

---

## SQL Nativo

### Role

La capa de persistencia utiliza consultas SQL directas en lugar de un ORM.

La decisión permite:

- Control total sobre consultas.
- Optimización específica.
- Transparencia del acceso a datos.
- Mejor comprensión del comportamiento real de la base.

Los repositories son responsables únicamente de:

- Obtener información.
- Persistir cambios.
- Ejecutar operaciones de almacenamiento.

Nunca contienen reglas de negocio.

---

# Frontend Stack

## Vue 3

### Role

Vue 3 es el framework principal utilizado para la aplicación frontend.

Fue elegido por:

- Flexibilidad.
- Curva de aprendizaje moderada.
- Ecosistema maduro.
- Excelente integración con aplicaciones empresariales.

La aplicación utiliza Vue como capa de presentación, manteniendo la lógica de negocio centralizada en backend.

---

## Composition API

### Role

La lógica reutilizable del frontend fue implementada utilizando Composition API.

Permite:

- Encapsular comportamiento.
- Crear composables reutilizables.
- Separar lógica de UI.
- Reducir componentes monolíticos.

Ejemplos:

useAuth()

useInvoices()

useProducts()

useCompany()

---

## Vite

### Role

Utilizado como herramienta de construcción y desarrollo.

Responsabilidades:

- Bundling.
- Desarrollo local.
- Hot Module Replacement.
- Optimización de producción.

Permite mantener tiempos rápidos de desarrollo y builds eficientes.

---

## Pinia

### Role

Pinia es utilizado como sistema centralizado de estado.

Se utiliza únicamente para estados globales necesarios:

- Usuario autenticado.
- Tenant Context.
- Permisos.
- Información compartida.

No se utiliza como reemplazo de estados locales de componentes.

---

## Tailwind CSS

### Role

Framework CSS utilizado para construcción de interfaces.

Permite:

- Diseño consistente.
- Desarrollo rápido.
- Componentización visual.
- Mantener estilos cercanos al componente.

---

# Authentication and Security Stack

## JWT

### Role

La autenticación utiliza JSON Web Tokens.

El token contiene información necesaria para identificar:

- Usuario.
- Empresa asociada.
- Cliente.
- Roles.
- Permisos.

El backend utiliza esta información para construir el Tenant Context.

---

## Bcrypt

### Role

Utilizado para almacenamiento seguro de contraseñas.

Responsabilidades:

- Hash de passwords.
- Comparación segura.
- Protección frente a exposición de credenciales.

---

# Infrastructure Stack

## Docker

### Role

Docker es utilizado para contenerizar servicios.

Permite:

- Entornos reproducibles.
- Configuración consistente.
- Separación de dependencias.
- Preparación para despliegues futuros.

---

## Render

### Role

Utilizado como plataforma de despliegue.

Actualmente soporta:

- Hosting frontend.
- Ejecución de servicios.
- Integración con repositorios.

Permite validar el comportamiento del sistema en un entorno cercano a producción.

---

## Supabase Storage

### Role

Utilizado para almacenamiento externo de recursos.

Principalmente:

- Imágenes.
- Assets.
- Archivos asociados.

Permite separar almacenamiento pesado del servidor principal.

---

# Development Tooling

## Git

### Role

Sistema de control de versiones utilizado para:

- Seguimiento de cambios.
- Organización mediante ramas.
- Historial del proyecto.

---

## GitHub

### Role

Plataforma utilizada para:

- Repositorios.
- Colaboración.
- Gestión del código fuente.

---

## Postman

### Role

Utilizado para pruebas manuales de API.

Permite validar:

- Endpoints.
- Respuestas.
- Autenticación.
- Estados HTTP.

---

## Swagger

### Role

Utilizado para documentación y exploración de API.

Permite:

- Definir contratos.
- Visualizar endpoints.
- Facilitar integración externa.

---

# Architectural Relationship

El stack completo fue seleccionado para trabajar como un ecosistema integrado:

Frontend
(Vue 3 + Pinia + Tailwind)
|
|
↓
REST API
(Go + Gorilla Mux)
|
|
↓
Business Layer
(Services + Workflows)
|
|
↓
Persistence Layer
(Repositories + PostgreSQL)
|
|
↓
Infrastructure
(Docker + Render + Storage)

Cada tecnología cumple una responsabilidad específica dentro del sistema.

No existen dependencias arbitrarias entre capas.

---

# Technology Selection Philosophy

La selección tecnológica de Nexora siguió los siguientes principios:

- Evitar sobreingeniería innecesaria.
- Mantener control sobre la arquitectura.
- Priorizar claridad sobre abstracción excesiva.
- Utilizar herramientas adecuadas para sistemas empresariales.
- Favorecer evolución futura.

El objetivo no fue solamente construir una aplicación funcional, sino una base tecnológica capaz de evolucionar hacia un producto empresarial completo.
