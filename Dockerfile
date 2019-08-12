FROM golang:stretch as base
ENV ENV=production
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY ./go.mod ./

FROM base as dev
COPY ./ ./
RUN go build
CMD ["go", "run", "main.go"]
