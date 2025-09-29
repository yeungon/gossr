# 🚀 GOSSR A Golang Template Clean Architecture with DDD and Server-Side Rendering

A Go web application implementing **Clean Architecture** (modular approach) and **Domain-Driven Design (DDD) Lite** principles with **server-side rendering** capabilities.

---

## 📂 Project Structure

### Features
- 🏛️ Clean Architecture with Domain-Driven Design  
- 🎨 Server-side rendering  
- 🐘 PostgreSQL with **SQLC** for type-safe queries  
- 🧩 Modular design with independent business domains  
- 🌐 **Chi** router for HTTP routing  
- 🔄 Database migrations support  

---

## ⚙️ Prerequisites
- Go **1.24+**  
- PostgreSQL  
- `golang-migrate`  
- `sqlc`  

---

## 🔧 Configuration
Configuration is managed through `config.go` with environment variables:

- `HTTP_ADDR`: Server address (default: `:8080`)  
- `DB_URL`: PostgreSQL connection string  

---

## 🚀 Getting Started

1. Clone the repository  
2. Copy `.env_example` → `.env` and configure  
3. Initialize the database:  
   ```bash
   make up
   ```
4. Start the development server:  
   ```bash
   make dev
   ```

---

## 🗄️ Database Migrations
Manage your database using the commands in the `Makefile`:

```bash
make up     # Apply migrations
make down   # Rollback migrations
```

---

## 🧩 Domain Modules

### 📦 Item Module
- `domain/item.go`: Item domain model  
- `business/service.go`: Business logic  
- `storage/postgres.go`: Data persistence  

### 📦 Order Module
- `domain/order.go`: Order domain model  
- `business/service.go`: Order processing  
- `storage/postgres.go`: Order persistence  

---

## 📡 API Endpoints

### Items
- `POST /items` → Create item  
- `GET /items?id={id}` → Get item by ID  

### Orders
- `POST /orders` → Create order  

---

## 💻 Development
The application uses:
- `app.NewServer` → HTTP server setup  
- `app.NewRouter` → Routing configuration  
- **Chi middleware** for logging and recovery  

---

## 📁 Directory Structure

- `cmd` → Application entrypoints  
- `internal` → Private application code  
- `config` → Configuration management  
- `pkg` → Shared utilities  
- `html` → Template files for server-side rendering  

---

## 🏗️ Project Layout
The project follows **Clean Architecture** principles with clear separation of layers:

- **Domain Layer** → Domain models and interfaces  
- **Business Layer** → Use cases and business rules  
- **Infrastructure Layer** → External interfaces  
- **Transport Layer** → HTTP handlers  

Each module is **self-contained** with its own layers following **DDD principles**.

module/
 ├── business/
 ├── transport/
 ├── domain/
 ├── storage/
 └── sqlc/   (generated)
infras/      (shared infrastructure)


---


---

## 🏛️ Clean Architecture Diagram

```mermaid
graph TD
    A[👤 User] --> B[🌐 Transport Layer<br/>(HTTP Handlers)]
    B --> C[⚙️ Business Layer<br/>(Use Cases & Rules)]
    C --> D[📦 Domain Layer<br/>(Entities & Interfaces)]
    C --> E[🗄️ Infrastructure Layer<br/>(DB, External Services)]
    E --> D
```

## References
https://github.com/golang-standards/project-layout
https://evrone.com/blog/go-clean-template

---

## License
This project is licensed under the **MIT License**.


