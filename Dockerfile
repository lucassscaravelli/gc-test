FROM node:10-alpine as build-front

WORKDIR /app

COPY  ./client .
RUN yarn
RUN yarn build

FROM golang:alpine as build

WORKDIR $GOPATH/src/gctest

COPY . .

# Instalar dep
RUN apk add --no-cache git
RUN go get -u github.com/golang/dep/cmd/dep

RUN dep ensure
RUN go build

COPY --from=build-front /app/dist ./client

EXPOSE 8080

CMD ["./gctest"]