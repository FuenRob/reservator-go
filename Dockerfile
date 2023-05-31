FROM golang

WORKDIR /app

COPY app/go.mod app/go.sum ./
RUN go mod download

COPY ./app/ ./

RUN cd cmd && CGO_ENABLED=0 GOOS=linux go build -o /server

EXPOSE 9999

CMD ["/server"]