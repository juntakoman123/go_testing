# go_testing

## テスト実行方法
```sh
# 全てのテストを実行
go test -v ./...

# カバレッジ取得
go test -cover ./...

# 詳しく手続きごとのカバレッジを取得
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out

# ベンチマーク実行
go test -bench .
```

## ドキュメンテーション
```sh
# pkgsite インストール
go install golang.org/x/pkgsite/cmd/pkgsite@latest

# ドキュメントサーバ起動
pkgsite
```
