# Banking Simulation API 🚀

## Overview

This project provides a simulated banking API, designed to mimic real banking operations in a controlled, read-only environment. The API is built using Go, leverages SQLite for data storage, and is containerized using Docker. All data generated and used by this API is entirely fake and does not represent any real banking information or operations.

> Totally experimental

## Features 🌟

- **Accounts**: Retrieve information about simulated bank accounts.
- **Transactions**: View transaction history associated with accounts.
- **Instruments**: Access information on various financial instruments.
- **Investments**: View investment details.
- **WebSocket**: Receive real-time updates on instrument prices.

## Important Notice ⚠️

**All data provided by this API is fake and for simulation purposes only. It does not represent any real banking data or operations.**

## API Documentation 📜

The API follows the Swagger 2.0 specification. Detailed API documentation can be accessed through the Swagger UI when running the application.

## Authentication 🔐

- Most endpoints require authentication.
- Use any email from the `/emails` endpoint with the password `password` to authenticate.
- The API uses Bearer Token for authentication, which should be included in the `Authorization` header.

## Rate Limiting ⏱️

- The API is rate-limited to 100 requests per second.

## WebSocket Endpoint 🌐

- For real-time updates, connect to the WebSocket endpoint using the `ws` protocol.
- Authentication is required. Pass the token obtained from the `/auth` endpoint as a query parameter `?token=[token]`.
- Messages will be sent at random intervals with updated prices of specified instruments.

## Docker Setup 🐳

A Dockerfile is provided to containerize the application. Follow the instructions below to build and run the Docker container.

### Building the Docker Image

```sh
docker build -t banking-simulation-api .
```

### Running the Docker Container

```sh
docker run -d -p 8080:8080 banking-simulation-api
```

## Local Development 🛠️

### Prerequisites

- Go 1.22 or later
- SQLite

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/banking-simulation-api.git
    cd banking-simulation-api
    ```

2. Install dependencies:
    ```sh
    go mod download
    ```

3. Run the application:
    ```sh
    go run main.go
    ```

## License 📄

This project is licensed under the Apache 2.0 License. See the [LICENSE](LICENSE) file for details.

## Contact 📬

For any queries or issues, please open an issue in the GitHub repository or contact the project maintainer. Feel free to reach out, we're here to help! 😊

---

**Disclaimer**: This API is for educational and simulation purposes only. All data and operations are simulated and do not represent any real banking system or data.

---