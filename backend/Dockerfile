FROM golang:1.11-stretch

COPY . /go/src/github.com/gustavohenrique/e2e-tests-using-docker-cypress-vuejs-golang-postgres/backend
WORKDIR /go/src/github.com/gustavohenrique/e2e-tests-using-docker-cypress-vuejs-golang-postgres/backend

ENV PORT 11080

RUN sh -c "go get -u github.com/gorilla/mux \
 && go get -u github.com/gorilla/handlers \
 && go get -u github.com/jmoiron/sqlx \
 && go get -u github.com/lib/pq"

RUN sh -c "go build -i -v -o api main.go \
 && mv api /usr/local/bin"

CMD ["api"]
