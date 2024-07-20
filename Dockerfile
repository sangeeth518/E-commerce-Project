FROM golang:1.21.3-alpine3.18 AS build-stage


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
# COPY go.mod go.sum ./
COPY ./ /app

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM gcr.io/distroless/static-debian11



# Copy the Pre-built binary file from the previous stage
COPY --from=build-stage /app/main .
COPY .env .env

# Expose port 8080 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./main"]
