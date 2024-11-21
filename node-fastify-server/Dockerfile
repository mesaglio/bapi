FROM node:22-alpine

WORKDIR /app
COPY *.json /app/
RUN npm install

COPY . /app/
CMD [ "npm", "start" ]