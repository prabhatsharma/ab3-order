#!/bin/bash

hey -n 100 -c 1 -q 1 -m POST -T application/json -D testdata.json http://localhost:8082/order
