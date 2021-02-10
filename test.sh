#!/bin/zsh

number=$1

# ./test.sh 問題番号
go test -v -cover ./coding/chapter1/question"${number}".go ./coding/chapter1/question"${number}"_test.go


# command一覧

# すべてのテスト実行
# go test -v -cover ./...

# 特定のテスト関数のみ実行　正規表現可 ex) TestList*
# go test -v -cover ./... -run テスト関数名

# 特定ファイルのみ実行
# go test -v -cover ファイル名　テストファイル名