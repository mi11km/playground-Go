# playground with Golang

- coding/ : 
[世界で闘うプログラミング力を鍛える本 コーディング面接189問とその解法](https://www.amazon.co.jp/dp/B071GN3JN2/)
の例題を解いたもの[(参考回答URL)](https://github.com/careercup/CtCI-6th-Edition)
  
### 実行
- `main.go`の実行
```shell
docker-compose run --rm app go run cmd/main.go
```
- すべてのテストの実行
```shell
docker-compose run --rm app go test -v -cover ./...
```
上記の`./...`をパッケージの相対パスにすることでそのパッケージのみのテストにできる。
```shell
# example
docker-compose run --rm app go test -v -cover ./coding/chapter1
```