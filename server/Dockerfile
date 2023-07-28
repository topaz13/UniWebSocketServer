FROM golang:1.20

RUN echo "HOGE"
WORKDIR /go/app
RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]