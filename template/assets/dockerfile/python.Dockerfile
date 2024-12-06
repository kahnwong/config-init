FROM python:3.13-slim

WORKDIR /opt/app

# app deps
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# app
COPY . .

## fastapi
EXPOSE 8080
CMD uvicorn project.main:app --port 8080 --host 0.0.0.0

### streamlit
#EXPOSE 8501
#CMD streamlit run module/frontend.py --server.port 8501
