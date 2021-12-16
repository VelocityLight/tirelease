# Only copy website/ and build React
FROM node:14.15-alpine as webbuilder
WORKDIR /webapp
COPY website/ ./website/
WORKDIR /webapp/website/
RUN yarn install
RUN yarn build


# Copy whole project and build Golang
FROM golang:alpine as serverbuilder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOPROXY=https://goproxy.cn,direct \
    GOARCH=amd64
WORKDIR /goapp
COPY . .
RUN go mod tidy
RUN go build -o ./bin/tirelease ./cmd/tirelease/*.go


# Copy whole directory website/ into goapp/bin/
# Aggregate multiple sources into one: finally /goapp/ contains everything
COPY --from=webbuilder /webapp/website/ ./bin/website/

# Set the image's default run command
CMD ["/goapp/bin/tirelease"]

