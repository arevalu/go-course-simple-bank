# Registro de Cambios

Todos los cambios notables de este proyecto serán documentados en este archivo.

El formato está basado en [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

---

## 2025-11-01

### Agregado
- **Pruebas unitarias para operaciones CRUD** (e5e4084)
  - Suite de tests completa para operaciones CRUD de cuentas en `db/sqlc/account_test.go`:
    - `TestCreateAccount`: verifica creación de cuentas con datos válidos
    - `TestGetAccount`: valida obtención de cuenta por ID
    - `TestUpdateAccount`: prueba actualización de balance
    - `TestListAccounts`: verifica paginación con LIMIT/OFFSET
    - `TestDeleteAccount`: comprueba eliminación de cuentas
  - Archivo `main_test.go` con configuración de test suite:
    - Setup de conexión a PostgreSQL para tests
    - Variable global `testQueries` para reutilizar en todos los tests
    - Constantes de configuración (dbDriver, dbSource)
  - Paquete de utilidades `util/random.go` para generación de datos de prueba:
    - `RandomOwner()`: genera nombres aleatorios de 6 caracteres
    - `RandomMoney()`: genera balances entre 0 y 1000
    - `RandomCurrency()`: selecciona aleatoriamente entre USD, EUR, ARS
    - `RandomInt()` y `RandomString()`: funciones auxiliares
  - Dependencias agregadas a `go.mod`:
    - `github.com/stretchr/testify v1.11.1`: framework de assertions
    - `github.com/lib/pq v1.10.9`: driver PostgreSQL
  - Comando `make test` agregado al Makefile para ejecutar tests con cobertura

- **Generación de código CRUD con SQLC** (1b3de07)
  - Configuración de SQLC para generación automática de código Go desde queries SQL
  - Archivo `sqlc.yaml` con configuración para PostgreSQL apuntando a `db/migration` (schema) y `db/query` (queries)
  - Queries SQL definidas para operaciones CRUD de la tabla `accounts`:
    - `CreateAccount`: inserta nueva cuenta con owner, balance y currency
    - `GetAccount`: obtiene cuenta por ID
    - `ListAccounts`: lista cuentas con paginación (LIMIT/OFFSET)
    - `UpdateAccount`: actualiza balance de cuenta por ID
    - `DeleteAccount`: elimina cuenta por ID
  - Código Go generado en `db/sqlc/`:
    - `account.sql.go`: implementación de las 5 operaciones CRUD tipadas
    - `db.go`: interfaz DBTX y struct Queries para ejecutar queries
    - `models.go`: modelo Go de la estructura Account con tags JSON
  - Comando `make sqlc` agregado al Makefile para regenerar código

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
