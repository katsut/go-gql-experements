# Migration

Using [pressly/goose](https://github.com/pressly/goose)

## Migrate to the latest

`docker-compose run --rm migration`

## Other uses

Operate in the container.

`docker-compose run --rm migration bash`

- Downgrade migrations.

`goose down` or `goose down-to {VERSION}`

- Reset all migrations.

`goose reset`

- Add a new migration query.

`goose create {migration name} sql`

Edit the Up and Down queries in the generated SQL file, making sure that goose up and goose down succeed.
