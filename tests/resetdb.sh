#!/bin/sh
echo "Reseting database..."
psql "${DATABASE_URL}" <<EOF
DROP TABLE IF EXISTS tasks;
CREATE TABLE tasks (id SERIAL NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, description VARCHAR, done BOOL NOT NULL DEFAULT false, CONSTRAINT id_pk PRIMARY KEY (id));
INSERT INTO tasks (description) VALUES ('Create a TODO app as example of e2e testing');
EOF