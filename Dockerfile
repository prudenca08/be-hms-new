FROM golang:1.17-alpine3.14

WORKDIR /be-hms-new

COPY . .

RUN go mod download


RUN go build -o cicd

EXPOSE 8080

CMD ["./cicd"]

FROM golang:1.17-alpine AS builder
WORKDIR /be-hms-new
COPY . .
RUN go mod download
RUN go build -o main main.go


FROM alpine:3.14 
WORKDIR /be-hms-new
COPY --from=builder /be-hms-new/config/config.json . 
COPY --from=builder /be-hms-new/main .
EXPOSE 8000

CMD ["/be-hms-new/main"]