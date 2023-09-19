# Use an official Go runtime as the base image
FROM golang:1.20.4

# Set the working directory inside the container
WORKDIR /code

# Copy the entire app directory to the container
COPY ./app /code

# Download and install dependencies
RUN go mod download

# Run the application using go run
CMD ["go", "run", "main.go"]

# RUN go build 
# CMD ["./app.exe"]