FROM node:8.11-alpine as build

WORKDIR /app

COPY package.json ./package.json

RUN npm install

COPY . ./

RUN npm run build

FROM nginx:1.14.0

COPY --from=build /app/build/ /usr/share/nginx/html/

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]


