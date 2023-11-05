FROM golang:1.18

WORKDIR /go/src/app

RUN mkdir static

COPY ./backend .

COPY ./frontend/dist ./static

RUN go build main.go

EXPOSE 8080

CMD ["./main"]
