# meliCoupon

## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Documentation](#documentation)
* [References](#references)
* [Setup](#setup)

## General info
Mercado Libre is implementing a new benefit for users who most use the
platform with a free coupon of a certain amount that will allow them to buy as many items marked as favorites that do not exceed the total amount.

To solve this challenge, it was necessary to implement a REST API based on the backpack algorithm and using recursive functions, it returns the list of products that can be purchased and the total amount of these products.

## Technologies
Project is created with:
* playground v9.31.0
* gorilla mux v1.8.0
* godotenv v1.4.0
* go-cache v2.1.0
* golang:1.19.0
* Docker
* docker-compose
* Swagger Open Api 3.0.0
* Heroku

## API documentation
Swagger documentation: 
* https://app.swaggerhub.com/apis-docs/LUISRODRIGUEZ031/coupon/1.0.0

Postman Documentation:

* [`postman json`](meliCoupon.postman_collection.json)

## References
* https://www.youtube.com/watch?v=vdVpRjO7g84&t=181s&ab_channel=UCAMUniversidadCat%C3%B3licadeMurcia
* https://www.youtube.com/watch?v=dHca8GaINXM&ab_channel=OCILabs
* https://www.youtube.com/watch?v=EH6h7WA7sDw&ab_channel=Mindez
* https://www.youtube.com/watch?v=Gd3Zy9fb0hU&ab_channel=Matem%C3%A1ticasParaNi%C3%B1osGeniales
* https://www.youtube.com/watch?v=W_KoJmM3f2Q&ab_channel=CamilaVertiz
* https://github.com/FabianNorbertoEscobar/Grafos
* https://github.com/mattschofield/go-knapsack/blob/aaf424030800/knapsack.go
* https://parzibyte.me/blog/2018/11/01/golang-algoritmo-busqueda-binaria/
* https://pkg.go.dev/github.com/mattschofield/go-knapsack#Knapsack

## Setup
How to run this project?
1. [Install Docker Compose](https://docs.docker.com/compose/install/)
2. Clone this repository
3. Create an .env file from .env.example changing the values ​​based on your environment
4. Run all containers with `docker-compose up`