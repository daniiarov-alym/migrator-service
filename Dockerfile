FROM golang:alpine as build 
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY . ./
RUN go build -o /main src/main.go

FROM alpine:latest
WORKDIR /
COPY --from=build /main /
ENTRYPOINT [ "/main" ]