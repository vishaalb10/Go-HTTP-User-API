# Go HTTP User API

This is a small program I wrote while learning Go networking and backend development.

I'm trying to understand how things actually work behind the scenes, especially how an HTTP server receives requests, processes data, and sends responses back to clients.

It's a simple implementation, and there's still a lot for me to learn. I built it mainly to get some hands-on experience with Go's standard library and HTTP communication.

## What This Program Does

- Creates a user.
- Retrieves a user using their ID.
- Deletes a user using their ID.
- Handles JSON request and response data.
- Uses mutexes to manage concurrent access to shared data.

User information is stored in a Go map, so the data is temporary and disappears when the server stops.

## Packages Used

| Package | Purpose |
|---|---|
| `net/http` | Creating the HTTP server, registering routes, and handling requests and responses. |
| `encoding/json` | Encoding and decoding JSON data. |
| `fmt` | Printing messages and writing responses. |
| `strconv` | Converting URL IDs from strings into integers. |
| `sync` | Using `sync.RWMutex` to safely access shared data. |

## Functions and Concepts

- **`main()`** – Creates the HTTP server, registers routes, and starts listening on port 8080.
- **`hellohandler()`** – Handles requests to the root endpoint and returns a welcome message.
- **`createUser()`** – Receives JSON data, validates the name, and stores a new user.
- **`getUser()`** – Retrieves a user using their ID and returns the information in JSON format.
- **`deleteUser()`** – Removes a user from the map using their ID.

I also used `sync.RWMutex` to understand how Go handles concurrent access to shared data.

## API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| GET | `/` | Welcome message |
| POST | `/users` | Create a user |
| GET | `/users/{id}` | Retrieve a user |
| DELETE | `/users/{id}` | Delete a user |

## Testing with Postman

I used **Postman** to send HTTP requests and check how the server responds.

For creating a user, I used a POST request with a JSON body:

```json
{
    "name": "John"
}
```

For retrieving and deleting users, I used GET and DELETE requests respectively.

## How to Run

**Requirements:**
- Go 1.22 or newer.
- Postman (for testing).

Clone the repository:

```bash
git clone <your-repository-url>
```

Navigate to the directory:

```bash
cd <directory-name>
```

Run the program:

```bash
go run main.go
```

The server will be available at:

```text
http://localhost:8080
```

You can now test the endpoints using Postman.

## A Little Note

This is just a small attempt from my side to understand networking and backend development using Go.

I'm still learning, experimenting, making mistakes, and trying to understand how each piece works rather than just getting the code to run.

There's definitely more to explore, and this is one of my small steps along the way. if there is any suggestions plz let me know
