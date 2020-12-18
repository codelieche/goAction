FROM busybox:latest

WORKDIR /app

COPY ./smallweb /app/smallweb

EXPOSE 8080

CMD ["/bin/sh", "-c", "/app/smallweb --host=0.0.0.0 --port=8080"]
