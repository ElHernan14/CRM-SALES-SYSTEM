# Convenciones de Desarrollo — Nexora

## Introducción

Nexora fue construido siguiendo una serie de convenciones arquitectónicas y de desarrollo cuyo objetivo principal es mantener la consistencia del código a medida que el sistema crece.

Estas convenciones no buscan imponer una única forma de programar, sino establecer reglas claras para que nuevos desarrolladores puedan comprender rápidamente la estructura del proyecto y agregar funcionalidades sin degradar la arquitectura existente.

La premisa principal es:

> La evolución del sistema debe mantener las mismas responsabilidades y límites definidos desde el inicio.

---

# Principios generales

Durante el desarrollo se priorizaron los siguientes principios:

- Separación de responsabilidades.
- Bajo acoplamiento entre módulos.
- Alta cohesión dentro de cada dominio.
- Lógica de negocio centralizada en services.
- Persistencia aislada mediante repositories.
- Controllers delgados.
- Infraestructura desacoplada.
- Código orientado a contratos mediante interfaces.

---

# Organización por módulos

Nexora organiza su backend siguiendo una estructura orientada a dominios.

Cada módulo representa una responsabilidad del negocio.

Ejemplo conceptual:

```text
invoice/

├── controller

├── service

├── repository

├── model

├── dto

└── interfaces
```

Cada módulo mantiene únicamente la lógica relacionada con su propio dominio.

---

# Responsabilidad de cada capa

## Controller

El controller representa la entrada HTTP al sistema.

Responsabilidades:

- recibir requests.
- validar estructura básica.
- obtener parámetros.
- invocar servicios.
- devolver responses.

No debe contener:

- reglas de negocio.
- consultas SQL.
- lógica de autorización compleja.

El controller conoce el caso de uso, pero no cómo se ejecuta internamente.

---

## Service

El service contiene la lógica principal del negocio.

Responsabilidades:

- validar reglas.
- coordinar procesos.
- ejecutar casos de uso.
- controlar workflows.
- interactuar con repositories.

Ejemplo:

```text
Crear invoice

↓

Validar cliente

↓

Validar productos

↓

Persistir información

↓

Actualizar estado
```

Toda decisión de negocio debe vivir en esta capa.

---

## Repository

El repository representa la comunicación con persistencia.

Responsabilidades:

- ejecutar consultas.
- obtener datos.
- almacenar información.
- mapear resultados.

No debe contener:

- reglas de negocio.
- decisiones funcionales.
- validaciones de dominio.

El repository responde:

"¿Cómo obtengo o guardo información?"

No:

"¿Cuándo debería hacerlo?"

---

# Servicios encapsulados

Cuando una operación involucra múltiples dominios, no se fuerza la lógica dentro de un único módulo.

Se utilizan servicios de workflow.

Ejemplo:

```text
Submit Invoice

Invoice

+

Client

+

Inventory

+

Permissions
```

En lugar de mezclar responsabilidades, el flujo se encapsula en un servicio específico.

Esto permite mantener:

- módulos independientes.
- procesos complejos organizados.
- reglas claras.

---

# Interfaces

Las interfaces se utilizan para reducir dependencia entre componentes.

Principio utilizado:

> Los módulos deben depender de contratos, no de implementaciones concretas.

Ejemplo:

Un service no necesita conocer una implementación específica de repository.

Solamente conoce la interfaz requerida.

Beneficios:

- testing más simple.
- menor acoplamiento.
- reemplazo de implementaciones.

---

# Manejo de errores

Los errores siguen una estrategia centralizada.

Los servicios pueden devolver:

- errores de negocio controlados.
- errores técnicos.

La capa superior interpreta el tipo de error y genera la respuesta correspondiente.

No se deben crear respuestas HTTP manualmente desde cada service.

---

# Validaciones

Las validaciones se dividen según responsabilidad.

## Validaciones de transporte

Responsabilidad:

Controller.

Ejemplos:

- JSON válido.
- parámetros existentes.
- formato básico.

---

## Validaciones de negocio

Responsabilidad:

Service.

Ejemplos:

- estado permitido.
- permisos.
- reglas del dominio.
- restricciones funcionales.

---

## Validaciones de persistencia

Responsabilidad:

Database / Repository.

Ejemplos:

- claves únicas.
- integridad referencial.

---

# Naming conventions

Se mantiene una nomenclatura consistente.

## Archivos

Ejemplos:

```text
invoice.service.go

invoice.repository.go

invoice.controller.go
```

El nombre debe indicar:

- dominio.
- responsabilidad.
- capa.

---

## Servicios

Los nombres deben representar acciones de negocio.

Ejemplo:

Correcto:

```text
SubmitInvoice()

PayInvoice()

EnsureCart()
```

Evitar:

```text
Process()

Handle()

Execute()
```

cuando no representan claramente la intención.

---

# DTOs

Los DTOs representan contratos de entrada y salida.

No deben exponer directamente modelos internos.

Beneficios:

- desacoplamiento.
- seguridad.
- evolución independiente.

---

# Uso del Tenant Context

Toda operación relacionada con datos multi-tenant debe utilizar el contexto generado desde autenticación.

No se deben obtener identificadores críticos desde parámetros externos cuando ya existen dentro del contexto.

Ejemplo:

Incorrecto:

```text
company_id enviado por frontend
```

Correcto:

```text
company_id obtenido desde TenantContext
```

---

# Reglas para nuevos módulos

Antes de crear un módulo nuevo se debe responder:

- ¿Representa un dominio real?
- ¿Tiene ciclo de vida propio?
- ¿Posee reglas independientes?
- ¿Debe interactuar con otros dominios?

Si la respuesta es positiva, probablemente corresponda crear un nuevo módulo.

---

# Reglas para mantener la arquitectura

Al agregar código nuevo se debe evitar:

- lógica de negocio en controllers.
- consultas SQL dentro de services.
- servicios gigantes.
- dependencia circular entre módulos.
- duplicación de validaciones.
- acceso directo a datos sin contexto tenant.

---

# Filosofía general del desarrollo

La arquitectura de Nexora fue diseñada para que agregar funcionalidades sea una extensión del sistema, no una modificación constante de código existente.

Cada nueva funcionalidad debe incorporarse respetando:

- dominio correspondiente.
- responsabilidad correcta.
- capa adecuada.
- contratos definidos.

El objetivo final es mantener un sistema donde el crecimiento no implique perder claridad ni mantenibilidad.
