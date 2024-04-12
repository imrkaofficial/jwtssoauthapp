# JWT and SSO Authentication with Go (Gin/Gorm) and MySQL

This is a simple project demonstrating JWT (JSON Web Token) and SSO authentication in Go using Gin for the backend and Gorm as the ORM (Object Relational Mapper) with MySQL database. The project includes basic frontend HTML/CSS for login, signup, and forgot password functionality.

## Features
- User authentication with JWT
- Signup, login, forgot password, and reset password functionality
- MySQL database integration with Gorm
- Basic HTML/CSS frontend for user interaction

## Prerequisites

- Go installed on your machine
- MySQL installed and running
- Basic understanding of Go, Gin, Gorm, JWT, HTML, and CSS

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/imrkaofficial/jwtssoauthapp.git
   cd jwtssoauthapp

2. Install Dependencies:

    ```bash
    go mod download

3. Set up MySQL database:

    - Update the MySQL DSN in `main.go` with your database connection details.

## Usage

1. Run the application:

    ```bash
    
    go run main.go
The server will start and listen on port 9000.

2. To visit on pages
    - Login page:           `http://localhost:9000/login`
    - Signup page:          `http://localhost:9000/signup`
    - Forgot Password page: `http://localhost:9000/forgotpwd`

## Author

- Rajat Kumar Goyal

## License
This project is licensed under the [MIT LICENSE](LICENSE.md).
