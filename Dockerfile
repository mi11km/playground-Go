FROM golang:alpine as builder
WORKDIR /go/src/app
#COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download
COPY . .
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/main ./cmd
#RUN go build -o /go/bin/main ./cmd
#CMD ["go", "run", "cmd/main.go"]

#FROM scratch as runner
# EXPOSE 8080
#COPY --from=builder /go/bin/main /app/main
#CMD ["/app/main"]