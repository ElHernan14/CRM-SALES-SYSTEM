# Backend API Design

## Overview

La API de Nexora fue diseñada siguiendo principios REST, priorizando consistencia, legibilidad y escalabilidad.

Antes de implementar endpoints individuales, se definió primero la infraestructura necesaria para garantizar que cada request tuviera un recorrido claro y controlado:

```

Request

↓

Middleware

↓

Controller

↓

Service

↓

Repository

↓

Database

```

El diseño de la API no fue pensado únicamente como una colección de endpoints, sino como una capa de entrada organizada sobre los diferentes dominios del sistema.

Cada endpoint representa una acción concreta dentro del negocio, manteniendo una separación clara entre operaciones simples y procesos complejos.

---

# API Design Principles

Los principales objetivos durante el diseño fueron:

- Mantener una estructura REST consistente.
- Separar dominios funcionales.
- Evitar endpoints genéricos con demasiada responsabilidad.
- Exponer únicamente información necesaria.
- Mantener compatibilidad futura.
- Separar transporte HTTP de lógica de negocio.

La API actúa como una puerta de entrada hacia los casos de uso del sistema, pero no contiene reglas de negocio.

---

# Route Organization

Las rutas fueron organizadas por módulos funcionales.

Ejemplo conceptual:

```

/api/auth

/api/users

/api/company

/api/products

/api/invoices

/api/inventory

/api/store

```

Cada módulo mantiene sus propios endpoints relacionados.

Esto evita mezclar responsabilidades y permite que nuevos dominios puedan incorporarse sin afectar módulos existentes.

---

# Public and Private Routes

La API diferencia entre rutas públicas y privadas.

## Public Routes

Son endpoints que no requieren autenticación.

Ejemplos:

- Login.
- Registro.
- Información pública.

Flujo:

```

Request

↓

Controller

↓

Service

```

---

## Private Routes

Requieren autenticación y autorización.

Flujo:

```

Request

↓

JWT Middleware

↓

Permission Middleware

↓

Controller

↓

Service

```

Las rutas privadas siempre operan bajo un Tenant Context válido.

---

# CRUD vs Business Operations

Una decisión importante fue diferenciar operaciones CRUD simples de procesos de negocio.

## CRUD Operations

Representan acciones directas sobre entidades.

Características:

- Una entidad principal.
- Operación simple.
- Pocas reglas.
- Sin coordinación entre dominios.

Ejemplo:

```

GET /products

POST /clients

```

---

## Business Operations

Representan procesos reales del negocio.

Características:

- Involucran múltiples entidades.
- Contienen reglas de negocio.
- Requieren validaciones.
- Pueden necesitar transacciones.

Ejemplos:

```

POST /invoice/submit

POST /invoice/pay

POST /store/checkout

```

Estos endpoints no representan simplemente una modificación de datos, sino una acción del dominio.

---

# Endpoint Design Decisions

Los endpoints complejos fueron diseñados como acciones explícitas.

Ejemplo:

Incorrecto:

```

PUT /invoice

```

para cambiar cualquier estado.

Correcto:

```

POST /invoice/submit

POST /invoice/pay

```

La razón es que cada operación representa una transición específica dentro del ciclo de vida de la entidad.

Esto permite:

- Mejor trazabilidad.
- Validaciones específicas.
- Menor ambigüedad.
- Mayor claridad del dominio.

---

# Controller Responsibility

Los controllers funcionan como capa de traducción entre HTTP y la aplicación.

Sus responsabilidades son:

- Recibir request.
- Parsear parámetros.
- Validar estructura de entrada.
- Ejecutar el caso de uso correspondiente.
- Construir response HTTP.

Los controllers no contienen:

- Reglas de negocio.
- Consultas SQL.
- Validaciones complejas del dominio.

---

# Service Exposure

La API expone casos de uso mediante servicios.

El controller delega la operación:

```

HTTP Request

↓

Controller

↓

Service

↓

Business Logic

```

Esto permite que la lógica principal pueda reutilizarse desde diferentes puntos de entrada.

Ejemplo:

El servicio `SubmitInvoice` puede ser utilizado desde:

- Endpoint HTTP.
- Store Checkout.
- Marketplace Checkout.

Sin duplicar lógica.

---

# Response Design

La API utiliza respuestas consistentes en toda la aplicación.

Todas las respuestas siguen estructuras comunes.

Ejemplo conceptual:

Success:

```json
{
  "success": true,
  "data": {}
}
```

Error:

```json
{
  "success": false,
  "code": 400,
  "message": "Invalid invoice state"
}
```

Esto permite que frontend y backend mantengan un contrato estable.

---

# Error Responses

Los errores se clasifican en dos grupos principales.

## Business Errors

Errores esperados del dominio.

Ejemplos:

- Invoice no encontrada.
- Estado inválido.
- Pago excedido.
- Permiso insuficiente.

Estos errores retornan mensajes controlados.

---

## Technical Errors

Errores inesperados de infraestructura.

Ejemplos:

- Fallos de base de datos.
- Errores internos.
- Problemas de conexión.

Estos errores devuelven mensajes genéricos al usuario mientras mantienen información detallada en logs.

---

# Data Exposure

La API fue diseñada siguiendo el principio de mínima exposición.

No se devuelve información innecesaria.

Ejemplos de datos ocultos:

- Passwords.
- Información interna.
- Datos técnicos.
- Campos no requeridos por el caso de uso.

Cada respuesta está adaptada al consumidor del endpoint.

---

# DTO and Data Contracts

Los contratos de entrada y salida se mantienen separados del modelo interno.

Esto permite:

- Evitar exponer entidades completas.
- Evolucionar internamente sin romper clientes.
- Controlar información enviada.

La API representa casos de uso, no tablas de base de datos.

---

# API Evolution Strategy

Para mantener compatibilidad futura se consideran diferentes estrategias:

## Versionado

Ejemplo:

```
/api/v1/invoices
```

Permite evolucionar contratos sin romper consumidores existentes.

---

## Nuevos Endpoints

Agregar funcionalidades nuevas mediante nuevos recursos.

Evitar modificar comportamiento existente.

---

## Deprecación Gradual

Los endpoints antiguos pueden mantenerse temporalmente mientras los consumidores migran.

---

# Future Improvements

Para una versión productiva de mayor escala podrían incorporarse:

- OpenAPI completamente automatizado.
- Versionado formal.
- Rate limiting.
- API Gateway.
- Pagination estándar.
- Filtering avanzado.
- HATEOAS en recursos públicos.
- Documentación automática.

---

# Final Considerations

El diseño de API de Nexora busca mantener una interfaz clara entre el mundo externo y el dominio interno.

Los endpoints representan acciones reales del negocio, mientras que la lógica permanece encapsulada en los servicios correspondientes.

La separación entre transporte HTTP, casos de uso y persistencia permite que la API pueda evolucionar sin comprometer la estabilidad del sistema.

Esta decisión arquitectónica mantiene una base preparada para crecimiento, nuevos módulos y futuros consumidores externos.

```

```
