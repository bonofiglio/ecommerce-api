FROM golang:alpine3.16

RUN mkdir /api
WORKDIR /api

# Install git
RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

COPY . .
COPY .env .

# Download the dependencies
RUN go get -d -v ./...

# Build the application
RUN go install -v ./... && go build -o /server

RUN chmod g=rx /server

# Expose the set port
EXPOSE $PORT

# Run the executable
CMD ["/server"]