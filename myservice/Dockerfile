FROM golang:1.13.4-alpine3.10
RUN apk add inotify-tools
ADD build/package/cmd.sh /
RUN chmod +x /cmd.sh
CMD ["/cmd.sh"]
