Backend application created by golang

## run

```bash
make docker/build
make docker/run
```

```bash
curl -X POST \
-H "Content-Type: application/json" \
-d '{"title":"焼肉について","content":"焼肉は美味しい"}' \
localhost:8080/posts

curl -X GET \
localhost:8080/posts

curl -X POST \
-H "Content-Type: application/json" \
-d '{
	"identifier": "mail@example.com",
	"credencial": "password"
}' \
localhost:8081/login
```

## migration

```bash
go install github.com/golang-migrate/migrate@latest
```

```bash
URL="mysql://root:@tcp(127.0.0.1)/zero_system?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true"
migrate -database $URL -path ./migrations up
```

```bash
URL="mysql://root:@tcp(127.0.0.1)/zero_system?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true"
migrate -database $URL -path ./migrations down
```

## 初期ユーザー

```sql
INSERT INTO `user`
    
```
