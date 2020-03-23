FROM centos:latest

Copy chat chat
COPY config/docker.yaml config/app.yaml
COPY static static

CMD [ "./chat" ]