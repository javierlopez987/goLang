# GoLang
## Project Layout
[Project layout](github.com/golang-standards/project-layout)

### Servicios de configuracion
Se accede con usuario y clave y se obtiene el set de configuracion de determinado ambiente, por ejemplo.

### Service
Crea una variable del tipo service y devolverla como si fuese su interface ChatService

La struct service implementa las tres funciones de la interface ChatService

Injeccion a chat service una configuracion
```js
    service, _ := chat.New(cfg)
```
utilizo New como FactoryMethod y polimorfismo
```js
    func New(c *config.Config) (ChatService, error) {
	    return service{c}, nil
    }
```