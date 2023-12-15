# Americas Tech - User API

This API has been developed for the Golang Developer position at Americas Technology.

You can find the User API built for the same test [here](https://github.com/sousair/americastech-user).

## Description

This API purpose is managing [User](https://github.com/sousair/americastech-user) cryptocurrencies (or Token) trades, with a few functionalities like create orders (market or limit), list them and cancel open orders. 

## [Postman Documentation](https://documenter.getpostman.com/view/31834520/2s9Ykn7gnf)

## How to run

### Prerequisites

- Go (version 1.21.3 or later)
- Docker & Docker Compose

### Installing

1. Clone the repository:
  ``` bash
    git clone https://github.com/sousair/americastech-exchange.git
    cd americastech-exchange
  ```

2. Create a `.env` file and fill it with the credentials.

3. Build and run the application:

``` bash
  docker compose -f build/docker-compose.yml up --build
```

## Running the tests
--
