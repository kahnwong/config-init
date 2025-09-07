# ----------- docker ----------- #
start-db:
	docker compose -f docker-compose-postgres.yaml up -d

# ----------- caddy ----------- #
start: build
	caddy file-server --root dist/spa --listen :4000

# -----------golang ----------- #
build:
	goreleaser build --skip validate --snapshot --clean

release:
	bash -c "GITHUB_TOKEN=$(gh auth token) goreleaser release --skip validate --clean"

test:
	go test ./...

# ----------- mlflow ----------- #
start:
	mlflow server --port 8080

# ----------- fastapi ----------- #
start:
	uvicorn python_api_template.main:app --port 8080 --reload
test:
	pytest tests/single.py
test-full:
	pytest -n 8 tests/full.py
codecov:
	pytest --cov=python_api_template tests/single.py

# ----------- streamlit ----------- #
start:
	streamlit run article_summarizer/frontend.py --server.port 8501

# ----------- rust ----------- #
build:
	cargo build --release
run:
	./target/release/qa-api-rs
dev:
	systemfd --no-pid -s http::3000 -- cargo watch -x run
