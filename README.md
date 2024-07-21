# Full-Fledged REST API

## Overview

This project is designed to build a **full-fledged** REST API with a range of advanced features. The API incorporates modern authentication methods, flexible database configurations, and caching mechanisms to provide a robust and scalable solution.

### Key Features

- **Email and Password Authentication**: Secure authentication using JWT and password hashing.
- **OAuth Configuration**: Integration with OAuth providers for additional authentication options.
- **Database Support**: Configuration for both MongoDB and PostgreSQL to handle diverse data storage needs.
- **Caching**: Efficient caching mechanisms to enhance performance and scalability.

## Getting Started

### Prerequisites

- Go 1.18 or later
- MongoDB
- PostgreSQL
- Redis (for caching, if applicable)
- Git

### Installation

1. **Clone the Repository**

    ```sh
    git clone https://github.com/your-username/your-project.git
    cd your-project
    ```

2. **Set Up Environment Variables**

    Create a `.env` file in the root directory and configure the following environment variables:

    ```env
    MONGODB_URI=mongodb://localhost:27017
    MONGODB_NAME=your-db-name
    POSTGRES_URI=postgres://user:password@localhost:5432/your-db-name
    PORT=8080
    JWT_SECRET=your_jwt_secret
    ```

3. **Install Dependencies**

    ```sh
    go mod tidy
    ```

4. **Run Migrations**

    Ensure your databases are set up by running the necessary migrations. For example, if you are using a migration tool, you might run:

    ```sh
    ./scripts/migrate.sh
    ```

5. **Start the Application**

    ```sh
    go run main.go
    ```

## Usage

### API Endpoints

- **Authentication**
  - `POST /auth/login` - Login and receive a JWT token.
  - `POST /auth/register` - Register a new user.

- **Listings**
  - `GET /listings` - Retrieve all listings.
  - `POST /listings` - Create a new listing.
  - `PUT /listings/{id}` - Update a listing.
  - `DELETE /listings/{id}` - Delete a listing.
  - `GET /listings/{id}` - Get a specific listing by ID.

- **Users**
  - `GET /users` - Retrieve all users.
  - `POST /users` - Create a new user.
  - `PUT /users/{id}` - Update a user.
  - `DELETE /users/{id}` - Delete a user.
  - `GET /users/{id}` - Get a specific user by ID.

## Configuration

The project supports configuration through environment variables. Make sure to set the required variables in your `.env` file.

### Configuration Details

- **MongoDB**: Configurable via `MONGODB_URI` and `MONGODB_NAME`.
- **PostgreSQL**: Configurable via `POSTGRES_URI`.
- **Port**: Configurable via `PORT`.
- **JWT Secret**: Configurable via `JWT_SECRET`.

## Backlog

- [ ] Implement Caching with Redis
- [ ] Add OAuth integration for social login
- [ ] Develop frontend interface for the API
- [ ] Create comprehensive API documentation
- [ ] Optimize for performance and scalability
- [ ] Enable email/password authentication

## Licene

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact 

Feel free to send me a message.