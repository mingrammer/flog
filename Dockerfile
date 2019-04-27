FROM golang

ENV CGO_ENABLED=0
ENV GO111MODULE=on

WORKDIR /go/src/flog

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /bin/flog

FROM scratch
COPY --from=0 /bin/flog /bin/flog
ENTRYPOINT ["flog"]
