FROM node:20 as builder
WORKDIR /app
COPY . .
RUN npm install
RUN npm run build

FROM nginx:1.24
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=builder /app/build .

ENTRYPOINT [ "nginx", "-g", "daemon off;" ]

# NOTE: exposes port 80 as default. Run with -p <port>:80
