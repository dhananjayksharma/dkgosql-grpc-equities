# Build stage
FROM golang:1.19.9-alpine3.18 AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN ls -ltr
RUN go mod download

ENV MICROSERVICECDEMONEWAPI "local" 
RUN GOOS=linux GOARCH=amd64 go build -o grpcequitiesapi

# Run stage
FROM builder
WORKDIR /app
COPY --from=builder /app/grpcequitiesapi/ /app/
EXPOSE 8989
CMD ["./grpcequitiesapi"]