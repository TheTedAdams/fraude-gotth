APP_NAME ?= app

.PHONY: vet
vet:
	go vet 

.PHONY: staticcheck
staticcheck:
	staticcheck ./..../...
	
.PHONY: test
test:
	go test -race -v -timeout 30s ./...

.PHONY: dev
dev:
	docker compose up
