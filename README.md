# octohack-back
October '18 GCP Hack Week - Back End!

# Running the server

## Local
```
go run main.go
```

## Docker
```
docker build -t octohack-back .
docker run -d -p 8080:8080 octochack-back
```

# Requests
The server will run on `http://localhost:8080/` by default.

## Hello
`GET /hello?name=<YOUR_NAME>`
```
{
    "name": <YOUR_NAME>,
    "message": "Hello <YOUR_NAME>!"
}
```
