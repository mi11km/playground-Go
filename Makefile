N=0
CHAPTER_N=1

file=./coding/chapter${CHAPTER_N}/question${N}*

# make t N=問題番号
t:
	go test -v -cover ${file}

# make sc N=問題番号
# show coverage in browser
sc:
	./test.sh ${N}

ta:
	 go test -v -cover ./...

f:
	go fmt ./...



# 実行した回数
# １回あたりの実行に掛かった時間(ns/op)
# １回あたりのアロケーションで確保した容量(B/op)
# 1回あたりのアロケーション回数(allocs/op)

b:
	go test -bench ${file} -benchmem