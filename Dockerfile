FROM alpine:latest AS builder

WORKDIR /

COPY . /

RUN apk add --no-cache git make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

RUN go build -ldflags "-s -w" -o api cmd/main.go

RUN rm -rf go/pkg go/bin

FROM alpine:latest
WORKDIR /
COPY --from=builder /api ./

CMD [ "./" ]