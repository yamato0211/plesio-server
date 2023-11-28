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