FROM golang:alpine as builder

WORKDIR /go/build/
RUN mkdir app
COPY main.go .
COPY app/*.go app/
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN GOARCH=wasm GOOS=js go build -o web/app.wasm app/*.go
RUN go build -o hambach

FROM alpine:3.6
RUN apk --no-cache add ca-certificates

WORKDIR /app/
RUN mkdir -p web/css
RUN mkdir -p web/images

COPY web/css/main.css web/css/.
COPY web/images/* web/images/
COPY --from=builder /go/build/hambach .
COPY --from=builder /go/build/web/app.wasm web/.

EXPOSE 8080

CMD ["./hambach"]