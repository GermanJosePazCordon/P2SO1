#FROM node:16 as node
#WORKDIR /app
#COPY . .
#RUN npm install
#RUN npm run build --prod

#FROM nginx:alpine
#COPY --from=node /app/dist/client /usr/share/nginx/html

# Stage 0, "build-stage", based on Node.js, to build and compile the frontend
FROM node:16-alpine as build-stage
WORKDIR /app
COPY package*.json /app/
RUN npm install
COPY ./ /app/
ARG configuration=production
RUN npm run build -- --output-path=./dist/client --configuration $configuration

# Stage 1, based on Nginx, to have only the compiled app, ready for production with Nginx
FROM nginx:1.15
#Copy ci-dashboard-dist
COPY --from=build-stage /app/dist/client /usr/share/nginx/html
#Copy default nginx configuration
COPY ./nginx-custom.conf /etc/nginx/conf.d/default.conf