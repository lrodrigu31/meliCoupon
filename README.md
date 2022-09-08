# meliCoupon

## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Documentation](#documentation)
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
* redis v3.6.4
* driver postgres v1.3.9
* gorm v1.23.8
* golang:1.19.0
* Docker
* docker-compose
* Swagger Open Api 3.0.0

## API documentation
https://app.swaggerhub.com/apis-docs/LUISRODRIGUEZ031/coupon/1.0.0

## Setup
How to run this project?

1. [Install Docker Compose](https://docs.docker.com/compose/install/)
2. Clone this repository
3. Create an .env file from .env.example changing the values ​​based on your environment
4. Run all containers with `docker-compose up`