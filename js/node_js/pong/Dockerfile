FROM node:10.15-alpine

COPY . /app/

RUN cd /app && npm install

WORKDIR /app

EXPOSE 80

CMD ["npm", "start"]
