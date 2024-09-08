# Use a Go image to build the application
FROM golang:1.21.6 AS builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download
ARG ENV_FILE

# Copy the entire application source code
COPY . .

COPY ${ENV_FILE} ./

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o todo-api ./cmd/api/main.go

# Use a minimal image to run the compiled Go binary
FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=builder /app/todo-api/ .
# RUN chmod +x /app/todo-api

ARG ENV_FILE
COPY --from=builder /app/${ENV_FILE} ${ENV_FILE}

# Set the application to listen on port 8080
EXPOSE 8080

# Command to run the API
CMD ["./todo-api"]
