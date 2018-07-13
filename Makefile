.PHONY: deploy
deploy:
	rm -rf dist/
	mkdir -p dist
	cd app && dep ensure
	go build -o dist/server app/cmd/server/main.go
	go build -o dist/worker app/cmd/worker/main.go
	vagrant scp dist/server default:server
	vagrant scp dist/worker default:worker

.PHONY: install
install:
	vagrant plugin install vagrant-scp
