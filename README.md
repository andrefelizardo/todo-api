# Todo List API

**Todo List API** built with **Golang** following the principles of **Clean Architecture**.

![Golang](https://img.shields.io/badge/Go-1.22.1-blue.svg)
![CloudRun](https://img.shields.io/badge/Cloud%20Run-Deployed-brightgreen)
![License](https://img.shields.io/badge/license-MIT-brightgreen)

## 📋 About the Project

This is a **Todo List API** built with **Golang**, following the principles of **Clean Architecture**. The goal of this project is to provide a robust and scalable solution for task management, allowing users to authenticate, create, update, list, and delete tasks. Additionally, this project aims to practice best development and deployment practices using cloud environments.

The API is fully containerized and deployed on **Google Cloud Run**, with a database hosted on **Cloud SQL** (PostgreSQL).

## 🚀 Features

- User authentication using JWT (Json Web Token).
- Full CRUD for tasks:
  - Create, list, update, and delete tasks.
  - Filter tasks by status (pending, in-progress, completed).
  - Filter tasks by due date.
- Route protection via authentication middleware.
- Pagination and sorting configuration.
- Deployment on **Google Cloud Platform**.

## 🛠️ Technologies Used

- **Language**: Golang (1.18)
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin)
- **Database**: PostgreSQL (Cloud SQL)
- **Authentication**: JWT
- **Deployment**: Google Cloud Run
- **ORM**: GORM
- **Others**: Docker, GitHub Actions, Google Cloud Platform

## 📂 Project Structure

The project follows the **Clean Architecture** structure, promoting a clear separation between business logic, infrastructure, and interfaces.

```
/todo-api
│
├── /cmd
│   └── /api                  # Application entry point
│       └── main.go
│
├── /internal
│   ├── /controllers           # Controllers (HTTP Handlers)
│   ├── /domain                # Models and entities
│   ├── /repositories          # Repositories for database access
│   ├── /usecases              # Business logic (use cases)
│   ├── /services              # External services (JWT, email, etc.)
│   ├── /middlewares           # Middlewares (Authentication)
│   ├── /routes                # Route definitions
│   └── /configs               # Application configurations
└── /scripts                   # Helper scripts (migrations, etc.)
```

## 🔧 Running the Project Locally

### Prerequisites

Before running the project, ensure you have the following installed:

- [Golang](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- A PostgreSQL instance (local or remote)

### Steps to run locally

1. **Clone the repository**

   ```bash
   git clone https://github.com/andrefelizardo/todo-api.git
   cd todo-api
   ```

2. **Configure environment variables**
   Create a `.env` file in the project root with the following variables:

   ```bash
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your-username
   DB_PASSWORD=your-password
   DB_NAME=todo_db
   JWT_SECRET=your-secret-key
   ```

3. **Run with Docker**
   Build and run the Docker container:

   ```bash
   docker build -t todo-api .
   docker run -d -p 8080:8080 --env-file .env todo-api
   ```

4. **Test the API**
   After starting the server, you can access the API endpoints, for example:
   - `POST /tasks` to create a task.
   - `GET /tasks` to list tasks.

## 🌐 Deployment on Google Cloud Platform

The API is deployed using **Google Cloud Run**, with the database hosted on **Cloud SQL**.

### Quick Deployment Tutorial:

1. **Containerize the application**
   Ensure your project is containerized with Docker.

2. **Create a Cloud SQL instance**
   In the Google Cloud Console, create a PostgreSQL instance and configure the credentials.

3. **Deploy to Cloud Run**
   Use the following command to deploy:

   ```bash
   gcloud run deploy todo-api --source . --platform managed --region us-central1
   ```

4. **Set up Cloud SQL**
   Connect your Cloud SQL instance to the Cloud Run service by following the [official documentation](https://cloud.google.com/sql/docs/postgres/connect-run).

## 🔑 Authentication

The API uses JWT authentication to secure routes. After creating a user and logging in, you will receive a JWT token, which must be included in the headers of all authenticated requests:

```
Authorization: Bearer <token>
```

## 🧪 Testing

To run the tests, use the following command:

```bash
go test ./...
```

The tests ensure that all core functionalities (task CRUD, authentication, etc.) work correctly.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

### 🌟 Contributions

Contributions are welcome! Feel free to open issues and submit pull requests for improvements or bug fixes.

---

**Made with 💙 by [André Felizardo](https://github.com/andrefelizardo)**
