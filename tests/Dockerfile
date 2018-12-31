FROM cypress/base:10

ADD . /home/node/tests
WORKDIR /home/node/tests

RUN bash -c 'apt-get update \
 && apt-get install -y postgresql-client \
 && npm install -g update \
 && chown -Rf node:node /home/node/tests'

USER node
RUN bash -c 'npm install \
 && npx cypress install'

CMD ["npm", "test"]
