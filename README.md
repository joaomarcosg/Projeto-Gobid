<h1 align="center" style="font-weight:bold;">Websocket API Real-Time Auction</h1>

<p align="center">
 <a href="#tech">Technologies</a> • 
 <a href="#started">Getting Started</a> • 
  <a href="#routes">API Endpoints</a>
</p>

<p align="center">
    <b>
    A WebSocket API developed in Go to allow users to track and participate in auctions in real time.
    This real-time communication application was created as a way to put my knowledge of concurrency using Go into practice.
    </b>
</p>

<h2 id="technologies">💻 Technologies</h2>

| Technology | Description |
| ---------- | ----------- |
| Go | Statically typed programming language |
| Chi | Go framework for creating HTTP servers |
| Postgres | Relational database |
| Docker | Software platform for deploying containerized applications |
| Gorilla Websocket | Library for implementing real-time communication |
| SCS - Session Manager | Session-based authentication |

<h2 id="started">🚀 Getting started</h2>

<h3>Prerequisites</h3>

- Go 1.20+

<h3>clone the repository</h3>

```bash
git clone https://github.com/joaomarcosg/Gobid-Project.git
```

<h3>Install the dependencies.</h3>

```bash
go mod tidy
```

<h3>Config .env variables</h2>

Use the `.env.example` as reference to create your configuration file `.env`

```yaml
GOBID_APP_PORT=8080
GOBID_DATABASE_PORT=5432
GOBID_DATABASE_NAME=bid
GOBID_DATABASE_USER=postgres
GOBID_DATABASE_PASSWORD=123
GOBID_DATABASE_HOST=localhost
GOBID_CSRF_KEY=abcdefghijlmnopqrstuvwxyz1234567
```

📌 **The environment variable ```GOBID_CSRF_KEY``` is a 32 bits key. Use [random.org](https://www.random.org/) to generate a string.**

<h3>Starting</h3>

    ```bash
    cd gobid
    go run /cmd/api/main.go
    ```

<h2 id="routes">📍 API Endpoints</h2>

| Route | Description |
| ----- | ----------- |
| <kbd>POST /api/v1/users/signupuser | User registration [request details](#post-signup-user) |
| <kbd>GET /api/v1/csrftoken</kbd> | Get authentication token [response details](#get-auth-detail) |
| <kbd>POST /api/v1/users/loginuser | User login [request details](#post-login-user)  |
| <kbd>POST /api/v1/users/logout | User logout [response details](#post-logout-user)  |
| <kbd>POST /api/v1/products/ | Create product [response details](#create-product) |
| <kbd>GET /api/v1/products/ws/subscribe{product_id} | WebSocket connection upgrade [response details](#websocket)  |

<h3 id="post-signup-user">POST /api/v1/users/signupuser</h3>

**REQUEST**
```json
{
    "user_name": "johndoe",
	"email": "johndoe@email.com",
	"password": "Password123",
	"bio": "testing my api"
}
```

**RESPONSE**
```json
{
	"user_id": "0db87e24-895b-4d11-8c35-95b2387dd211"
}
```

<h3 id="get-auth-detail">GET /api/v1/csrftoken</h3>

**RESPONSE**
```json
{
	"csrf_token": "tOZEaiWTtM2ZcxiteUuNmdob3ZFshZ7a1XWJuwxeE0UZE32nXjsXeHHfoid0GKNTIqXs7O4/tNs+v3FEIIgzUg=="
}
```

<h3 id="post-login-user">POST /api/v1/users/loginuser</h3>

**REQUEST**
```json
{
    "email": "johndoe@email.com",
	"password": "Password123"
}
```

**RESPONSE**
```json
{
    "message": "logged in sucessfully"
}
```

<h3 id="create-product">POST /api/v1/products/</h3>

**REQUEST**
```json
{
    "product_name": "test",
    "product_description": "testing",
    "base_price": 99.99,
    "auction_end": "YYYY-MM-DDTHH:MM:SSZ"
}
```

**RESPONSE**
```json
{
    "product_id": "0db87e24-895b-4d11-8c35-95b2387dd211"
}
```

<h3 id="post-logout-user">POST /api/v1/users/logout</h3>

**RESPONSE**
```json
{
    "message": "logged out sucessfully"
}
```










