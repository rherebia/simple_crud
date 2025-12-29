# Simple CRUD API

This project is a simple CRUD (Create, Read, Update, Delete) API for managing a collection of vintage vinyl albums. It is built using Go and the [Gin Web Framework](https://github.com/gin-gonic/gin).

## Prerequisites

- **Go**: Version 1.16 or later.
- **Curl**: A tool for transferring data with URLs (often pre-installed on macOS/Linux).

## Installation

1.  Clone the repository or navigate to the project directory.
2.  Install dependencies:

    ```bash
    go mod tidy
    ```

## Running the Application

Start the server using the following command:

```bash
go run .
```

The server will start on `localhost:8080`.

## API Endpoints

The API provides the following endpoints:

### 1. Get All Albums
Retrieves a list of all available albums.

- **URL**: `/albums`
- **Method**: `GET`
- **Response**: JSON array of albums.

**Example Request:**
```bash
curl http://localhost:8080/albums
```

### 2. Get Album by ID
Retrieves details of a specific album by its ID.

- **URL**: `/albums/:id`
- **Method**: `GET`
- **Response**: JSON object of the album.

**Example Request:**
```bash
curl http://localhost:8080/albums/2
```

### 3. Add a New Album
Adds a new album to the collection.

- **URL**: `/albums`
- **Method**: `POST`
- **Body**: JSON object representing the new album.

**Example Request:**
```bash
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

### 4. Update an Album
Updates an existing album's details.

- **URL**: `/albums/:id`
- **Method**: `PUT`
- **Body**: JSON object with updated album data.

**Example Request:**
```bash
curl http://localhost:8080/albums/1 \
    --include \
    --header "Content-Type: application/json" \
    --request "PUT" \
    --data '{"id": "1","title": "Blue Train (Remastered)","artist": "John Coltrane","price": 59.99}'
```

### 5. Delete an Album
Removes an album from the collection.

- **URL**: `/albums/:id`
- **Method**: `DELETE`

**Example Request:**
```bash
curl -X DELETE http://localhost:8080/albums/1
```

## Data Structure

Each album has the following structure:

```json
{
  "id": "string",
  "title": "string",
  "artist": "string",
  "price": number
}
```