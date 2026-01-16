# Simple CRUD API

This project is a simple CRUD (Create, Read, Update, Delete) API for managing a collection of vintage vinyl albums. It is built using Go and the [Gin Web Framework](https://github.com/gin-gonic/gin). The project folder hierarchy is inspired by the recommended [GoLang community layout](https://github.com/golang-standards/project-layout).

## Prerequisites

- **Go**: Version 1.16 or later.
- **Curl**: A tool for transferring data with URLs (often pre-installed on macOS/Linux).

## Installation

1.  Clone the repository or navigate to the project directory.
2.  Install dependencies:

    ```bash
    go mod tidy
    ```

## Running the Application {#running-application}

Start the server using the following command:

```bash
go run cmd/api/main.go
```

The server will start on `localhost:8080`.

### Environment Variables

There is a configuration expecting the definition of an environment variable when instantiating the album repository. The variable STORE_TYPE could be defined using the .env file and might be defined with SQLITE value to activate the respective repository. In absense of that variable, the implementation will fallback to use json in memory repository.

### Makefile

Another way to run the project is using the Makefile tasks definition. The following ones are available:

- just `make` or `make run`: executes the same command explained before inside [Running the Application](#running-application);
- `make build`: creates the binary file inside bin folder;
- `make test`: runs tests in all project subfolders

## API Endpoints

The API provides the following endpoints:

### 1. Get All Albums
Retrieves a list of all available albums.

- **URL**: `/v1/albums`
- **Method**: `GET`
- **Response**: JSON array of albums.

**Example Request:**
```bash
curl http://localhost:8080/v1/albums
```

### 2. Get Album by ID
Retrieves details of a specific album by its ID.

- **URL**: `/v1/albums/:id`
- **Method**: `GET`
- **Response**: JSON object of the album.

**Example Request:**
```bash
curl http://localhost:8080/v1/albums/2
```

### 3. Add a New Album
Adds a new album to the collection.

- **URL**: `/v1/albums`
- **Method**: `POST`
- **Body**: JSON object representing the new album.

**Example Request:**
```bash
curl http://localhost:8080/v1/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

### 4. Update an Album
Updates an existing album's details.

- **URL**: `/v1/albums/:id`
- **Method**: `PUT`
- **Body**: JSON object with updated album data.

**Example Request:**
```bash
curl http://localhost:8080/v1/albums/1 \
    --include \
    --header "Content-Type: application/json" \
    --request "PUT" \
    --data '{"id": "1","title": "Blue Train (Remastered)","artist": "John Coltrane","price": 59.99}'
```

### 5. Delete an Album
Removes an album from the collection.

- **URL**: `/v1/albums/:id`
- **Method**: `DELETE`

**Example Request:**
```bash
curl -X DELETE http://localhost:8080/v1/albums/1
```

### API test calls

All the api calls is defined inside api-test directory using .http files with plain text declaration. IDEs like VS Code (with REST Client Extension) and JetBrains (with HTTP Client support) offer simple way to test the api endpoints.

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