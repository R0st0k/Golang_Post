#!/bin/bash

mongoimport --db post --collection employees --file /samples/employees.json --jsonArray && \
mongoimport --db post --collection postOffices --file /samples/post_offices.json --jsonArray && \
mongoimport --db post --collection sendings --file /samples/sendings.json --jsonArray