FROM golang:1.17.6

WORKDIR /app
EXPOSE 1323
RUN go install github.com/cosmtrek/air@latest

CMD ["air"]