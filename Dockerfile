FROM golang as builder
WORKDIR /app/src/gin_docker
ADD . ./
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"
RUN go build -o qqbot-RSS_amd64_linux .

FROM ubuntu
WORKDIR /app/src/gin_docker
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /app/src/gin_docker/qqbot-RSS_amd64_linux /app/src/gin_docker
COPY --from=builder /app/src/gin_docker/config.yaml /app/src/gin_docker
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
ENTRYPOINT  ["./qqbot-RSS_amd64_linux"]