# Post Connect

REST API for user management built with Go.

## Technologies

- **Go 1.19** — application runtime
- **MySQL 8** — database
- **Docker & Docker Compose** — containerized development and deployment

## Project Structure

```
.
├── api/
│   ├── main.go
│   └── src/
│       ├── config/
│       ├── controllers/
│       ├── database/
│       │   ├── database.go
│       │   └── migrations/
│       │       └── 001_create_users.sql
│       ├── models/
│       ├── repositories/
│       └── router/
│           └── routes/
├── Dockerfile
├── Dockerfile.dev
├── docker-compose.yml
├── .dockerignore
├── go.mod
├── go.sum
└── .env
```

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

For local development without Docker:

- Go 1.19+
- MySQL 8

## Getting Started

### 1. Configure environment variables

Copy or edit the `.env` file in the project root:

```env
APP_PORT=5000
API_URL=root:root@tcp(post_db:3306)/post
HASH_KEY=change-me-hash-key-32chars!!
BLOCK_KEY=change-me-block-key-32chars!

MYSQL_ROOT_PASSWORD=root
MYSQL_DATABASE=post
MYSQL_PORT=3306
```

When running locally (outside Docker), update `API_URL` to point to your MySQL host:

```env
API_URL=root:root@tcp(localhost:3306)/post
```

### 2. Build (Docker)

Generates the API image. **Does not start any container.**

The compose file defines two API services (`api-dev` and `api-prod`). Build only the one you need.

**Development** (`Dockerfile.dev`):

```bash
docker compose build api-dev
```

**Production** (`Dockerfile`):

```bash
docker compose build api-prod
```

Force a full rebuild without cache:

```bash
docker compose build api-dev --no-cache
```

> `post_db` uses the official `mysql:8` image — there is nothing to build for the database.

### 3. Run (Docker)

Starts the containers. **Requires a previous build** (step 2).

Do not use `--build` here — build and run are separate steps.

**Development** (hot reload with Air):

```bash
docker compose up api-dev post_db
```

API available at `http://localhost:5000`.

**Production**:

```bash
docker compose up api-prod post_db
```

API available at `http://localhost:5001`.

Run in the background by adding `-d`:

```bash
docker compose up api-dev post_db -d
```

Stop the containers:

```bash
docker compose down
```

On the first startup, MySQL automatically runs the SQL files in `api/src/database/migrations/` and creates the `users` table.

If you already started the database before adding migrations, reset the volume and start again:

```bash
docker compose down -v
docker compose build api-dev
docker compose up api-dev post_db
```

**Typical workflow:**

```bash
docker compose build api-dev    # step 1 — build
docker compose up api-dev post_db   # step 2 — run
```

### 4. Useful Docker commands

#### Manage Go dependencies (without Go installed on the host)

The project uses Go 1.19. When adding packages, pin versions compatible with that release (e.g. `mysql@v1.7.1`).

**Container not running** — use `run`:

```bash
docker compose run --rm api-dev go get github.com/go-sql-driver/mysql@v1.7.1
docker compose run --rm api-dev go mod tidy
```

**Container already running** — use `exec`:

```bash
docker compose exec api-dev go get github.com/go-sql-driver/mysql@v1.7.1
docker compose exec api-dev go mod tidy
```

Changes are saved to `go.mod` and `go.sum` on the host via the volume mount (`.:/app`).

#### Access containers

- **`exec -it`** — entra no container que já está rodando
- **`run --rm -it`** — cria um container temporário (quando o serviço está parado)
- **`attach`** — reconecta aos logs do processo principal (Air), não abre shell

**Go dev** (`post_connect_dev`):

```bash
docker compose exec -it api-dev sh          # rodando
docker compose run --rm -it api-dev sh      # parado
```

**MySQL** (`post_db`):

```bash
docker compose up post_db -d
docker compose exec -it post_db mysql -u root -proot post
```

Senha: valor de `MYSQL_ROOT_PASSWORD` no `.env`.

**Logs do Air** (api em background com `-d`):

```bash
docker compose attach api-dev
```

Sair sem parar o container: `Ctrl+P`, `Ctrl+Q`.

### 5. Run locally (without Docker)

Make sure MySQL is running and the `API_URL` in `.env` matches your local database connection. Run the SQL from `api/src/database/migrations/001_create_users.sql` manually on your local database.

## API Endpoints

| Method | Endpoint          | Description       |
|--------|-------------------|-------------------|
| POST   | `/users`          | Create a user     |
| GET    | `/users`          | List users        |
| GET    | `/users/{userId}` | Get a user by ID  |
| PUT    | `/users/{userId}` | Update a user     |
| DELETE | `/users/{userId}` | Delete a user     |

### Create user

**Route:** `POST /users`

**Body:**

```json
{
  "name": "John Doe",
  "nickname": "johnd",
  "email": "john@example.com",
  "password": "secret123"
}
```

## Environment Variables

| Variable              | Description                              |
|-----------------------|------------------------------------------|
| `APP_PORT`            | Port the API listens on (default: 5000)  |
| `API_URL`             | MySQL connection string (DSN)            |
| `HASH_KEY`            | Secret key for password hashing          |
| `BLOCK_KEY`           | Secret key for encryption                |
| `APP_ENV`             | Env (`dev` or `prod`)                    |
| `MYSQL_ROOT_PASSWORD` | MySQL root password (Docker only)        |
| `MYSQL_DATABASE`      | MySQL database name (Docker only)        |
| `MYSQL_PORT`          | MySQL port exposed on the host           |
