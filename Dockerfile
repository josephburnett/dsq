FROM golang
ADD . /go/src/github.com/josephburnett/dsq-golang
RUN go install github.com/josephburnett/dsq-golang/cmd
ENTRYPOINT /go/bin/cmd
EXPOSE 8080