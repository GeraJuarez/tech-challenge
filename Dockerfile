# Test stage
FROM golang:1.16 as test
COPY . /src
WORKDIR /src
RUN go mod download
RUN go test ./...

# Build stage
FROM golang:1.16-alpine as build
COPY . /src
COPY --from=test /go/pkg/mod/ /go/pkg/mod/ 
RUN apk --no-cache add ca-certificates
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -o api

# Copy binary to container
FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/api .
EXPOSE 8080
CMD ["/api"]
