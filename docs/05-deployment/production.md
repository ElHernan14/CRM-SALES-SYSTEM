# Despliegue en Producción — Nexora

## Introducción

Nexora fue diseñado con una separación clara entre aplicación, infraestructura y configuración.

El entorno productivo mantiene la misma filosofía utilizada durante el desarrollo:

- Backend independiente.
- Frontend independiente.
- Configuración externa mediante variables de entorno.
- Servicios administrados para infraestructura crítica.
- Separación entre código y recursos persistentes.

El objetivo principal del despliegue es garantizar que la aplicación pueda evolucionar sin depender de configuraciones específicas del entorno de desarrollo.

---

# Arquitectura de Producción

La arquitectura productiva está compuesta por los siguientes servicios principales:

```text
                    Usuario

                       │

                       ▼

              Frontend Vue
              Render Static Site

                       │

                 HTTP API Requests

                       │

                       ▼

              Backend Go API
              Docker Container

                       │

        ┌──────────────┴──────────────┐

        ▼                             ▼

 PostgreSQL Database             Storage Service

     Supabase                 Local / Supabase Storage
```

Cada componente posee una responsabilidad específica.

---

# Componentes principales

## Frontend

Tecnología:

- Vue 3
- Vite
- TypeScript

Responsabilidad:

- Renderizado de interfaz.
- Gestión de navegación.
- Manejo de estado visual.
- Comunicación con la API.

El frontend no contiene lógica de negocio crítica.

Toda regla de negocio permanece en el backend.

---

## Backend

Tecnología:

- Go
- REST API
- Docker

Responsabilidad:

- Casos de uso.
- Seguridad.
- Autorización.
- Multi-tenancy.
- Validaciones.
- Persistencia.
- Reglas de dominio.

El backend funciona como única fuente de verdad del sistema.

---

## Base de datos

Tecnología:

- PostgreSQL
- Supabase

Responsabilidad:

- Persistencia de entidades.
- Relaciones entre dominios.
- Estados de negocio.
- Información multi-tenant.

El acceso se realiza únicamente desde la API backend.

El frontend nunca accede directamente a la base de datos.

---

## Storage

Nexora utiliza almacenamiento externo para manejar recursos asociados a la aplicación.

Actualmente contempla:

- almacenamiento local
- Supabase Storage

La estrategia utilizada permite cambiar el proveedor sin modificar la lógica de negocio.

---

# Flujo de despliegue

El proceso general es:

```text
Código fuente

↓

Repositorio Git

↓

Render

↓

Build

↓

Deploy

↓

Aplicación disponible
```

Cada servicio se despliega de forma independiente.

---

# Deploy del Backend

El backend se despliega como un servicio Docker.

Flujo:

```text
Repositorio Backend

↓

Render detecta Dockerfile

↓

Construcción de imagen

↓

Creación de container

↓

Ejecución de API Go
```

Durante el proceso:

1. Render obtiene el código fuente.
2. Construye la imagen Docker.
3. Inyecta variables de entorno.
4. Inicia el contenedor.
5. Expone la API mediante HTTP.

---

# Deploy del Frontend

El frontend utiliza un flujo diferente.

Proceso:

```text
Repositorio Frontend

↓

Render Build

↓

npm install

↓

npm run build

↓

Publicación de archivos estáticos
```

La aplicación genera recursos optimizados:

- HTML
- CSS
- JavaScript
- Assets

Estos archivos son servidos por Render.

---

# Variables de entorno en producción

La configuración productiva se administra mediante variables externas.

Esto permite:

- evitar información sensible en código.
- cambiar infraestructura sin recompilar.
- manejar múltiples ambientes.

---

# Backend Environment

Las principales categorías son:

## Servidor

Configuración del servidor HTTP:

```text
SERVER_PORT
```

---

## Base de datos

Conexión PostgreSQL:

```text
DB_HOST

DB_PORT

DB_USER

DB_PASSWORD

DB_NAME

DB_SSLMODE
```

---

## Storage

Configuración de almacenamiento:

```text
STORAGE_DRIVER

UPLOAD_DIR
```

---

## Servicios externos

Integración con Supabase:

```text
SUPABASE_URL

SUPABASE_SERVICE_KEY

SUPABASE_BUCKET
```

Los valores sensibles deben mantenerse únicamente dentro del sistema de configuración del proveedor.

---

# Frontend Environment

El frontend utiliza variables propias del entorno Vue:

```text
VITE_API_URL

VITE_ASSETS_URL
```

Estas variables determinan:

- URL del backend.
- Ubicación de recursos estáticos.

El frontend no conoce detalles internos de infraestructura.

---

# Separación entre ambientes

Nexora contempla diferentes configuraciones según el ambiente:

## Desarrollo

Características:

- Backend local.
- Frontend local.
- Variables orientadas a pruebas.
- Debug habilitado.

---

## Producción

Características:

- Servicios desplegados.
- Variables seguras.
- Base de datos externa.
- HTTPS.
- Configuración optimizada.

---

# Seguridad en producción

Algunas medidas aplicadas:

## Secretos fuera del código

Nunca se almacenan:

- passwords.
- tokens.
- keys privadas.

en el repositorio.

---

## HTTPS

La comunicación productiva utiliza conexiones seguras mediante los servicios de hosting.

---

## Backend como única autoridad

Aunque el frontend realiza validaciones preventivas, toda autorización real ocurre en backend.

Esto evita:

- manipulación del cliente.
- acceso indebido.
- inconsistencias.

---

# Persistencia y datos

Los datos importantes no dependen del ciclo de vida del container.

El contenedor puede ser recreado sin perder información.

La persistencia pertenece a:

- PostgreSQL.
- Storage externo.

---

# Manejo de archivos

Los archivos generados por la aplicación utilizan una estrategia desacoplada.

El backend no asume una ubicación física permanente.

La configuración define el mecanismo utilizado.

Esto permite evolucionar hacia:

- almacenamiento cloud.
- CDN.
- servicios externos.

---

# Consideraciones operativas

En un entorno productivo se deben monitorear:

## Disponibilidad

Verificar:

- API funcionando.
- Frontend accesible.
- Base de datos disponible.

---

## Errores

Utilizar:

- logs del backend.
- errores HTTP.
- request tracking.

---

## Rendimiento

Observar:

- tiempos de respuesta.
- consultas lentas.
- consumo de recursos.

---

# Estrategia de escalabilidad futura

La arquitectura actual permite evolucionar progresivamente.

Posibles mejoras:

## Backend

- múltiples instancias.
- balanceador de carga.
- cache.
- colas de procesamiento.

---

## Base de datos

- optimización de índices.
- replicas de lectura.
- separación por servicios.

---

## Infraestructura

- CI/CD automático.
- infraestructura como código.
- contenedores orquestados.

---

# Estado actual de la arquitectura

La infraestructura actual representa una solución equilibrada entre simplicidad y escalabilidad.

No introduce complejidad innecesaria, pero mantiene principios utilizados en sistemas profesionales:

- separación de responsabilidades.
- configuración externa.
- servicios independientes.
- despliegue reproducible.
- preparación para crecimiento futuro.

La arquitectura productiva de Nexora permite operar el sistema actualmente y al mismo tiempo mantener una base preparada para evolucionar hacia una infraestructura empresarial más avanzada.
