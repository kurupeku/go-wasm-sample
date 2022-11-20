FROM  --platform=linux/amd64 node:19-bullseye

ENV WORKDIR /app

RUN apt-get update && \
  npm i -g pnpm

WORKDIR ${WORKDIR}

COPY ./ ${WORKDIR}

RUN pnpm install
