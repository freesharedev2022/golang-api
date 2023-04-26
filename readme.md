golang version: 1.18 

nếu chưa tạo file go.mod thì có thể tạo bằng cách:
tạo file .mod :
```
go mod init golang-api
```

cài đặt thư viện: 
```
go mod tidy
```

run app
```
go run main.go
```

docs API
```
http://localhost:8000/swagger/index.html
```