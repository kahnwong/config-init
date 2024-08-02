# docker
start-db:
	docker compose -f docker-compose-postgres.yaml up -d

# caddy
start: build
	caddy file-server --root dist/spa --listen :4000

# golang
build:
	goreleaser release --skip publish --skip validate --clean -f .goreleaser-linux.yaml

test:
	go test ./...

# mlflow
start:
	mlflow server --port 8080

# fastapi
start:
	uvicorn python_api_template.main:app --port 8080 --reload
test:
	pytest tests/single.py
test-full:
	pytest -n 8 tests/full.py
codecov:
	pytest --cov=python_api_template tests/single.py

# streamlit
start:
	streamlit run article_summarizer/frontend.py --server.port 8501
