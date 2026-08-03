# Production Checklist

## Introducción

Este documento contiene la lista final de verificaciones necesarias antes de desplegar Nexora en un entorno productivo.

El objetivo es asegurar que la aplicación mantenga estabilidad, seguridad y consistencia antes de ser utilizada por usuarios reales.

La checklist está organizada por áreas:

- Backend.
- Frontend.
- Base de datos.
- Seguridad.
- Infraestructura.
- Documentación.
- Operación.

---

# Backend

## Configuración

- [ ] Variables de entorno productivas configuradas.
- [ ] Puerto del servidor correctamente definido.
- [ ] Configuración de base de datos validada.
- [ ] Configuración de almacenamiento revisada.
- [ ] Modo producción habilitado cuando corresponda.

---

## API

- [ ] Todos los endpoints principales funcionan correctamente.
- [ ] Respuestas HTTP utilizan códigos adecuados.
- [ ] Errores están controlados.
- [ ] Validaciones de entrada implementadas.
- [ ] Permisos revisados por endpoint.

---

## Seguridad

- [ ] JWT funcionando correctamente.
- [ ] Secretos fuera del repositorio.
- [ ] CORS configurado correctamente.
- [ ] Middleware de autenticación activo.
- [ ] Middleware de permisos activo.
- [ ] Información sensible no expuesta.

---

## Multi-Tenant

- [ ] Todas las operaciones validan tenant.
- [ ] Los usuarios solo acceden a información permitida.
- [ ] Company context correctamente cargado.
- [ ] No existen consultas sin filtro de contexto cuando corresponde.

---

# Base de Datos

## Configuración

- [ ] Migraciones aplicadas correctamente.
- [ ] Estructura de tablas actualizada.
- [ ] Índices revisados.
- [ ] Relaciones verificadas.
- [ ] Datos iniciales cargados cuando corresponda.

---

## Integridad

- [ ] Constraints funcionando.
- [ ] Relaciones consistentes.
- [ ] Soft delete funcionando donde aplica.
- [ ] Auditoría funcionando donde corresponde.

---

# Frontend

## Build

- [ ] Build productivo generado correctamente.
- [ ] Variables `VITE_*` configuradas.
- [ ] No existen errores de compilación.
- [ ] No existen errores críticos en consola.

---

## Aplicación

- [ ] Login funcionando.
- [ ] Navegación protegida funcionando.
- [ ] Permisos visuales funcionando.
- [ ] Manejo de errores validado.
- [ ] Estados de carga funcionando.

---

# Docker

## Backend Container

- [ ] Dockerfile actualizado.
- [ ] Imagen construye correctamente.
- [ ] Container inicia correctamente.
- [ ] Variables de entorno cargadas.
- [ ] Volúmenes configurados correctamente.

---

## Docker Compose

- [ ] Servicios definidos correctamente.
- [ ] Puertos configurados.
- [ ] Reinicio automático configurado.
- [ ] Configuración compatible con producción.

---

# Storage

## Archivos

- [ ] Directorios de almacenamiento existen.
- [ ] Permisos correctos.
- [ ] Archivos subidos funcionan.
- [ ] No se almacenan secretos accidentalmente.

---

# Deployment

## Backend

- [ ] Servicio desplegado correctamente.
- [ ] Health check funcionando.
- [ ] Logs accesibles.
- [ ] Variables productivas cargadas.

---

## Frontend

- [ ] Aplicación desplegada.
- [ ] URL correcta configurada.
- [ ] Comunicación con backend validada.
- [ ] Assets cargando correctamente.

---

# Monitoring

## Logs

- [ ] Logs importantes disponibles.
- [ ] Errores identificables.
- [ ] Información sensible excluida.

---

## Observabilidad

- [ ] Errores pueden rastrearse.
- [ ] Problemas críticos generan información suficiente.
- [ ] Estado del sistema puede verificarse.

---

# Backup y Recuperación

- [ ] Base de datos con backups configurados.
- [ ] Información crítica protegida.
- [ ] Procedimiento de recuperación definido.

---

# Documentación

- [ ] README actualizado.
- [ ] Documentación técnica completa.
- [ ] Variables documentadas.
- [ ] Diagramas actualizados.
- [ ] Instrucciones de instalación verificadas.

---

# Revisión Final

Antes de considerar el sistema listo:

- [ ] Usuario puede registrarse o iniciar sesión.
- [ ] Flujos principales funcionan correctamente.
- [ ] ERP operativo.
- [ ] Store operativo.
- [ ] Facturación funcionando.
- [ ] Inventario consistente.
- [ ] Permisos funcionando.
- [ ] Deploy estable.

---

# Resultado Esperado

Una vez completada esta checklist, Nexora debe encontrarse en condiciones de operar como una aplicación productiva, con:

- arquitectura documentada;
- configuración validada;
- seguridad revisada;
- despliegue preparado;
- procesos principales funcionando.
