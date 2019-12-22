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