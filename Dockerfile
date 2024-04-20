FROM golang

WORKDIR .

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8080

COPY . .

RUN go build -o stat4market ./cmd/main.go

CMD ["./stat4market"]