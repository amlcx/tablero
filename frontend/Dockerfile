FROM oven/bun:1.3.2-alpine AS builder
WORKDIR /app
COPY package*.json bun.lock* ./

ARG DATABASE_URL
ARG BETTER_AUTH_SECRET
ARG BETTER_AUTH_URL
ARG UPLOAD_DIR
ARG BACKEND_URL

ENV PORT=${PORT}
ENV DATABASE_URL=${DATABASE_URL}
ENV BETTER_AUTH_SECRET=${BETTER_AUTH_SECRET}
ENV BETTER_AUTH_URL=${BETTER_AUTH_URL}
ENV ORIGIN=${ORIGIN}
ENV UPLOAD_DIR=${UPLOAD_DIR}
ENV BACKEND_URL=${BACKEND_URL}

RUN bun install --frozen-lockfile
COPY . .
RUN bun --bun run build

FROM oven/bun:1.3.2-alpine
WORKDIR /app
COPY --from=builder /app/build build/
COPY --from=builder /app/node_modules node_modules/
COPY package.json .
COPY drizzle.config.ts .
COPY migrations migrations/
ENV NODE_ENV=production
VOLUME [ "/app/uploads" ]
CMD ["bun", "start"]