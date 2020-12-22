FROM golang:1.14.3-alpine
WORKDIR /app
ENV CGO_ENABLED=0
COPY . /app
ARG TARGETOS
ARG TARGETARCH
RUN GO111MODULE=on go get github.com/minio/minio-go/v7
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build

EXPOSE 8080
CMD ["./sparagn-media-service"]
