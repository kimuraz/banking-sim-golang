FROM golang:1.22-alpine

RUN apk add --no-cache gcc g++ git openssh-client

WORKDIR /go/src/app

COPY . .

ENV CGO_ENABLED=1

ENV GIN_MODE=release

RUN go build

CMD ["./banking_sim"]

EXPOSE "8080:8080"

