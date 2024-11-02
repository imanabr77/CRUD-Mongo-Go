# User Management System

This project demonstrates CRUD (Create, Read, Update, Delete) operations using MongoDB with Golang, specifically for managing user information.

## Table of Contents
- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Setup](#setup)
- [Usage](#usage)
  - [Creating Users](#creating-users)
  - [Reading Users](#reading-users)
  - [Updating Users](#updating-users)
  - [Deleting Users](#deleting-users)
- [API Endpoints](#api-endpoints)
- [Configuration](#configuration)
- [Contributing](#contributing)

## Introduction

This application provides a simple RESTful API for managing user records in a MongoDB database using Golang. It covers the core CRUD operations, allowing users to interact with the system through HTTP requests.

## Prerequisites

Before running this application, ensure you have the following installed:

- Go (version 1.16+)
- MongoDB (version 4.0+)
- Docker and Docker-compose (for development environment setup)

## Setup

To set up the project:

1. Clone the repository:
git clone [https://github.com/yourusername/user-management-system.git](https://github.com/imanabr77/CRUD-Mongo-Go) cd CRUD-Mongo-Go

2. Set up the MongoDB connection string in the `.env` file:
MONGO_URI=mongodb://localhost:27017/users_db

3. Install dependencies:
go mod tidy

4. Build the application:
go build

5. Run the application:
./main.go

6. Start the MongoDB server (if not already running):

Example curl command:
bash curl -X POST http://localhost:9000/users
-H "Content-Type: application/json"
-d '{"name":"John Doe","email":"john@example.com"}'


bash curl -X DELETE http://localhost:9000/users/{id}

Replace `{id}` with the actual user ID.

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/users` | GET | List all users |
| `/users` | POST | Create a new user |
| `/users/{id}` | GET | Get user by ID |
| `/users/{id}` | PUT | Update user |
| `/users/{id}` | DELETE | Delete user |

## Configuration

The application uses environment variables for configuration. You can set these in your `.env` file:

- `MONGO_URI`: MongoDB connection string
- `PORT`: Port to listen on (default: 8080)

## Contributing

Contributions are welcome! Please feel free to submit pull requests or issues.

[Your GitHub username]: iman.abrehdari


