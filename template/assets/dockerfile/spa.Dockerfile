FROM node:24 AS build

WORKDIR /opt/build

COPY package.json .
COPY yarn.lock .
RUN yarn install

COPY . .
ARG FOO
RUN yarn build

FROM caddy:2-alpine AS deploy

COPY Caddyfile /etc/caddy/Caddyfile
COPY --from=build /opt/build/dist/spa /app/dist/spa

EXPOSE 80
