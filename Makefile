#* ---------- VARS ----------

go-tidy = go mod tidy -e

npx = npx --no --
npm-ci = npm ci --legacy-peer-deps

release-it = ${npx} release-it -V
release-it-dry = ${npx} release-it -V -d --no-git.requireCleanWorkingDir

compose = docker compose
compose-env = ${compose} --env-file

#* ---------- RULES ----------

genesis:
	@git clean -e .env* -fdx
	@git reset --hard HEAD
	@make init

from-zero:
	@git init
	@make init
	@${npx} husky install

upgrade-manager:
	@npm upgrade -g
	@brew update && brew upgrade

upgrade-node:
	@${npx} ncu -u
	@rm -f package-lock.json
	@npm i --legacy-peer-deps

upgrade-reset:
	@git reset --hard HEAD
	@${npm-ci}

upgrade:
	@go run scripts/upgrade.go

init: upgrade-manager
	@${npm-ci}
	@brew install staticcheck upx
	@curl -sSfL https://raw.githubusercontent.com/trufflesecurity/trufflehog/main/scripts/install.sh | sudo sh -s -- -b /usr/local/bin v3.63.11

lint:
	@gofmt -l -s -w src/contexts/crud src/apps/crud/backend tests/
	@${npx} prettier --ignore-unknown --write .
	@rm -f go.work.sum
	@cd src/contexts/crud && ${go-tidy}
	@cd src/apps/crud/backend && ${go-tidy}
	@cd tests/ && ${go-tidy}

lint-check:
	@staticcheck ./src/apps/crud/backend/... ./src/contexts/crud/...
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
	@${compose-env} .env.example.test up --exit-code-from backend

compose-prod-down:
	@${compose-env} .env.example.prod down

compose-prod: compose-prod-down
	@${compose-env} .env.example.prod up

compose-down: compose-dev-down compose-test-down compose-prod-down

test:
	@go clean -testcache
	@cd tests/ && mkdir -p reports && go test -v -cover ./... > reports/report.txt

build:
	@rm -rf dist/
	@go build -o dist/codexgo ./src/apps/**/backend/cmd/web

build-upx: build
	@upx dist/codexgo

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
