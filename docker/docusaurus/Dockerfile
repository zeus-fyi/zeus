#syntax=docker/dockerfile:1.4
FROM node:lts as base

WORKDIR /app

ARG DIR
# Copy the package.json and package-lock.json files into the container
COPY apps/docusaurus/zeusfyi/package.json apps/docusaurus/zeusfyi/yarn.lock ./
COPY .openapi/mockingbird.yaml ./static/
# Install any needed packages
RUN yarn install --frozen-lockfile

# Copy the rest of the application code into the container
COPY apps/docusaurus/zeusfyi/ ./


# Build the React app for production
RUN yarn run build

ENV NODE_ENV=production

EXPOSE 3000

CMD ["yarn", "run", "serve", "--host", "0.0.0.0"]

