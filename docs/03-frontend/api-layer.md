# API Layer — Frontend Nexora

## Introducción

La capa API del frontend de Nexora fue diseñada como una capa de abstracción entre la interfaz de usuario y la comunicación con el backend.

El objetivo principal fue evitar que los componentes Vue conocieran detalles de implementación relacionados con HTTP, endpoints, headers o manejo directo de respuestas.

La comunicación sigue un flujo definido:

```
View Component

↓

Composable

↓

API Module

↓

HTTP Client (Axios)

↓

Backend API
```

Esta separación permite mantener una arquitectura desacoplada donde la interfaz consume casos de uso y no implementaciones concretas de infraestructura.

---

# Objetivos de Diseño

La capa API fue construida teniendo en cuenta los siguientes objetivos:

- Centralizar todas las comunicaciones HTTP.
- Evitar llamadas directas desde componentes visuales.
- Separar lógica de presentación e infraestructura.
- Facilitar cambios futuros en endpoints.
- Mantener consistencia en autenticación y manejo de errores.
- Permitir reutilización de servicios entre diferentes módulos.

La idea principal fue que ningún componente Vue necesitara conocer cómo se comunica la aplicación con el backend.

---

# Cliente HTTP Centralizado

La aplicación utiliza una única instancia de Axios configurada globalmente.

Esta instancia concentra configuraciones comunes para todas las requests:

- URL base del backend.
- Credenciales necesarias.
- Headers globales.
- Interceptores.
- Configuración común de comunicación.

Conceptualmente:

```
Componentes

↓

API Modules

↓

Axios Instance

↓

Backend
```

Esto evita crear múltiples clientes HTTP independientes dentro del proyecto.

---

# Configuración del Cliente

La URL del backend no se encuentra escrita directamente dentro del código.

La aplicación utiliza una capa de configuración basada en variables de entorno.

Ejemplo conceptual:

```text
config

↓

env.apiUrl

↓

Axios baseURL
```

Esto permite utilizar la misma aplicación en distintos ambientes:

- Desarrollo local.
- Testing.
- Docker.
- Producción.

Sin necesidad de modificar código fuente.

---

# Organización por Módulos

Cada dominio posee su propio módulo API.

Ejemplo conceptual:

```
api/

├── auth.api.ts
├── invoice.api.ts
├── product.api.ts
├── client.api.ts
├── dashboard.api.ts
└── inventory.api.ts
```

Cada archivo contiene únicamente las operaciones relacionadas con su dominio.

Esto mantiene una organización consistente con la arquitectura del backend.

---

# Responsabilidad de los API Modules

Un módulo API tiene como responsabilidad:

- Definir endpoints utilizados.
- Enviar parámetros necesarios.
- Ejecutar requests HTTP.
- Devolver resultados al composable correspondiente.

No debe contener:

- Lógica de negocio.
- Validaciones complejas.
- Decisiones de UI.
- Manejo de estado global.

La capa API únicamente representa el contrato de comunicación con el backend.

---

# Separación entre API y Composables

La lógica de interacción con la aplicación se mantiene en composables.

El flujo completo es:

```
Usuario interactúa

↓

Component

↓

Composable

↓

API Module

↓

Axios

↓

Backend
```

Ejemplo:

Un componente no realiza:

```javascript
axios.get("/products");
```

Sino que utiliza una abstracción:

```javascript
useProducts()

↓

productApi.getProducts()
```

Esto permite modificar la comunicación sin afectar la interfaz.

---

# Manejo de Autenticación

La capa API trabaja junto al sistema de autenticación definido en el frontend.

El flujo es:

```
Login

↓

Backend devuelve JWT

↓

Token almacenado

↓

Axios utiliza contexto autenticado

↓

Requests protegidas
```

El frontend no genera permisos ni reglas de autorización.

Solamente transporta el contexto entregado por el backend.

---

# Manejo de Errores

La capa API participa en la normalización de errores provenientes del backend.

Los errores pueden clasificarse en:

## Errores de negocio

Ejemplos:

- Factura no válida.
- Estado incorrecto.
- Falta de permisos.
- Stock insuficiente.

Estos errores deben conservar el mensaje enviado por backend para permitir una correcta comunicación con el usuario.

---

## Errores técnicos

Ejemplos:

- Servidor caído.
- Error interno.
- Problemas de red.

Estos errores deben manejarse sin exponer información sensible al usuario.

---

# Responsabilidad del Frontend frente a Errores

El frontend interpreta los errores recibidos, pero no reemplaza la validación del backend.

La filosofía aplicada fue:

```
Frontend:

Previene errores conocidos

Backend:

Valida reglas reales del negocio
```

Ejemplo:

El frontend puede impedir visualmente enviar una factura sin productos.

Pero la validación definitiva pertenece al backend.

---

# Ventajas de la Arquitectura API

Esta separación permite:

## Bajo acoplamiento

Los componentes no dependen directamente de HTTP.

## Fácil mantenimiento

Si cambia un endpoint, solamente se modifica su módulo API correspondiente.

## Reutilización

Un mismo servicio puede ser utilizado por múltiples vistas.

## Escalabilidad

Agregar un nuevo módulo implica crear su propia capa API sin modificar módulos existentes.

---

# Integración con Arquitectura General

La capa API mantiene la misma filosofía utilizada en todo Nexora:

- Responsabilidades separadas.
- Módulos independientes.
- Contratos claros.
- Bajo acoplamiento.
- Organización por dominio.

El frontend actúa como consumidor de la API, mientras que el backend continúa siendo la única fuente de verdad del negocio.

---

# Evaluación Arquitectónica

La implementación de una capa API dedicada permite que Nexora mantenga una arquitectura frontend escalable y preparada para evolución futura.

La separación entre componentes, composables y comunicación HTTP evita mezclar responsabilidades y permite que nuevas funcionalidades puedan incorporarse manteniendo la estructura existente.

Esta decisión fue fundamental para mantener consistencia entre frontend y backend, logrando que ambos sistemas evolucionen de manera independiente pero bajo los mismos principios arquitectónicos.
