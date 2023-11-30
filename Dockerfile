FROM golang:1.20.5-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o app ./cmd/main.go
#RUN chmod +x ./service

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/configs ./configs
EXPOSE 8080
CMD [ "/app/app" ]
