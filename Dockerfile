FROM golang

WORKDIR /app

COPY app/go.mod ./
RUN go mod download

COPY ./app/* ./

RUN go build -o /server

EXPOSE 9999

CMD ["/server"]