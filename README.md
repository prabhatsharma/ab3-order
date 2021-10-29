# Order microservice

Has following endpoint following:

POST /order
{
    "product": "Football",
    "quantity": 3
}

It will send the message to in-memory queue for processing


Data generation using hey

hey -n 2000 -q 1 -m POST -T application/json -D testdata.json http://localhost:8082/order


