# go-crud
 
netstat -ano | findstr :3000
taskkill /PID 20076 /F

go get github.com/githubnemo/CompileDaemon@latest
go install github.com/githubnemo/CompileDaemon@latest

go get github.com/joho/godotenv

go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgress
go run migrate/migrate.go
netstat -ano | findstr :3000
taskkill /PID 35232 /F