.PHONY: *

#* ~~~~~~~~~~~~VARS~~~~~~~~~~~~

#* ~~~~~~Go~~~~~~

go-tidy = go mod tidy -e

#* ~~~~~~Node~~~~~~

npx = npx --no --
npm-ci = npm ci --legacy-peer-deps

release-it = ${npx} release-it -V
release-it-dry = ${npx} release-it -V -d --no-git.requireCleanWorkingDir

#* ~~~~~~Git~~~~~~

git-reset-hard = git reset --hard HEAD

#* ~~~~~~Docker~~~~~~

compose = cd deployments/ && docker compose
compose-env = ${compose} --env-file

#* ~~~~~~Bash~~~~~~

bash = bash -o pipefail -c

#* ~~~~~~~~~~~~RULES~~~~~~~~~~~~

#* ~~~~~~Upgrades~~~~~~

upgrade-managers:
	#? sudo apt update && sudo apt upgrade -y
	npm upgrade -g

upgrade-go:
	go get -t -u ./...

upgrade-node:
	${npx} ncu --root -ws -u
	rm -f package-lock.json
	npm i --legacy-peer-deps

upgrade-reset:
	${git-reset-hard}
	${npm-ci}

upgrade:
	go run ./scripts/upgrade

#* ~~~~~~Dependencies~~~~~~

install-deps:
	go mod download
	${npm-ci}
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/a-h/templ/cmd/templ@latest
	curl -sSfL https://raw.githubusercontent.com/trufflesecurity/trufflehog/main/scripts/install.sh | sudo sh -s -- -b /usr/local/bin v3.63.11

copy-deps:
	go run ./scripts/copydeps

#* ~~~~~~Generators~~~~~~

generate-required:
	go generate ./...
	templ generate

#* ~~~~~~Initializations~~~~~~

init: upgrade-managers install-deps copy-deps generate-required

init-zero:
	git init
	$(MAKE) init
	${npx} husky install

#* ~~~~~~Linters/Formatters~~~~~~

lint: generate-required
	go mod tidy
	gofmt -l -s -w .
	${npx} prettier --ignore-unknown --write .
	templ fmt .

lint-check:
	staticcheck ./...
	${npx} prettier --check .

#* ~~~~~~Tests~~~~~~

test-sut:
	air

test-clean: generate-required
	go clean -testcache
	cd test/ && mkdir -p report

test-codegen:
	${npx} playwright codegen http://localhost:8080

test-sync: upgrade-go
	${npx} concurrently -s first -k --names 'SUT,TEST' '$(MAKE) test-sut' '${npx} wait-on -l http-get://localhost:8080 && $(TEST_SYNC)'

test-unit: test-clean
	${bash} 'go test -v -cover ./pkg/context/... -run TestUnit.* |& tee test/report/unit.report.log'

test-integration: test-clean
	${bash} 'go test -v -cover ./pkg/context/... -run TestIntegration.* |& tee test/report/integration.report.log'

test-acceptance-sync: 
	${bash} 'TEST_URL="http://localhost:8080" go test -v -cover ./pkg/cmd/... -run TestAcceptance.* |& tee test/report/acceptance.report.log'

test-acceptance: test-clean
	TEST_SYNC="$(MAKE) test-acceptance-sync" $(MAKE) test-sync

tests-sync:
	${bash} 'TEST_URL="http://localhost:8080" go test -v -cover ./... |& tee test/report/report.log'

tests: test-clean
	TEST_SYNC="$(MAKE) tests-sync" $(MAKE) test-sync

#* ~~~~~~Releases~~~~~~

release:
	${release-it}

release-alpha:
	${release-it} --preRelease=alpha
	
release-beta:
	${release-it} --preRelease=beta

release-ci:
	${release-it} --ci $(OPTIONS)

release-dry:
	${release-it-dry}
 
release-dry-version:
	${release-it-dry} --release-version

release-dry-changelog:
	${release-it-dry} --changelog

#* ~~~~~~Builds~~~~~~

build: generate-required lint
	rm -rf build/
	go build -ldflags="-s -w" -o build/codexgo ./cmd/codexgo

#* ~~~~~~ENVs~~~~~~

sync-env-reset:
	${git-reset-hard}

sync-env:
	cd deployments && go run ../scripts/syncenv

#* ~~~~~~Git~~~~~~

commit:
	${npx} cz

WARNING-git-forget:
	git rm -r --cached .
	git add .

WARNING-git-genesis:
	git clean -e .env* -fdx
	${git-reset-hard}
	$(MAKE) init

#* ~~~~~~Docker~~~~~~

docker-usage:
	docker system df

docker-it:
	docker exec -it $(ID) bash

compose-dev-down:
	${compose-env} .env.dev down
	docker volume rm codexgo-database-dev -f

compose-dev: compose-dev-down
	${compose-env} .env.dev up

compose-test-down:
	${compose-env} .env.test down
	docker volume rm codexgo-database-test -f

compose-test-integration: compose-test-down
	${compose-env} .env.test --env-file .env.demo.test.integration up --exit-code-from server

compose-test-acceptance: compose-test-down
	${compose-env} .env.test --env-file .env.demo.test.acceptance up --exit-code-from server

compose-tests: compose-test-down
	${compose-env} .env.test up --exit-code-from server

compose-prod-down:
	${compose-env} .env.prod down

compose-prod: compose-prod-down
	${compose-env} .env.prod up

demo-down:
	${compose-env} .env.demo down

demo: demo-down
	${compose-env} .env.demo up

compose-down: compose-dev-down compose-test-down compose-prod-down demo-down

WARNING-docker-prune-soft:
	docker system prune
	$(MAKE) compose-down
	$(MAKE) docker-usage

WARNING-docker-prune-hard:
	docker system prune --volumes -a
	$(MAKE) compose-down
	$(MAKE) docker-usage

#* ~~~~~~Fixes~~~~~~

fix-local-playwright:
	go get -u github.com/playwright-community/playwright-go
	go run github.com/playwright-community/playwright-go/cmd/playwright@latest install chromium --with-deps
