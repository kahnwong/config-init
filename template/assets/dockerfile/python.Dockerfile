FROM ghcr.io/astral-sh/uv:python3.13-bookworm-slim

# install deps
# hadolint ignore=DL3045
COPY pyproject.toml uv.lock ./
RUN uv export --no-hashes --no-dev --no-emit-project --output-file=requirements.txt && \
  pip install --no-cache-dir -r requirements.txt

# app
WORKDIR /opt/app
COPY . .

# entrypoint
COPY entrypoint.sh .
EXPOSE 8080
CMD bash entrypoint.sh
