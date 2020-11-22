# GoLang - Estructura de datos
## Array
Arreglo de tamaño fijo.
    ```js
        var arr [2]int
        arr[0] = 1
        arr[1] = 2
    ```

## Slice
Similar a ArrayList de Java.
Estructura que contiene tres propiedades: 
1 - Array (arreglo de tamaño fijo).
2 - Capacity (capacidad máxima de un slice)
3 - Length (cuánto se está ocupando de la capacidad)

Contiene un arreglo de slices.
    ```js
        var l []int
        l = append(l, 10)
    ```

## Map
Similiar al HashMap de Java