# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

- Mejoras Leonhard Euler : más funciones matemáticas

  - Refactorizar y cambiar el atributo **rune de OperationRequest** que es un carácter, por un type que lo represente. La idea es poder permitir más operaciones, 1 SUM | 2 MINUS | 3 DIV | 4 MULT | 5 SIN | 6 LOG |7 EXP |8 SQR
  - Las operaciones indicadas deben quedar implementas en el server
  - Se debe agregar el manejo de errores , junto con su manejo de errores . ej 10 div 0 => error no se puede dividir por cero

- Mejora George-Boole
  - Se desea poder resolver operación lógicas **AND, OR, NOT, XOR, NAND**
  - Las operaciones indicadas deben quedar implementas en el server
  - Se debe agregar el manejo de errores , junto con su manejo de errores. ej 1 AND true => los parámetros ingresados deben ser boolean (TRUE ó FALSE)

- Mejora Claude-Elwood-Shannon:
  - Hacer un mecanismo de login entre cliente y servidor
  - El **server debe tener un registro de los clientes**, para conocer quien envía operaciones, se puede usar un archivo plano, un json, un array persistido, una base de datos, etc...
  - Los datos que se envían encriptados, un cypher entre cliente y servidor, algoritmo de clave public / privada o similar 

- Mejora Agner-Krarup-Erlang,

  - El cliente lee las operaciones desde el archivo **operation.txt** y las envia una por una al server
  - El server debe tener una bitacoria con las operaciones procesadas, se puede utilizar biblioteas externas

## [0.0.3] - 2024-04-02

### Removed

### Changed
La implementación de este programa funciona con una base de datos. Si bien vienen integradas las migraciones, les dejo los script para la reacrion de las tablas, dado que como en mi caso, las migraciones no funcionaban como lo esperado.

Conexion a la DB var dsn = "root:root@tcp(localhost:3306)/aritmetica?charset=utf8mb4&parseTime=True&loc=Local"

Scripts para la creación de las tablas:
MySQL, gestor DBeaver
tabla user: 
CREATE TABLE users (
    user_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    num1 DOUBLE,  
    num2 DOUBLE, 
    op int,
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    operations VARCHAR(255)
);

tabla user_operations
CREATE TABLE user_operations (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT UNSIGNED NOT NULL,
    operation VARCHAR(100) NOT NULL,
    result DOUBLE NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

Server
  Función main(): 

Esta es la función principal del programa.
Se encarga de establecer una conexión TCP con el servidor y enviar solicitudes de operación.
Proporciona un menú para que el usuario seleccione el tipo de operación a realizar y solicite los números necesarios según la operación seleccionada.
Maneja las respuestas del servidor e imprime los resultados o mensajes de error.
Estructuras OperationRequest y OperationResponse:

OperationRequest representa la solicitud del cliente al servidor e incluye los números a operar y el tipo de operación.
OperationResponse representa la respuesta del servidor e incluye el resultado de la operación y un código de error si corresponde.

CreateUser. Representa una estructura para que el usuario se pueda registrar y loguear
UserOperation. Contiene el tipo de operacion que hace el usuario y su resultado.

Estas 2 estructuras a su vez estan representadas en los modelos para la interacción con la base de datos

Función operationSymbol(op OperationType) string:

Devuelve el símbolo correspondiente a cada tipo de operación, utilizado para mostrar la operación realizada en los resultados.

Client
Estructuras y constantes:
OperationType: Es un tipo enumerado que define los diferentes tipos de operaciones disponibles, como suma, resta, división, multiplicación, seno, logaritmo, exponencial y raíz cuadrada.

OperationRequest: Estructura que representa la solicitud del cliente al servidor, incluyendo los números a operar y el tipo de operación.

OperationResponse: Estructura que representa la respuesta del servidor, incluyendo el resultado de la operación y un código de error si corresponde.

CreateUser. Representa una estructura para que el usuario se pueda registrar y loguear
UserOperation. Contiene el tipo de operacion que hace el usuario y su resultado

Funciones:
operationSymbol(op OperationType) string: Esta función toma un tipo de operación como entrada y devuelve el símbolo correspondiente de la operación, utilizado para mostrar la operación realizada en los resultados.

//Funcionalidades de OPERADORES LOGICOS AND, OR, NOT, XOR, NAND
Quedo implementado en la version 0.0.2

//Funcionalidades de registro
Se implemento una base de datos y un registro de usuario y contraseña
Ahora se almacena cada operacion que realiza el usuario. Evita usuarios repetidos
Se almacenan las operaciones que hace el usuario. El register y login estan implementados en la misma funcion. El sistema detecta si el usuario ya esta o no registrado
### Fixed

### Added

- Begin a project
- Create basic server and client