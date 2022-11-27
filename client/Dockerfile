# syntax=docker/dockerfile:1

FROM node:18-slim as builder

WORKDIR /app
COPY package.json package-lock.json ./

# Установить зависимости только из package-lock.json
RUN npm ci
COPY . .
RUN npm run build

FROM nginx:1.23.2-alpine
COPY --from=builder /app/build /usr/share/nginx/html

RUN rm /etc/nginx/conf.d/default.conf
COPY nginx/nginx.conf /etc/nginx/conf.d

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
