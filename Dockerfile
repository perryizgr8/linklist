FROM golang:alpine as builder
WORKDIR /src/app
COPY . .
RUN go build -o linklist cmd/linklist/linklist.go

FROM alpine
WORKDIR /root
COPY --from=builder /src/app ./app

WORKDIR /root/app
CMD [ "./linklist" ]
