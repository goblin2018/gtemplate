FROM nginx:latest

LABEL author="goblin"

COPY ./dist /www
COPY ./scripts/nginx.conf /etc/nginx/nginx.conf
RUN echo 'web load ok'