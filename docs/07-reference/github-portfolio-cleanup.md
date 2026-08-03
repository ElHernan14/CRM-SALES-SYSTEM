# GitHub & Portfolio Cleanup Guide

## Introducción

Este documento contiene la lista final de tareas necesarias para preparar Nexora como proyecto de portfolio profesional.

El objetivo no es modificar la arquitectura del sistema, sino asegurar que el repositorio comunique correctamente la calidad técnica del proyecto.

Una buena presentación permite que una persona externa pueda entender:

- qué problema resuelve el sistema;
- cómo está construido;
- qué decisiones técnicas fueron tomadas;
- cómo ejecutar el proyecto.

---

# Revisión General del Repositorio

Antes de publicar o presentar el proyecto, verificar:

- nombre del repositorio;
- descripción;
- README principal;
- estructura de carpetas;
- documentación;
- archivos innecesarios.

---

# Nombre del Repositorio

El nombre debe ser claro y profesional.

Ejemplo:

```text
nexora-platform
```

Evitar nombres temporales como:

```text
test-project
backend-final-v2
new-version-final
```

---

# Descripción del Repositorio

La descripción debe explicar rápidamente el propósito.

Ejemplo:

```text
Multi-tenant SaaS ERP + Commerce platform built with Go, Vue 3 and PostgreSQL.
```

Debe incluir:

- tipo de sistema;
- arquitectura principal;
- tecnologías relevantes.

---

# README

El README debe ser la puerta de entrada principal.

Debe contener:

- descripción del proyecto;
- características principales;
- arquitectura;
- stack;
- ejecución local;
- documentación.

Evitar:

- README vacío;
- exceso de detalles internos;
- instrucciones desactualizadas.

---

# Limpieza de Código

Antes de mostrar el proyecto revisar:

## Código comentado

Eliminar:

- código muerto;
- pruebas antiguas;
- comentarios temporales.

Ejemplo:

```go
// TODO: arreglar esto después
```

si ya no corresponde debe eliminarse.

---

## Logs Temporales

Eliminar:

```text
fmt.Println()
console.log()
debug prints
```

salvo aquellos necesarios para observabilidad.

---

## Archivos Temporales

Eliminar:

```text
tmp/

coverage/

dist/

build/

logs/
```

si no forman parte del proyecto.

---

# Variables de Entorno

Nunca subir:

```text
.env
.env.production
.env.local
```

El repositorio debe contener únicamente ejemplos:

```text
.env.example
```

sin valores sensibles.

---

# Información Sensible

Verificar que no existan:

- passwords;
- tokens;
- API keys;
- secretos;
- claves privadas.

Especial atención en:

```text
.env

config files

docker-compose

scripts
```

---

# Documentación

Verificar que todos los documentos:

- estén actualizados;
- tengan nombres consistentes;
- no tengan información contradictoria.

Estructura esperada:

```text
docs/

├── backend
├── frontend
├── database
├── deployment
├── development-guide
└── diagrams
```

---

# Diagramas

Los diagramas deben:

- abrir correctamente;
- utilizar sintaxis Mermaid válida;
- representar la arquitectura actual.

Revisar especialmente:

- nombres de módulos;
- nombres de tecnologías;
- flujos de autenticación;
- relaciones entre dominios.

---

# Backend Review

Antes de publicar:

Verificar:

- arquitectura limpia;
- endpoints documentados;
- errores controlados;
- permisos funcionando;
- middleware correcto.

---

# Frontend Review

Verificar:

- rutas funcionando;
- variables de entorno correctas;
- ausencia de errores de consola;
- componentes sin código muerto.

---

# Git History

El historial debe transmitir orden.

Evitar:

```text
fix
fix2
final
final-final
working now
```

Preferir mensajes descriptivos:

```text
Add invoice lifecycle validation

Implement tenant middleware

Create product management module
```

---

# Branches

Antes de presentar:

Eliminar ramas innecesarias.

Mantener únicamente:

```text
main
```

o las ramas necesarias para desarrollo activo.

---

# Screenshots

Para portfolio es recomendable incluir capturas:

Ejemplos:

- Login.
- Dashboard.
- Gestión ERP.
- Productos.
- Facturación.
- Store.

Las imágenes deben mostrar:

- interfaz limpia;
- datos coherentes;
- estados reales.

---

# Deploy

Verificar:

- frontend accesible;
- backend funcionando;
- variables productivas configuradas;
- conexión con base de datos correcta.

---

# Presentación Profesional

El proyecto debe poder explicarse en pocos minutos:

## Problema

¿Qué resuelve Nexora?

## Arquitectura

¿Cómo está construido?

## Decisiones

¿Por qué se eligió esta arquitectura?

## Complejidad

¿Qué desafíos técnicos resolvió?

---

# Checklist Final

Antes de compartir:

- [ ] README actualizado.
- [ ] Documentación completa.
- [ ] Sin secretos en Git.
- [ ] Sin archivos temporales.
- [ ] Sin código muerto.
- [ ] Variables documentadas.
- [ ] Deploy funcionando.
- [ ] Screenshots preparados.
- [ ] Repositorio limpio.
- [ ] Historial de commits ordenado.

---

# Resultado Esperado

Luego de completar esta lista, Nexora queda preparado como proyecto profesional de portfolio, mostrando no solamente código funcional sino también:

- arquitectura;
- documentación;
- decisiones técnicas;
- buenas prácticas de desarrollo.
