FROM golang
ADD . /go/src/github.com/janitor/linker
WORKDIR /go/src/github.com/janitor/linker
RUN go get && go install && rm -rf /go/src/
ENTRYPOINT /go/bin/linker -mongo-host mongodb://mongo:27017
EXPOSE 8000
