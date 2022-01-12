# stage I - khusus build dengan envinroment yang sama
# FROM golang:1.16-alpine AS builder
# RUN mkdir /app
# ADD . /app
# WORKDIR /app
# RUN go clean --modcache
# RUN go build -o main
# # EXPOSE 8080
# # CMD ["/app/main"]

# # stage 2
# FROM alpine:3.14
# WORKDIR /root/
# COPY --from=builder /app/config.json .
# COPY --from=builder /app/main .
# EXPOSE 8080
# CMD ["./main"]

# FROM golang:1.17-alpine3.14

# WORKDIR /minpro_arya

# COPY . .

# RUN go mod download


# RUN go build -o mainfile

# EXPOSE 8080

# CMD ["./mainfile"]
#BUILD STAGE
FROM golang:1.17-alpine AS builder
WORKDIR /finalproject-be-hms
COPY . .
COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go build -o main main.go

#RUN STAGE
FROM alpine:3.14 
WORKDIR /finalproject-be-hms
COPY --from=builder /finalproject-be-hms/config/config.json . 
COPY --from=builder /finalproject-be-hms/main .
EXPOSE 8000

CMD ["/finalproject-be-hms/main"]