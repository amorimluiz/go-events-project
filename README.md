# Go Events Project

This project is a Go application that manages events, allowing user registration, event creation, and querying.

## Table of Contents

- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Contributing](#contributing)
- [License](#license)

## Technologies Used

- Go (Golang)
- GORM (ORM for Go)
- MySQL
- Docker
- Docker Compose
- Gin (Web framework for Go)

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)

### Steps for Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/amorimluiz/go-events-project.git
   cd go-events-project
   ```

2. Create a `.env` file in the root of the project with the required environment variables:

   ```bash
   DB_USERNAME=
   DB_ROOT_PASSWORD=
   DB_HOST=
   DB_PORT=
   DB_NAME=
   JWT_SECRET=
   ```

3. Start the containers using Docker Compose:

   ```bash
   docker-compose up --build
   ```

4. The application will be available at `http://localhost:8080`.

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
