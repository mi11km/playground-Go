FROM golang:alpine as builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/main -ldflags '-s -w'
CMD ["go", "run", "main.go"]

#FROM scratch as runner
#EXPOSE 8080
#COPY --from=builder /go/bin/main /app/main
#COPY --from=builder /go/src/app/chat/templates /app/chat/templates
#CMD ["/app/main"]