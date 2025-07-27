# app/Dockerfile

FROM golang:1.24.2

WORKDIR /app

COPY go.mod go.sum ./
ENV GOPROXY=direct 
RUN go mod download

COPY . .

RUN go build -o main .

CMD ["./main"]
