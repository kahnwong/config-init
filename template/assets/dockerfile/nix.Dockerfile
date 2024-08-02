# hadolint ignore=DL3007
FROM nixos/nix:latest

RUN nix-channel --update
RUN nix-env -iA nixpkgs.bash && \
  nix-env -iA nixpkgs.gnutar && \
  nix-env -iA nixpkgs.gzip && \
  nix-env -iA nixpkgs.curl && \
  nix-env -iA nixpkgs.postgresql_16 && \
  nix-env -iA nixpkgs.awscli2

# set entrypoint
WORKDIR /opt/backup
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

ENTRYPOINT ["bash", "entrypoint.sh"]
