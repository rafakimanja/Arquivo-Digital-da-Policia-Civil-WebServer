FROM golang:1.23.1

WORKDIR /server

COPY . .

RUN go mod tidy

RUN GOOS=linux go build -o main

EXPOSE 5000

CMD [ "./main" ]