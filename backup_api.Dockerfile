########################################################################

FROM node:20-bullseye-slim as base
RUN mkdir /app 
WORKDIR /app
RUN groupadd groupuserfe
RUN useradd -g groupuserfe userfe
COPY --chown=root:root ./backup_api/package*.json ./backup_api/yarn*.lock ./
RUN npm ci --only=production && npm cache clean --force
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    vim \
    tini \
    htop \
    && rm -rf /var/lib/apt/lists/*

########################################################################

FROM base as source
COPY --chown=root:root . .

########################################################################

FROM base as dev
ENV NODE_ENV=development
ENV PATH=/app/node_modules/.bin:$PATH
RUN npm install && npm cache clean --force
CMD ["nodemon", "test.js"]
RUN chmod 700 * -Rf
USER userfe

########################################################################

FROM source as prod
ENV NODE_ENV=production 
ENTRYPOINT ["/usr/bin/tini", "--"]
CMD ["node", "test.js"]
USER userfe

########################################################################

