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

## [0.0.0] - 2023-03-05

### Removed

### Changed

### Fixed

### Added

- Begin a project
- Create basic server and client