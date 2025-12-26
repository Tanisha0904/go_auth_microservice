<img width="354" height="195" alt="{DC5E4409-48DE-4B2D-B99F-A56DA14B00E2}" src="https://github.com/user-attachments/assets/dbeec686-f0c8-4b27-a94a-42e80bcaebbd" />

### How to run:
> go mod tidy
> go run cmd/main.go
### POST: localhost:8080/login
request: {
    "username": "",
    "password": ""
}
response: {"token":""}
### GET: localhost:8080/api/home
request: add the token in Bearer Token for authorization
response: {
    "message": "Welcome to Home",
    "user": ""
}
