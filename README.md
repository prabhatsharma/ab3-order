# Order microservice

Has following endpoint following:

POST /order
{
    "product": "Football",
    "quantity": 3
}

It will send the message to in-memory queue for processing


Add orders

hey -n 2000 -q 1 -c 1 -m POST -T application/json -D testdata.json http://localhost:8082/order


Read from queue

hey -n 20 -q 1 -c 1 -m GET -T application/json -D testdata.json http://localhost:8080/queue/Swimming
