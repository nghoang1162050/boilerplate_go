# boilerplate_go

## Overview
boilerplate_go is a sample Go web application that demonstrates a RESTful API built using the Echo framework. It utilizes GORM for MySQL database interactions, go-redis for caching, and integrates Swagger for API documentation. The project also supports Prometheus monitoring with periodic metric pushes.

## Features
- **REST API Endpoints**: CRUD operations for products.
- **Pagination**: Built-in support for paginated search results.
- **Caching**: Redis caching for improved performance.
- **Monitoring**: Integrated Prometheus middleware with metrics pushed every 5 minutes.
- **Swagger Documentation**: API docs generated using swag.
- **Environment-Based Configuration**: Uses a `.env` file for configuring database, caching, and other settings.

## Getting Started

### Prerequisites
- [Go (>=1.18)](https://golang.org/dl/)
- MySQL Database
- Redis Server
- [godotenv](https://github.com/joho/godotenv) for loading environment variables.
- [swag](https://github.com/swaggo/swag) for generating Swagger docs.
- Prometheus and (optionally) a Pushgateway for metrics collection.

### Installation

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/your-username/boilerplate_go.git
   ```

2. **Navigate to the Project Directory:**
   ```bash
   cd boilerplate_go
   ```

3. **Download Dependencies:**
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

# Redis Configuration
REDIS_URL=redis://127.0.0.1:6379/0?protocol=3
# Alternatively, use:
# REDIS_HOST=127.0.0.1
# REDIS_PORT=6379

# HTTP Metrics Port (if needed)
HTTP_METRIC_PORT=your_metric_port
```

### Running the Application

Run the application with:

```bash
go run main.go
```

Your server will start (by default on port 8080) and will expose Swagger documentation at:
```
http://localhost:8080/swagger/index.html
```

### Prometheus Monitoring

The project includes Prometheus integration with a middleware for scraping metrics from the Echo application. Additionally, a ticker pushes metrics to a Pushgateway every 5 minutes. Ensure you have a Pushgateway running (default URL in the code: `http://pushgateway:9091`) or adjust the URL as necessary.

Your [Prometheus configuration](./config/prometheus.yml) might look like this:

```yaml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'echo-app'
    static_configs:
      - targets: ['host.docker.internal:8080']

  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']

  - job_name: 'grafana'
    static_configs:
      - targets: ['grafana:3000']
```

## Project Structure

```
boilerplate_go/
├── internal/
│   ├── controller/        # HTTP controllers handling API endpoints
│   ├── database/          # Database connection and initialization
│   ├── dto/               # Data Transfer Objects (DTOs)
│   ├── helper/            # Utility functions (e.g., automapper)
│   ├── model/             # GORM models and auto-generated entities
│   ├── repository/        # Data access layer with CRUD operations
│   ├── router/            # HTTP route definitions
│   ├── usecase/           # Business logic layer (use cases)
│   └── utils/             # Utility functions including Redis cache and Prometheus setup
├── config/
│   └── prometheus.yml     # Prometheus configuration file
├── docs/                  # Swagger documentation (generated via swag init)
├── main.go                # Application entry point
├── .env                   # Environment variable configurations
└── README.md
```

## Generating Swagger Documentation

Generate the API docs by running:

```bash
swag init
```

Then visit:
```
http://localhost:8080/swagger/index.html
```

## Contributing

Contributions are welcome! Please open issues for any major changes or improvements and submit pull requests accordingly.

## License

Distributed under the MIT License. See `LICENSE` for more information.