#FROM scratch
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache curl
VOLUME /data/
COPY ./static/ /data/

ENV HOST="0.0.0.0" PORT=8080 DURATION=35 VERSION=1 \
 TRACE_HEADERS="X-B3-TraceId,X-B3-ParentSpanId,X-B3-SpanId,X-B3-Sampled"
EXPOSE $PORT
COPY ./app /app/app

# 如果FROM其它的系统就从下面的命令启动
CMD ["/bin/sh", "-c", "/app/app --host=$HOST --port=$PORT --duration=$DURATION --version=$VERSION"]

# FROM scratch的话，用这个命令启动
#CMD ["/app/app"]
