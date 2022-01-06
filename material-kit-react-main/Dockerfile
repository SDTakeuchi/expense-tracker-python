FROM node:16-slim

WORKDIR /usr/src/app

COPY ./material-kit-react-main/package.json .

RUN yarn install

COPY ./material-kit-react-main/ .

CMD [ "yarn", "start" ]
