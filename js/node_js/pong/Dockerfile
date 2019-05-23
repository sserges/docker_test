FROM node:10.15-alpine

COPY package.json /app/package.json

RUN cd /app/ && npm install

COPY . /app/

WORKDIR /app

EXPOSE 80

CMD ["npm", "start"]
