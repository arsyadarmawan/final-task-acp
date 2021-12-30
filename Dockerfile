FROM golang:1.17-alpine

WORKDIR ./Sites/acp141
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o acp14 .

CMD ["./Sites/acp141"]     