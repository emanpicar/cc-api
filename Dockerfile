FROM golang:1.14

# ARG PROXY_URI="http://10.158.100.8:8080"

WORKDIR $GOPATH/src/github.com/emanpicar/cc-api

COPY . $GOPATH/src/github.com/emanpicar/cc-api

# RUN export http_proxy=$PROXY_URI && \
#     export https_proxy=$PROXY_URI && \
#     git config --global http.proxy $PROXY_URI && \
#     go get -d -v ./...; exit 0

RUN go get -d -v ./...; exit 0

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o cc-api .

FROM scratch

WORKDIR /root/

COPY --from=0 /go/src/github.com/emanpicar/cc-api .

CMD ["./cc-api"]