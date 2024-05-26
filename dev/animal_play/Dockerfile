# Currently not used.
# Dockerfile will be used after docker-compose.yml is setup

FROM golang:1.22-alpine

WORKDIR /app

# Caching dependencies 
COPY golang/go.mod golang/go.sum ./

RUN go mod download

COPY golang .

# WORKDIR /cmd

RUN go build -o main ./cmd/main.go

EXPOSE 8080
 
CMD [ “./main” ]