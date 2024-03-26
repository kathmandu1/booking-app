FROM golang:1.18 as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/app/api
# RUN go install github.com/pilu/fresh@latest

# CMD ["air", "fresh"]

CMD ["air"]