## Start a database via docker

```
docker run --name pg1 \
  -e POSTGRES_PASSWORD=welcome \
  -p 5432:5432 \
  -v /Users/nzhong/dev/docker/pg1:/var/lib/postgresql/data \
  -d postgres:12.0
```

then, from a GUI tool, execute command

```
CREATE DATABASE hello;
SELECT datname FROM pg_database;
```

or alternatively, 

```
minimal-go-postgres% docker exec -it pg1 /bin/bash
root@2b095d5bb4d2:/# psql -U postgres
psql (12.0 (Debian 12.0-2.pgdg100+1))
Type "help" for help.

postgres=# CREATE DATABASE hello;
postgres=# \l
```




```
CREATE TABLE books (
  isbn    char(14) NOT NULL,
  title   varchar(255) NOT NULL,
  author  varchar(255) NOT NULL,
  price   decimal(5,2) NOT NULL
);

INSERT INTO books (isbn, title, author, price) VALUES
('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
('978-1503379640', 'The Prince', 'Niccolò Machiavelli', 6.99);

ALTER TABLE books ADD PRIMARY KEY (isbn);
```
https://www.alexedwards.net/blog/practical-persistence-sql
