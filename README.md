# Go TCP Web Server

A custom HTTP web server built from scratch using Go's raw TCP sockets. This project demonstrates how to create a web server without using any HTTP framework, handling raw socket connections and manually parsing HTTP requests.

## 🚀 Features

- **Raw TCP Socket Implementation**: Built using Go's `net` package for direct socket handling
- **Manual HTTP Parsing**: Custom HTTP request parsing without external libraries
- **Static File Serving**: Serves HTML files from the `public` directory
- **Concurrent Request Handling**: Uses goroutines for handling multiple client connections
- **Keep-Alive Connections**: Supports persistent connections with timeout management
- **Error Handling**: Custom 404 error pages for missing files

## 📁 Project Structure

```
golang/
├── go.mod                 # Go module definition
├── main.go               # Main server entry point
├── server/
│   └── server.go         # HTTP request parsing logic
├── public/               # Static web files
│   ├── home.html         # Homepage with server status
│   ├── error.html        # Error page template
│   ├── contact/
│   │   └── contact.html  # Contact page
│   └── parser.go         # File serving logic
└── README.md            # This file
```

## 🛠️ Prerequisites

- **Go 1.24.4** or later
- Basic understanding of TCP sockets and HTTP protocol

## 🚀 Quick Start

### 1. Run the Server

```bash
go run main.go
```

The server will start and listen on `localhost:3000`

### 3. Access the Website

Open your browser and navigate to:

- **Homepage**: http://localhost:3000/home.html
- **Contact Page**: http://localhost:3000/contact/contact.html

## 📋 Available Endpoints

| Endpoint                | Description                 | File Served                   |
| ----------------------- | --------------------------- | ----------------------------- |
| `/home.html`            | Homepage with server status | `public/home.html`            |
| `/contact/contact.html` | Contact page                | `public/contact/contact.html` |
| `/error.html`           | Error page                  | `public/error.html`           |
| Any other path          | 404 Not Found               | Custom error message          |

## 🔧 How It Works

### Server Architecture

1. **TCP Listener**: Creates a TCP listener on port 3000
2. **Connection Handling**: Accepts incoming connections and spawns goroutines
3. **Request Parsing**: Manually parses HTTP requests to extract the requested path
4. **File Serving**: Maps paths to files in the `public` directory
5. **Response Generation**: Crafts HTTP responses with proper headers

### Key Components

#### `main.go`

- Sets up TCP listener on `localhost:3000`
- Handles incoming connections with goroutines
- Manages connection timeouts (30 seconds)
- Implements keep-alive connections

#### `server/server.go`

- `ParsePathName()`: Extracts the requested path from HTTP request
- Parses raw HTTP request bytes to find the path between `/` and `HTTP`

#### `public/parser.go`

- `ParsePublic()`: Maps requested paths to actual files
- Handles directory traversal and file serving
- Returns 404 errors for missing files
- Case-insensitive file matching

## 🧪 Testing

### Manual Testing

1. Start the server: `go run main.go`
2. Use curl to test endpoints:

```bash
# Test homepage
curl http://localhost:3000/home.html

# Test contact page
curl http://localhost:3000/contact/contact.html

# Test 404
curl http://localhost:3000/nonexistent.html
```

### Browser Testing

- Open http://localhost:3000/home.html in your browser
- The homepage includes a beautiful dashboard showing server status
- Test navigation between different pages

## 🎯 Example Usage

### Starting the Server

```bash
$ go run main.go
Received New User 127.0.0.1:54321
```

### Making Requests

```bash
$ curl -v http://localhost:3000/home.html
* Connected to localhost (127.0.0.1) port 3000
> GET /home.html HTTP/1.1
> Host: localhost:3000
>
< HTTP/1.1 200 OK
< Content-Type: text/html
< Content-Length: 12345
< Connection: keep-alive
< Keep-Alive: timeout=30, max=100
```

## 🔍 Code Examples

### Adding New Pages

1. Create HTML file in `public/` directory:

```html
<h1>New Page</h1>
```

## 🚨 Error Handling

- **Connection Errors**: Server continues running if individual connections fail
- **File Not Found**: Returns custom 404 messages for missing files
- **Timeout Handling**: Connections timeout after 30 seconds of inactivity
- **Directory Traversal**: Safely handles requests for non-existent directories

## 🎨 Customization

### Changing Port

Edit `main.go` line 37:

```go
conn, err := net.Listen("tcp", "localhost:3000") // Change port here
```

### Adding Static Assets

1. Create directories in `public/` as needed
2. Add HTML, CSS, JS files
3. Access via URL path matching directory structure

### Modifying Response Headers

Edit `main.go` line 28 to customize HTTP response headers:

```go
ctx := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nCustom-Header: value\r\n\r\n%s", len(pageFile), pageFile)
```

## 📚 Learning Resources

This project demonstrates:

- Raw TCP socket programming in Go
- Manual HTTP protocol implementation
- Concurrent programming with goroutines
- File system operations
- Error handling patterns


## 📄 License

This project is for educational purposes. Feel free to use and modify as needed.

---

**Happy Coding! 🚀**

_Built with ❤️ using Go's raw TCP sockets_
