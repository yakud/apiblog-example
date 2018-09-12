FROM golang:1.11.0-stretch

ENV GOPATH="/go"
ENV GOPROJECT=${GOPATH}/src/github.com/yakud/apiblog-example

ADD entrypoint.sh /
RUN mkdir -p ${GOPROJECT}

VOLUME ${GOPATH}
VOLUME ${GOPROJECT}
WORKDIR ${GOPROJECT}

EXPOSE 8080

# Запукаем go-шное приложение
CMD ["/entrypoint.sh"]