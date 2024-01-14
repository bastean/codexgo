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
	@husky install
	@make init

init:
	@curl -sSfL https://raw.githubusercontent.com/trufflesecurity/trufflehog/main/scripts/install.sh | sudo sh -s -- -b /usr/local/bin
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@npm ci --legacy-peer-deps

lint:
	@gofmt -l -w src/apps/crud/backend src/contexts/crud
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

compose-production:
	@${compose} up

compose-production-down:
	@${compose} down

test:
	@cd tests/ && go clean -testcache && go test -v -cover ./...

WARNING-docker-prune:
	@docker system prune --volumes -fa
