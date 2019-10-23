FROM golang:latest as builder
RUN mkdir /go/src/build
ADD . /go/src/build/
WORKDIR /go/src/build
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go test
RUN go build -o go-web-tools .
FROM debian:latest
LABEL maintainer="@themoah" 
LABEL version="0.1"
LABEL description="go-web-tools"
COPY --from=builder /go/src/build/go-web-tools /app/
ENV PATH /app:$PATH
WORKDIR /app
EXPOSE 8080
CMD ["./go-web-tools"]