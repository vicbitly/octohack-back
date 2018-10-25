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

## User
`GET /users/:username`
```
{
    "id": 1,
    "username": "user",
    "email": "user@bit.ly"
}
```

`POST /users` with:
```
{
    "username": "user",
    "email": "user@bit.ly"
}
```
