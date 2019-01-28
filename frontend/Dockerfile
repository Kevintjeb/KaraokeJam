FROM node:11.8.0-alpine

WORKDIR /karaoke-app

COPY package*.json ./

RUN yarn install

ADD . .

EXPOSE 4040
CMD ["yarn", "start"]
