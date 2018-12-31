FROM gustavohenrique/vue-cli:10-slim
ADD . /app
ADD Caddyfile /etc/Caddyfile
ARG VUE_APP_API_URL

RUN cd /app \
 && npm install \
 && VUE_APP_API_URL=${VUE_APP_API_URL} npm run build

WORKDIR /app/dist
CMD ["caddy", "-conf", "/etc/Caddyfile"]
