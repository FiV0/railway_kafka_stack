FROM golang:1.22
WORKDIR /app

# Deps
COPY go.mod go.sum .
RUN go mod download

# Source code
COPY *.go .
RUN go build -o /app.out

ENTRYPOINT ["/app.out"]
