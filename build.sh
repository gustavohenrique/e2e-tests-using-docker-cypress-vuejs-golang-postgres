#!/bin/sh
git clone https://github.com/gustavohenrique/e2e-tests-using-docker-cypress-vuejs-golang-postgres
cd e2e-tests-using-docker-cypress-vuejs-golang-postgres
docker-compose up --exit-code-from tests --force-recreate
