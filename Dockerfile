FROM golang:1.13-alpine as build
WORKDIR /emailValidator
ADD go.mod  go.sum ./
RUN apk add git
RUN go mod download
ADD . .


RUN CGO_ENABLED=0 GOOS=linux go install cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /opt/
COPY --from=build /go/bin/main emailValidator
CMD ["/opt/emailValidator"]