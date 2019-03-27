FROM golang
ENV CGO_ENABLED=0
RUN go get github.com/mingrammer/flog

FROM scratch
COPY --from=0 /go/bin/flog /bin/flog
ENTRYPOINT ["flog"]
