FROM golang:1.12

EXPOSE 12001

ENV GOPATH=/go:/pedafy

COPY . /pedafy/src/pedafy-assignments

WORKDIR /pedafy/src/pedafy-assignments

CMD ["go", "run", "src/main.go"]