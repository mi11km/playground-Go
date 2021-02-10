N=0
CHAPTER_N=1

file=./coding/chapter${CHAPTER_N}/question${N}.go
test_file=./coding/chapter${CHAPTER_N}/question${N}_test.go

# make t N=問題番号
t:
	go test -v -cover "${file}" "${test_file}"

# make sc N=問題番号
# show coverage in browser
sc:
	./test.sh "${N}"

ta:
	 go test -v -cover ./...

f:
	go fmt ./...