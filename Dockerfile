FROM golang:1.11.1
WORKDIR $GOPATH/src/github.com/hwsc-org/
RUN git clone https://github.com/hwsc-org/hwsc-grpc-sample-svc.git
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
WORKDIR $GOPATH/src/github.com/hwsc-org/hwsc-grpc-sample-svc
RUN dep ensure -v
RUN go install
ENTRYPOINT ["/go/bin/hwsc-grpc-sample-svc"]
EXPOSE 50051