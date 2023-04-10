FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o dist

EXPOSE 1323

ENTRYPOINT [ "./dist" ]