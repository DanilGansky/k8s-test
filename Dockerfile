FROM golang:1.16

WORKDIR /go/src/k8s-test
ADD . .

RUN make deps
RUN make build

ENTRYPOINT [ "./bin/server" ]