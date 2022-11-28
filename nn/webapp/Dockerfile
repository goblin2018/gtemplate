FROM nginx:alpine
LABEL author="goblin"
LABEL date="2022/11/28"

WORKDIR /usr/share/nginx/html

# Remove default nginx static assets
RUN rm -rf ./*
COPY ./dist . 
# copy config file
COPY ./nginx/nginx.conf /etc/nginx/conf.d/default.conf

# Containers run nginx with global directives and daemon off
ENTRYPOINT ["nginx", "-g", "daemon off;"]