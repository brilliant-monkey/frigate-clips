
FROM golang:1.20 AS build

WORKDIR /app

COPY . .

RUN go build -o app cmd/main.go

FROM linuxserver/ffmpeg

WORKDIR /app

COPY --from=build /app/app .

ENTRYPOINT [ "./app" ]