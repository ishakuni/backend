# backend

### Local development

**WIP**

```bash
$ cp .env.example .env
$ docker volume create ishakuni
$ docker run \
    -v ishakuni:/var/lib/postgresql/data \
    -e POSTGRES_PASSWORD=1234 \
    -e POSTGRES_DB=ishakuni-dev\
    -p 5432:5432 \
    -d postgres:10.4-alpine
```
