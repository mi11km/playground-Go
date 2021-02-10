#!/bin/zsh

# ./test.sh 引数　で実行して引数にテストしたい関数名を書く
go test -v -cover ./coding/chapter1/... -run $1