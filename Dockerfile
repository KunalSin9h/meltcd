FROM golang:alpine

RUN apk add gcc musl-dev

WORKDIR /meltcd

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o meltcd main.go

FROM alpine:latest

COPY --from=0 /meltcd/meltcd /bin/meltcd

EXPOSE 11771

ENV MELTCD_HOST=0.0.0.0

ENTRYPOINT [ "/bin/meltcd" ]
CMD ["serve", "--verbose"]