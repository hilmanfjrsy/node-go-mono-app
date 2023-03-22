# Start with the official Golang image as a base
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o fetch-app .

# Expose the port that the application will listen on
EXPOSE 5002

# Start the application when the container starts
CMD ["./fetch-app"]
