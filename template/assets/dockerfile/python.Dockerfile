FROM python:3.12-slim

WORKDIR /opt/app

# app deps
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# app
COPY . .
EXPOSE 8080
CMD uvicorn project.main:app --port 8080 --host 0.0.0.0
