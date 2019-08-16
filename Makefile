build: # イメージビルド
	docker image build -t mygo:latest .

run: # コンテナ起動
	docker container run --rm -dt -v $(PWD)/app:/go/app -p 8081:8081 -p 5002:5002 --name mygo mygo:latest

ssh: # コンテナへ入る
	docker container exec -it mygo /bin/sh
