
# ContactsAPI

ContactsAPI is a simple RESTful API for managing contact information. It is developed in Go and follows best practices in API design and Go programming. 

## Functional Requirements

This API includes the following endpoints:
- **POST /contacts**: Create a new contact.
- **GET /contacts/{id}**: Retrieve a contact by ID.
- **DELETE /contacts/{id}**: Delete a contact by ID.
- **PUT /contacts/{id}**: Update an existing contact.

## Getting Started

### Prerequisites

You need to have Go installed on your machine. To install Go, you can follow the instructions on the [official Go website](https://golang.org/doc/install).

### Installing

To set up the project, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/andrew221293/contacts.git
   ```
2. Navigate to the project directory:
   ```bash
   cd contacts
   ```
3. Build the project:
   ```bash
   go build
   ```
4. Run the server:
   ```bash
   ./contacts
   ```

## Usage

To use the API, you can start the server and use a tool like Postman or curl to make requests:

### Creating a Contact

```bash
curl -X POST http://localhost:8080/contacts -d '{"name":"John Doe", "email":"johndoe@example.com", "phone":"1234567890"}'
```

### Getting a Contact

```bash
curl http://localhost:8080/contacts/{id}
```

## Design and Implementation

The API has been designed with simplicity and rapid development in mind. Here are some key aspects of its implementation:
- **In-Memory Storage**: We opted to store contact data in an in-memory list to simplify the architecture and avoid external dependencies, which accelerates both development and testing phases.
- **No Authorization Logic**: To make the API easily accessible and straightforward for demonstration purposes, we did not implement any form of authorization or authentication.
- **Best Practices**: Throughout the development, we adhered to Go's best practices for naming conventions, code structuring, and error handling to ensure code quality and maintainability.
- **Code Documentation**: Comprehensive documentation has been provided within the code to explain its functionality and important design decisions clearly.
- **Development Tools**: Utilization of advanced IDE features and code generation tools like GoLand and Visual Studio Code has been essential in expediting the development process.
- **Postman Collection**: A Postman collection is available to demonstrate how to interact with each endpoint effectively.

## Development Time Estimate

With the assistance of artificial intelligence tools like ChatGPT, which facilitated rapid code generation and design decision-making, the development of this project was highly efficient. This approach significantly reduced the estimated development time, allowing the project to be completed about 2 hours. This efficiency was achieved while maintaining high standards of quality and adhering to Go's best practices.
