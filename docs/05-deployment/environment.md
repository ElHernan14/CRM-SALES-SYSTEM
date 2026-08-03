# Gestión de Entornos y Configuración

## Introducción

Nexora fue diseñado considerando la separación entre código fuente y configuración de ejecución.

La configuración del sistema no se encuentra embebida dentro del código de la aplicación, sino que se administra mediante variables de entorno independientes según el ambiente donde sea ejecutado.

Esta decisión permite mantener una misma base de código funcionando en distintos escenarios:

- Desarrollo local
- Testing
- Producción

sin modificar lógica interna ni realizar cambios manuales sobre la aplicación.

La configuración externa permite además una integración más segura con servicios externos como bases de datos, almacenamiento de archivos y plataformas de despliegue.

---

# Principios de configuración utilizados

Durante el desarrollo se siguieron los siguientes principios:

## Separación de código y configuración

Los valores sensibles o dependientes del ambiente no forman parte del repositorio.

Ejemplos:

- Credenciales de base de datos
- Secretos JWT
- URLs externas
- Configuración de servicios
- Claves privadas

El código solamente conoce contratos de configuración, no valores concretos.

---

## Configuración centralizada

Cada aplicación posee una capa encargada de cargar y exponer configuración.

La responsabilidad de esta capa es:

- Leer variables de entorno.
- Validar existencia de valores requeridos.
- Proveer configuración al resto del sistema.

Las demás capas de la aplicación no acceden directamente al entorno.

---

# Ambientes contemplados

## Desarrollo local

El ambiente de desarrollo permite ejecutar Nexora desde una máquina local.

Su objetivo principal es:

- Desarrollo activo.
- Debug.
- Pruebas funcionales.
- Validación de nuevas funcionalidades.

Características principales:

Backend:

- Servidor Go ejecutándose localmente.
- Conexión a PostgreSQL local.
- Variables de entorno propias.

Frontend:

- Aplicación Vue ejecutada mediante Vite.
- API apuntando al backend local.

---

## Producción

El ambiente productivo utiliza configuraciones independientes.

Su objetivo es:

- Ejecutar servicios estables.
- Mantener seguridad de credenciales.
- Separar infraestructura real del entorno de desarrollo.

En producción:

- Las variables son proporcionadas por la plataforma de despliegue.
- No existen valores sensibles dentro del repositorio.
- El frontend utiliza URLs productivas.
- El backend utiliza conexiones externas configuradas.

---

# Variables de entorno Backend

El backend en Go utiliza variables externas para definir su comportamiento.

Las principales configuraciones corresponden a:

## Base de datos

Permiten establecer la conexión con PostgreSQL.

Incluyen información como:

- Host
- Puerto
- Usuario
- Password
- Nombre de base de datos

La aplicación no conoce valores concretos, únicamente consume la configuración cargada.

---

## Seguridad

La configuración relacionada con autenticación incluye:

- JWT secret.
- Tiempo de expiración del token.
- Parámetros utilizados para validación.

Esto permite modificar políticas de seguridad sin recompilar la aplicación.

---

## Servidor

Incluye parámetros como:

- Puerto HTTP.
- Ambiente de ejecución.
- Configuración general del runtime.

---

# Variables de entorno Frontend

El frontend utiliza variables de entorno administradas por Vite.

Principalmente:

## URL del Backend

Permite definir contra qué API debe comunicarse la aplicación.

Ejemplo conceptual:

```

Desarrollo:

[http://localhost:8080](http://localhost:8080)

Producción:

[https://api-produccion.com](https://api-produccion.com)

```

Esto permite generar diferentes builds sin modificar código fuente.

---

## Recursos externos

También se contemplan configuraciones relacionadas con servicios externos.

Ejemplo:

- URLs de almacenamiento.
- Recursos estáticos.
- Servicios auxiliares.

---

# Seguridad de configuración

Las variables sensibles nunca son almacenadas dentro del repositorio.

El proyecto evita:

- Secrets hardcodeados.
- Credenciales visibles.
- Configuración dependiente del ambiente dentro del código.

Los archivos locales de configuración permanecen fuera del control de versiones.

Ejemplo:

```

.env

.env.production

.env.local

```

son considerados archivos de ambiente y no forman parte del código compartido.

---

# Integración con despliegue

La arquitectura de configuración permite que las plataformas de despliegue puedan inyectar valores externos.

Actualmente Nexora utiliza:

- Render para despliegue.
- Supabase Storage para almacenamiento externo.
- PostgreSQL como base de datos.

Cada servicio recibe únicamente las variables necesarias para funcionar.

---

# Decisiones arquitectónicas

## No almacenar configuración dentro del código

Se decidió separar configuración porque:

- Reduce riesgos de seguridad.
- Facilita despliegues.
- Permite múltiples ambientes.
- Evita modificaciones manuales antes de publicar.

---

## No acoplar la aplicación al proveedor de infraestructura

La aplicación no depende directamente de Render u otro proveedor.

La infraestructura solamente entrega:

- Variables.
- Recursos.
- Runtime.

La aplicación mantiene independencia del ambiente donde sea ejecutada.

---

# Evolución futura

Para una versión enterprise podrían incorporarse herramientas adicionales:

- Secret Managers.
- Gestión automática de configuración.
- Rotación de credenciales.
- Configuración distribuida.

Ejemplos:

- AWS Secrets Manager.
- Hashicorp Vault.
- Servicios equivalentes.

---

# Conclusión

La gestión de ambientes en Nexora fue diseñada bajo el principio de separación entre aplicación e infraestructura.

La configuración externa permite mantener un sistema flexible, seguro y preparado para evolucionar desde desarrollo local hacia ambientes productivos sin alterar la base de código.

```

```
