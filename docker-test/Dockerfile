FROM golang:1.24.2-alpine3.21
WORKDIR  /test
COPY . /test
RUN go build /test
EXPOSE 8080
ENTRYPOINT [ "./dockergin" ]