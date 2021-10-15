FROM node:14-alpine

RUN apk update && apk add bash
RUN apk add --no-cache git make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

# Install Glide
RUN go get -u github.com/Masterminds/glide/...

WORKDIR /usr/src/app
COPY . .

EXPOSE 80
EXPOSE 443
RUN go build

CMD ./deer-woman-dezigns
