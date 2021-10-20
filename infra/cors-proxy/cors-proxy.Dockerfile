FROM ghcr.io/deerwomandezigns/nginx_proxy:latest

WORKDIR /etc/nginx
COPY ./nginx.conf ./conf.d/default.conf
ENTRYPOINT [ "nginx" ]
CMD [ "-g", "daemon off;" ]
