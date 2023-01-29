FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

ENV CONFIG_FILE_PATH=/app/config/config.yaml

RUN go build -o /toronto-pizza ./cmd/toronto-pizza

EXPOSE 3000

CMD [ "/toronto-pizza" ]
