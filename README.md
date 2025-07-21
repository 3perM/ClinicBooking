# 🏥 ClinicBooking

**ClinicBooking** is a full-stack, event-driven appointment booking system built for clinics and healthcare providers.

It features:

- A user-friendly **frontend and admin panel using Laravel (Blade)**
- A robust **Go backend** powered by **CQRS + Event Sourcing**
- **gRPC** for high-performance communication between PHP and Go services
- PostgreSQL as both **event store** and **read model database**
- Docker-based development environment for easy setup

---

## 🔧 Tech Stack

| Layer         | Technology                         |
|--------------|------------------------------------|
| Frontend      | Laravel 11 + Blade                 |
| Backend Logic | Go (v1.22+)                        |
| API           | gRPC with Protocol Buffers         |
| Database      | PostgreSQL                         |
| Architecture  | CQRS + Event Sourcing              |
| DevOps        | Docker, Docker Compose, NGINX      |

---

## 📐 Architecture Overview

```plaintext
User (Browser)
     ↓ Blade Form
Laravel Controller (PHP)
     ↓
gRPC Client (PHP)
     ↓
Go gRPC Server
     ↓
Command Handling → Event Store (Postgres)
     ↓
Read Model DB
     ↑
Laravel reads for display
