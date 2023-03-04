# Use the official Go image as the base image
FROM golang:1.15

# Set the working directory in the container
WORKDIR /app

# Copy the application files into the working directory
COPY . /app

# Build the application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Define the entry point for the container
CMD ["./main"]
