FROM node as build
COPY . /src
WORKDIR /src
RUN npm install
RUN npm run build

FROM nginx
COPY --from=build /src/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
