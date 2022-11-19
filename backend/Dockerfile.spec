# syntax=docker/dockerfile:1

FROM swaggerapi/swagger-ui:v4.15.5

COPY ./docs/openapi.yaml /spec/openapi.yaml

EXPOSE 8080
