init:
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@npm ci

prepare: 
	@git init
	@husky install

commit:
	@npx cz

dev:
	@docker compose -f docker-compose.dev.yml up

dev-down:
	@docker compose -f docker-compose.dev.yml down

docker-prune:
	@docker system prune --volumes -fa
