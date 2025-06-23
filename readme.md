# Golang BMS

A Book Management System (BMS) built with Go (Golang). This project demonstrates CRUD operations, RESTful APIs, and structuring a Go web service for managing books, authors, and related entities.

## Features

- RESTful API for managing books and authors
- CRUD operations: Create, Read, Update, Delete
- Organized project structure for scalability
- Example usage of Go’s powerful standard library
- Simple error handling and logging
- Easy to extend and integrate

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.18 or later installed

### Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/devhimal/golang-bms.git
   cd golang-bms
   ```

2. **Install dependencies:**
   ```sh
   go mod tidy
   ```

3. **Run the application:**
   ```sh
   go run main.go
   ```

## Usage

The API will start at `http://localhost:8080` (default).  
You can interact with the API using tools like [curl](https://curl.se/) or [Postman](https://www.postman.com/).

### Example Endpoints

- `GET /books` – List all books
- `GET /books/{id}` – Get details of a book
- `POST /books` – Add a new book
- `PUT /books/{id}` – Update a book
- `DELETE /books/{id}` – Remove a book

(Adjust endpoints based on your actual implementation.)

## Project Structure

```
golang-bms/
├── main.go
├── handlers/
│   └── book_handlers.go
├── models/
│   └── book.go
├── storage/
│   └── storage.go
└── README.md
```

## Contributing

Contributions are welcome! Please open issues or submit pull requests.

## License

This project is licensed under the MIT License.

## Contact

Created by [devhimal](https://github.com/devhimal)  
Feel free to reach out with questions or suggestions.
