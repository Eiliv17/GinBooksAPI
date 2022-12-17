# GinBooksAPI

This is an example of a bookshelf REST API that provides book data and performs CRUD operations.

Here are the available endpoints for now:

/books
- GET - Get a list of all the books, returned as JSON.
- POST - Add a new book from request data sent as JSON.

/books/:id
- GET - Get a book from its ID, returning the book data as JSON.

## Setup

By default you'll need to provide a .env file or set up the environment variables for your operating system, these are the required variables that you need to setup:

Variable        | Description
---             | ---
PORT            | The HTTP server port
MONGODB_URI     | The URI for connecting to your MongoDB server
DB_NAME         | The name of the database (inside MongoDB) where you want to store the books data.
COLL_NAME       | The name of the collection where you want to store the books data.


Here's an example of the .env file:
```
PORT=3000
MONGODB_URI=mongodb://127.0.0.1:27017/?maxPoolSize=20&w=majority
DB_NAME=bookshelf
COLL_NAME=books
```

## Book structure

Here's a sample of the basic structure of a book:
```json
{
    "title":"Build Systems With Go: Everything a Gopher must know",
    "authors":["Juan M. Tirado"],
    "publicationDate": "2021-05-10",
    "publisher":"Independently published",
    "language":"English",
    "isbn13":"979-8502040150"
}
```

There is also a small script inside the **populate** folder where you can import some books as a sample into MongoDB.


## HTTP Requests Examples

Here you can find some examples of the possible HTTP requests and responses you can perform.

### Get all the books

#### Request
**GET** `http://localhost:3000/books`


#### Response
```json
[
    {
        "title": "Build Systems With Go: Everything a Gopher must know",
        "ID": "639e0300a4e6eb079a4ed7ac",
        "authors": [
            "Juan M. Tirado"
        ],
        "publicationDate": "2021-05-10T00:00:00Z",
        "publisher": "Independently published",
        "language": "English",
        "isbn13": "979-8502040150"
    },
    {
        "title": "Learning Go: An Idiomatic Approach to Real-World Go Programming",
        "ID": "639e0300a4e6eb079a4ed7ad",
        "authors": [
            "Jon Bodner"
        ],
        "publicationDate": "2021-04-06T00:00:00Z",
        "publisher": "O'Reilly Media",
        "language": "English",
        "isbn13": "978-1492077213"
    }
]
```

### Get a specific book

#### Request
**GET** `localhost:3000/books/639e0300a4e6eb079a4ed7ad`


#### Response
```json
{
    "title": "Learning Go: An Idiomatic Approach to Real-World Go Programming",
    "ID": "639e0300a4e6eb079a4ed7ad",
    "authors": [
        "Jon Bodner"
    ],
    "publicationDate": "2021-04-06T00:00:00Z",
    "publisher": "O'Reilly Media",
    "language": "English",
    "isbn13": "978-1492077213"
}
```

### Create a book

#### Request
**POST** `127.0.0.1:3000/books`
```json
{
    "title": "My Book",
    "authors": [ "Myself" ],
    "publicationdate": "2019-01-12",
    "publisher": "Independently published",
    "language": "English",
    "isbn13": "924-0434291440"
}
```

#### Response
```json
{
  "result": "book inserted successfully"
}
```