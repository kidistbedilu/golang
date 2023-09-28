From golang:1-alpine3.18

WORKDIR /app

COPY . .

RUN go mod init sayhello && go mod tidy

RUN go build -o hello

CMD ["./hello"]

# docker build -t hello .
# docker run --rm -d --name sayhello -p 8080:1337 hello
# docker stop sayhello
