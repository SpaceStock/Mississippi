FROM golang:alpine3.11

RUN mkdir -p /app
WORKDIR /go/src/build_mssp

COPY . .
RUN rm -rf /vendor/
RUN apk add --no-cache ca-certificates git
RUN go get -u github.com/gobuffalo/packr/packr
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -vendor-only

WORKDIR /go/src/build_mssp/cmd
RUN GOOS=linux GOARCH=amd64 packr build -o testsuite

RUN mv testsuite ../../../../app/
RUN rm -rf /../../../../build_mssp

ENTRYPOINT ["/app/testsuite"]