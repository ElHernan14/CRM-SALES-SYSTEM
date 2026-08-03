# Dockerización — Nexora Backend

## Introducción

Nexora utiliza Docker como mecanismo de empaquetado y despliegue del backend API.

La dockerización fue aplicada exclusivamente sobre el servicio backend desarrollado en Go, mientras que el frontend Vue se despliega mediante un servicio independiente de tipo Web Service/Static Site utilizando la infraestructura del proveedor de hosting.

La separación permite mantener responsabilidades independientes:

```
Frontend Vue

↓

API Backend Go

↓

Base de datos PostgreSQL

↓

Servicios externos (Storage)
```

Docker tiene como objetivo garantizar que el backend pueda ejecutarse de manera consistente en distintos entornos, evitando diferencias entre desarrollo, testing y producción.

---

# Arquitectura de despliegue

La distribución general del sistema es:

```
                 Usuario

                   │

                   ▼

            Frontend Vue
              (Render)

                   │
                   │ HTTP API
                   ▼

          Nexora Backend API
              (Docker)

                   │

        ┌──────────┴──────────┐

        ▼                     ▼

 PostgreSQL              Storage

(Supabase)           (Local/Supabase)
```

El frontend consume únicamente la API pública del backend mediante variables de entorno.

El backend contiene toda la lógica de negocio, seguridad, validaciones y acceso a datos.

---

# Dockerfile Backend

El backend utiliza una construcción multi-stage para separar la compilación del entorno final de ejecución.

La estructura utilizada es:

```
Builder Stage

↓

Runtime Stage
```

Esto permite generar una imagen final más pequeña y segura.

---

# Stage 1 — Build

La primera etapa utiliza una imagen oficial de Go:

```dockerfile
FROM golang:1.26-alpine AS builder
```

Responsabilidades:

- Descargar dependencias.
- Compilar el proyecto.
- Generar el binario ejecutable.

---

## Instalación de dependencias necesarias

Durante la etapa de compilación se agregan certificados y herramientas necesarias:

```dockerfile
RUN apk add --no-cache ca-certificates git
```

Los certificados son necesarios para conexiones seguras mediante TLS, por ejemplo:

- conexiones HTTPS
- servicios externos
- Supabase

---

## Optimización del cache de Docker

Antes de copiar todo el proyecto se copian únicamente los archivos de dependencias:

```dockerfile
COPY go.mod go.sum ./
RUN go mod download
```

Esto permite aprovechar la caché de Docker.

Mientras las dependencias no cambien, Docker evita descargarlas nuevamente en cada build.

---

## Compilación del backend

La aplicación se compila como binario Linux:

```dockerfile
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o nexora-api main.go
```

Características:

- Binario independiente.
- Sin dependencia de CGO.
- Compatible con ejecución dentro de Alpine Linux.

---

# Stage 2 — Runtime

La segunda etapa utiliza una imagen mínima:

```dockerfile
FROM alpine:3.22
```

Su objetivo es ejecutar únicamente el binario generado.

No contiene:

- código fuente
- herramientas de desarrollo
- dependencias innecesarias

Esto reduce:

- tamaño de imagen
- superficie de ataque
- tiempo de despliegue

---

# Archivos incluidos en la imagen final

La imagen contiene:

```
/app

├── nexora-api
│
└── docs/
```

El binario generado es el único componente necesario para iniciar la API.

---

# Puerto expuesto

El backend escucha mediante:

```
8081
```

Configurado mediante:

```dockerfile
EXPOSE 8081
```

El puerto real puede variar según el entorno utilizando variables de configuración.

---

# Ejecución del contenedor

El contenedor inicia directamente el binario:

```dockerfile
CMD ["./nexora-api"]
```

No existe un servidor intermedio.

La aplicación Go es el proceso principal del contenedor.

---

# Docker Compose

El archivo `docker-compose.yml` permite ejecutar el backend localmente utilizando una configuración similar a producción.

Servicio principal:

```yaml
services:
  nexora-api:
```

---

## Build automático

Compose construye la imagen utilizando:

```yaml
build:
  context: .
  dockerfile: Dockerfile
```

Esto mantiene centralizada la definición de construcción dentro del Dockerfile.

---

## Variables de entorno

Las variables son cargadas mediante:

```yaml
env_file:
  - .env
```

Esto permite separar:

- código fuente
- configuración
- secretos

Nunca se almacenan valores sensibles dentro de la imagen.

---

# Persistencia de archivos

El backend utiliza un volumen para archivos generados:

```yaml
volumes:
  - ./uploads:/app/uploads
```

Esto permite mantener archivos persistentes fuera del ciclo de vida del contenedor.

La aplicación utiliza:

```
UPLOAD_DIR=/app/uploads
```

como directorio interno.

---

# Política de reinicio

El contenedor utiliza:

```yaml
restart: unless-stopped
```

Esto permite:

- recuperación automática ante fallos
- reinicio luego de reiniciar el servidor
- operación más estable en ambientes productivos

---

# Docker Ignore

El archivo `.dockerignore` evita copiar archivos innecesarios durante la construcción.

Se excluyen:

## Archivos de desarrollo

```
.vscode
.idea
tmp/
```

## Variables locales

```
.env.*
```

## Archivos generados

```
uploads/
coverage/
*.log
```

## Dependencias innecesarias

```
vendor/
```

## Código no requerido en runtime

```
internal/scripts/
deployment/
```

Esto mejora:

- velocidad de build
- tamaño de contexto enviado a Docker
- seguridad

---

# Variables de entorno

La aplicación utiliza configuración externa mediante variables de entorno.

El contenedor no contiene valores específicos de ningún ambiente.

La configuración cambia según:

- desarrollo
- staging
- producción

---

## Backend

Principales grupos de configuración:

### Servidor

Ejemplo:

```
SERVER_PORT
```

Define el puerto donde inicia la API.

---

### Base de datos

Configuración PostgreSQL:

```
DB_HOST
DB_PORT
DB_USER
DB_PASSWORD
DB_NAME
DB_SSLMODE
```

Permite conectar el backend con diferentes instancias sin modificar código.

---

### Storage

Configuración de archivos:

```
STORAGE_DRIVER
UPLOAD_DIR
```

Permite seleccionar la estrategia de almacenamiento.

Actualmente soporta:

- almacenamiento local
- integración con Supabase Storage

---

### Servicios externos

Configuración para proveedores externos:

```
SUPABASE_URL
SUPABASE_SERVICE_KEY
SUPABASE_BUCKET
```

Estos valores son administrados únicamente mediante variables seguras del entorno.

---

# Integración con Render

El backend se despliega en Render utilizando Docker como método de construcción.

Flujo:

```
Git Repository

↓

Render

↓

Docker Build

↓

Docker Image

↓

Container Running

↓

Nexora API
```

Render utiliza automáticamente el Dockerfile presente en el repositorio.

No requiere ejecutar manualmente:

```
go build
```

ni instalar dependencias del sistema.

---

# Frontend y Docker

El frontend no utiliza Docker actualmente.

Su despliegue es independiente:

```
Vue Application

↓

Build automático

↓

Render Static Site
```

La configuración se realiza mediante variables de entorno del propio servicio.

Ejemplo:

```
VITE_API_URL
VITE_ASSETS_URL
```

Esto permite apuntar el frontend hacia diferentes instancias del backend.

---

# Decisiones arquitectónicas

## Multi-stage build

Se decidió utilizar multi-stage build para:

- reducir tamaño final
- separar compilación y ejecución
- mejorar seguridad

---

## Configuración externa

Las variables de entorno permiten:

- cambiar infraestructura sin recompilar
- mantener secretos fuera del código
- soportar múltiples ambientes

---

## Contenedor como unidad de ejecución

El backend se ejecuta como una unidad independiente.

Esto facilita:

- despliegue
- escalado horizontal futuro
- migración de infraestructura

---

# Consideraciones futuras

Para una versión enterprise podrían incorporarse:

- imágenes privadas en registry
- pipelines CI/CD automáticos
- escaneo de vulnerabilidades
- health checks de Docker
- Kubernetes u otro sistema de orquestación
- separación de servicios adicionales mediante contenedores independientes

La implementación actual mantiene una base simple pero preparada para evolucionar hacia una infraestructura más compleja.
