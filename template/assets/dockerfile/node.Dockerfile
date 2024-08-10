FROM node:18 AS build

WORKDIR /opt/build

COPY package.json .
COPY yarn.lock .
RUN yarn install

COPY . .
RUN yarn build

FROM node:18-alpine AS deploy

WORKDIR /opt/app

COPY --from=build /opt/build/package.json ./
COPY --from=build /opt/build/node_modules ./node_modules
COPY --from=build /opt/build/dist ./dist

EXPOSE 8000

CMD [ "yarn", "start", "-h", "0.0.0.0" ]
