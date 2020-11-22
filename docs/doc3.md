# GoLang
## Estructuras
Estructuras como contenedor de atributos.
Se le asocian funciones a esa estructura.
Funciones dentro de la estructura (Metodos)

## Herencia
Heredar conteniendo cosas

## Actividad 3
Creación de una estructura
Asignar una función a esa estructura
    ``` js
        type person struct {
            name string
        }

        func (p person) print()  {
            fmt.Println("[person]", p.name)
        }
    ```
## Receiver
*func **(p person)** print() {}*
Se trata de la estructura (**person**) a la cual está asociada una función

Por defecto realiza una copia de p.
Anteponiendo asterisco a person, genera una copia del puntero que apunta a p, 
por lo que se tiene un nuevo puntero que apunta a la misma dirección de memoria.

## Interface
La comunidad recomienda generar interfaces de pocas funciones para ser reutilizadas fácilmente.
