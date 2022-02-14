# Go- catApi

## Motivo de la aplicación

Es un simple programa creado en Go que emplea el API de [cat facts](https://alexwohlbruck.github.io/cat-facts/) para mostrar los datos según los parámetros en la línea de comando.

Este programa fue creado sólamente para entender los siguientes conceptos:

* Uso de flags sin librerías de terceros
* Consumo de Rest Apis
* Comprender como activar los errores que puedan producirse al usar APIS


Para compilar el programa usar: 

`go build`

## Parametrización

Por defecto el programa sólo muestra una información si no hay ningún parametro.

* Ejemplo de uso sin compilación previa:

`go run gurl`

* Ejemplo de uso con 3 resultados:

`go run gurl -c 3`

