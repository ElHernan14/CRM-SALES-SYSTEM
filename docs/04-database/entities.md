# Database Entities — Nexora

## Introducción

Las entidades de Nexora representan los conceptos principales del dominio del negocio y forman la base persistente sobre la cual se construyen los diferentes módulos de la aplicación.

Cada entidad fue diseñada teniendo en cuenta:

- Responsabilidad dentro del dominio.
- Ciclo de vida.
- Relaciones con otras entidades.
- Reglas de negocio asociadas.
- Aislamiento multi-tenant.

El objetivo no fue crear tablas aisladas, sino construir un modelo coherente donde cada entidad tenga una función definida dentro del sistema.

---

# User

## Descripción

La entidad `User` representa la identidad autenticable dentro de Nexora.

Es la base del sistema de seguridad y acceso, ya que a partir de ella se construye el contexto de sesión utilizado durante cada request.

---

## Responsabilidades

User maneja:

- Identidad del usuario.
- Credenciales de acceso.
- Asociación con empresa.
- Asociación con cliente.
- Roles.
- Permisos.

---

## Relaciones principales

```text
User

├── Roles

├── Permissions

├── Company

└── Client
```

---

## Rol dentro de arquitectura

User pertenece principalmente al dominio de seguridad.

No contiene reglas comerciales del negocio.

Su responsabilidad es identificar quién realiza una operación y qué capacidades posee.

---

# Company

## Descripción

La entidad `Company` representa una organización dentro del modelo multi-tenant.

Cada empresa funciona como un espacio aislado dentro del sistema.

---

## Responsabilidades

Company define:

- Identidad empresarial.
- Tenant principal.
- Agrupación de recursos.
- Contexto organizacional.

---

## Recursos asociados

Una compañía puede poseer:

- Usuarios.
- Productos.
- Clientes.
- Facturas.
- Compras.
- Inventario.

---

## Rol dentro del sistema

Company es una de las entidades principales del aislamiento multi-tenant.

La mayoría de operaciones empresariales utilizan `company_id` como filtro principal.

---

# Client

## Descripción

La entidad `Client` representa una relación comercial con un comprador.

Puede participar tanto dentro del ecosistema empresarial como dentro de flujos Store.

---

## Responsabilidades

Client representa:

- Compradores.
- Relaciones comerciales.
- Datos necesarios para facturación.

---

## Relaciones principales

```text
Client

↓

Invoices
```

---

## Consideración arquitectónica

Client posee ciclo de vida propio.

Aunque participa en procesos de facturación, no depende completamente de Invoice.

Puede existir antes de generar operaciones comerciales.

---

# Product

## Descripción

La entidad `Product` representa los productos o servicios ofrecidos dentro del sistema.

Es una entidad central porque participa en múltiples dominios:

- Ventas.
- Compras.
- Inventario.
- Store.
- Marketplace.

---

## Responsabilidades

Product contiene:

- Información comercial.
- Precio.
- Tipo de producto.
- Categoría.
- Identificación dentro del catálogo.

---

## Relaciones principales

```text
Product

├── Category

├── Product Type

├── Inventory

└── Invoice Items
```

---

## Consideración arquitectónica

Product posee ciclo de vida propio.

Aunque interactúa con inventario e invoices, sus reglas principales pertenecen al dominio producto.

---

# Category

## Descripción

Representa una clasificación lógica de productos.

Permite organizar el catálogo y facilitar búsquedas y agrupaciones.

---

## Responsabilidades

- Organización de productos.
- Clasificación comercial.
- Agrupación de tipos.

---

# Product Type

## Descripción

Representa una especialización dentro de una categoría de producto.

Permite diferenciar productos según su naturaleza.

Ejemplo:

```text
Categoría:

Electrónica

Tipos:

Notebook

Monitor

Accesorio
```

---

## Validaciones importantes

Un tipo de producto debe pertenecer a la categoría correspondiente.

La relación debe mantenerse consistente:

```text
Product Type

belongs to

Category
```

---

# Invoice

## Descripción

`Invoice` es la entidad central del dominio comercial.

Representa un documento comercial que puede originarse desde distintos flujos:

- ERP.
- Store.
- Marketplace.

---

## Responsabilidades

Invoice controla:

- Estado del documento.
- Cliente asociado.
- Empresa relacionada.
- Totales.
- Flujo comercial.

---

## Relaciones principales

```text
Invoice

├── Client

├── Company

├── Invoice Items

└── Payments
```

---

## Importancia dentro del dominio

Invoice funciona como Aggregate Root principal.

Su estado determina qué operaciones pueden realizarse sobre entidades relacionadas.

---

# Invoice Item

## Descripción

Representa cada línea dentro de una invoice.

No posee sentido independiente fuera del documento principal.

---

## Responsabilidades

Contiene:

- Producto asociado.
- Cantidad.
- Precio.
- Subtotal.

---

## Ciclo de vida

Invoice Item depende directamente del estado de Invoice.

Ejemplo:

- Una invoice cancelada no debería permitir modificaciones de items.
- Una invoice enviada mantiene sus items asociados.

---

# Invoice Payment

## Descripción

Representa movimientos de pago asociados a una invoice.

Permite registrar el historial financiero del documento.

---

## Responsabilidades

Maneja:

- Monto pagado.
- Fecha.
- Asociación con invoice.

---

## Relación

```text
Invoice

   |

   └── Payments
```

---

## Reglas principales

Los pagos deben respetar:

- Estado actual de invoice.
- Monto restante.
- Reglas del workflow de pago.

---

# Inventory

## Descripción

Inventory representa la gestión de disponibilidad de productos.

Aunque inicialmente deriva del producto, fue separado debido a que posee reglas propias.

---

## Responsabilidades

Maneja:

- Stock disponible.
- Reservas.
- Confirmación de movimientos.

---

## Relación con otros dominios

Inventory colabora principalmente con:

- Product.
- Invoice.
- Checkout.

---

## Consideración arquitectónica

Inventario representa un dominio independiente porque posee procesos propios y reglas específicas.

---

# Relaciones Generales del Modelo

La estructura conceptual puede resumirse:

```text
Company

├── User

├── Product

│      └── Inventory

│
├── Client

│
└── Invoice

       ├── Invoice Item

       └── Payment
```

---

# Entidades con Ciclo de Vida Propio

Poseen independencia:

- User.
- Company.
- Client.
- Product.
- Category.
- Product Type.
- Inventory.

---

# Entidades Dependientes

Dependen de otras entidades:

- Invoice Item → Invoice.
- Invoice Payment → Invoice.

---

# Consideraciones de Diseño

El modelo evita concentrar toda la lógica dentro de las entidades persistentes.

La lógica compleja permanece en la capa de servicios.

Las entidades representan:

- Estado.
- Información.
- Relaciones.

Mientras que los servicios representan:

- Procesos.
- Reglas.
- Workflows.

---

# Evaluación Arquitectónica

El modelo de entidades de Nexora refleja una separación clara entre conceptos del negocio.

Las entidades principales poseen responsabilidades definidas y relaciones explícitas, permitiendo que los distintos módulos evolucionen sin generar dependencia excesiva.

La decisión de utilizar Invoice como núcleo comercial, separar Inventory como dominio propio y mantener entidades independientes para seguridad, catálogo y relaciones comerciales permite construir una base preparada para futuras extensiones del sistema.
