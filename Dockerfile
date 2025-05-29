# syntax=docker/dockerfile:1

FROM golang:1.23 AS dev

WORKDIR /app

# Install air for hot reload
RUN go install github.com/air-verse/air@latest

# Copy go mod files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy the full source code
COPY . .

EXPOSE 8080

# Run Air (auto reload on file changes)
CMD ["air"]
