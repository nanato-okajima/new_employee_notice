FROM golang:1.18-alpine3.16

WORKDIR /go/app/

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@v1.40.4

CMD [ "air", "-c", ".air.toml" ]
