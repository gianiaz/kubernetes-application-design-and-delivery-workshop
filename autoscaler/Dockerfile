FROM golang:1.11.1
WORKDIR /home/exporter
COPY main.go go.mod ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /home/exporter/app .
CMD ["./app"]
