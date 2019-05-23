FROM node:10.13.0-alpine

COPY . /app/

RUN cd /app && npm i

EXPOSE 80

WORkDIR /app

CMD ["npm", "start"]
