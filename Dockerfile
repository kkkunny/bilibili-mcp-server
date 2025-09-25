FROM golang:1.24.7-alpine3.22 AS builder

#ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY . .
RUN go build --tags release -o bin/bilibili_mcp .


FROM alpine:3.22 AS final

RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone \
ENV TZ Asia/Shanghai
RUN apk add --no-cache ca-certificates && update-ca-certificates
WORKDIR /app
COPY --from=builder /app/bin/* /app
EXPOSE 80
ENTRYPOINT ["/app/bilibili_mcp"]