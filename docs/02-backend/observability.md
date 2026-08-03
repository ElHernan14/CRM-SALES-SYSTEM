# Backend Observability Architecture

## Overview

La observabilidad de Nexora fue considerada como una parte fundamental de la arquitectura desde las primeras etapas del desarrollo.

El objetivo principal fue poder comprender qué ocurre dentro del sistema durante el ciclo completo de una request, facilitando:

- Debugging.
- Análisis de errores.
- Seguimiento de operaciones.
- Diagnóstico de problemas.
- Mantenimiento futuro.

La observabilidad no fue agregada como una solución posterior, sino diseñada como una capacidad transversal del backend.

---

# Observability Goals

Los principales objetivos fueron:

- Identificar cada request individualmente.
- Diferenciar errores funcionales de errores técnicos.
- Mantener información útil para debugging.
- Evitar exponer información sensible al usuario.
- Centralizar logs y manejo de errores.
- Facilitar investigación de problemas en producción.

La idea principal fue:

> El usuario debe recibir información clara y segura, mientras que el sistema debe conservar información suficiente para diagnosticar internamente cualquier problema.

---

# Request Tracking

Cada request posee un identificador único denominado:

```

request_id

```

Este identificador permite rastrear una operación completa dentro del sistema.

Ejemplo:

```

Request

↓

Request ID generado

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

El request_id permite reconstruir el recorrido de una operación incluso cuando existen múltiples usuarios ejecutando acciones simultáneamente.

---

# Request Metadata

La información asociada a una request permite identificar el contexto de ejecución.

Información relevante:

- Request ID.
- HTTP Method.
- Path.
- Status Code.
- Timestamp.
- Tiempo de ejecución.
- Resultado final.

Ejemplo conceptual:

```

REQUEST_ID=8a2f91

METHOD=POST

PATH=/api/invoice/pay

STATUS=200

TIME=120ms

```

---

# Middleware-Based Observability

La observabilidad está implementada mediante middleware transversal.

El middleware tiene acceso al ciclo completo de la request:

```

Before Request

↓

Execute Handler

↓

After Request

```

Esto permite registrar:

- Inicio de operación.
- Resultado.
- Tiempo total.
- Errores generados.

---

# Error Tracking Architecture

El manejo de errores está integrado con la observabilidad.

Existen dos categorías principales:

---

## Business Errors

Son errores esperados del dominio.

Ejemplos:

- Invoice inexistente.
- Estado inválido.
- Pago incorrecto.
- Falta de permisos.

Estos errores:

- Son controlados.
- Tienen mensajes específicos.
- Poseen códigos HTTP definidos.

Ejemplo:

```

409 Conflict

"Invoice cannot be paid"

```

---

## Technical Errors

Son errores inesperados del sistema.

Ejemplos:

- Fallos de base de datos.
- Problemas internos.
- Excepciones inesperadas.

Estos errores:

- No exponen información técnica al usuario.
- Son registrados internamente.
- Devuelven mensajes genéricos.

Ejemplo:

Respuesta usuario:

```json
{
  "success": false,
  "message": "Internal server error"
}
```

Log interno:

```
database connection timeout
```

---

# Centralized Error Middleware

Los errores fueron centralizados para evitar repetir lógica en cada controller o service.

El flujo es:

```
Service

↓

Return Error

↓

Error Middleware

↓

Classify Error

↓

HTTP Response
```

Responsabilidades:

- Detectar errores de negocio.
- Detectar errores técnicos.
- Generar respuesta consistente.
- Registrar información relevante.

---

# Panic Recovery

El sistema cuenta con un middleware de recuperación global.

Su objetivo es capturar errores inesperados provocados por:

- Panics.
- Fallos no controlados.
- Excepciones internas.

Flujo:

```
Panic

↓

Recovery Middleware

↓

Log Error

↓

HTTP 500 Response
```

Esto evita que un error inesperado provoque una caída completa del servidor.

---

# Logging Strategy

Los logs fueron pensados para ser útiles sin generar información redundante.

Información registrada:

- Timestamp.
- Request ID.
- Endpoint.
- Método HTTP.
- Tipo de error.
- Mensaje interno.
- Resultado.

Información que debe evitarse:

- Passwords.
- Tokens completos.
- Datos sensibles.
- Información privada innecesaria.

---

# Debugging Workflow

Ante un problema operativo, el proceso recomendado sería:

## 1. Identificar request

Buscar:

```
request_id
```

---

## 2. Revisar contexto

Analizar:

- Endpoint.
- Usuario involucrado.
- Timestamp.
- Resultado.

---

## 3. Revisar lógica ejecutada

Seguir el flujo:

```
Controller

↓

Service

↓

Repository
```

---

## 4. Validar persistencia

Revisar:

- Estado de entidades.
- Transacciones.
- Datos almacenados.

---

# Business Observability

Además de errores técnicos, los logs permiten analizar comportamiento del negocio.

Ejemplos:

- Facturas rechazadas.
- Pagos fallidos.
- Problemas de stock.
- Errores de checkout.

Esto permite detectar problemas funcionales y mejorar procesos.

---

# Current Observability Level

Actualmente Nexora posee:

- Request tracking.
- Error centralizado.
- Recovery global.
- Logging estructurado básico.
- Separación entre errores funcionales y técnicos.

Estas capacidades permiten un mantenimiento efectivo durante desarrollo y primeras etapas productivas.

---

# Future Improvements

Para una versión enterprise podrían incorporarse:

## Métricas

Ejemplos:

- Requests por segundo.
- Latencia promedio.
- Error rate.
- Tiempo de respuesta por endpoint.

---

## Distributed Tracing

Permitiría seguir operaciones entre múltiples servicios.

Ejemplo:

```
API

↓

Inventory

↓

Payment

↓

Notification
```

---

## Monitoring Platforms

Integraciones posibles:

- Prometheus.
- Grafana.
- OpenTelemetry.
- Sistemas centralizados de logs.

---

## Alerting

Alertas automáticas sobre:

- Caídas del sistema.
- Incremento de errores.
- Saturación de recursos.
- Latencias elevadas.

---

# Production Considerations

En un entorno productivo sería recomendable complementar la observabilidad actual con:

- Logs estructurados en formato JSON.
- Correlación mediante request_id.
- Retención de logs.
- Métricas históricas.
- Dashboards operativos.
- Alertas automáticas.

---

# Final Considerations

La observabilidad de Nexora fue diseñada siguiendo un principio fundamental:

> Un sistema mantenible no solo debe funcionar correctamente, también debe permitir entender qué ocurrió cuando algo falla.

La incorporación temprana de request tracking, manejo centralizado de errores y recuperación global permite que el backend sea más fácil de diagnosticar, mantener y evolucionar.

Esta base deja preparado al sistema para incorporar herramientas avanzadas de monitoreo cuando el producto alcance una escala mayor.

```

```
