# Email Notification System

A simple Go application for sending email notifications using SMTP. This application demonstrates how to structure an email with HTML content and securely handle sensitive information using a `.env` file.

## Features

- Sends HTML-formatted emails.
- Supports SMTP configuration.
- Uses `.env` for secure and flexible configuration.

## Prerequisites

- Go 1.18 or later
- An SMTP server (e.g., Gmail SMTP)
- `godotenv` library for loading `.env` files:
  ```
  go get github.com/joho/godotenv
  ```

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/your-username/email-notification-system.git
   cd email-notification-system
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Create a `.env` file:
   Copy the provided `.env.example` file and update it with your own SMTP configuration.

   ```
   cp .env.example .env
   ```

   Update the values in `.env`:

   ```
   SMTP_EMAIL=your-email@example.com
   SMTP_PASSWORD=your-email-password
   SMTP_HOST=smtp.example.com
   SMTP_PORT=587
   TO_EMAIL=recipient@example.com
   ```

## Usage

Run the application:
```
go run main.go
```

If configured correctly, the program will send an email to the recipient specified in the `.env` file.

## Project Structure

```
├── .env              # Your environment variables (ignored by Git)
├── .env.example      # Example environment variables for reference
├── go.mod            # Go module file
├── go.sum            # Go dependencies file
├── main.go           # The main application logic
└── README.md         # Project documentation
```

## Environment Variables

The following environment variables are required:

| Variable       | Description                                    |
|----------------|------------------------------------------------|
| `SMTP_EMAIL`   | The sender's email address                    |
| `SMTP_PASSWORD`| The sender's email password or app-specific password |
| `SMTP_HOST`    | The SMTP server (e.g., `smtp.gmail.com`)       |
| `SMTP_PORT`    | The SMTP server port (e.g., `587`)             |
| `TO_EMAIL`     | The recipient's email address                 |

## Notes

- **Security**: Avoid hardcoding sensitive credentials in the codebase. Use `.env` for this purpose.
- **Gmail Users**: If using Gmail, ensure that "Allow less secure apps" is enabled or use an App Password if 2FA is enabled.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgements

- [godotenv](https://github.com/joho/godotenv) for managing environment variables.
- [Gmail SMTP](https://support.google.com/mail/answer/7126229) for email delivery.

---

**Happy coding!**
