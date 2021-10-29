#!/bin/bash

VERSION=v4
docker build . -t ab3-order:$VERSION -t ab3-order:latest 

docker tag ab3-order:latest 525158249545.dkr.ecr.us-west-2.amazonaws.com/ab3-order:latest
docker tag ab3-order:latest 525158249545.dkr.ecr.us-west-2.amazonaws.com/ab3-order:$VERSION 

aws ecr get-login-password --region us-west-2 | docker login --username AWS --password-stdin 525158249545.dkr.ecr.us-west-2.amazonaws.com

docker push 525158249545.dkr.ecr.us-west-2.amazonaws.com/ab3-order:latest
docker push 525158249545.dkr.ecr.us-west-2.amazonaws.com/ab3-order:$VERSION 


