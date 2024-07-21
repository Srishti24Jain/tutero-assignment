# Use an official Golang runtime as a parent image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy the current directory contents into the container at /usr/src/app
COPY . .

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
