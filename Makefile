all:linux

linux:
	GOOS=linux GOARCH=amd64 revel build github.com/zhifeiji/leanote  ./bin

run:
	revel run -a .
