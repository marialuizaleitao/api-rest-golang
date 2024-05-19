# API-Rest-Golang

## Overview

This is a RESTful API built with Go (Golang) for managing pilots. It allows you to perform CRUD operations (Create, Read, Update, Delete) on pilot entities. The application uses Docker for containerization and connects to a PostgreSQL database for data storage. Additionally, it includes integration with pgAdmin for database management.

## Technologies Used

- **Go (Golang)**: Main programming language used for backend development.
- **Docker**: Containerization platform used to manage application dependencies and environment consistency.
- **PostgreSQL**: Relational database management system used to store pilot data.
- **GORM**: Object-Relational Mapping (ORM) library for Go, used to interact with the PostgreSQL database.
- **pgAdmin**: Web-based administration tool for PostgreSQL, used for database management and visualization.

## Setup

### Prerequisites

- Docker installed on your machine ([Install Docker](https://docs.docker.com/get-docker/))

### Instructions

1. Clone the repository:
   ```sh
   git clone <repository-url>
2. Navigate to the project directory:
   ```sh
   cd api-rest-golang
3. Start the Docker containers:
   ```sh
   docker-compose up
4. Access the API at http://localhost:8000/


## Endpoints

- **GET /pilots**: Retrieve all pilots
- **GET /pilots/{id}**: Retrieve a pilot by ID
- **POST /pilots**: Create a new pilot
- **PUT /pilots/{id}**: Update an existing pilot
- **DELETE /pilots/{id}**: Delete a pilot

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.
