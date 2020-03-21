FROM golang:latest
WORKDIR /go/src/github.com/AlexanderAsmakov/qapps
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o main ./cmd/qapps/
CMD ["qapps"]