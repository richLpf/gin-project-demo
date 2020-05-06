FROM golang:1.12

WORKDIR /go/src/app

COPY . /go/src/app

ADD ./testproject_linux_amd64 /go/src/app

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

EXPOSE 9000

CMD ["./testproject_linux_amd64 -mode=prod"]