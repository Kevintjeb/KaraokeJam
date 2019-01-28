# Stage 0, "build-stage", based on Node.js, to build and compile the frontend
FROM node:11.8.0 as build-stage
WORKDIR /karaoke-app
COPY package*.json ./
RUN yarn
COPY . .
RUN yarn run build

# Stage 1, based on Nginx, to have only the compiled app, ready for production with Nginx
FROM nginx:1.15.8

COPY --from=build-stage /karaoke-app/dist/ /usr/share/nginx/html