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

dev:
	@docker compose -f docker-compose.dev.yml up

dev-down:
	@docker compose -f docker-compose.dev.yml down

test:
	@echo "Test"

docker-prune:
	@docker system prune --volumes -fa
