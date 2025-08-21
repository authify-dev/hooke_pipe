# Etapa build
FROM golang:1.24 AS build
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o ingress ./cmd/api/main.go

# Etapa runtime
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /app/ingress /app/ingress
ENV LISTEN_ADDR=:9002
USER nonroot:nonroot
EXPOSE 9002
ENTRYPOINT ["/app/ingress"]
