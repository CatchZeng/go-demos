FROM centos:7
 
MAINTAINER CatchZeng <catchzenghh@gmail.com>

# install golang
RUN yum install -y gcc
RUN yum -y install wget
RUN wget https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
RUN rpm -ivh epel-release-latest-7.noarch.rpm
RUN yum install -y golang

# config GOROOT
ENV GOROOT /usr/lib/golang
ENV PATH=$PATH:/usr/lib/golang/bin

# config GOPATH
RUN mkdir -p /root/gopath
RUN mkdir -p /root/gopath/src
RUN mkdir -p /root/gopath/pkg
RUN mkdir -p /root/gopath/bin
ENV GOPATH /root/gopath

# copy source files
RUN mkdir -p /root/gopath/src/server
COPY src/* /root/gopath/src/server/

# build the server
WORKDIR /root/gopath/src/server
RUN go build -o server.bin main.go

# startup the server
CMD /root/gopath/src/server/server.bin