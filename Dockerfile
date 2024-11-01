# Use an official Golang runtime as a parent image
FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -gcflags "all=-N -l" -o main .

# Expose ports for the app and the debugger
EXPOSE 8089
ENV PORT 8089

# Command to run the executable
CMD ["./main"]
