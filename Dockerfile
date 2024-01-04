FROM golang:alpine

RUN apk add gcc musl-dev

WORKDIR /meltcd

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_CFLAGS="-D_LARGEFILE64_SOURCE" go build -o meltcd-bin main.go

# Runtime
FROM alpine:latest

COPY --from=0 /meltcd/meltcd-bin /bin/meltcd

EXPOSE 11771

ENV MELTCD_HOST=0.0.0.0

ENTRYPOINT [ "/bin/meltcd" ]
CMD ["serve", "--verbose"]