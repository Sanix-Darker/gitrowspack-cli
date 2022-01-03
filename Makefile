docker-build:
	docker build --no-cache -t sanixdarker/grp:latest -f Dockerfile .

docker-run:
	docker run -it sanixdarker/grp:latest

run:
	go run main.go

build:
	go build -o ./grp

exec: build
	./grp

