FROM python:3.9-alpine

WORKDIR /app
COPY requirements.txt /app/
RUN pip install --no-cache-dir -r requirements.txt

COPY . /app
CMD ["uvicorn", "--factory", "--host", "0.0.0.0", "--port", "8080", "src:create_app"]