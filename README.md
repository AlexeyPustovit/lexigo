# LexiGo

LexiGo is a self‑educational sandbox project aimed at mastering a modern production‑grade Go backend stack while building an English‑learning service. It is **not** intended to be a polished end‑user product; the primary goal is to explore tools and best practices commonly used in production services.

## Key Technologies

| Category       | Stack                           |
| -------------- | ------------------------------- |
| Web/API        | Gin, gRPC, WebSockets, Swagger  |
| Data Layer     | PostgreSQL (sqlx), Redis, Kafka |
| Infrastructure | Docker, Kubernetes (Helm)       |
| Observability  | Prometheus, Grafana, Loki       |
| Testing        | Testify, Mockery                |

## Service Landscape

| # | Service                      | Responsibility (single sentence)                                                      |
| - | ---------------------------- | ------------------------------------------------------------------------------------- |
| 1 | **API Gateway**              | Gin REST/WS + gRPC entry point; auth middleware; Redis-based rate limiting            |
| 2 | **Auth & User Profile**      | Registration and JWT authentication; session cache in Redis                           |
| 3 | **Vocabulary**               | Word CRUD and search                                                                  |
| 4 | **Learning Engine**          | SM-2 spaced-repetition algorithm, WebSocket study sessions; publishes events to Kafka |
| 5 | **Scheduler & Notification** | gocron jobs, SMTP e-mail dispatch; Kafka consumer; Telegram CD alerts                 |
| 6 | **Observability Stack**      | Prometheus exporters, Grafana dashboards, Loki log aggregation                        |
