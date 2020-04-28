FROM golang:latest

WORKDIR /go/src/gcsgo

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

RUN go get -u github.com/golang/dep/cmd/dep

COPY . /go/src/gcsgo

RUN dep ensure -v
RUN go build -ldflags "-linkmode external -extldflags -static" -o bin/gcsgo

CMD ["./gcsgo"]
