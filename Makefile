#* ---------- VARS ----------

npx = npx --no --

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

init:
	@npm ci --legacy-peer-deps
	@brew install trufflehog staticcheck

lint:
	@gofmt -l -s -w src/apps/crud/backend src/contexts/crud
	@${npx} prettier --ignore-unknown --write .

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

compose-dev:
	@${compose-env} .env.example.dev up

compose-dev-down:
	@${compose-env} .env.example.dev down
	@docker volume rm codexgo-database-dev -f

compose-test:
	@${compose-env} .env.example.test up --exit-code-from backend

compose-test-down:
	@${compose-env} .env.example.test down
	@docker volume rm codexgo-database-test -f

compose-prod:
	@${compose-env} .env.example.prod up

compose-prod-down:
	@${compose-env} .env.example.prod down

test:
	@go clean -testcache
	@cd tests/ && mkdir -p reports && go test -v -cover ./... > reports/report.txt

build:
	@rm -rf dist/
	@go build -o dist/codexgo ./src/apps/**/backend/cmd/web

docker-usage:
	@docker system df

WARNING-docker-prune-soft:
	@docker system prune

WARNING-docker-prune-hard:
	@docker system prune --volumes -a
