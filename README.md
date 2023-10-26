# Days to 

**Days to 2025 Calculator** is a simple Go program that calculates the number of days remaining until January 1, 2025, when provided with any token as a URL query parameter.

## Prerequisites

Before you begin, make sure you have the following requirements:

- Go (Golang) must be installed on your system.

## Usage

1. Clone or download this repository to your local machine.

2. Open a terminal and navigate to the directory containing the code.

3. Run the Go program using the following command:

    ```shell
    go run cmd/main.go
    ```

4. The server will start and listen on port 8080.

5. To calculate the days remaining until January 1, 2025, access the following URL in your web browser :

    ```bash
    http://localhost:8080/?token=your_token
    ```

    Replace `your_token` with any token.

6. If a token is provided, the server will respond with the number of days remaining. If the token is missing, a 403 Forbidden response will be returned.

