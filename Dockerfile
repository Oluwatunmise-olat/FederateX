FROM golang:1.26-alpine

RUN apk add --no-cache iproute2 iptables

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o /tun-playground ./cmd/playground

CMD [ "/tun-playground" ]