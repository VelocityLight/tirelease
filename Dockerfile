# -- Build react web static files using yarn
FROM node:14.15-alpine as webbuilder
WORKDIR /webapp
COPY web/package.json web/yarn.lock ./
RUN yarn install
COPY web .
RUN yarn build

# -- Build Golang server
FROM golang:alpine as serverbuilder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOPROXY=https://goproxy.cn,direct \
    GOARCH=amd64
WORKDIR /goapp
COPY . .
RUN go mod tidy
RUN go build -o ./bin/tirelease cmd/tirelease/*.go

# -- Combine & Set the default run command for the container
COPY --from=webbuilder /webapp/build ./web/build
CMD ["/goapp/bin/tirelease"]
