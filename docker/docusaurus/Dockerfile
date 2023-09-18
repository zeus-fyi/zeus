#syntax=docker/dockerfile:1.4
FROM node:lts as base

WORKDIR /app

# Copy the package.json and package-lock.json files into the container
COPY . .

# Install any needed packages
RUN npm ci

# Copy the rest of the application code into the container
COPY apps/docusaurus/zeusfyi/ ./

# Build the React app for production
RUN npm run build --nomaps

ENV NODE_ENV=production

EXPOSE 3000

CMD ["node", "/app/node_modules/.bin/serve", "-s", "/app/build", "-l", "3000"]
