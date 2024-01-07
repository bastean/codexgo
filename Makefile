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
	@npx --no -- prettier --ignore-unknown --write .

lint-check:
	@staticcheck ./src/apps/crud/backend/... ./src/contexts/crud/...
	@npx --no -- prettier --check .

commit:
	@npx cz

dev:
	@docker compose -f docker-compose.dev.yml up

dev-down:
	@docker compose -f docker-compose.dev.yml down

docker-prune:
	@docker system prune --volumes -fa
