# hadolint ignore=DL3007
FROM nixos/nix:latest AS build

# hadolint ignore=DL3059
RUN nix-channel --update && \
    nix-env -iA nixpkgs.bash nixpkgs.gnutar nixpkgs.gzip nixpkgs.curl nixpkgs.postgresql_18 nixpkgs.awscli2

# Use nix-store to identify exactly what is needed and copy it to a specific path
# hadolint ignore=SC2046
RUN mkdir -p /tmp/nix-store-closure && \
    cp -R $(nix-store -qR $(which bash gnutar gzip curl psql aws)) /tmp/nix-store-closure

# hadolint ignore=DL3007
FROM alpine:latest
COPY --from=build /tmp/nix-store-closure /nix/store
COPY --from=build /root/.nix-profile/bin /usr/local/bin

COPY entrypoint.sh /
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
