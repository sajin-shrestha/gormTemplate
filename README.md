# Go Project with GORM, JWT Authentication, and Mux

This project is a Go application that demonstrates CRUD operations using GORM with PostgreSQL, JWT authentication, and routing with Gorilla Mux.

## Features

- **GORM**: Object-Relational Mapping library for database interactions.
- **PostgreSQL**: Database backend for storing user data.
- **JWT Authentication**: Secure endpoints using JSON Web Tokens.
- **Gorrila Mux**: Powerful HTTP router and dispatcher.

## Getting Started

To run the application locally, follow these steps:

### Prerequisites

- Go (version 1.16 or higher)
- PostgreSQL
- Postman (for testing APIs)

### Installation

1. **Clone the repository:**

```bash
git clone https://github.com/your-username/gormTemplate.git
cd goromTemplate

```

2. **Create a .env file in root folder with following data:**

   ```bash
    DB_HOST=localhost
    DB_USER=<database-username>
    DB_PASSWORD=<database-password>
    DB_NAME=<database-name>
    DB_PORT=5432

    JWT_SECRET_KEY=<your-custom-jwt-secret-key>
   ```

3. **Install Dependencies:**

   ```bash
   go mod tidy

   ```

4. **Start the backend server:**
   ```bash
   make build
   make run
   ```

### Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b <new-branch>`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin <your-branch-name>`)
5. Open a Pull Request
