FROM golang
WORKDIR $GOPATH/src/gin_docker
ADD . ./
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"
RUN go build -o qqbot-RSS_amd64_linux .
EXPOSE 8080
ENTRYPOINT  ["./qqbot-RSS_amd64_linux"]