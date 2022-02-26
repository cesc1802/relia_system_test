FROM golang:1.14.4 AS build
LABEL maintainer="thuocnv1802@gmail.com"

# Switches to /tmp/app as the working directory, similar to 'cd'
WORKDIR /build/app

## If you have a go.mod and go.sum file in your project, uncomment lines 13, 14, 15

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Builds the current project to a binary file.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o relia_system .

#########################################################

# The project has been successfully built and we will use a
# lightweight ubuntu image to run the server
FROM alpine:3.9

# Copies the binary file from the BUILD container to /app folder
COPY --from=build /build/app/relia_system /app/relia_system

# Switch working directory
WORKDIR /app

# Exposes the 8088 port from the container
EXPOSE 8088

# Runs the binary once the container starts
ENTRYPOINT ["./relia_system"]