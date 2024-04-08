.PHONY: *

#* ---------- VARS ----------

#* Go
go-tidy = go mod tidy -e

#* Node
npx = npx --no --
npm-ci = npm ci --legacy-peer-deps

release-it = ${npx} release-it -V
release-it-dry = ${npx} release-it -V -d --no-git.requireCleanWorkingDir

#* Git
git-reset-hard = git reset --hard HEAD

#* Docker
compose = cd deployments/ && docker compose
compose-env = ${compose} --env-file

#* ---------- RULES ----------

#* Upgrades
upgrade-managers:
	#? @sudo apt update && sudo apt upgrade -y
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

#* Installations
install-deps:
	@go mod download
	@${npm-ci}
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install github.com/a-h/templ/cmd/templ@latest
	@curl -sSfL https://raw.githubusercontent.com/trufflesecurity/trufflehog/main/scripts/install.sh | sudo sh -s -- -b /usr/local/bin v3.63.11

#* Generators
generate-required:
	@templ generate

#* Initializations
init: upgrade-managers install-deps generate-required
	
init-from-zero:
	@git init
	@make init
	@${npx} husky install

#* Linters/Formatters
lint: generate-required
	@go mod tidy
	@gofmt -l -s -w .
	@${npx} prettier --ignore-unknown --write .
	@templ fmt .

lint-check:
	@staticcheck ./...
	@${npx} prettier --check .

#* Tests
test-sut:
	@air

test-clean:
	@go clean -testcache
	@cd test/ && mkdir -p report

test-sync: upgrade-go
	@${npx} concurrently -s first -k --names 'SUT,TEST' 'make test-sut' '${npx} wait-on -l http-get://localhost:8080 && $(TEST_SYNC)'

test-unit: test-clean
	@go test -v -cover ./pkg/context/... -run TestUnit.* > test/report/unit-report.txt

test-integration: test-clean
	@go test -v -cover ./pkg/context/... -run TestIntegration.* > test/report/integration-report.txt

test-acceptance-sync: 
	@TEST_URL='http://localhost:8080' go test -v -cover ./pkg/cmd/... -run TestAcceptance.* > test/report/acceptance-report.txt

test-acceptance: test-clean
	@TEST_SYNC="make test-acceptance-sync" make test-sync

test-all-sync:
	@TEST_URL='http://localhost:8080' go test -v -cover ./... > test/report/report.txt

test-all: test-clean
	@TEST_SYNC="make test-all-sync" make test-sync

#* Releases
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

#* Builds
build: generate-required lint
	@rm -rf build/
	@go build -o build/codexgo ./cmd/codexgo

build-upx: build
	#? @upx build/codexgo

#* ENVs
sync-env-reset:
	@${git-reset-hard}

sync-env:
	@cd deployments && go run ../scripts/sync-env/**

#* Git
commit:
	@${npx} cz

WARNING-git-forget:
	@git rm -r --cached .
	@git add .

WARNING-git-genesis:
	@git clean -e .env* -fdx
	@${git-reset-hard}
	@make init

#* Docker
docker-usage:
	@docker system df

compose-dev-down:
	@${compose-env} .env.dev down
	@docker volume rm codexgo-database-dev -f

compose-dev: compose-dev-down
	@${compose-env} .env.dev up

compose-test-down:
	@${compose-env} .env.test down
	@docker volume rm codexgo-database-test -f

compose-test-integration: compose-test-down
	@${compose-env} .env.test --env-file .env.example.test.integration up --exit-code-from server

compose-test-acceptance: compose-test-down
	@${compose-env} .env.test --env-file .env.example.test.acceptance up --exit-code-from server

compose-test-all: compose-test-down
	@${compose-env} .env.test up --exit-code-from server

compose-prod-down:
	@${compose-env} .env.prod down

compose-prod: compose-prod-down
	@${compose-env} .env.prod up

compose-down: compose-dev-down compose-test-down compose-prod-down

WARNING-docker-prune-soft:
	@docker system prune
	@make compose-down
	@make docker-usage

WARNING-docker-prune-hard:
	@docker system prune --volumes -a
	@make compose-down
	@make docker-usage
