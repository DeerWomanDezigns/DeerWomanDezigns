FROM nginx:alpine

WORKDIR /etc/nginx
COPY ./nginx.conf ./conf.d/default.conf
EXPOSE 90
ENTRYPOINT [ "nginx" ]
CMD [ "-g", "daemon off;" ]
