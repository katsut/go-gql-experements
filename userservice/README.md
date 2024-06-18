# user-service

## コンテナ内での操作

`docker-compose run --rm userservice bash`

## モジュールの整理

`go mod tidy`

使用されていない依存関係の削除、必要な依存関係の追加、ライブラリの最新化を行い、go.mod, go.sumが書き換わる

## プロトコルバッファをコンパイルしてGoのコードを生成

`make generate`

