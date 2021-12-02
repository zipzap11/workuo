#BUILD STAGE
FROM golang:1.17-alpine AS builder
WORKDIR /workuo
COPY . .
RUN go build -o main main.go

#RUN STAGE
FROM alpine:3.14 
WORKDIR /workuo
COPY --from=builder /workuo/main .
COPY app.env .
EXPOSE 8000

CMD ["/workuo/main"]


