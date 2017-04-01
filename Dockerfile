FROM alpine

RUN apk update && apk add sshfs

RUN mkdir -p /run/docker/plugins /mnt/docme/state /mnt/docme/volumes

COPY docme docme

CMD ["docme agent"]