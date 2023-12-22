# CLIENTE SERVIDOR TCP 

El sistema **aritmética-X** permite realizar operaciones matemáticas a gran velocidad. Se debe implentar con una arquitectura cliente /  servidor.

## METADATA

- autor: damian mac dougall
- mail: <dmacdougall@teceng-gaming.com>
- company: Teceng gaming
- versión: 1.0.0

## DESCRIPCIÓN

El sistema **aritmética-X** permite realizar operaciones matemáticas a gran velocidad. Consta de dos servicios: un cliente y un servidor.

El **cliente** envia envia operaciones (por ejemplo suma)

```go
// OperationRequest representa la estructura de la solicitud del cliente
type OperationRequest struct {
	Num1 int
	Num2 int
	Op   rune //char que representa a la operación + ó -'
}
```

El Servidor analiza la request. esto implica verificar los parámetros ( que la operación exista, no sumar letras, etc..)

Una vez analizada la operación, se calcula y se genera una respuesta que contiene:

- El OperationRequest original,
- Un campo float con el resultado de la operación
- Un campo que indica el resultado de la operación (ejemplo 0 ok, 1 overflow de suma, 2 parametro de entrada invalido)

```go
// OperationResponse representa la estructura de la respuesta del servidor
type OperationResponse struct {
	Request   OperationRequest
	Result    float64
	ErrorCode int
}
```

La comunicación SE DEBE REALIZAR USANDO SOCKET TCP, NO WEBSOCKET NI REST

## Mejoras a implementar

Las mejoras a realizar se pueden hacer en cualquier orden, pero se debe documentar el avance en todas las ramas.
Si no se llega a implementar todo el desarrollo, generar un documento que indique hasta donde se alcanzó y que falta desarrollar, ejemplo actualizar el CHANGELOG.md
EJECUTAR TEST Y COVERTURA DEL CODIGO

### Mejoras Leonhard Euler : más funciones matemáticas

1. Refactorizar y cambiar el atributo **rune de OperationRequest** que es un carácter, por un type que lo represente. La idea es poder permitir más operaciones, 1 SUM | 2 MINUS | 3 DIV | 4 MULT | 5 SIN | 6 LOG |7 EXP |8 SQR
2. Las operaciones indicadas deben quedar implementas en el server
3. Se debe agregar el manejo de errores , junto con su manejo de errores . ej 10 div 0 => error no se puede dividir por cero

### Mejora George-Boole :

1. Se desea poder resolver operación lógicas **AND, OR, NOT, XOR, NAND**
2. Las operaciones indicadas deben quedar implementas en el server
3. Se debe agregar el manejo de errores , junto con su manejo de errores. ej 1 AND true => los parámetros ingresados deben ser boolean (TRUE ó FALSE)

### Mejora Claude-Elwood-Shannon:

1. Hacer un mecanismo de login entre cliente y servidor
2. El **server debe tener un registro de los clientes**, para conocer quien envía operaciones, se puede usar un archivo plano, un json, un array persistido, una base de datos, etc...
3. Los datos que se envían encriptados, un cypher entre cliente y servidor, algoritmo de clave public / privada o similar 

## Mejora Agner-Krarup-Erlang,

1. El cliente lee las operaciones desde el archivo **operation.txt** y las envia una por una al server
2. El server debe tener una bitacoria con las operaciones procesadas, se puede utilizar biblioteas externas

## Sugerencias y lineamientos del proyecto

- De ser posible crear tipos más específicos para los atributos, tratar de quitar los tipos primitivos. ejemplo:
    - la respuesta en lugar de ser float, debe ser de type OperationResult struct 
    - la operación no tiene que ser rune debe ser de type OperationFunction
- Tratar de aplicar el desarrollo con Test Driven Design. Desarrollar los test y luego codear
- Se puede refactorizar el código para facilitar el testing, además se recomienda hacer un análisis de cobertura del código
- Con respecto a **git**, no hacer un solo push con todos, mostrar los pasos intermedios. trabajar en la rama solution.
- Describir en texto y mantener el changelog
- Manejar la version de version del projecto con el esquima de golang

## Ejemplo de Uso

Luego de clonar el proyecto, entrar el directorio `Server` y arrancar el servidor y queda a la espera de mensajes

```bash
$ go run .
Servidor esperando conexiones en el puerto 8080...
```

Desde otra terminal entrar a la carpeta `Client` y arrancar el cliente, este va a enviar la operación `10 + 5` al server ni bien arranca

```bash
$ go run .
Resultado de 10 + 5 = 15.00
```
