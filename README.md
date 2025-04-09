# boilerplate_go

## Overview
boilerplate_go is a sample Go web application that demonstrates a RESTful API built with the Echo framework, using GORM for MySQL database interactions, go-redis for caching, and Swagger for API documentation.

## Features
- **REST API Endpoints**: CRUD operations for products.
- **Pagination**: Built-in support for paginated search results.
- **Caching**: Redis caching for improved response times.
- **Swagger Documentation**: Swagger docs generated with swaggo/swag.
- **Environment-Based Configuration**: Uses a `.env` file to configure DB and Redis connections.

## Getting Started

### Prerequisites
- Go (>=1.18)
- MySQL Database
- Redis Server
- [godotenv](https://github.com/joho/godotenv) for loading environment variables.
- [swag](https://github.com/swaggo/swag) for generating Swagger docs.

### Installation
1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/boilerplate_go.git
   ```
2. **Change to the project directory:**
   ```bash
   cd boilerplate_go
   ```
3. **Install dependencies:**
   ```bash
   go mod download
   ```

### Configuration
Create a `.env` file in the project root with the following content (adjust the values as needed):
```dotenv
# Database Configuration
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_HOST=your_db_host
DB_PORT=your_db_port
DB_NAME=your_db_name
HTTP_METRIC_PORT=your_metric_port

# Redis Configuration
REDIS_URL=redis://127.0.0.1:6379/0?protocol=3
```
*Alternatively, if not using `REDIS_URL`, set:*
```dotenv
REDIS_HOST=127.0.0.1
REDIS_PORT=6379
```

### Running the Application
Run the application using:
```bash
go run main.go
```
The server will start on the configured port (default is 8080).

### Generating Swagger Documentation
Generate Swagger docs by running:
```bash
swag init
```
The documentation will be generated under the `/docs` directory. You can then access it by navigating to:
```
http://localhost:8080/swagger/index.html
```

## Project Structure
```
boilerplate_go/
├── internal/
│   ├── controller/        # HTTP controllers handling API endpoints
│   ├── database/          # Database connection and initialization
│   ├── dto/               # Data Transfer Objects and response wrappers
│   ├── helper/            # Utility functions (e.g., automapper)
│   ├── model/             # GORM models and auto-generated entities
│   ├── repository/        # Data access layer with CRUD and search functions
│   ├── router/            # HTTP route definitions
│   ├── usecase/           # Business logic layer (use cases)
│   └── utils/             # Utility functions including Redis cache management
├── docs/                  # Swagger documentation (generated via swag init)
├── main.go                # Application entry point
├── .env                   # Environment variable configurations
└── README.md
```

## Contributing
Contributions are welcome! For major changes, please open an issue first to discuss what you would like to change. Pull requests are appreciated.

## License
Distributed under the MIT License. See `LICENSE` for more information.