FROM golang:1.15

WORKDIR /go/src/app
COPY . .

RUN go mod init && go mod tidy
RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 3000

CMD ["app"]