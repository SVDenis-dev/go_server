FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

#RUN go build ./cmd/server/server.go

EXPOSE 8080

ENV MONGO_USERNAME=myuser
ENV MONGO_HOST=mongo
ENV MONGO_PORT=27017
ENV MONGO_DBNAME=posts

CMD ["./server"]
