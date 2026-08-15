# Email Campaign Dispatcher

A backend service for creating email campaigns, importing large recipient lists, and dispatching campaign emails using Go and PostgreSQL.

The project focuses on backend engineering, database design, authentication, concurrent processing, and performance optimization for large-scale recipient ingestion.

---

##  Features

- JWT-based authentication
- Campaign creation and management
- Recipient management
- CSV-based recipient import
- Streaming CSV processing
- Batched PostgreSQL inserts
- PostgreSQL persistence
- Campaign-recipient relationships
- Concurrent email processing using goroutines and workers
- Mailpit integration for local email testing
- Docker Compose development environment
- REST API architecture

---

##  System Architecture

```text
                         ┌──────────────────────┐
                         │       Client         │
                         │   curl / Postman     │
                         └──────────┬───────────┘
                                    │
                                    │ HTTP
                                    ▼
                         ┌──────────────────────┐
                         │      Go API          │
                         │                      │
                         │  HTTP Handlers       │
                         │  JWT Authentication  │
                         │  Campaign Logic      │
                         └───────┬───────┬──────┘
                                 │       │
                    ┌────────────┘       └─────────────┐
                    │                                  │
                    ▼                                  ▼
          ┌──────────────────┐                ┌──────────────────┐
          │   PostgreSQL     │                │  Mail Dispatcher │
          │                  │                │                  │
          │ users            │                │ Goroutines       │
          │ campaigns        │                │ Worker-based     │
          │ recipients       │                │ email processing │
          │ campaign_        │                └────────┬─────────┘
          │ recipients       │                         │
          └──────────────────┘                         ▼
                                                      
                                             ┌──────────────────┐
                                             │     Mailpit      │
                                             │ Local SMTP/Test  │
                                             │ Email Interface  │
                                             └──────────────────┘
