# Feedback: plantilla de base de datos para microservicios Go en Kubernetes

Tu propuesta va en una dirección **muy buena** para un template base: explícita, simple y con bajo acoplamiento.

## Lo más sólido de la propuesta

1. **Postgres + Goose + SQLC**
   - Es una combinación madura y fácil de operar.
   - Evita magia de ORM y mejora la trazabilidad de cambios.

2. **SQL por dominio (`internal/<domain>/queries`)**
   - Favorece ownership por bounded context.
   - Hace más fácil refactorizar sin romper un paquete global gigante.

3. **Transacciones visibles**
   - Que el `tx` esté explícito en handler/service reduce sorpresas.
   - Funciona bien cuando quieres garantizar atomicidad en varios writes.

4. **Tests con transacción real y rollback**
   - Excelente balance entre realismo y aislamiento.
   - Menos mocks, más confianza en comportamiento real.

## Ajustes recomendados para producción en Kubernetes

1. **No abrir una transacción por request “por default”**
   - Úsala solo cuando realmente haya múltiples operaciones atómicas.
   - Para lecturas simples o inserts unitarios, usa pool directo para reducir lock time.

2. **Time-outs en todas las operaciones DB**
   - Usa `context.WithTimeout` en handlers/use-cases.
   - Define `statement_timeout` y `idle_in_transaction_session_timeout` a nivel DB.

3. **Migraciones como Job/Init step en deploy**
   - En Kubernetes, ejecutar Goose en un Job dedicado antes de promover tráfico.
   - Evita correr migraciones desde cada réplica de la app.

4. **Observabilidad por query**
   - Logging estructurado (duración, filas, error).
   - Métricas (p95 latency, errores por query, saturación de pool).
   - Trazas OpenTelemetry con span por operación SQL.

5. **Configurar pool explícitamente**
   - `max_conns`, `min_conns`, `max_conn_lifetime`, `health_check_period`.
   - Ajustarlo según CPU/memoria del pod y límites del Postgres.

6. **Versionado de SQLC por dominio**
   - Un `sqlc.yaml` por dominio o una config central multi-package, pero con límites claros.
   - Mantener nombres consistentes para evitar colisiones (`InsertAuditEvent`, etc.).

## Sugerencia de estructura final (simple, escalable)

- `internal/platform/db/`
  - pool, tx helpers, migraciones, test helpers comunes.
- `internal/<domain>/queries/sql/*.sql`
  - SQL fuente por dominio.
- `internal/<domain>/queries/*.go`
  - generado por SQLC, pegado al dominio.
- `cmd/migrations/main.go`
  - binario dedicado para Goose.

## Conclusión

Como **template base**, está muy bien orientado a simplicidad real:
- explícito,
- entendible por cualquier dev Go,
- fácil de operar en Kubernetes,
- y sin sobreingeniería temprana.

Si quieres, siguiente paso: convertir esto en un **RFC corto (1-2 páginas)** con decisiones, trade-offs y checklist de adopción para que el equipo lo apruebe rápido.
