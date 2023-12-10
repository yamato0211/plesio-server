# plesio-server

## 環境構築
### Env file

```.env.sample```を参考にして```.env```を作成してください

### Docker

Install Docker Desktop

https://docs.docker.jp/v1.12/engine/installation/toc.html


Dockerコマンドが入ったら以下を実行
```
make run
```
APIサーバ,DB,Adminer(DBを操作できるGUI)が立ち上がる

### logs
ログを見たい場合は以下のコマンドを実行
```
make logs
```

### lint
Install golangci-lint

https://golangci-lint.run/usage/install/

golangci-lintコマンドが入ったら以下を実行
```
make lint # lintをかける

make lint-fix # lintをかける、直せるものは直してくれる
```

## ディレクトリ構成
```
.
├── Makefile 
├── README.md
├── cmd
│   └── main.go //APIのエントリポイント
├── compose.yml
├── docker //Dockerfile群
│   ├── dev //dev(開発用)
│   │   └── Dockerfile
│   └── prod //prod(本番用)
│       └── Dockerfile
├── go.mod
├── go.sum
├── k6 //負荷テスト k6ファイル
│   ├── websocket.js
│   └── ws_scenario.js
├── pkg //clean architecture packages
│   ├── adapter //adapter層
│   │   ├── http //http用のハンドラー
│   │   │   └── handler
│   │   │       └── user.go
│   │   ├── router.go //ルーター集約
│   │   ├── schemas //スキーマ定義
│   │   │   └── user.go
│   │   └── ws //websocket用のハンドラー
│   │       └── handler
│   │           └── websocket.go
│   ├── domain //ドメイン層
│   │   ├── entity //ドメインエンティティ
│   │   │   ├── event.go
│   │   │   ├── push.go
│   │   │   └── user.go
│   │   └── repository //レポジトリのinterface
│   │       ├── redis.go
│   │       └── user.go
│   ├── infra //infrastrcture層
│   │   ├── mysql //mysql関連の実装
│   │   │   ├── init.go
│   │   │   └── user.go
│   │   └── redis //redis関連の実装
│   │       └── init.go
│   ├── injection // wireによるDI
│   │   ├── wire.go
│   │   └── wire_gen.go
│   ├── usecase //usecase層
│   │   └── user.go
│   ├── utils //その他関数群
│   │   └── config //環境変数
│   │       └── config.go
│   └── web //webのその他
│       └── ws //wsの実装
│           ├── client.go
│           └── hub.go
└── tmp //バイナリ
    └── main
```