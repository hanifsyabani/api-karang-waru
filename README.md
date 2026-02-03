# 🏘️ API Karang Waru

A robust RESTful API for Karang Waru Village Management System built with Go, providing comprehensive endpoints for managing village data, residents, services, and administrative information.

![Go Version](https://img.shields.io/badge/Go-1.24-00ADD8?style=flat-square&logo=go)
![Gin Framework](https://img.shields.io/badge/Gin-1.10-00ADD8?style=flat-square)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-336791?style=flat-square&logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat-square&logo=docker)

---

## 📋 Table of Contents

- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Project Structure](#-project-structure)
- [Getting Started](#-getting-started)
- [Environment Variables](#-environment-variables)
- [API Endpoints](#-api-endpoints)
- [Authentication](#-authentication)
- [Docker Deployment](#-docker-deployment)
- [License](#-license)

---

## ✨ Features

- **🔐 Authentication** — JWT-based user authentication with secure login/logout
- **👥 User Management** — Complete CRUD operations for user accounts
- **🏘️ Village Profile** — Manage village information, demographics, history, and vision/mission
- **📰 News/Berita** — Village news and announcements management with slug support
- **🏪 UMKM** — Small and medium enterprise directory management
- **📄 Services/Layanan** — Village administrative services catalog
- **💰 APBD** — Village budget and financial information
- **👨‍👩‍👧‍👦 Residents/Penduduk** — Resident data management with demographics
- **🎓 Education** — Educational programs, institutions, statistics, achievements, and documentation
- **🏥 Health** — Health services and facilities management

---

## 🛠️ Tech Stack

| Technology | Description |
|------------|-------------|
| **Go 1.24** | Backend programming language |
| **Gin** | High-performance HTTP web framework |
| **GORM** | ORM library for Go |
| **PostgreSQL** | Primary database |
| **JWT** | JSON Web Token for authentication |
| **Docker** | Containerization platform |

---

## 📁 Project Structure

```
api-karang-waru/
├── config/             # Database and environment configuration
├── errors/             # Custom error definitions
├── handlers/           # HTTP request handlers
├── helpers/            # Utility helper functions
├── middlewares/        # Authentication and other middlewares
├── migrations/         # Database migration files
├── models/             # GORM data models
├── repositories/       # Database access layer
├── requests/           # Request validation structs
├── responses/          # Response formatting structs
├── services/           # Business logic layer
├── utils/              # General utilities
├── main.go             # Application entry point
├── Dockerfile          # Docker configuration
├── go.mod              # Go module dependencies
└── go.sum              # Dependency checksums
```

---

## 🚀 Getting Started

### Prerequisites

- Go 1.24 or higher
- PostgreSQL 15+
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/hanifsyabani/api-karang-waru.git
   cd api-karang-waru
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080` by default.

---


---

## 📡 API Endpoints

Base URL: `/api/karang-waru`

### Public Routes

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/` | API welcome message |
| `GET` | `/health` | Health check |
| `POST` | `/api/karang-waru/register` | User registration |
| `POST` | `/api/karang-waru/login` | User login |
| `POST` | `/api/karang-waru/logout` | User logout |

### Protected Routes (Requires Authentication)

#### 👤 Users
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/me` | Get current user profile |
| `GET` | `/users` | Get all users |
| `GET` | `/users/:id` | Get user by ID |
| `POST` | `/users` | Create new user |
| `PUT` | `/users/:id` | Update user |
| `DELETE` | `/users/:id` | Delete user |

#### 🏘️ Village Profile
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/village-profile` | Get village profile |
| `POST` | `/village-profile` | Create village profile |
| `PUT` | `/village-profile` | Update village profile |
| `DELETE` | `/village-profile` | Delete village profile |

#### 📊 Demographics
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/demographics` | Get demographics data |
| `POST` | `/demographics` | Create demographics |
| `PUT` | `/demographics` | Update demographics |
| `DELETE` | `/demographics` | Delete demographics |

#### 📜 History
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/history-village` | Get village history |
| `POST` | `/history-village` | Create history |
| `PUT` | `/history-village` | Update history |
| `DELETE` | `/history-village` | Delete history |

#### 🎯 Vision & Mission
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/visi-misi` | Get vision & mission |
| `POST` | `/visi-misi` | Create vision & mission |
| `PUT` | `/visi-misi` | Update vision & mission |
| `DELETE` | `/visi-misi` | Delete vision & mission |

#### 📰 News
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/news` | Get all news |
| `GET` | `/news/:id` | Get news by ID |
| `GET` | `/news/slug/:slug` | Get news by slug |
| `GET` | `/news/category/count` | Get news count by category |
| `POST` | `/news` | Create news article |
| `PUT` | `/news/:id` | Update news article |
| `DELETE` | `/news/:id` | Delete news article |

#### 🏪 UMKM
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/umkm` | Get all UMKM |
| `GET` | `/umkm/:id` | Get UMKM by ID |
| `GET` | `/umkm/slug/:slug` | Get UMKM by slug |
| `GET` | `/umkm/count-status` | Get UMKM status count |
| `POST` | `/umkm` | Create UMKM |
| `PUT` | `/umkm/:id` | Update UMKM |
| `DELETE` | `/umkm/:id` | Delete UMKM |

#### 📄 Services
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/services` | Get all services |
| `GET` | `/service/:id` | Get service by ID |
| `GET` | `/service/slug/:slug` | Get service by slug |
| `POST` | `/service` | Create service |
| `PUT` | `/service/:id` | Update service |
| `DELETE` | `/service/:id` | Delete service |

#### 💰 APBD (Village Budget)
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/apbd` | Get all APBD records |
| `GET` | `/apbd/:id` | Get APBD by ID |
| `POST` | `/apbd` | Create APBD record |
| `PUT` | `/apbd/:id` | Update APBD record |
| `DELETE` | `/apbd/:id` | Delete APBD record |

#### 👨‍👩‍👧‍👦 Residents
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/residents` | Get all residents |
| `GET` | `/resident/:id` | Get resident by ID |
| `GET` | `/resident/count` | Get resident count |
| `POST` | `/resident` | Create resident |
| `PUT` | `/resident/:id` | Update resident |
| `DELETE` | `/resident/:id` | Delete resident |

#### 🎓 Education
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET/POST/PUT/DELETE` | `/education/program` | Education programs |
| `GET/POST/PUT/DELETE` | `/education/institution` | Education institutions |
| `GET/POST/PUT/DELETE` | `/education/statistic` | Education statistics |
| `GET/POST/PUT/DELETE` | `/education/achievements` | Education achievements |
| `GET/POST/PUT/DELETE` | `/education/documentation` | Education documentation |

#### 🏥 Health
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET/POST/PUT/DELETE` | `/health/service` | Health services |
| `GET/POST/PUT/DELETE` | `/health/facility` | Health facilities |

---

## 🔐 Authentication

This API uses **JWT (JSON Web Token)** for authentication.

### Login
```bash
curl -X POST http://localhost:8080/api/karang-waru/login \
  -H "Content-Type: application/json" \
  -d '{"email": "admin@example.com", "password": "password123"}'
```

### Using the Token
Include the token in the `Authorization` header:
```bash
curl -X GET http://localhost:8080/api/karang-waru/me \
  -H "Authorization: Bearer <your-jwt-token>"
```

---

## 🐳 Docker Deployment

### Build the Image
```bash
docker build -t api-karang-waru .
```

### Run the Container
```bash
docker run -d \
  --name api-karang-waru \
  -p 8080:8080 \
  -e DATABASE_URL="postgres://user:password@host:5432/karang_waru" \
  -e APP_PORT="8080" \
  api-karang-waru
```

### Using Docker Compose (Optional)
Create a `docker-compose.yml`:
```yaml
version: '3.8'
services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DATABASE_URL=postgres://postgres:postgres@db:5432/karang_waru?sslmode=disable
      - APP_PORT=8080
    depends_on:
      - db

  db:
    image: postgres:15-alpine
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_DB=karang_waru
    volumes:
      - postgres_data:/var/lib/postgresql/data
    ports:
      - "5432:5432"

volumes:
  postgres_data:
```

Run with:
```bash
docker-compose up -d
```

---

## 📝 License

This project is licensed under the MIT License.

---

## 👨‍💻 Author

**Hanif Syabani**

---

<p align="center">Made with ❤️ for Karang Waru Village</p>
