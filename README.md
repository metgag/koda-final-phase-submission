## 📋 Project Overview

### Technologies Used

- **Go (Golang)** - Main programming language
- **Gin** - Web framework
- **PostgreSQL** - Database (via pgx/v5)
- **Redis** - Caching and session management
- **Docker** - Containerization and deployment

### Features

- ✅ User may create a post
- ✅ Following another user 
- ✅ View recently followed user's post
- ✅ Interraction between user's post, like, comment

## 🚀 Installation

### Prerequisites

- Go 1.25
- PostgreSQL
- Redis

### Environment Variables

Create a `.env` file in the root directory:

```env
# env for pg database
PG_USER=root
PG_PWD=root
PG_HOST=localhost
PG_PORT=5432
PG_DB=social

# env for jwt
JWT_SECRET=a-string-secret-at-least-256-bits-long
JWT_ISSUER=issuer_name

# env for redis
RDB_ADDR=localhost:6379
```

#### Setup Instructions Migrate
1. make sure no no existing generated database
``` bash
make migrate-down
```
2. create database tables
``` bash
make migrate-up
``` 
3. insert seed to tables
``` bash
make seed
```

The server will start on `http://localhost:8090`

## 🔐 Authentication

Most endpoints require authentication using JWT tokens. Include the token in the Authorization header:

```
Authorization: Bearer <your_jwt_token>
```