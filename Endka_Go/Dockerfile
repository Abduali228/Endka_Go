FROM golang
COPY . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build /app/calculator_server && chmod +x /app/calculator_server
RUN CGO_ENABLED=0 GOOS=linux go build /app/calculator_client && chmod +x /app/calculator_client

EXPOSE 4000
