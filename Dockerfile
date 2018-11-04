FROM golang:alpine as builder

RUN apk --update upgrade \
    && apk --no-cache --no-progress add git mercurial bash gcc musl-dev curl tar \
    && rm -rf /var/cache/apk/*
RUN mkdir -p /usr/local/bin \
    && curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 \
    && chmod +x /usr/local/bin/dep

ADD . $GOPATH/src/kubewheel
WORKDIR $GOPATH/src/kubewheel

RUN export GIN_MODE=release

RUN dep ensure
RUN rm -rf ./vendor/k8s.io/
RUN go get k8s.io/client-go/...
RUN go get -u k8s.io/apimachinery/...; exit 0

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -tags prod -o kubewheel .
RUN ls
RUN pwd

#COPY /go/src/kubewheel/kubewheel /app/
#WORKDIR /app
#CMD ["./kubewheel"]

FROM scratch
COPY --from=builder ./kubewheel /app/
WORKDIR /app
CMD ["./kubewheel"]