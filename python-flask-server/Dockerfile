FROM python:3.9-alpine

WORKDIR /app
COPY requirements.txt /app/
RUN pip install --no-cache-dir -r requirements.txt

COPY . /app
CMD ["waitress-serve", "--call", "main:create_app"]