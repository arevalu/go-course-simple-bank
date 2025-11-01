# Registro de Cambios

Todos los cambios notables de este proyecto serán documentados en este archivo.

El formato está basado en [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

---

## 2025-11-01

### Agregado
- **Sistema de transacciones CRUD para entries y transfers** (255ca07)
  Implementación completa del sistema de transacciones bancarias con generación automática de código CRUD vía SQLC para las entidades `entries` y `transfers`. Incluye queries SQL para creación, consulta y listado de movimientos contables y transferencias, junto con el patrón Store que encapsula operaciones transaccionales mediante `TransferTx`, garantizando consistencia ACID en transferencias entre cuentas con creación automática de entries de débito y crédito.

- **Pruebas unitarias para operaciones CRUD** (a2a8490)
  Suite completa de tests unitarios para validar operaciones CRUD de cuentas bancarias, implementando helpers de generación de datos aleatorios para testing y configuración de ambiente de pruebas con PostgreSQL. Integración de dependencias `testify` para assertions y `lib/pq` como driver de base de datos, con comando `make test` para ejecución automatizada de tests con cobertura.

- **Generación de código CRUD con SQLC** (1b3de07)
  Integración de SQLC como generador automático de código Go type-safe a partir de queries SQL, implementando el patrón de acceso a datos para la entidad `accounts` con operaciones CRUD completas (Create, Get, List con paginación, Update, Delete). Generación de modelos Go con serialización JSON y abstracción de interfaz `DBTX` para soportar tanto conexiones directas como transacciones.

---

## 2025-10-24

### Agregado
- **Migraciones iniciales de base de datos** (724e26f)
  - Creación del esquema de base de datos con tres tablas principales: `accounts`, `entries` y `transfers`
  - Tabla `accounts`: almacena cuentas bancarias con owner, balance, currency y timestamps
  - Tabla `entries`: registra movimientos de cuentas (débitos/créditos)
  - Tabla `transfers`: registra transferencias entre cuentas
  - Índices agregados para optimizar consultas en owner, account_id y relaciones de transferencias
  - Relaciones de claves foráneas establecidas entre las tablas
  - Constraints y comentarios para validación de datos (amounts positivos/negativos)

- **Inicialización del proyecto** (290e45b)
  - Configuración de la estructura del módulo Go
  - Inicialización del repositorio git
  - Configuración de Docker Compose para base de datos PostgreSQL
  - Creación de Makefile para tareas comunes de desarrollo

---

## Leyenda

- **Agregado**: Nuevas características o funcionalidades
- **Cambiado**: Cambios a funcionalidades existentes
- **Deprecado**: Características que serán removidas en versiones futuras
- **Removido**: Características removidas
- **Corregido**: Correcciones de bugs
- **Seguridad**: Mejoras o correcciones de seguridad
