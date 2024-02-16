
# FROM golang:1.18
# RUN mkdir /app 
# WORKDIR /app
# COPY ./webservice_api/* ./
# RUN go mod download
# COPY . .
# RUN go build -o /web
# EXPOSE 8085
# CMD ["/web"]



########################################################################
FROM ubuntu:22.04 as base
# FROM node:20-bullseye-slim as base
RUN mkdir /app 
WORKDIR /app
COPY ./webservice_api/* ./
RUN groupadd groupuser
RUN useradd -g groupuser user

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    vim \
    tzdata \
    # nodejs \
    # npm \
    htop && \
    ln -fs /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    rm -rf /var/lib/apt/lists/*
########################################################################
FROM base as prod
CMD ["/app/web"]
EXPOSE 8085
USER user
########################################################################
# FROM base as dev
# RUN npm install -g nodemon
# CMD ["nodemon","--exec ./web --signal SIGTERM"]
# EXPOSE 8085
# USER user
########################################################################

