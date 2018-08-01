FROM alpine:3.7
# 因为需要访问htts的网页，需要add ca-certificates
RUN apk add --no-cache ca-certificates

WORKDIR /data/tools/
# 设置环境变量
ENV LANG='C.UTF-8' PYTHONIOENCODING='UTF-8' TZ='Asia/Shanghai' \
USERNAME=admin PASSWORD=admin INTERVAL=30 PORT=9000 \
DOMAIN="http://127.0.0.1:8080/" DBUSER=admin DBPWD=admin DB=devops \
DBHOST="127.0.0.1" DBPORT=8086 LOGLENGTH=50

EXPOSE $PORT

COPY ./webMonitor /data/tools/monitor

CMD /data/tools/monitor --username $USERNAME --password $PASSWORD --domain $DOMAIN \
--dbhost $DBHOST --dbuser $DBUSER --dbpwd $DBPWD --db $DB \
--port $PORT

