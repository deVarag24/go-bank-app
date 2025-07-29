FROM golang:1.24.2

WORKDIR /app

# Copy go.mod and go.sum from root
COPY go.mod go.sum ./

ENV GOPROXY=direct
RUN go mod download

# Copy app source code into the container
COPY ./app .

RUN go build -o main .

CMD ["./main"]
