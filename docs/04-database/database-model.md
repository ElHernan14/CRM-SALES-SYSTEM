# Database Model — Nexora

## Introducción

La base de datos de Nexora fue diseñada como una estructura relacional orientada al dominio del negocio, buscando mantener consistencia, integridad de datos y soporte para una arquitectura multi-tenant.

El modelo fue construido alrededor de las principales entidades del sistema, teniendo como núcleo el concepto de `Invoice`, sobre el cual se construyen gran parte de los procesos comerciales:

- Ventas.
- Compras.
- Pagos.
- Inventario.
- Store.
- Marketplace.
- Gestión empresarial.

La base de datos no fue pensada únicamente como almacenamiento de información, sino como una representación persistente del modelo de negocio.

---

# Principios de Diseño

Durante el diseño del modelo se priorizaron los siguientes principios:

## Separación por dominio

Cada entidad representa una responsabilidad concreta dentro del negocio.

Ejemplos:

- Usuarios y seguridad.
- Empresas.
- Clientes.
- Productos.
- Facturación.
- Inventario.
- Pagos.

Esto permite que cada módulo pueda evolucionar sin mezclar responsabilidades.

---

## Integridad relacional

La base utiliza relaciones explícitas entre entidades mediante claves foráneas.

Esto permite garantizar:

- Existencia de registros relacionados.
- Consistencia entre entidades.
- Integridad del modelo.

---

## Multi-Tenancy

El modelo fue preparado para soportar múltiples empresas utilizando aislamiento lógico mediante `company_id`.

Conceptualmente:

```text id="d1x0m4"
Company A

    ├── Users
    ├── Products
    ├── Invoices
    └── Clients


Company B

    ├── Users
    ├── Products
    ├── Invoices
    └── Clients
```

Los datos permanecen dentro de una misma base, pero separados mediante contexto de empresa.

---

# Modelo Central: Invoice

La entidad principal del sistema es `invoice`.

Esta tabla representa el documento comercial alrededor del cual se construyen múltiples procesos.

Una invoice puede representar:

- Venta.
- Compra.
- Operación interna.
- Flujo generado desde Store.
- Flujo generado desde Marketplace.

Su ciclo de vida controla múltiples procesos relacionados.

---

# Entidades Principales

## User

Representa los usuarios autenticados dentro del sistema.

Responsabilidades:

- Identidad.
- Seguridad.
- Acceso.
- Asociación con empresa o cliente.

Relacionada con:

- Roles.
- Permisos.
- Company.
- Client.

---

## Company

Representa una organización dentro del sistema empresarial.

Responsabilidades:

- Definir el tenant.
- Agrupar usuarios.
- Asociar recursos empresariales.

Una empresa posee información propia de:

- Productos.
- Clientes.
- Facturas.
- Inventario.

---

## Client

Representa clientes comerciales asociados al negocio.

Puede existir tanto dentro de un contexto empresarial como dentro del flujo Store.

Responsabilidades:

- Relaciones comerciales.
- Datos de comprador.
- Asociación con invoices.

---

## Product

Representa los productos o servicios disponibles.

Responsabilidades:

- Catálogo.
- Precio.
- Categoría.
- Tipo.
- Información comercial.

Se relaciona con:

- Inventario.
- Invoice items.
- Categorías.

---

## Invoice Item

Representa los elementos individuales dentro de una factura.

Su existencia depende de una invoice.

Relación:

```text id="9l4m4n"
Invoice

   |

   └── Invoice Items
```

Responsabilidades:

- Producto vendido.
- Cantidad.
- Precio.
- Subtotal.

---

## Invoice Payment

Representa los pagos realizados sobre una invoice.

Permite soportar:

- Pagos parciales.
- Pagos completos.
- Historial de pagos.

Relación:

```text id="g8f4fp"
Invoice

   |

   └── Payments
```

---

## Inventory

El inventario representa la gestión de disponibilidad y movimientos de stock.

Aunque inicialmente nace asociado al producto, fue separado como módulo propio debido a que representa reglas de negocio independientes.

Responsabilidades:

- Reserva de stock.
- Confirmación de stock.
- Actualización de cantidades.

---

# Relaciones Principales

El modelo general puede representarse:

```text id="7j4c9m"
Company

 ├── Users

 ├── Products

 ├── Clients

 └── Invoices

        ├── Invoice Items

        └── Payments
```

---

# Estados y Persistencia

Las entidades principales poseen estados que representan su ciclo de vida.

Especialmente:

## Invoice State Machine

Ejemplo conceptual:

```text id="x7s9kf"
Draft

↓

Pending

↓

Paid

↓

Completed
```

El estado de invoice condiciona operaciones posteriores.

Ejemplos:

- Una invoice pagada no debe modificarse.
- Una invoice pendiente puede continuar el flujo.
- Un pago debe respetar el estado actual.

---

# Consistencia de Datos

La consistencia se mantiene mediante diferentes mecanismos:

## Validaciones de aplicación

La capa service valida:

- Estados.
- Permisos.
- Reglas de negocio.
- Relaciones necesarias.

---

## Restricciones de base de datos

La base protege:

- Relaciones.
- Tipos.
- Integridad referencial.

---

## Transacciones

Las operaciones críticas utilizan transacciones para garantizar atomicidad.

Ejemplo:

Pago de una invoice:

```text id="q7a3de"
Actualizar pago

+

Crear payment

+

Actualizar estado invoice

+

Actualizar inventario

=

Una única operación atómica
```

Si una parte falla, toda la operación vuelve atrás.

---

# Decisiones Técnicas

## Uso de PostgreSQL

Se eligió PostgreSQL debido a:

- Robustez.
- Soporte transaccional.
- Integridad relacional.
- Consultas complejas.
- Escalabilidad.

---

## Sin ORM

El acceso fue implementado utilizando repositorios con consultas SQL explícitas.

Ventajas:

- Control completo sobre queries.
- Optimización manual.
- Transparencia del acceso a datos.

Desventajas:

- Mayor cantidad de código.
- Mayor esfuerzo de mantenimiento.
- Migraciones más costosas hacia ORM.

---

# Acceso desde la Arquitectura Backend

La base de datos no es accedida directamente desde controllers.

El flujo definido es:

```text id="6m5q88"
Controller

↓

Service

↓

Repository

↓

Database
```

Los repositorios son la única capa responsable de persistencia.

---

# Escalabilidad del Modelo

El diseño permite incorporar nuevas áreas de negocio.

Ejemplos:

## Logística

Podría agregarse sobre:

- Inventario.
- Productos.
- Invoices.

## Contabilidad

Podría extender:

- Pagos.
- Facturación.
- Movimientos financieros.

## Proveedores

Podría evolucionar desde:

- Clientes.
- Productos.
- Compras.

---

# Evaluación Arquitectónica

El modelo de base de datos de Nexora fue diseñado como una representación persistente del dominio, no simplemente como un conjunto de tablas aisladas.

La combinación de PostgreSQL, relaciones explícitas, aislamiento multi-tenant y una arquitectura basada en repositorios permite mantener consistencia y escalabilidad.

El modelo actual proporciona una base sólida para evolucionar desde un MVP hacia una plataforma empresarial, manteniendo separación de responsabilidades y control sobre los procesos críticos del negocio.
