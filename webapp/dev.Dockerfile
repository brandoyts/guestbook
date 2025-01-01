FROM node:22.12-alpine AS base

WORKDIR /app

COPY package*.json .

RUN npm ci

COPY . .

EXPOSE 3000

ENV NODE_ENV development

ENTRYPOINT ["npm","run", "dev"]

