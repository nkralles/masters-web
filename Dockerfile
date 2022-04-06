FROM node:lts-bullseye

RUN mkdir -p /srv/masters && chown node:node /srv/masters
USER node
WORKDIR /srv/masters
COPY --chown=node:node package.json package-lock.json /srv/masters/
RUN npm install
COPY --chown=node:node ./ /srv/masters/
RUN NODE_ENV=production npm run build

FROM golang:1.18-bullseye

WORKDIR /srv/masters
COPY ./ /srv/masters/
RUN go install -mod=vendor ./...

COPY --chown=root:root --from=0 /srv/masters/ /srv/masters/

EXPOSE 5454
ENV ENV=production
ENV LISTEN_ADDRESS=0.0.0.0
ENV PORT=5454
ENTRYPOINT ["/go/bin/masters-web"]
