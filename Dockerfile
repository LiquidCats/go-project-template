FROM golang:1.23-alpine as build

WORKDIR /app

ADD ./ /app

RUN go build -o main ./cmd/rater/main.go

FROM scratch

EXPOSE 8088

COPY --from=build /app/main /main

CMD ["/main"]



