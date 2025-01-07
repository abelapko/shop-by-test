FROM golang:1.23-alpine

# Install required tools
RUN apk add --no-cache curl bash mysql-client

# Copy the wait-for-mysql.sh script to the container
COPY wait-for-mysql.sh /usr/local/bin/wait-for-mysql.sh
RUN chmod +x /usr/local/bin/wait-for-mysql.sh

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Install dependencies
RUN go mod tidy

# Install migration tool
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Build the application (for prod)
#RUN go build -o currency-service .

# Expose port 8080
EXPOSE 8080

# Set environment variables for MySQL connection
ENV MYSQL_HOST mysql
ENV MYSQL_PORT 3306

# Wait for MySQL and then start the application
CMD ["sh", "-c", "/usr/local/bin/wait-for-mysql.sh $MYSQL_HOST $MYSQL_PORT 60 && go run main.go"]
