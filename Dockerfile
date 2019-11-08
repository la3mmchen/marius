FROM golang:alpine as builder
ARG version
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go build -o main -ldflags "-X main.GitCommit=$version" . 
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
ENTRYPOINT [ "./main" ]
CMD ["--help"]