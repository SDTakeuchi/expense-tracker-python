FROM golang:1.17-alpine

WORKDIR /go/src/app

RUN export GO111MODULE=auto \
&& apk update && apk add git

COPY ./app/ .

RUN go mod init main \
&& go get . \
&& go build

EXPOSE 8080

CMD ["go", "run", "main.go", "0.0.0.0:8080"]