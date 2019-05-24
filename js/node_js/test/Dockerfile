FROM node:10.13.0-alpine

COPY package.json /app/package.json

RUN cd /app && npm i

COPY . /app/

WORKDIR /app

EXPOSE 80


CMD ["npm", "start"]
