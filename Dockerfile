# Use the official Golang image as a base
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project to the working directory inside the container
COPY . .

# Build the app
RUN go build -o main .

# Exposes the port on which the app runs
EXPOSE 8080

# Run the executable
CMD ["./main"]

# Docker object labels
LABEL maintainer="Elhadj Abdoul Diallo <elhhabdouldiallo@gmail.com>"
LABEL version="1.0"
LABEL description="Dockerized ascii-art-web which is a simple web application that allows users to generate ASCII art from text input using various banner styles."
