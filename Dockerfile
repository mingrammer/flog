FROM golang
ENV CGO_ENABLED=0
ENV GO111MODULE=on
COPY . $GOPATH/src/github.com/flog
WORKDIR $GOPATH/src/github.com/flog
RUN go build -mod=vendor

FROM scratch
COPY --from=0 /go/src/github.com/flog/flog /bin/flog
ENTRYPOINT ["flog"]
