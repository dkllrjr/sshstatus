FROM golang:alpine3.15 as gobuilder
WORKDIR /app
COPY *.go ./
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go build -o /app/sshstatus

FROM alpine:latest
ENV PORT=9429
ENV CHECK_TIMEOUT=60
ENV TCP_TIMEOUT=10
COPY config.yml ./
COPY --from=gobuilder /app/sshstatus ./
EXPOSE $PORT
CMD ./sshstatus $PORT $CHECK_TIMEOUT $TCP_TIMEOUT

