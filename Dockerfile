FROM golang
ADD . /go/src/github.com/josephburnett/dsq
RUN go install github.com/josephburnett/dsq/cmd
ENTRYPOINT /go/bin/cmd
EXPOSE 8080