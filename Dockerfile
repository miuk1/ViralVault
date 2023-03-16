# Build a docker image of postgres
FROM postgres:13.4-alpine as db

# Set up environment variables
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword
ENV POSTGRES_DB=viralvault

# Create a directory to store initialization scripts
RUN mkdir /docker-entrypoint-initdb.d

# Copy the initialization script
COPY init.sql /docker-entrypoint-initdb.d/

# Build the API
FROM golang:1.16-alpine AS api

# Set up environment variables
ENV DB_HOST=db
ENV DB_PORT=5432
ENV DB_USER=myuser
ENV DB_PASSWORD=mypassword
ENV DB_NAME=viralvault

# Set the working directory
WORKDIR /app

# Copy the source code
COPY api/ .

# Download and install dependencies
RUN apk add --no-cache git
RUN go mod download

# Run the tests and build the api
RUN go test -v ./tests/...
RUN go build -o cmd/viralvault/main .

# Frontend app
FROM node:16-alpine AS client

# Set the working directory
WORKDIR /app

# Copy the source code
COPY client/ .

# Install dependencies
RUN yarn install

# Build the app
RUN yarn build

# Copy the built app into the final image
FROM alpine

# Copy the built app
COPY --from=client /app/build /app

# Copy the API binary
COPY --from=api /app/main /app

# Copy the postgres configuration
COPY --from=db /usr/local/share/postgresql/pg_service.conf /usr/local/share/postgresql/pg_service.conf

# Set the entry point
ENTRYPOINT ["./main"]