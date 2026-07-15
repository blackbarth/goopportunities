# Projeto e API em Go



Bibliotecas

```
https://gin-gonic.com/
```

```shell
go get "github.com/gin-gonic/gin"    
```

```
go get github.com/glebarez/sqlite
```



Limpara o mod

```shell
go mod tidy
```



Swagger

https://github.com/swaggo/gin-swagger



instalar ele de forma global

```shell
go install github.com/swaggo/swag/cmd/swag@latest
```



```shell
go get -u github.com/swaggo/swag/cmd/swag
```



adicionar em routes

```go
   docs "github.com/go-project-name/docs"
   swaggerfiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
```

a justes de rotas

http://localhost:8080/swagger/index.html
