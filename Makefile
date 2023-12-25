run-dev:
	@docker compose -f docker-compose.dev.yml up

down-dev:
	@docker compose -f docker-compose.dev.yml down

docker-prune:
	@docker system prune --volumes -fa
