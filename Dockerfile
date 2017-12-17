FROM golang:alpine
RUN apk --no-cache add curl git && curl https://glide.sh/get | sh
RUN mkdir -p /gopath/bin
ENV GOROOT /usr/lib/go
ENV GOPATH /gopath
ENV GOBIN /gopath/bin
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin
ADD . /gopath
RUN cd /gopath
RUN glide install
RUN go build -o server
WORKDIR /gopath
EXPOSE 8080
ENTRYPOINT ./server
