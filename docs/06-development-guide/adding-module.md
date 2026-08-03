# Agregar un Nuevo Módulo — Nexora

## Introducción

Nexora está diseñado bajo una arquitectura modular orientada a dominios.

Esto significa que nuevas funcionalidades no deben agregarse simplemente creando nuevos endpoints o tablas, sino identificando primero si representan un nuevo dominio dentro del sistema.

La incorporación de un módulo debe respetar las decisiones arquitectónicas existentes:

- separación de responsabilidades.
- organización por dominio.
- bajo acoplamiento.
- lógica de negocio encapsulada.
- comunicación mediante contratos definidos.

El objetivo es que agregar nuevas capacidades sea una extensión natural del sistema y no una modificación estructural.

---

# ¿Cuándo crear un nuevo módulo?

Antes de crear un módulo se debe evaluar si la nueva funcionalidad representa realmente un dominio independiente.

Algunas preguntas importantes:

## ¿Tiene reglas de negocio propias?

Si la funcionalidad posee sus propias validaciones, estados o comportamientos, probablemente representa un dominio.

Ejemplo:

Inventario posee reglas propias:

- stock disponible.
- reservas.
- actualización de cantidades.
- movimientos.

Por lo tanto, representa un módulo independiente.

---

## ¿Tiene ciclo de vida propio?

Una entidad que evoluciona independientemente suele justificar un módulo.

Ejemplo:

Una empresa puede:

- crearse.
- actualizarse.
- desactivarse.

Posee un ciclo propio.

---

## ¿Interactúa con otros dominios?

Un módulo puede colaborar con otros sin perder independencia.

Ejemplo:

Invoice utiliza Inventory:

```text
Invoice

↓

Inventory
```

Pero Inventory mantiene sus propias responsabilidades.

---

# Estructura esperada de un módulo

Un módulo debe seguir la organización establecida:

```text
module/

├── controller/

├── service/

├── repository/

├── model/

├── dto/

├── interfaces/

└── routes/
```

Cada capa mantiene una responsabilidad específica.

---

# Model

Representa las entidades utilizadas por el dominio.

Responsabilidades:

- definir estructuras.
- representar datos persistentes.
- expresar conceptos del dominio.

No debe contener lógica compleja de aplicación.

---

# Repository

Responsable exclusivamente de persistencia.

Debe encargarse de:

- consultas SQL.
- inserciones.
- actualizaciones.
- eliminación lógica.
- recuperación de información.

No debe decidir reglas del negocio.

Ejemplo:

Correcto:

```text
Buscar producto por ID
```

Incorrecto:

```text
Determinar si el producto puede venderse
```

Esa decisión pertenece al service.

---

# Service

Es la capa principal del módulo.

Responsabilidades:

- implementar casos de uso.
- validar reglas.
- coordinar repositories.
- controlar estados.

Ejemplo:

```text
CreateProduct()

UpdateInventory()

DisableCompany()
```

Los métodos deben representar acciones del negocio.

---

# Controller

El controller expone el módulo hacia HTTP.

Responsabilidades:

- recibir requests.
- obtener parámetros.
- validar estructura.
- llamar servicios.
- devolver respuestas.

No debe contener lógica del dominio.

---

# Interfaces

Cada módulo debe definir contratos claros.

Ejemplo:

```text
ProductRepository interface

ProductService interface
```

Esto permite:

- desacoplamiento.
- testing.
- reemplazo de implementaciones.

---

# Integración con Dependency Injection

Un nuevo módulo debe incorporarse al contenedor central de dependencias.

Flujo conceptual:

```text
Repository

↓

Service

↓

Controller

↓

Router
```

La creación de dependencias debe realizarse en un único punto de inicialización.

Esto evita:

- instancias duplicadas.
- dependencias ocultas.
- dificultad de mantenimiento.

---

# Integración con rutas

Las rutas deben mantenerse separadas por dominio.

Ejemplo:

```text
/api/products

/api/inventory

/api/companies
```

Evitar endpoints mezclados entre dominios.

---

# Integración con seguridad

Si el módulo requiere autenticación:

Debe pasar por:

```text
Request

↓

JWT Validation

↓

Tenant Context

↓

Permission Middleware

↓

Controller
```

El módulo no debe implementar nuevamente lógica de autenticación.

---

# Integración con Multi-Tenant

Todo módulo que almacene información asociada a empresas debe respetar aislamiento mediante tenant context.

Regla principal:

> Ninguna consulta debe depender de un company_id recibido directamente desde el cliente.

El flujo correcto es:

```text
JWT

↓

Tenant Context

↓

Service

↓

Repository

↓

Database filter
```

---

# Creación de permisos

Si el nuevo módulo requiere autorización específica:

Se deben agregar:

- permisos.
- asociaciones con roles.
- protección de endpoints.

Ejemplo conceptual:

```text
products.read

products.create

products.update
```

El módulo debe integrarse al sistema RBAC existente.

---

# ¿Cuándo utilizar un Workflow Service?

Si una funcionalidad involucra múltiples dominios, no se debe forzar dentro de un único módulo.

Ejemplo:

Un proceso de logística:

```text
Order

+

Inventory

+

Shipping

+

Payment
```

No pertenece completamente a ningún módulo.

Debe crearse un servicio de workflow encargado de orquestarlo.

---

# Ejemplo conceptual de nuevo módulo

Supongamos agregar:

```text
Supplier Management
```

El análisis sería:

## Dominio

¿Tiene identidad propia?

Sí.

---

## Entidades

```text
Supplier

SupplierContact

SupplierPaymentTerms
```

---

## Módulo:

```text
supplier/

├── controller

├── service

├── repository

├── model

└── dto
```

---

## Relaciones

Supplier puede colaborar con:

```text
Product

Purchase

Invoice
```

pero mantiene sus propias reglas.

---

# Checklist para agregar un módulo

Antes de finalizar un módulo nuevo verificar:

## Arquitectura

- [ ] Existe un dominio claramente definido.
- [ ] Las responsabilidades están separadas.
- [ ] No existe lógica duplicada.

---

## Backend

- [ ] Controller delgado.
- [ ] Service contiene reglas.
- [ ] Repository solo maneja persistencia.
- [ ] Interfaces definidas.
- [ ] Dependencias registradas.

---

## Seguridad

- [ ] Tenant Context utilizado correctamente.
- [ ] Permisos definidos.
- [ ] Endpoints protegidos.

---

## Calidad

- [ ] Errores utilizan el sistema centralizado.
- [ ] Responses mantienen formato estándar.
- [ ] Logs y observabilidad contemplados.

---

# Filosofía de extensión

Agregar un módulo en Nexora no significa agregar archivos aislados.

Significa incorporar un nuevo dominio respetando el lenguaje arquitectónico existente.

La regla principal es:

> Nuevos dominios deben integrarse al sistema, no modificar la identidad arquitectónica del sistema.

Esta filosofía permite que Nexora pueda crecer incorporando nuevas capacidades sin perder mantenibilidad, claridad ni escalabilidad.
