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

# .PHONY: tailwind-watch
# tailwind-watch:
# 	tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch
	
# .PHONY: tailwind-build
# tailwind-build:
# 	tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify
	
# .PHONY: templ-watch
# templ-watch:
# 	templ generate --watch

# .PHONY: templ-build
# templ-build:
# 	templ generate
	
# .PHONY: dev
# dev:
# 	go build -o ./tmp/main ./cmd/main.go && air

# .PHONY: build
# build:
# 	make tailwind-build
# 	make templ-generate
# 	go build -ldflags "-X main.Environment=production" -o ./bin/$(APP_NAME) ./cmd/main.go
	
# .PHONY: docker-build
# docker-build:
# 	docker-compose build

# .PHONY: docker-up
# docker-up:
# 	docker-compose up

# .PHONY: docker-dev
# docker-dev:
# 	docker-compose up

# .PHONY: docker-down
# docker-down:
# 	docker-compose down

# .PHONY: docker-clean
# docker-clean:
# 	docker-compose down -v --rmi all

# Trying non-docker better DX?

# run templ generation in watch mode to detect all .templ files and 
# re-create _templ.txt files on change, then send reload event to browser. 
# Default url: http://localhost:7331
# Optional --open-browser=false 
.PHONY: live/templ
live/templ:
	templ generate --watch --proxy="http://localhost:3001" --proxyport=3000

# run air to detect any go file changes to re-build and re-run the server.
live/server:
	go run github.com/air-verse/air@v1.61.7 \
	--build.cmd "go build -o tmp/main ./cmd" --build.bin "./tmp/main" --build.delay "100" \
	--tmp_dir "./tmp" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--build.post_cmd "kill -15 $(lsof -ti:3001)" \
	--misc.clean_on_exit true

# run tailwindcss to generate the styles.css bundle in watch mode.
.PHONY: live/tailwind
live/tailwind:
	npx --yes @tailwindcss/cli -i ./static/css/input.css -o ./static/css/style.min.css --minify --watch

# run esbuild to generate the index.js bundle in watch mode.
# live/esbuild:
# 	npx --yes esbuild static/js/index.ts --bundle --outdir=static/js/ --watch

# watch for any js or css change in the assets/ folder, then reload the browser via templ proxy.
.PHONY: live/sync_assets
live/sync_assets:
	go run github.com/air-verse/air@v1.61.7 \
	--build.cmd "templ generate --notify-proxy --proxyport=3000" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "static" \
	--build.include_ext "js,css"

# start all 4 watch processes in parallel.
.PHONY: live
live: 
	make -j4 live/templ live/server live/tailwind live/sync_assets
