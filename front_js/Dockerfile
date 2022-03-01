FROM node:16-slim

WORKDIR /usr/src/app/expense-tracker

COPY ./expense-tracker/package.json .

RUN yarn install \
&&  yarn install --save react-chartjs-2 chart.js

COPY ./expense-tracker/ .

ENTRYPOINT [ "yarn", "dev"]
