FROM golang:1.6
COPY . /go/src/github.com/alpagap/c14-rclone
WORKDIR /go/src/github.com/alpagap/c14-rclone
RUN go install -v ./cmd/c14
ENTRYPOINT ["c14"]
CMD ["help"]
