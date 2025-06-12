# Pós Go-Expert

## Desafio Clean Architecture
Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

## Para executar o projeto

1. Subindo os containers
``` shell
make up

```

2. Testes: REST API server
- faça uma chamada POST para criar uma nova order via rest-client usando o arquivo api/api.http no vscode
- faça uma chamada GET para listar as orders via rest-client usando o arquivo api/api.http no vscode

3. Testes: gRPC server
``` shell
## crie uma nova order
evans --proto internal/infra/grpc/protofiles/order.proto --host localhost --port 8082
=> call CreateOrder
=> 2
=> 10.5
=> 0.5

## liste as orders
evans --proto internal/infra/grpc/protofiles/order.proto --host localhost --port 8082
=> call ListOrders
```

4. Testes: GraphQL server

GraphQL Playground: http://localhost:8081
``` shell
## crie uma nova order
mutation createOrder {
  createOrder(input: {id:"3", Price: 10.5, Tax: 0.5}) {
    id
    Price
    Tax
  }
}

## liste as orders
query queryOrders {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

5. Baixando os containers 
``` shell
make down

```
