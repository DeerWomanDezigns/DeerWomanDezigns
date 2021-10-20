FROM ghcr.io/deerwomandezigns/deerwomandezigns/nginx_proxy

WORKDIR /etc/nginx
COPY ./nginx.conf ./conf.d/default.conf
ENTRYPOINT [ "nginx" ]
CMD [ "-g", "daemon off;" ]
