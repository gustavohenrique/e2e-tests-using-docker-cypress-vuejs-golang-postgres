## Running local

Clone this repo:

```
git clone https://github.com/gustavohenrique/e2e-tests-using-docker-cypress-vuejs-golang-postgres.git
mkdir -p $GOPATH/src/github.com/gustavohenrique
mv e2e-tests-using-docker-cypress-vuejs-golang-postgres $GOPATH/src/github.com/gustavohenrique/
```

Create the database container:

```
docker run -d --name todos_test -p 5434:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -e POSTGRES_DB=todos gustavohenrique/postgis:11
```

Create the table:

```
docker exec todos_test bash -c 'psql -U admin todos_test -c "DROP TABLE IF EXISTS tasks;"'
docker exec todos_test bash -c 'psql -U admin todos_test -c "CREATE TABLE tasks (id SERIAL NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, description VARCHAR, done BOOL NOT NULL DEFAULT false, CONSTRAINT id_pk PRIMARY KEY (id));"'
```

Run the server:

```
cd $GOPATH/src/github.com/gustavohenrique/e2e-tests-using-docker-cypress-vuejs-golang-postgres/backend
DATABASE_URL=postgres://admin:password@127.0.0.1:5434/todos_test?sslmode=disable go run main.go
```

## Unit Tests

```
cd $GOPATH/src/github.com/gustavohenrique/e2e-tests-using-docker-cypress-vuejs-golang-postgres/backend
go test -cover -v github.com/gustavohenrique/e2e-tests-using-docker-vuejs-golang-postgres/backend/app/todolist