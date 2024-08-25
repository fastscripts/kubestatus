##
## Build
##


FROM golang:1.22.6-alpine AS build





RUN mkdir /build
COPY . /build
WORKDIR /build



RUN go mod tidy && CGO_ENABLED=0 go build -o /build/kubestatus /build/cmd/kubestatus/*.go


##
## Deploy
##

#FROM swm-docker-registry.artifactory.intra.swm.de/ubuntu:22.04

#COPY containerfiles/apt/sources.list /etc/apt
#RUN  apt-get update && \
#     apt-get upgrade -y && \
#     apt-get install -y \
#     curl \
#     wget \
#     vim

FROM alpine:latest
#FROM scratch
RUN apk update


WORKDIR /
COPY --from=build /build/kubestatus /kubestatus 
COPY /web/assets/ /web/assets 
COPY /web/app/templates/ /web/app/templates
#COPY   configuration.json / 

USER 1000

ENV GOMEMLIMIT=250MiB GOGC=50
ENTRYPOINT [ "/kubestatus" ]