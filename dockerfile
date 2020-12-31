FROM golang:1.15.5-alpine3.12 as builder

RUN apk update && apk add --no-cache git

WORKDIR /go/src/app

COPY / /go/src/app/

RUN go get -d -v

RUN go build -o /go/src/app/letsgoapp

FROM alpine AS final

COPY --from=builder /go/src/app/letsgoapp /app/letsgoapp

RUN apk -U upgrade \
    && apk add ffmpeg \
    && mkdir /convert

WORKDIR /convert

VOLUME  /convert

USER 1000

CMD /app/letsgoapp

# docker build -t letsgo-ffmpeg:x.0 .
# docker run -it --rm -v %cd%/in:/convert -e "to=mp3" letsgo-ffmpeg:2.0