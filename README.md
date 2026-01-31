# GPI Backend - Core Rules

- Metrics are append-only (never overwritten)
- All metrics are time-bound
- Estimated data must include confidence score
- AI never generates raw facts
- Domain modules cannot access each other's repositories
- Data ingestion is async-first