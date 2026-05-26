FROM alpine:latest
MAINTAINER zhufuyi "g.zhufuyi@gmail.com"

# add curl, used for http service checking, can be installed without it if deployed in k8s
RUN apk add curl

COPY configs/ /app/configs/
COPY pledge_be /app/pledge_be
RUN chmod +x /app/pledge_be

# http port
EXPOSE 8080


WORKDIR /app

CMD ["./pledge_be", "-c", "configs/pledge_be.yml"]
# if you use the Configuration Center, pledge_be.yml is changed to the Configuration Center configuration.
#CMD ["./pledge_be", "-c", "configs/pledge_be.yml", "-enable-cc"]
