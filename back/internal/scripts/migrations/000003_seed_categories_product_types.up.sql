-- NEXORA - SEED MAESTRO DE CATEGORÍAS Y TIPOS
-- PostgreSQL
-- Incluye category_company, category_product y product_type.
-- No incluye productos, empresas ni asignaciones a productos existentes.
-- Script idempotente: evita duplicados por nombre.

BEGIN;

-- ================================================================
-- 1. CATEGORÍAS DE EMPRESA
-- ================================================================
WITH seed(name, description) AS (
    VALUES
        ('Tecnología y Software', 'Empresas dedicadas al desarrollo de software, plataformas digitales y soluciones tecnológicas.'),
        ('Comercio Electrónico', 'Empresas dedicadas a la venta digital, marketplaces y soluciones para operaciones comerciales en línea.'),
        ('Fintech y Pagos', 'Empresas que ofrecen servicios financieros digitales, pagos, crédito y tecnología bancaria.'),
        ('Productividad y Colaboración', 'Empresas que desarrollan herramientas para organización, comunicación y trabajo colaborativo.'),
        ('CRM y Ventas', 'Empresas enfocadas en gestión de clientes, automatización comercial y procesos de ventas.'),
        ('Movilidad y Logística', 'Empresas dedicadas al transporte, distribución, entregas y gestión de operaciones logísticas.'),
        ('Turismo y Alojamiento', 'Empresas que ofrecen servicios de viajes, reservas, hospedaje y experiencias turísticas.'),
        ('Entretenimiento Digital', 'Empresas dedicadas a contenidos digitales, streaming, medios y experiencias de entretenimiento.'),
        ('Construcción e Infraestructura', 'Empresas dedicadas a la construcción, infraestructura, ingeniería, obras civiles y maquinaria pesada.'),
        ('Tecnología e Ingeniería', 'Empresas dedicadas al desarrollo tecnológico, ingeniería, innovación y soluciones industriales.')
)
INSERT INTO category_company (name, description)
SELECT s.name, s.description
FROM seed s
WHERE NOT EXISTS (
    SELECT 1 FROM category_company cc
    WHERE LOWER(TRIM(cc.name)) = LOWER(TRIM(s.name))
);

-- ================================================================
-- 2. CATEGORÍAS DE PRODUCTO
-- ================================================================
WITH seed(name, description) AS (
    VALUES
        ('Desarrollo de Software', 'Productos y servicios relacionados con desarrollo, consultoría, diseño e implementación de software.'),
        ('Cloud y DevOps', 'Soluciones de infraestructura, nube, despliegue, observabilidad y automatización técnica.'),
        ('IA y Automatización', 'Soluciones basadas en inteligencia artificial, automatización de procesos y análisis avanzado.'),
        ('Pagos y Fintech', 'Productos y servicios financieros digitales, pagos, facturación, crédito y prevención de fraude.'),
        ('Comercio Electrónico', 'Plataformas, herramientas y servicios para tiendas digitales, marketplaces y operaciones comerciales.'),
        ('Productividad y Colaboración', 'Herramientas para organización, documentación, comunicación y trabajo en equipo.'),
        ('CRM y Marketing', 'Soluciones para gestión de clientes, ventas, campañas, analítica y automatización comercial.'),
        ('Movilidad y Delivery', 'Servicios de transporte, entregas, distribución, movilidad urbana y operaciones de última milla.'),
        ('Turismo y Alojamiento', 'Servicios y plataformas de reservas, hospedaje, experiencias, viajes y gestión turística.'),
        ('Streaming y Entretenimiento', 'Plataformas y servicios de contenido audiovisual, transmisión digital y entretenimiento.'),
        ('Tecnología e Ingeniería', 'Productos tecnológicos, sistemas industriales, hardware avanzado y soluciones de ingeniería.'),
        ('Maquinaria Pesada', 'Equipos y maquinaria para construcción, minería, movimiento de suelo y obras de infraestructura.'),
        ('Servicios de Construcción', 'Servicios técnicos, operativos y profesionales para maquinaria, obras civiles e infraestructura.')
)
INSERT INTO category_product (name, description)
SELECT s.name, s.description
FROM seed s
WHERE NOT EXISTS (
    SELECT 1 FROM category_product cp
    WHERE LOWER(TRIM(cp.name)) = LOWER(TRIM(s.name))
);

-- ================================================================
-- 3. TIPOS DE PRODUCTO
-- ================================================================
WITH seed(category_name, name, description) AS (
    VALUES
        -- Desarrollo de Software
        ('Desarrollo de Software', 'Desarrollo Web', 'Aplicaciones y plataformas web modernas, escalables y orientadas al negocio.'),
        ('Desarrollo de Software', 'Desarrollo Mobile', 'Aplicaciones móviles nativas o multiplataforma para Android e iOS.'),
        ('Desarrollo de Software', 'Software Empresarial', 'Sistemas empresariales personalizados para procesos internos y operaciones críticas.'),
        ('Desarrollo de Software', 'Consultoría Tecnológica', 'Asesoramiento técnico, análisis de arquitectura y definición de soluciones digitales.'),
        ('Desarrollo de Software', 'Equipos Dedicados', 'Equipos de desarrollo asignados de forma exclusiva a proyectos y productos digitales.'),
        ('Desarrollo de Software', 'Diseño UX/UI', 'Diseño de interfaces, experiencias digitales y sistemas visuales centrados en el usuario.'),
        ('Desarrollo de Software', 'Plataforma Low-Code', 'Herramientas para crear aplicaciones y automatizaciones con menor desarrollo manual.'),
        ('Desarrollo de Software', 'API e Integraciones', 'Diseño de APIs e integración entre sistemas, plataformas y servicios externos.'),

        -- Cloud y DevOps
        ('Cloud y DevOps', 'Migración Cloud', 'Migración de aplicaciones, servicios y datos hacia infraestructuras de nube.'),
        ('Cloud y DevOps', 'Consultoría DevOps', 'Diseño y mejora de prácticas DevOps, automatización y entrega continua.'),
        ('Cloud y DevOps', 'Infraestructura Cloud', 'Configuración y administración de infraestructura escalable en la nube.'),
        ('Cloud y DevOps', 'CI/CD', 'Automatización de integración, pruebas, compilación y despliegue continuo.'),
        ('Cloud y DevOps', 'Contenedores', 'Implementación y administración de entornos basados en Docker y orquestadores.'),
        ('Cloud y DevOps', 'Observabilidad', 'Monitoreo, métricas, trazas y logs para sistemas y aplicaciones.'),
        ('Cloud y DevOps', 'Cloud Security', 'Protección de infraestructura, identidades, redes y datos alojados en la nube.'),
        ('Cloud y DevOps', 'Soporte de Infraestructura', 'Mantenimiento y asistencia técnica para infraestructura y servicios críticos.'),

        -- IA y Automatización
        ('IA y Automatización', 'Inteligencia Artificial', 'Modelos y soluciones inteligentes aplicados a procesos empresariales.'),
        ('IA y Automatización', 'Asistente Inteligente', 'Asistentes virtuales para atención, soporte y productividad.'),
        ('IA y Automatización', 'Automatización QA', 'Automatización de pruebas, validaciones y control de calidad de software.'),
        ('IA y Automatización', 'Workflow Automation', 'Automatización de flujos de trabajo y procesos operativos repetitivos.'),
        ('IA y Automatización', 'Machine Learning', 'Modelos predictivos y sistemas que aprenden a partir de datos.'),
        ('IA y Automatización', 'Procesamiento de Datos', 'Procesamiento, clasificación y enriquecimiento automatizado de información.'),
        ('IA y Automatización', 'Analítica Predictiva', 'Predicciones y recomendaciones basadas en datos históricos y patrones.'),
        ('IA y Automatización', 'Automatización Robótica', 'Automatización de tareas administrativas mediante bots y procesos programados.'),

        -- Pagos y Fintech
        ('Pagos y Fintech', 'Pasarela de Pagos', 'Procesamiento seguro de pagos digitales para comercios y plataformas.'),
        ('Pagos y Fintech', 'Facturación Recurrente', 'Gestión automatizada de suscripciones, cobros periódicos y renovaciones.'),
        ('Pagos y Fintech', 'Terminal de Pago', 'Dispositivos y soluciones para aceptar pagos de forma presencial.'),
        ('Pagos y Fintech', 'Crédito y Financiación', 'Productos de crédito, préstamos y financiación para personas o empresas.'),
        ('Pagos y Fintech', 'Tarjetas y Emisión', 'Emisión, administración y procesamiento de tarjetas físicas o virtuales.'),
        ('Pagos y Fintech', 'Gestión Fiscal', 'Herramientas para impuestos, comprobantes y cumplimiento fiscal.'),
        ('Pagos y Fintech', 'Prevención de Fraude', 'Detección y mitigación de operaciones sospechosas y riesgos financieros.'),
        ('Pagos y Fintech', 'Transferencias Digitales', 'Envío y recepción de dinero mediante cuentas y redes digitales.'),

        -- Comercio Electrónico
        ('Comercio Electrónico', 'Tienda Online', 'Plataforma para crear, administrar y operar comercios digitales.'),
        ('Comercio Electrónico', 'Marketplace', 'Plataforma que conecta múltiples vendedores con compradores.'),
        ('Comercio Electrónico', 'Gestión de Catálogo', 'Administración de productos, variantes, precios y disponibilidad.'),
        ('Comercio Electrónico', 'Checkout', 'Flujo optimizado para confirmar pedidos y completar pagos.'),
        ('Comercio Electrónico', 'Publicidad Digital', 'Herramientas de promoción, campañas y visibilidad de productos.'),
        ('Comercio Electrónico', 'Gestión de Pedidos', 'Seguimiento y administración del ciclo completo de pedidos.'),
        ('Comercio Electrónico', 'Fulfillment', 'Preparación, almacenamiento, empaquetado y despacho de pedidos.'),
        ('Comercio Electrónico', 'Integración Comercial', 'Conexión de tiendas con pagos, logística, ERP y canales externos.'),

        -- Productividad y Colaboración
        ('Productividad y Colaboración', 'Gestión de Proyectos', 'Planificación y seguimiento de tareas, objetivos y proyectos.'),
        ('Productividad y Colaboración', 'Documentación Colaborativa', 'Creación y edición compartida de documentos y conocimiento.'),
        ('Productividad y Colaboración', 'Mensajería Empresarial', 'Comunicación interna mediante canales, mensajes y grupos de trabajo.'),
        ('Productividad y Colaboración', 'Videoconferencia', 'Reuniones virtuales, llamadas y colaboración remota.'),
        ('Productividad y Colaboración', 'Gestión del Conocimiento', 'Organización y consulta centralizada de información empresarial.'),
        ('Productividad y Colaboración', 'Automatización de Tareas', 'Automatización de acciones repetitivas y flujos internos.'),
        ('Productividad y Colaboración', 'Calendario y Agenda', 'Coordinación de reuniones, eventos y disponibilidad de equipos.'),
        ('Productividad y Colaboración', 'Espacio de Trabajo Digital', 'Entorno unificado para documentos, tareas, comunicación y recursos.'),

        -- CRM y Marketing
        ('CRM y Marketing', 'CRM de Ventas', 'Gestión de clientes, oportunidades, contactos y procesos comerciales.'),
        ('CRM y Marketing', 'Automatización de Marketing', 'Automatización de campañas, segmentos y comunicaciones comerciales.'),
        ('CRM y Marketing', 'Atención al Cliente', 'Gestión de consultas, soporte, tickets y experiencia del cliente.'),
        ('CRM y Marketing', 'Analítica Comercial', 'Métricas, reportes y análisis para ventas y rendimiento comercial.'),
        ('CRM y Marketing', 'Email Marketing', 'Creación, envío y seguimiento de campañas de correo electrónico.'),
        ('CRM y Marketing', 'Gestión de Leads', 'Captura, clasificación y seguimiento de prospectos comerciales.'),
        ('CRM y Marketing', 'Publicidad Digital', 'Administración y optimización de campañas publicitarias digitales.'),
        ('CRM y Marketing', 'Customer Data Platform', 'Unificación y análisis de datos de clientes provenientes de múltiples canales.'),

        -- Movilidad y Delivery
        ('Movilidad y Delivery', 'Transporte Urbano', 'Servicios de movilidad de pasajeros dentro de ciudades y áreas metropolitanas.'),
        ('Movilidad y Delivery', 'Entrega de Última Milla', 'Distribución rápida de productos hasta el destino final.'),
        ('Movilidad y Delivery', 'Logística Empresarial', 'Gestión de transporte y distribución para organizaciones.'),
        ('Movilidad y Delivery', 'Gestión de Flotas', 'Administración, monitoreo y optimización de vehículos y conductores.'),
        ('Movilidad y Delivery', 'Micromovilidad', 'Soluciones de movilidad mediante bicicletas, monopatines y vehículos livianos.'),
        ('Movilidad y Delivery', 'Transporte Programado', 'Reservas y coordinación anticipada de viajes o traslados.'),
        ('Movilidad y Delivery', 'Distribución de Paquetes', 'Recolección, clasificación y entrega de paquetes y encomiendas.'),
        ('Movilidad y Delivery', 'Plataforma de Conductores', 'Herramientas para administrar conductores, recorridos y asignaciones.'),

        -- Turismo y Alojamiento
        ('Turismo y Alojamiento', 'Reserva de Alojamiento', 'Búsqueda y reserva de hoteles, departamentos y propiedades.'),
        ('Turismo y Alojamiento', 'Experiencias Turísticas', 'Actividades, excursiones y experiencias ofrecidas a viajeros.'),
        ('Turismo y Alojamiento', 'Gestión Hotelera', 'Administración de habitaciones, huéspedes, reservas y operaciones hoteleras.'),
        ('Turismo y Alojamiento', 'Alquiler Temporal', 'Alojamiento por períodos cortos en propiedades residenciales.'),
        ('Turismo y Alojamiento', 'Viajes Corporativos', 'Gestión de traslados, reservas y gastos de viajes empresariales.'),
        ('Turismo y Alojamiento', 'Paquetes Turísticos', 'Combinación de alojamiento, transporte y actividades de viaje.'),
        ('Turismo y Alojamiento', 'Gestión de Anfitriones', 'Herramientas para publicar, administrar y operar propiedades.'),
        ('Turismo y Alojamiento', 'Servicios para Huéspedes', 'Servicios complementarios para mejorar la experiencia del viajero.'),

        -- Streaming y Entretenimiento
        ('Streaming y Entretenimiento', 'Video bajo Demanda', 'Catálogo de películas, series y contenido audiovisual bajo demanda.'),
        ('Streaming y Entretenimiento', 'Streaming en Vivo', 'Transmisión de eventos y contenido audiovisual en tiempo real.'),
        ('Streaming y Entretenimiento', 'Suscripción Digital', 'Acceso recurrente a contenidos mediante planes de suscripción.'),
        ('Streaming y Entretenimiento', 'Contenido Infantil', 'Contenido audiovisual y experiencias orientadas a público infantil.'),
        ('Streaming y Entretenimiento', 'Producción Original', 'Desarrollo y distribución de películas, series y contenido exclusivo.'),
        ('Streaming y Entretenimiento', 'Recomendación de Contenido', 'Sistemas de personalización y sugerencias según preferencias del usuario.'),
        ('Streaming y Entretenimiento', 'Distribución Digital', 'Publicación y distribución de contenidos en múltiples plataformas.'),
        ('Streaming y Entretenimiento', 'Entretenimiento Interactivo', 'Experiencias digitales participativas, juegos y contenido interactivo.'),

        -- Tecnología e Ingeniería
        ('Tecnología e Ingeniería', 'Drones Industriales', 'Vehículos aéreos no tripulados para inspección, vigilancia y monitoreo.'),
        ('Tecnología e Ingeniería', 'Robótica Industrial', 'Robots y sistemas automatizados para procesos industriales y manufactura.'),
        ('Tecnología e Ingeniería', 'Infraestructura Inteligente', 'Soluciones para edificios, ciudades e instalaciones inteligentes.'),
        ('Tecnología e Ingeniería', 'Sistemas de Seguridad', 'Equipamiento y plataformas para vigilancia y protección avanzada.'),
        ('Tecnología e Ingeniería', 'Comunicaciones Satelitales', 'Equipos y plataformas para conectividad y comunicaciones satelitales.'),
        ('Tecnología e Ingeniería', 'Energía Inteligente', 'Tecnologías para generación, almacenamiento y gestión energética.'),
        ('Tecnología e Ingeniería', 'Automatización Industrial', 'Soluciones para automatizar procesos industriales y operativos.'),
        ('Tecnología e Ingeniería', 'Hardware Empresarial', 'Equipamiento tecnológico de alto rendimiento para organizaciones.'),
        ('Tecnología e Ingeniería', 'Inteligencia Artificial', 'Plataformas y soluciones basadas en inteligencia artificial.'),
        ('Tecnología e Ingeniería', 'Investigación y Desarrollo', 'Productos y tecnologías para innovación e investigación aplicada.'),

        -- Maquinaria Pesada
        ('Maquinaria Pesada', 'Excavadora', 'Maquinaria para excavación, movimiento de tierra y trabajos de gran escala.'),
        ('Maquinaria Pesada', 'Retroexcavadora', 'Equipo versátil para excavación, carga y tareas de obra.'),
        ('Maquinaria Pesada', 'Topadora', 'Maquinaria de gran potencia para empuje y nivelación de terrenos.'),
        ('Maquinaria Pesada', 'Minicargadora', 'Equipo compacto para carga, movimiento de materiales y tareas en espacios reducidos.'),
        ('Maquinaria Pesada', 'Motoniveladora', 'Maquinaria para nivelación, perfilado y mantenimiento de superficies.'),
        ('Maquinaria Pesada', 'Compactador', 'Equipo para compactación de suelo, asfalto y materiales de obra.'),
        ('Maquinaria Pesada', 'Grúa', 'Equipo para elevación y desplazamiento de cargas pesadas.'),
        ('Maquinaria Pesada', 'Cargadora Frontal', 'Maquinaria para carga, transporte corto y manipulación de materiales.'),

        -- Servicios de Construcción
        ('Servicios de Construcción', 'Alquiler de Maquinaria', 'Alquiler temporal de maquinaria y equipos para obras y proyectos.'),
        ('Servicios de Construcción', 'Mantenimiento Preventivo', 'Revisión programada para reducir fallas y prolongar la vida útil de los equipos.'),
        ('Servicios de Construcción', 'Reparación de Equipos', 'Diagnóstico y reparación de maquinaria, componentes y sistemas industriales.'),
        ('Servicios de Construcción', 'Capacitación de Operadores', 'Formación técnica para el uso seguro y eficiente de maquinaria.'),
        ('Servicios de Construcción', 'Ingeniería de Obra', 'Planificación técnica, diseño y soporte profesional para proyectos constructivos.'),
        ('Servicios de Construcción', 'Inspección Técnica', 'Evaluación del estado de maquinaria, estructuras e instalaciones.'),
        ('Servicios de Construcción', 'Gestión de Proyectos', 'Coordinación de alcance, recursos, tiempos y costos de obras.'),
        ('Servicios de Construcción', 'Soporte en Obra', 'Asistencia técnica y operativa durante la ejecución de proyectos.')
)
INSERT INTO product_type (category_id, name, description)
SELECT cp.id, s.name, s.description
FROM seed s
JOIN category_product cp
  ON LOWER(TRIM(cp.name)) = LOWER(TRIM(s.category_name))
WHERE NOT EXISTS (
    SELECT 1
    FROM product_type pt
    WHERE pt.category_id = cp.id
      AND LOWER(TRIM(pt.name)) = LOWER(TRIM(s.name))
);

COMMIT;

-- ================================================================
-- 4. VERIFICACIÓN
-- ================================================================
SELECT id, name, description
FROM category_company
ORDER BY name;

SELECT
    cp.id,
    cp.name,
    COUNT(pt.id) AS total_types
FROM category_product cp
LEFT JOIN product_type pt ON pt.category_id = cp.id
GROUP BY cp.id, cp.name
ORDER BY cp.name;

SELECT
    cp.name AS category,
    pt.id,
    pt.name AS product_type,
    pt.description
FROM product_type pt
JOIN category_product cp ON cp.id = pt.category_id
ORDER BY cp.name, pt.name;
