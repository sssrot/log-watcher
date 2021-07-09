FROM golang:latest

WORKDIR /app
COPY go.mod go.sum main.go /app/
RUN go build -o watcher .
CMD [ "./watcher", "-conf", "/etc/logwatcher/conf.json" ]

# docker build --tag logwatcher:0.1 .
# docker run -d -e TG_TOKEN=$yourTgToken --mount type=bind,src=$confDir,target=/etc/logwatcher --mount type=bind,src=$logDir,target=/var/logs logwatcher:0.1 