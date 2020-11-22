# GoLang - Actividad de introducción al lenguaje
## Entrypoint
Funcion main del package main
```js
    package main

    func main() {

    }
```
## Funciones
Es posible retornar más de un valor, por ej. *(int, error)*, devuelve una tupla.
    ```js
        func suma(a, b int) (int, error) {
            if a < b {
                return 0, errors.New("el primer valor es menor que el segundo valor")
            }
            return a + b, nil
        }
    ```
Buena práctica realizar funciones genéricas que otro puede reutilizar.

### private / public
Función en minúscula indica nivel package
    ```js
        func suma(a, b int) (int, error) {
            if a < b {
                return 0, errors.New("el primer valor es menor que el segundo valor")
            }
            return a + b, nil
        }
    ```
Función en Mayúscula indica nivel público (y requiere documentación)
    ```js
        //Documentacion de la funcion Suma
        func Suma(a, b int) (int, error) {
            if a < b {
                return 0, errors.New("el primer valor es menor que el segundo valor")
            }
            return a + b, nil
        }
    ```
### Manejo de retorno de valores
    ```js
        r, err := suma(1, 2)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(r)
        }
    ```
## Librerias/ Paquetes
La comunidad utiliza las que vienen en el lenguaje (no tanto las externas) 
## Manejo de errores
Distinto a otros lenguajes(prestar atención).
Se tratan los errores como parte del ciclo de vidad normal de un programa.
Es decir, no se trata como excepción sino como parte de la regla en la lógica de negocio del software.
    ```js
        if error != nil
    ```
### Mensaje de error
Es buena práctica escribir los mensajes comenzando con minúscula y no termine con punto.
Porque se pueden concatenar con otros mensajes.
## Lenguaje restrictivo
### Tiempo de compilación y estructura del código
En GoLang se destaca la importancia que otorga el lenguaje al tiempo de compilación y la estructura del código.
Por ello, es un lenguaje restrictivo en cuanto a estructura del código.
### Tratamiento de errores fuera de la función
Es buena práctica realizar el tratamiento de errores fuera de la función que lo retorna.
## Variables
No es posible tener varibles declaradas sin utilizar, no permite la compilación.
*:=* infiere el tipo