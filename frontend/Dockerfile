FROM oven/bun:1.3.2 AS builder
WORKDIR /app
COPY package*.json bun.lock* ./

ARG BETTER_AUTH_SECRET
ARG BETTER_AUTH_URL
ARG POSTGRES_HOSTNAME
ARG POSTGRES_PORT
ARG POSTGRES_DB
ARG POSTGRES_USER
ARG POSTGRES_PASSWORD

ENV BETTER_AUTH_SECRET=${BETTER_AUTH_SECRET}}
ENV BETTER_AUTH_URL=${BETTER_AUTH_URL}}
ENV POSTGRES_HOSTNAME=${POSTGRES_HOSTNAME}}
ENV POSTGRES_PORT=${POSTGRES_PORT}}
ENV POSTGRES_DB=${POSTGRES_DB}}
ENV POSTGRES_USER=${POSTGRES_USER}}
ENV POSTGRES_PASSWORD=${POSTGRES_PASSWORD}}

RUN apt-get update && apt-get install -y python3 build-essential
RUN bun install --frozen-lockfile
COPY . .
RUN bunx --bun @better-auth/cli migrate
RUN bun --bun run build

FROM oven/bun:1.3.2-alpine
WORKDIR /app
RUN mkdir -p static/avatars
COPY --from=builder /app/build build/
COPY --from=builder /app/node_modules node_modules/
COPY package.json .
EXPOSE 3001
ENV NODE_ENV=production
ENV PORT=3001
ENV ORIGIN=${ORIGIN}
CMD [ "bun", "./build/index.js" ]