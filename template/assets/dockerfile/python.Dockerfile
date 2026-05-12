FROM python:3.14-slim
COPY --from=ghcr.io/astral-sh/uv:latest /uv /usr/local/bin/uv

WORKDIR /app
COPY pyproject.toml uv.lock README.md ./
COPY src/ ./src/

RUN uv sync --frozen --no-dev

EXPOSE 8080
CMD ["uv", "run", "--no-sync", "uvicorn", "app.server:app", "--host", "0.0.0.0", "--port", "8080"]
