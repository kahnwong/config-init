# hadolint ignore=DL3007
FROM caddy:latest

WORKDIR /usr/share/caddy

COPY . .

EXPOSE 80
