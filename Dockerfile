FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GIN_MODE=release

WORKDIR /build

COPY go.mod .
# COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o gocalc .

WORKDIR /dist
RUN cp /build/gocalc .

EXPOSE 8080

CMD ["/dist/gocalc"]