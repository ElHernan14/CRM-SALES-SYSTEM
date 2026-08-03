# Local Development Guide — Nexora

## Introducción

Este documento describe el flujo recomendado para ejecutar Nexora en entorno local durante el desarrollo.

El objetivo es permitir que nuevos colaboradores puedan comprender rápidamente:

- requisitos del entorno.
- estructura de ejecución.
- configuración necesaria.
- flujo recomendado de desarrollo.

Nexora está compuesto por dos aplicaciones independientes:

```
Nexora Backend

+

Nexora Frontend

```

Cada una posee su propio ciclo de desarrollo.

---

# Requisitos Previos

Antes de ejecutar el proyecto se requiere:

## Backend

- Go instalado.
- PostgreSQL disponible.
- Variables de entorno configuradas.

Versión utilizada durante desarrollo:

```
Go 1.26+

```

---

## Frontend

- Node.js instalado.
- npm instalado.

Versión utilizada durante desarrollo:

```
Node.js 22+

npm 10+

```

---

# Estructura General

El proyecto está organizado como:

```
Nexora

├── backend

└── frontend

```

Ambas aplicaciones se comunican mediante API REST.

---

# Configuración Backend

El backend requiere un archivo:

```
.env

```

con las variables necesarias para:

- servidor.
- base de datos.
- almacenamiento.
- servicios externos.

Ejemplo:

```
SERVER_PORT=8081

DB_HOST=

DB_PORT=

DB_USER=

DB_PASSWORD=

DB_NAME=

DB_SSLMODE=

UPLOAD_DIR=

```

---

# Ejecutar Backend

Desde la carpeta backend:

Instalar dependencias:

```
go mod download

```

Ejecutar aplicación:

```
go run main.go

```

El servidor quedará disponible en:

```
http://localhost:8081

```

---

# Ejecutar Backend con Docker

El backend posee soporte Docker.

Construcción:

```
docker compose build

```

Ejecución:

```
docker compose up

```

Esto permite levantar el servicio utilizando la misma configuración definida para despliegue.

---

# Configuración Frontend

El frontend utiliza variables de entorno mediante Vite.

Archivo:

```
.env

```

Ejemplo:

```
VITE_API_URL=http://localhost:8081/api

VITE_ASSETS_URL=http://localhost:8081

```

Estas variables permiten modificar endpoints sin cambiar código fuente.

---

# Ejecutar Frontend

Desde la carpeta frontend:

Instalar dependencias:

```
npm install

```

Ejecutar servidor de desarrollo:

```
npm run dev

```

La aplicación estará disponible en:

```
http://localhost:5173

```

---

# Flujo de Desarrollo Recomendado

El flujo recomendado es:

```
Frontend

↓

API Request

↓

Backend Middleware

↓

Controller

↓

Service

↓

Repository

↓

Database

```

Al desarrollar nuevas funcionalidades se debe respetar la separación de responsabilidades.

---

# Agregar Nuevos Cambios

Antes de implementar una nueva funcionalidad se recomienda definir:

1. ¿Pertenece a un dominio existente?

2. ¿Es un nuevo módulo?

3. ¿Es lógica de negocio o infraestructura?

4. ¿Debe existir un nuevo endpoint?

5. ¿Debe incorporarse un nuevo permiso?

---

# Desarrollo Backend

Al agregar lógica:

Correcto:

```
Controller

↓

Service

↓

Repository

```

Incorrecto:

```
Controller

↓

SQL directo

```

Los controllers no deben contener reglas de negocio.

---

# Desarrollo Frontend

Al agregar funcionalidades:

La estructura recomendada es:

```
View

↓

Composable

↓

API Layer

↓

Backend

```

Evitar:

- llamadas HTTP dentro de componentes.
- lógica compleja dentro de templates.
- duplicación de estado global.

---

# Manejo de Errores

Todos los errores deben respetar la estrategia definida.

Backend:

```
Service

↓

Error Wrapper

↓

Error Middleware

↓

Response

```

Frontend:

```
API Error

↓

Composable

↓

UI Feedback

```

---

# Variables Sensibles

Nunca incluir:

- contraseñas.
- tokens.
- claves privadas.
- archivos .env reales.

en:

- Git.
- documentación.
- commits públicos.

---

# Base de Datos

Los cambios de estructura deben realizarse mediante scripts controlados.

Evitar modificaciones manuales sin registrar.

Toda modificación importante debe quedar documentada.

---

# Debugging

Cuando ocurre un error:

Orden recomendado:

1. Revisar consola frontend.

2. Revisar logs backend.

3. Identificar Request ID.

4. Revisar flujo del endpoint.

5. Validar contexto del usuario.

6. Revisar base de datos.

---

# Buenas Prácticas

Durante el desarrollo mantener:

- commits descriptivos.
- módulos separados.
- código limpio.
- responsabilidades claras.
- documentación actualizada.

---

# Flujo Git Recomendado

Ejemplo:

```
main

↓

feature/nueva-funcionalidad

↓

Pull Request

↓

Merge

```

Los cambios importantes deben revisarse antes de incorporarse a producción.

---

# Objetivo Arquitectónico

El entorno local debe mantener la misma filosofía del sistema completo:

- separación de responsabilidades.
- configuración externa.
- módulos independientes.
- bajo acoplamiento.
- reproducibilidad.

El objetivo no es solamente ejecutar Nexora, sino mantener una forma consistente de evolucionar el sistema.
