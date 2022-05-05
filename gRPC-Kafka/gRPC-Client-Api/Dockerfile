FROM node:17-alpine

WORKDIR /app
COPY . .
RUN npm ci
EXPOSE 3000

CMD ["npm", "start"]