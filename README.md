# Bookstore Go API

A small, modular Bookstore REST API written in Go using only the standard library. Data is stored in memory and resets when the app restarts.

Endpoints:
- GET  /books            - list all books
- GET  /books?search=Q   - search by title, author, or ISBN
- POST /books            - create a book (JSON body)
- GET  /books/{id}       - get book by id
- DELETE /books/{id}     - delete book by id
- GET  /health           - health check

Run:
```
go run .
```
The server listens on port 8080 by default.
