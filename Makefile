#* ---------- VARS ----------

git-reset-hard = git reset --hard HEAD

go-tidy = go mod tidy -e

npx = npx --no --
npm-ci = npm ci --legacy-peer-deps

release-it = ${npx} release-it -V
release-it-dry = ${npx} release-it -V -d --no-git.requireCleanWorkingDir

compose = cd deployments/ && docker compose
compose-env = ${compose} --env-file

#* ---------- RULES ----------

genesis:
	@git clean -e .env* -fdx
	@${git-reset-hard}
	@make init

from-zero:
	@git init
	@make init
	@${npx} husky install

upgrade-manager:
	@sudo apt update && sudo apt upgrade -y
	@npm upgrade -g

upgrade-go:
	@go get -t -u ./...

upgrade-node:
	@${npx} ncu -u
	@rm -f package-lock.json
	@npm i --legacy-peer-deps

upgrade-reset:
	@${git-reset-hard}
	@${npm-ci}

upgrade:
	@go run scripts/upgrade/**

init: upgrade-manager
	@go mod download
	@${npm-ci}
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install github.com/a-h/templ/cmd/templ@latest
	@curl -sSfL https://raw.githubusercontent.com/trufflesecurity/trufflehog/main/scripts/install.sh | sudo sh -s -- -b /usr/local/bin v3.63.11

lint:
	@go mod tidy
	@gofmt -l -s -w .
	@${npx} prettier --ignore-unknown --write .
	@templ generate
	@templ fmt .

lint-check:
	@staticcheck ./...
	@${npx} prettier --check .

commit:
	@${npx} cz

release:
	@${release-it}

release-alpha:
	@${release-it} --preRelease=alpha
	
release-beta:
	@${release-it} --preRelease=beta

release-ci:
	@${release-it} --ci $(OPTIONS)

release-dry:
	@${release-it-dry}
 
release-dry-version:
	@${release-it-dry} --release-version

release-dry-changelog:
	@${release-it-dry} --changelog

compose-dev-down:
	@${compose-env} .env.example.dev down
	@docker volume rm codexgo-database-dev -f

compose-dev: compose-dev-down
	@${compose-env} .env.example.dev up

compose-test-down:
	@${compose-env} .env.example.test down
	@docker volume rm codexgo-database-test -f

compose-test: compose-test-down
	@${compose-env} .env.example.test up --exit-code-from server

compose-prod-down:
	@${compose-env} .env.example.prod down

compose-prod: compose-prod-down
	@${compose-env} .env.example.prod up

compose-down: compose-dev-down compose-test-down compose-prod-down

test-server:
	@air

test-start:
	@go clean -testcache
	@cd test/ && mkdir -p report
	@TEST_URL='http://localhost:8080' go test -v -cover ./... > test/report/report.txt

test-run: upgrade-go
	@${npx} concurrently -s first -k --names 'SUT,TEST' 'make test-server' '${npx} wait-on -l http-get://localhost:8080 && make test-start'

build:
	@rm -rf dist/
	@templ generate
	@go build -o dist/codexgo ./cmd/codexgo

build-upx: build
	#? @upx dist/codexgo

sync-env-reset:
	@${git-reset-hard}

sync-env:
	@cd deployments && go run ../scripts/sync-env/**

docker-usage:
	@docker system df

WARNING-docker-prune-soft:
	@docker system prune
	@make compose-down
	@make docker-usage

WARNING-docker-prune-hard:
	@docker system prune --volumes -a
	@make compose-down
	@make docker-usage
