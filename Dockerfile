## Generic Go dockerfile
# Stage: build

# https://hub.docker.com/_/golang
FROM golang:alpine AS buildStage

ENV GOBIN=/go/bin
WORKDIR /go/src

COPY ./ ./

RUN go install -mod=vendor ./...

# Stage: final
FROM alpine:latest AS final

COPY --from=buildStage /go/bin/* /usr/local/bin/

WORKDIR /usr/local/bin

# Sets command to run the first executable in the current dir. This will be the just built golang binary
CMD find ./ -maxdepth 1 -perm -111 -type f -exec {} \;
