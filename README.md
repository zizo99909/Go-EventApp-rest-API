# Go EventApp REST API

![Go](https://img.shields.io/badge/Go-1.20+-blue) ![License](https://img.shields.io/badge/License-MIT-green)

A RESTful API for managing events, users, and bookings, built using **Go** with **Gin**, **GORM**, and a **PostgreSQL** database.

## Features
- **User Authentication** (JWT-based login/signup)
- **Event Management** (Create, Read, Update, Delete events)
- **User Bookings** (Register & manage event bookings)
- **Role-based Access Control**
- **Pagination & Filtering** for events

## Tech Stack
- **Backend:** Golang, Gin
- **Database:** PostgreSQL
- **Authentication:** JWT

## Getting Started

### Prerequisites
Make sure you have the following installed:
- [Go 1.20+](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)

### Installation
1. **Clone the repository:**
   ```sh
   git clone https://github.com/zizo99909/Go-EventApp-rest-API.git
   cd Go-EventApp-rest-API
   ```
2. **Create a `.env` file** and configure your database settings:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_NAME=eventapp_db
   JWT_SECRET=your_secret_key
   ```
3. **Run the application:**
   ```sh
   go mod tidy
   go run main.go
   ```

### Using Docker
```sh
docker-compose up --build
```

## API Endpoints

| Method | Endpoint            | Description                  | Auth Required |
|--------|---------------------|------------------------------|---------------|
| POST   | `/auth/signup`      | Register a new user          | No            |
| POST   | `/auth/login`       | Login and get a JWT token    | No            |
| GET    | `/events`           | Get a list of events         | No            |
| POST   | `/events`           | Create a new event           | Yes (Admin)   |
| GET    | `/events/:id`       | Get details of an event      | No            |
| PUT    | `/events/:id`       | Update an event              | Yes (Admin)   |
| DELETE | `/events/:id`       | Delete an event              | Yes (Admin)   |
| POST   | `/bookings`         | Register for an event        | Yes           |
| GET    | `/bookings`         | View user bookings           | Yes           |

For a full API reference, check the **Swagger Docs** at `/docs`.

## Contributing
Contributions are welcome! To contribute:
1. Fork the repository
2. Create a new branch (`feature/your-feature`)
3. Commit your changes
4. Push to your branch and open a Pull Request

## License
This project is licensed under the MIT License.

## Contact
For questions or support, contact [Ziyad Zaki](https://github.com/zizo99909).

