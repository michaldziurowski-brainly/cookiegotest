# Cookie Test Server

A simple Go HTTP server that demonstrates cookie handling and token refresh functionality.

## Features
- Sets initial cookie (x-b-sess) with value "123"
- Provides a button to refresh the token
- Updates cookie value to "234" when refreshing
- Cookie expiration set to 2025-10-06T13:28:00.000Z

## Running the Server

1. Make sure you have Go installed on your system
2. Run the server:
   ```bash
   go run main.go
   ```
3. Open your browser and visit: http://localhost:8080

## Testing
1. When you first load the page, a cookie named "x-b-sess" with value "123" will be set
2. Click the "Refresh Token" button to update the cookie value to "234"
3. You can verify the cookie values in your browser's developer tools (usually under Application > Cookies) 