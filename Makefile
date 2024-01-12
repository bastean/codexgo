npx = npx --no --
release-it = ${npx} release-it -V
release-it-dry = ${npx} release-it -V -d --no-git.requireCleanWorkingDir

init:
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@npm ci --legacy-peer-deps

prepare: 
	@git init
	@husky install

genesis:
	@git clean -e .env* -fdx
	@git reset --hard HEAD
	@make init

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
	@docker compose --env-file .env.example.dev up

compose-dev-down:
	@docker compose --env-file .env.example.dev down

compose-test:
	@docker compose --env-file .env.example.test up --exit-code-from backend

compose-test-down:
	@docker compose --env-file .env.example.test down

compose-production:
	@docker compose up

compose-production-down:
	@docker compose down

test:
	@cd tests/ && go clean -testcache && go test -v -cover ./...

WARNING-docker-prune:
	@docker system prune --volumes -fa
