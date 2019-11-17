# hexample

Go っぽいプロジェクト構成でヘキサゴナルアーキテクチャ的なことをやるリポジトリ

## wire

### インストール

```sh
go get github.com/google/wire/cmd/wire
```

### ビルド

```sh
cd cmd/httpwire
wire
```

### アプリケーションの起動

```sh
cd cmd/httpwire
go run main.go wire_gen.go
```
