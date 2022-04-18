# Simple Ecommerce in Microservices Architetcture using GRPC

## Instalation

Clone this repo
```bash
git clone https://github.com/kbiits/simple-ecommerce-microservices-grpc
```

Go to the project root directory
```bash
cd simple-ecommerce-microservices-grpc
```

Run start script
```bash
chmod u+x ./start.sh
./start.sh
```

The API is running on localhost:3000

To stop the app, just use docker compose
```bash
docker-compose down
```

## Routes
Authentication : JWT as Bearer Token

| http method | routes                | desc                              |
| ----------- | --------------------- | --------------------------------- |
| POST        | /auth/register        | register user                     |
| POST        | /auth/login           | login user                        |
| GET         | /product              | list products (require auth)      |
| GET         | /product/{product_id} | get product by id (require auth)  |
| POST        | /product              | create product (require auth)     |
| POST        | /order                | create order (require auth)       |

## Design
![backend high level design](https://user-images.githubusercontent.com/67781184/163797975-c233c3f7-8a14-4534-ac0c-762085839cbc.png)


