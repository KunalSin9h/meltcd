FROM golang:alpine

RUN apk add gcc musl-dev

WORKDIR /meltcd

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o meltcd main.go

# Runtime
FROM alpine:latest

WORKDIR /meltcd

COPY --from=0 /meltcd/meltcd /meltcd/meltcd
COPY --from=0 /meltcd/ui/dist/ /meltcd/ui/dist/

EXPOSE 11771

ENV MELTCD_HOST=0.0.0.0

ENTRYPOINT [ "/meltcd/meltcd" ]
CMD ["serve", "--verbose"]