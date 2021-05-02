# playground with Golang

- coding/ : 
[世界で闘うプログラミング力を鍛える本 コーディング面接189問とその解法](https://www.amazon.co.jp/dp/B071GN3JN2/)
の例題を解いたもの[(参考回答URL)](https://github.com/careercup/CtCI-6th-Edition)
  
- web-app/ : [Go言語によるWebアプリケーション開発](https://www.amazon.co.jp/Go%E8%A8%80%E8%AA%9E%E3%81%AB%E3%82%88%E3%82%8BWeb%E3%82%A2%E3%83%97%E3%83%AA%E3%82%B1%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%E9%96%8B%E7%99%BA-Mat-Ryer/dp/4873117526)
  
### 実行
- `main.go`の実行
```shell
make run
```
- すべてのテストの実行
```shell
make test_all
```
fileにパッケージの相対パスを指定することでそのパッケージのみのテストにできる。
```shell
# example
make test file=./coding/chapter1
```