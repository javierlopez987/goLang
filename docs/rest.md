# GoLang
## REST
Libreria Gin-Gonic
Request-Response, Contexto
Cambia como se escribe el Handler respecto de otras librerias
## GinGonic
[Github](https://github.com/gin-gonic/gin)
### Installation
```js
    $ go get -u github.com/gin-gonic/gin
```
Luego, importarlo en el código
```js
    import "github.com/gin-gonic/gin"
```
Luego, importar *net/http*
```js
    import "net/http"
```
```js
    package main

    import "github.com/gin-gonic/gin"

    func main() {
        r := gin.Default()
        r.GET("/ping", func(c *gin.Context) {
            c.JSON(200, gin.H{
                "message": "pong",
            })
        })
        r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
    }
```

### Registrar Handlers
Indicar el router quien maneja temas de seguridad, interceptores, etc.
Usamos *gin.Default()*

Registrar handler pasando un puntero de *gin.Context* por parametro

### Ejemplo basico
```js
    package main

    import "github.com/gin-gonic/gin"

    func main()  {
        r := gin.Default()
        r.GET("/users", getUsersHandler)
        r.Run()
    }

    func getUsersHandler(c *gin.Context)  {
        
        c.JSON(200, gin.H{
            "status": "ok",
        })
    }
```
