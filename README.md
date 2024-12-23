
## Prerequisites

- Docker
- Docker Compose
- Golang 1.22

## Getting Started

### Build and Run with Docker Compose

1. Build and start the services:

    ```sh
    docker-compose up --build
    ```

2. The server will be available at [http://localhost:3001](http://_vscodecontentref_/4).

### Endpoints

- `GET /ping`: Returns a pong message.
- `GET /hello/:name`: Greets the user with the provided name.
- `GET /user/:id`: Returns user information for the given ID.
- `GET /product/:id`: Returns product information for the given ID.
- `GET /cart/:id`: Returns cart information for the given ID.
- `GET /order/:id`: Returns order information for the given ID.
- `POST /user`: Creates a new user.
- `GET /products`: Returns a list of all products.

### Example Requests

- `GET /ping`

    ```sh
    curl http://localhost:3001/ping
    ```

- `POST /user`

    ```sh
    curl -X POST http://localhost:3001/user -H "Content-Type: application/json" -d '{"name":"John Doe","email":"john@example.com","password":"password123"}'
    ```

## Jenkins Pipeline

The project includes a Jenkins pipeline defined in the [Jenkinsfile](http://_vscodecontentref_/5). The pipeline performs the following stages:

1. Clone Repository
2. Build Docker Image
3. Run Tests
4. Push to Docker Hub
5. Deploy Golang to DEV

## License

This project is licensed under the MIT License.