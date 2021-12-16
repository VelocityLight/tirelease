# -- Tips: When 'FROM', many are allowed, and '--from=number' or '--from=name' is widely used to aggregate sources
# -- Tips: When 'WORKDIR', it is widely used in dockerfile to switch directories. And if the directory not exist, will also create it
# -- Tips: When 'COPY', only the files inside the source folder are copied(not entire folder),
# And if you want to copy the entire folder, then keep the dest folder name is same as the source, make it look like a complete copy.

# Copy 'website/' contents(not entire folder) into .(/webapp/) and build React 
FROM node:14.15-alpine as webbuilder
WORKDIR /webapp
COPY website/ .
RUN yarn install
RUN yarn build

# Copy whole project's contents(not entire folder) into .(/goapp/) and build Golang
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

# Copy whole folder /webapp/build/ into goapp/website/build
# Aggregate multiple sources into one: finally /goapp/ contains everything
COPY --from=webbuilder /webapp/build/ ./website/build/

# Set the image's default run command
# Tips: 'docker run' absolute path is: /goapp/, so "cmd/tirelease/main.go:http.Dir" == '/goapp/website/build/'
CMD ["/goapp/bin/tirelease"]

