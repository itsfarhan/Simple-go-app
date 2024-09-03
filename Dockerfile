# Use the official Golang image to build the application
FROM golang:1.23 as base

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app with static linking to avoid dependencies on libc
RUN go build -o main .

# Final minimal image
FROM gcr.io/distroless/base

# Set the working directory inside the distroless container

# Copy the executable from the builder stage
COPY --from=base /app/main .

# Copy the HTML templates to the runtime container
COPY --from=base /app/templates /templates

# Expose port 8081
EXPOSE 8081

# Command to run the application
CMD ["./main"]
