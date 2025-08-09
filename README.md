## Websocket API - Leilão em Tempo Real

### 👨‍💻 Tecnologias e ferramentas

| Tecnologia | Descrição |
| ---------- | --------- |
| Go | Linguagem de programação estaticamente tipada |
| Chi | Framework Go para criação de servidores HTTP |
| Postgres | Banco de dados relacional |
| Docker | Plataforma de software para implantar aplicativos em containers |
| Gorilla Websocket | Biblioteca para implementação de comunicação em tempo real |
| SCS - Session Manager | Autentição baseada em sessões |

### 📝 Descrição do projeto

Uma API WebSocket desenvolvida em Go para permitir que usuários acompanhem e participem de leilões em tempo real. Ideal para plataformas de e-commerce, marketplaces ou qualquer aplicação com necessidade de comunicação em tempo real.

### ⚡ Funcionalidades do projeto

- Comunicação bidirecional via Websocket
- Autenticação de usuários via sessão
- Inscrição para receber eventos de produtos/leilões específicos
- Notificações em tempo real sobre lances atualizados

### ⚙ Endpoints

- Obter token de autenticação: ```/api/v1/csrftoken```
- Inscrição de usuários: ```/api/v1/users/signupuser```
- Login de usuários: ```/api/v1/users/loginuser```
- Logout de usuários: ```/api/v1/users/logout```
- Criar produto: ```/api/v1/products/```
- Upgrade de conexão websocket: ```/api/v1/products/ws/subscribe{product_id}```

### 📂 Estrutura de pastas

```shell
├───cmd                      # pontos de entrada da aplicação
│   ├───api                  # API aplicada
│   └───terndotenv           # roda migrations
└───internal                 # lógica de negócio
    ├───api                  # handlers, rotas e lógica de comunicação HTTP/Websocket
    ├───jsonutils            # funções que lidam com requisições HTTP
    ├───services             # funções de tratamento das requisições
    ├───store                # banco de dados
    │   └───pgstore          # funcionalidades do banco de dados postgres
    │       ├───migrations   # migrações do banco de dados
    │       └───queries      # queries para manipulação do banco de dados
    ├───usecase              # casos de uso
    │   ├───product          # funções para caso de uso de produtos
    │   └───user             # funções para caso de uso de usuários
    └───validator            # funções de validação do JSON
```

### 🚀 Rodando localmente

**Pré-requisitos:**
- Go 1.20+

**Clone o repositório**
```bash
git clone https://github.com/joaomarcosg/Projeto-Gobid.git
```

**Instale as dependências**
```bash
go mod tidy
```

**Rode o servidor**
```bash
go run /cmd/api/main.go
```

**Testes**
```bash
go test ./...
```

