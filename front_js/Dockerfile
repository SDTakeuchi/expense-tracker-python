FROM node:16-slim

WORKDIR /usr/src/app/expense-tracker

COPY ./expense-tracker/package.json .

RUN yarn install

COPY ./expense-tracker/ .

ENTRYPOINT [ "yarn", "dev"]
