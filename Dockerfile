FROM alpine:latest AS builder

WORKDIR /app

COPY . /app

RUN apk add --no-cache git make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

ENV PORT 1323
ENV MYSQL_HOST mysql
ENV MYSQL_USER eulabs
ENV MYSQL_PORT 3304
ENV MYSQL_DB_NAME eulabs
ENV MYSQL_PASSWORD eulabs

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

RUN go build -ldflags "-s -w" -o app cmd/main.go

RUN rm -rf go/pkg go/bin

FROM alpine:latest
WORKDIR /
COPY --from=builder /app ./

ENTRYPOINT [ "/app" ]