all:linux

pack:dep-prod
	rm -rf target target.tar.gz
	GOOS=linux GOARCH=amd64 revel build -a . -t target
	rm -rf target/app
	rm -rf target/src/github.com/zhifeiji/leanote/app/views
	cp -rf dist/views target/src/github.com/zhifeiji/leanote/app/
	tar zcf target.tar.gz target

run:dep-dev
	rm -rf bin
	revel build -a . -t bin
	rm -rf bin/src/github.com/zhifeiji/leanote/app/views
	cp -rf dist/views bin/src/github.com/zhifeiji/leanote/app/
	./bin/run.sh

dep-prod:
	gulp --env prod

dep-dev:
	gulp --env dev

