# Database Relationships — Nexora

## Introducción

Las relaciones entre entidades de Nexora fueron diseñadas buscando representar correctamente los vínculos del dominio de negocio y mantener la integridad de la información.

El modelo relacional no solamente define cómo se almacenan los datos, sino también cómo interactúan los diferentes módulos del sistema.

Las relaciones principales fueron construidas alrededor de:

- Seguridad e identidad.
- Multi-tenancy.
- Catálogo comercial.
- Facturación.
- Pagos.
- Inventario.

---

# Principios Relacionales

Durante el diseño del modelo se aplicaron los siguientes principios:

## Integridad referencial

Las relaciones importantes utilizan referencias explícitas entre entidades.

Esto permite evitar:

- Registros huérfanos.
- Datos inconsistentes.
- Relaciones inválidas.

---

## Dependencias claras

Cada relación representa una dependencia real del dominio.

Ejemplo:

```text
Invoice Item

depende de

Invoice
```

Mientras:

```text
Product

existe independientemente

de Invoice
```

---

## Separación de dominios

Aunque algunas entidades colaboran entre sí, cada una mantiene su propia responsabilidad.

Ejemplo:

```text
Invoice

colabora con

Inventory

pero Inventory mantiene sus propias reglas
```

---

# Relación User - Company

## Descripción

Un usuario puede pertenecer a una empresa dentro del modelo empresarial.

Esta relación permite construir el contexto multi-tenant.

Conceptualmente:

```text
Company

   |

   └── Users
```

---

## Responsabilidad

Esta relación permite determinar:

- A qué empresa pertenece un usuario.
- Qué recursos puede consultar.
- Qué datos puede modificar.

---

## Uso en arquitectura

El `company_id` obtenido desde autenticación forma parte del Tenant Context.

Ese contexto es utilizado durante toda la request.

---

# Relación User - Roles - Permissions

## Descripción

El sistema utiliza un modelo basado en permisos.

La relación conceptual es:

```text
User

↓

Role

↓

Permission
```

---

## Responsabilidad

Permite definir:

- Qué acciones puede ejecutar un usuario.
- Qué endpoints puede consumir.
- Qué operaciones están autorizadas.

---

## Diseño elegido

Los permisos representan capacidades concretas del sistema.

Esto permite mayor escalabilidad que un sistema basado únicamente en roles.

---

# Relación Company - Products

## Descripción

Una empresa posee productos propios dentro del sistema.

```text
Company

↓

Products
```

---

## Responsabilidad

Permite mantener aislamiento entre catálogos empresariales.

Ejemplo:

Una empresa no debe visualizar productos pertenecientes a otra.

---

## Validación

El aislamiento se garantiza mediante:

- Tenant Context.
- Filtros por company_id.
- Reglas en servicios.

---

# Relación Product - Category

## Descripción

Un producto pertenece a una categoría.

```text
Category

↓

Products
```

---

## Responsabilidad

Permite:

- Clasificación.
- Organización del catálogo.
- Filtrado.
- Navegación comercial.

---

# Relación Category - Product Type

## Descripción

Una categoría posee diferentes tipos de productos.

```text
Category

↓

Product Types
```

---

## Regla importante

Un tipo de producto solamente puede pertenecer a una categoría válida.

Ejemplo:

Correcto:

```text
Categoría:

Electrónica

Tipo:

Notebook
```

Incorrecto:

```text
Categoría:

Ropa

Tipo:

Notebook
```

---

# Relación Product - Inventory

## Descripción

El inventario representa la disponibilidad asociada a productos.

```text
Product

↓

Inventory
```

---

## Consideración arquitectónica

Aunque existe una relación directa, Inventory fue separado como dominio debido a que posee reglas propias.

Ejemplos:

- Reservas.
- Confirmaciones.
- Movimientos.

---

# Relación Client - Invoice

## Descripción

Un cliente puede participar en múltiples operaciones comerciales.

```text
Client

↓

Invoices
```

---

## Responsabilidad

Permite:

- Historial comercial.
- Seguimiento de compras.
- Asociación documental.

---

# Relación Company - Invoice

## Descripción

Dentro del modelo ERP, una invoice pertenece a una empresa.

```text
Company

↓

Invoices
```

---

## Importancia

Es una de las relaciones principales del modelo multi-tenant.

Todas las operaciones empresariales deben respetar esta asociación.

---

# Relación Invoice - Invoice Items

## Descripción

Una invoice posee múltiples items.

```text
Invoice

1

↓

N

Invoice Items
```

---

## Cardinalidad

Una invoice:

- Puede tener muchos items.

Un item:

- Pertenece a una única invoice.

---

## Dependencia

Invoice Item no posee significado fuera de Invoice.

Por eso pertenece al agregado principal.

---

# Relación Invoice - Payments

## Descripción

Una invoice puede registrar múltiples pagos.

```text
Invoice

1

↓

N

Payments
```

---

## Uso

Permite soportar:

- Pagos parciales.
- Historial financiero.
- Conciliación futura.

---

# Relación Product - Invoice Items

## Descripción

Cada item referencia un producto comercializado.

```text
Product

↓

Invoice Item
```

---

## Responsabilidad

Permite conocer:

- Qué producto fue vendido.
- Cantidad.
- Precio aplicado.

---

# Mapa General del Modelo

La estructura completa puede representarse:

```text
                    Company
                       |
        --------------------------------
        |              |               |
      Users        Products        Invoices
                       |               |
                  Inventory        ----------
                                     |      |
                                Items     Payments
                                     |
                                  Product


        Client ----------------> Invoice
```

---

# Relaciones por Dominio

## Seguridad

```text
User

↓

Role

↓

Permission
```

---

## Organización empresarial

```text
Company

↓

Users

↓

Business Resources
```

---

## Catálogo

```text
Category

↓

Product Type

↓

Product
```

---

## Comercial

```text
Client

↓

Invoice

↓

Invoice Items

↓

Payments
```

---

## Inventario

```text
Product

↓

Inventory

↓

Stock Operations
```

---

# Reglas de Integridad

Las principales garantías del modelo son:

## No existen invoices sin empresa cuando corresponde

El tenant debe estar definido para operaciones empresariales.

---

## No existen items sin invoice

Los items dependen del documento principal.

---

## No existen pagos sin invoice

Todo movimiento financiero debe estar asociado a una operación comercial.

---

## No existen productos fuera de su contexto permitido

El acceso depende del tenant y reglas del negocio.

---

# Evolución del Modelo

El diseño actual permite incorporar nuevos dominios.

Ejemplos:

## Proveedores

Nueva relación:

```text
Supplier

↓

Purchase Invoice
```

---

## Logística

Nueva relación:

```text
Invoice

↓

Shipment

↓

Delivery
```

---

## Contabilidad

Nueva relación:

```text
Invoice

↓

Accounting Entry
```

---

# Evaluación Arquitectónica

El modelo relacional de Nexora mantiene una separación clara entre entidades independientes y entidades dependientes.

Las relaciones fueron diseñadas para representar el dominio real evitando acoplamientos innecesarios.

La combinación de claves relacionales, aislamiento multi-tenant y reglas aplicadas desde la capa de servicios permite mantener consistencia mientras el sistema continúa creciendo.

Este diseño proporciona una base preparada para extender Nexora hacia nuevos dominios empresariales sin necesidad de modificar completamente la estructura existente.
genia
