FROM golang:latest as builder
RUN mkdir -p /go/src/github.com/themoah/go-web-tools
COPY . /go/src/github.com/themoah/go-web-tools
WORKDIR /go/src/github.com/themoah/go-web-tools 
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go test
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-web-tools .
FROM alpine:3.10
RUN apk --no-cache add ca-certificates
LABEL maintainer="@themoah" 
LABEL version="0.2"
LABEL description="go-web-tools"
RUN adduser -S -D -H -h /app appuser
WORKDIR /app
USER appuser
COPY --from=builder /go/src/github.com/themoah/go-web-tools/go-web-tools /app/
EXPOSE 8080
CMD ["./go-web-tools"]