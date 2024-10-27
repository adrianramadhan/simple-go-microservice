# Microservice Architecture with API Gateway

## Overview
This repository contains a microservice architecture that consists of an API Gateway, User Service, Order Service, and a separate database. 
The architecture follows the Clean Architecture principles, ensuring separation of concerns and scalability. 
This setup is designed to handle user and order management effectively while allowing for future expansion of services.

## Architecture Diagram
![diagram-export-10-27-2024-5_29_34-PM](https://github.com/user-attachments/assets/4503466d-1909-45da-b245-8b540608c854)

The system uses a microservice architecture. Each service is independent, with its own database:

1. API Gateway: Acts as the entry point.
2. User Service: A service for user management.
3. Order Service: A service to manage orders associated with users.
4. Each service is containerized and can be individually scaled or updated without affecting the others.

## Explanation of the Folder Structure
- **config**: Contains configuration files and settings, ensuring centralized configuration management.
- **internal/db**: Responsible for managing database connections and any related initialization tasks.
- **internal/dto**: Data Transfer Objects are defined here to facilitate data exchange between different layers of the application.
- **internal/entity**: This folder contains domain entities that represent the core business objects.
- **internal/handler**: The HTTP request handlers are implemented here, defining how incoming requests are processed.
- **internal/repository**: This layer handles all interactions with the database, abstracting data persistence logic.
- **internal/service**: Contains the business logic of the application, orchestrating interactions between the handler and repository layers.
- **main.go**: The entry point for the application, initializing the services and starting the server.

## Services
### Api Gateway
- **Description:** Acts as a single entry point for all client requests. It routes requests to the appropriate microservices (User Service and Order Service) and handles authentication, rate limiting, and response aggregation.
- **Technologies:** Go, Gin

### User Service
- **Description:** Manages user-related operations such as registration, authentication, and profile management.
- **Technologies:** Go, Gin, GORM
  
### Order Service
- **Description:** Handles order processing, including order creation, retrieval, and management.
- **Technologies:** Go, Gin, GORM
  
### Database
- **Description:** A database that stores user and order information. This setup simplifies data management but can be scaled out in the future to a more decoupled architecture.
- **Technologies:** PostgreSQL (or your choice of RDBMS)

## Getting Started
### Prerequisites
- Docker
- Docker Compose
  
### Running the Application
1. Clone the repository:
```
  git clone https://github.com/adrianramadhan/simple-go-microservice.git
  cd simple-go-microservice

```
2. Build and run the services using Docker Compose:
```
  docker-compose up --build
```
3. Accessing Services:
- API Gateway: http://localhost:8080
- User Service: http://localhost:8081
- Order Service: http://localhost:8082
- pgAdmin (Database Management): http://localhost:8083 (if pgAdmin setup added)

## API Endpoints

### API Gateway

#### User Service
- **POST /api/users**: Register a new user.
- **GET /api/users**: List all users
- **GET /api/users/{id}**: Retrieve user details.

#### Order Service
- **POST /api/orders**: Create a new order.
- **GET /api/orders**: List all orders
- **GET /api/orders/{id}**: Retrieve order details.

## Architecture Principles

### Clean Architecture
This project follows Clean Architecture principles, promoting a clear separation of concerns:

- **Independence of Frameworks**: The architecture is designed to be independent of any specific framework, enabling easy adoption of new technologies.
- **Testability**: Business logic can be tested independently of external systems.
- **Independence of UI**: The user interface can be developed separately from the underlying business logic.

### Microservices
The microservices architecture allows for:

- **Scalability**: Each service can be scaled independently based on demand.
- **Flexibility**: Services can be developed, deployed, and maintained independently.
- **Resilience**: If one service fails, it does not affect the entire system.

## Future Enhancements
- Introduce additional services (e.g., Payment Service) to enhance functionality.
- Implement service discovery for improved service communication.
- Consider moving to a more decoupled database architecture to prevent service dependencies.

## Contributing
Contributions are welcome! Please read the [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
