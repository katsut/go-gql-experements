# Backend

## Start in loca environment

`docker-compose up`

## Develop the application

Develop in the vscode devcontainer.

### Graphql (backends for frontends)

Attach to running container > bff

- generate schema

Edit ./bff/

`gqlgen`

- for graphql

```sh
docker-compose run --rm bff bash
```

- for grpc

```sh
docker-compose run --rm userservice bash
```

## Migrations

See: [Migration](./migration/README.md)

## Deployment

### Install AWS CLI and Copilot CLI

- [AWS Cli](https://docs.aws.amazon.com/ja_jp/cli/latest/userguide/getting-started-install.html)
- [AWS Copilot](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/copilot-install.html)
