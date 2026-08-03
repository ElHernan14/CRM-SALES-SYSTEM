# Error Handling — Nexora

## Introducción

El manejo de errores de Nexora fue diseñado como una responsabilidad transversal de la arquitectura.

Desde el inicio del desarrollo se decidió evitar que cada módulo, controller o servicio implementara su propia lógica para:

- interpretar errores.
- definir códigos HTTP.
- construir respuestas.
- registrar logs.
- diferenciar errores funcionales de errores técnicos.

La estrategia adoptada fue centralizar el manejo de errores mediante una capa compartida, permitiendo mantener una convención uniforme en toda la aplicación.

El objetivo principal fue lograr:

- respuestas consistentes.
- menor duplicación de código.
- separación de responsabilidades.
- mejor trazabilidad.
- facilidad de debugging.

---

# Principios de diseño

El sistema sigue los siguientes principios:

## Las capas inferiores no conocen HTTP

Los servicios y repositorios no deben decidir:

- códigos HTTP.
- formatos de respuesta.
- estructura JSON.

Su responsabilidad es devolver errores representativos del problema ocurrido.

---

## El controller no maneja toda la lógica de error

El controller únicamente ejecuta el caso de uso y propaga el resultado.

No debe contener lógica repetitiva para:

- mapear errores.
- construir respuestas.
- registrar excepciones.

---

## El middleware es responsable de la traducción final

La capa HTTP interpreta los errores recibidos y decide:

- status code.
- mensaje visible.
- formato de respuesta.

---

# Arquitectura del flujo de errores

El flujo general es:

```

Repository

↓

Service

↓

Controller

↓

Error Middleware

↓

HTTP Response

```

Cada capa mantiene su responsabilidad.

---

# Tipos de errores

Nexora diferencia principalmente dos categorías:

## Errores de negocio

Representan situaciones esperadas dentro del dominio.

Ejemplos:

- invoice inexistente.
- usuario sin permisos.
- estado inválido.
- monto superior al permitido.
- operación no disponible.

Estos errores pueden mostrarse al usuario porque representan reglas funcionales conocidas.

---

## Errores técnicos

Representan fallos internos del sistema.

Ejemplos:

- error de base de datos.
- fallo de infraestructura.
- problemas inesperados.
- errores internos.

Estos errores no deben exponerse directamente al usuario.

---

# AppError

Para representar errores controlados se creó un wrapper común:

```go
type AppError struct {
	Code    int
	Message string
}
```

Este objeto encapsula:

- código HTTP asociado.
- mensaje funcional.

Ejemplo:

```go
NewAppError(
    404,
    "invoice not found",
)
```

El objetivo es transportar información controlada desde las capas internas hasta el middleware HTTP.

---

# Errores estándar

El sistema contempla errores comunes:

```go
var (
    ErrBadRequest   = NewAppError(400, "Bad request")
    ErrUnauthorized = NewAppError(401, "Unauthorized")
    ErrForbidden    = NewAppError(403, "Forbidden")
    ErrConflict     = NewAppError(409, "Conflict")
    ErrInternal     = NewAppError(500, "Internal server error")
)
```

Estas definiciones permiten mantener una convención global.

Aunque actualmente algunos mensajes son dinámicos por dominio, la estructura permite una futura evolución hacia:

- internacionalización.
- catálogo centralizado.
- códigos internos de error.

---

# Error Middleware

El middleware de errores funciona como punto único de interpretación.

Responsabilidades:

- capturar errores provenientes de controllers.
- identificar el tipo de error.
- construir respuesta HTTP.
- registrar información relevante.

---

# Flujo del middleware

Conceptualmente:

```
Controller retorna error

↓

Middleware recibe error

↓

¿Es AppError?

↓

Sí:
    devolver código y mensaje controlado

No:
    devolver Internal Server Error

↓

Registrar log
```

---

# Manejo de errores de negocio

Cuando el middleware detecta un `AppError`:

Obtiene:

- código HTTP.
- mensaje personalizado.

Luego responde:

```json
{
  "code": 400,
  "message": "invoice not payable"
}
```

Esto permite que el usuario reciba información útil sin exponer detalles internos.

---

# Manejo de errores técnicos

Cuando el error no pertenece a `AppError`:

El usuario recibe una respuesta genérica:

```json
{
  "code": 500,
  "message": "Internal server error"
}
```

Pero internamente se registra el error real.

Ejemplo:

```
Database connection failed
constraint violation
unexpected nil pointer
```

Esto mantiene una separación correcta entre:

- experiencia del usuario.
- información interna del sistema.

---

# Recovery Middleware

Además del middleware de errores existe un middleware de recuperación de panics.

Su objetivo es evitar que un error inesperado derribe completamente el servidor.

Flujo:

```
Request

↓

Recovery Middleware

↓

Application

↓

Panic

↓

Recover

↓

500 Response
```

---

# Responsabilidades del Recovery Middleware

Captura:

- panic inesperados.
- errores no controlados.

Luego:

- registra el incidente.
- evita caída del proceso.
- devuelve respuesta segura.

---

# Logging y trazabilidad

Los errores registrados incluyen información útil para debugging:

- timestamp.
- método HTTP.
- path.
- request id.
- mensaje del error.
- contexto disponible.

El objetivo es poder reconstruir el recorrido de una operación fallida.

---

# Relación con Request ID

Cada request posee un identificador único.

Este ID permite relacionar:

- logs del middleware.
- errores.
- operaciones internas.
- debugging concurrente.

Ejemplo conceptual:

```
Request ID:

a83f-92bd-7712
```

permite buscar toda la información asociada a una ejecución específica.

---

# Ventajas de centralizar errores

## Consistencia

Todas las respuestas mantienen la misma estructura.

---

## Menor duplicación

Los controllers no repiten lógica de manejo.

---

## Mayor seguridad

Los errores internos no son expuestos.

---

## Mejor mantenimiento

Modificar la estrategia global de errores requiere cambios en un único punto.

---

## Preparación para producción

La arquitectura permite agregar:

- observabilidad avanzada.
- métricas.
- alertas.
- códigos internos.
- integración con sistemas externos.

---

# Decisiones arquitectónicas relevantes

## Separar errores funcionales y técnicos

Fue una decisión importante para diferenciar:

"El usuario realizó una operación inválida"

de

"El sistema tiene un problema interno".

---

## Centralizar la respuesta HTTP

Evita que cada endpoint tenga comportamientos diferentes.

---

## Mantener servicios independientes del transporte

El dominio no conoce HTTP.

Esto permite reutilizar lógica en:

- APIs futuras.
- procesos batch.
- workers.
- servicios internos.

---

# Evaluación arquitectónica

El manejo de errores de Nexora sigue una estrategia alineada con arquitecturas empresariales modernas.

La aplicación evita mezclar responsabilidades y mantiene una separación clara entre:

- lógica de negocio.
- infraestructura.
- transporte HTTP.
- experiencia del usuario.

La centralización mediante middleware permite que el sistema sea más predecible, seguro y mantenible, especialmente a medida que aumenta la cantidad de módulos y casos de uso.

Esta decisión resulta fundamental para escalar el proyecto sin multiplicar comportamientos inconsistentes a través de la aplicación.

```

```
