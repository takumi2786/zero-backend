Backend application created by golang


## migration

```bash
go install github.com/golang-migrate/migrate@latest
```

```bash
URL="mysql://root:@tcp(127.0.0.1)/zero_system?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true"
migrate -database $URL -path ./migrations up
```
