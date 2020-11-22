# Iniciando GoLang
## Instalacion
* Una vez instalado siguiendo la [documentacion](https://golang.org/doc/install)
* Se verifica que Go ha sido instalado
  ```
  go version
  ```
  > go version go1.15.3 windows/amd64

## Codigo inicial
* Se crea la carpeta [src/hello](../src/hello)
* Se crea un archivo [hello.go](../src/hello/hello.go)
* Se utiliza `go run hello.go` para testear el funcionamiento
  > Hello World!!!
* También se utiliza `go build` desde la carpeta [src/hello](../src/hello)
* Se ejecuta el binario `./hello.exe`
  > Hello World!!!

## Call code in an external package
* Se crea un archivo [quote.go](../src/hello/quote.go)
* Se importa el package [quote](https://pkg.go.dev/rsc.io/quote) del modulo *rsc.io/quote*
* Se llama a la funcion *Go* propia del paquete
* Se crea un archivo *go.mod* a traves del comando `go mod init <module_name>`
* Se ejecuta `go run quote.go`
  > go: finding module for package rsc.io/quote
  > go: downloading rsc.io/quote v1.5.2
  > go: found rsc.io/quote in rsc.io/quote v1.5.2
  > go: downloading rsc.io/sampler v1.3.0
  > go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
  > Don't communicate by sharing memory, share memory by communicating.
