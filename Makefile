.PHONY: *

#*------------VARS------------

#*______URL______

url-server = http://localhost:8080
url-github = https://github.com/bastean/codexgo

#*______Go______

module = github.com/bastean/codexgo/v4

go-tidy = go mod tidy -e

#*______Node______

npx = npx --no --
npm-ci = npm ci --legacy-peer-deps

ncu = ${npx} ncu -ws -u

release-it = ${npx} release-it -V
release-it-dry = ${npx} release-it -V -d --no-git.requireCleanWorkingDir

#*______Bash______

bash = bash -o pipefail -c

#*______Git______

git-reset-hard = git reset --hard HEAD

#*______Docker______

docker-rm-vol = docker volume rm -f
docker-rm-img = docker rmi -f

compose = cd deployments && docker compose
compose-env = ${compose} --env-file

#*______cURL______

curl = curl -sSfL

#*------------RULES------------

#*______Upgrade______

upgrade-managers:
	#? sudo apt update && sudo apt upgrade -y
	npm upgrade -g

upgrade-go:
	go get -t -u ./cmd/... ./internal/... ./pkg/... ./scripts/...

copydeps:
	go run ./scripts/copydeps

upgrade-node:
	${ncu}
	${ncu} fomantic-ui -t @nightly
	npm i --legacy-peer-deps
	$(MAKE) copydeps

upgrade-reset:
	${git-reset-hard}
	${npm-ci}

upgrade:
	godo exec -c configs/upgrade.json

#*______Install______

install-scanners:
	go install github.com/google/osv-scanner/cmd/osv-scanner@latest
	${curl} https://raw.githubusercontent.com/trufflesecurity/trufflehog/main/scripts/install.sh | sudo sh -s -- -b /usr/local/bin v3.63.11
	${curl} https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sudo sh -s -- -b /usr/local/bin v0.52.2

install-linters:
	go install golang.org/x/tools/cmd/goimports@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	npm i -g prettier

install-debuggers:
	go install golang.org/x/tools/cmd/deadcode@latest

install-tools-dev: install-scanners install-linters install-debuggers
	go install github.com/bastean/godo/cmd/godo@v0.2.0
	go install github.com/air-verse/air@latest
	go install github.com/a-h/templ/cmd/templ@latest

install-tools-test:
	go run github.com/playwright-community/playwright-go/cmd/playwright@latest install chromium --with-deps
	npm i -g concurrently wait-on

install-tooling: install-tools-dev install-tools-test

install-tooling-ci: install-tools-dev

#*______Download______

download-dependencies:
	go mod download
	${npm-ci}

#*______Generate______

generate-required:
	go generate ./...
	find . -name "*_templ.go" -type f -delete
	templ generate

#*______Restore______

restore:
	${npx} husky init
	git restore .

#*______Init______

init: upgrade-managers install-tooling download-dependencies copydeps generate-required restore

init-ci: upgrade-managers install-tooling-ci download-dependencies generate-required restore

genesis:
	git init
	git add .
	$(MAKE) init

#*______ENV______

syncenv-reset:
	${git-reset-hard}

syncenv:
	cd deployments && go run ../scripts/syncenv

#*______Scan______

scan-leaks-local:
	sudo trufflehog git file://. --only-verified
	trivy repo --scanners secret .

scan-leaks-remote:
	sudo trufflehog git ${url-github} --only-verified
	trivy repo --scanners secret ${url-github}

scan-vulns-local:
	osv-scanner --call-analysis=all -r .
	trivy repo --scanners vuln .

scan-misconfigs-local:
	trivy repo --scanners misconfig .

scan-leaks: scan-leaks-local scan-leaks-remote

scan-vulns: scan-vulns-local

scan-misconfigs: scan-misconfigs-local

scans: scan-leaks scan-vulns scan-misconfigs

#*______Lint/Format______

lint:
	go mod tidy
	goimports -l -w -local ${module} .
	gofmt -l -s -w .
	${npx} prettier --no-config --ignore-unknown --write .
	templ fmt .
	$(MAKE) generate-required

lint-check:
	staticcheck ./...
	${npx} prettier --check .

#*______Debug______

debug-dead:
	deadcode -test ./...

debugs: debug-dead

#*______Test______

test-sut:
	air

test-clean: generate-required
	go clean -testcache
	mkdir -p test/report

test-codegen:
	${npx} playwright codegen ${url-server}

test-sync:
	${npx} concurrently -s first -k --names 'SUT,TEST' '$(MAKE) test-sut' '${npx} wait-on -l ${url-server}/health && $(TEST_SYNC)'

test-unit: test-clean
	${bash} 'go test -v -cover ./pkg/context/... -run TestUnit.* |& tee test/report/unit.report.log'

test-integration: test-clean
	${bash} 'go test -v -cover ./pkg/context/... -run TestIntegration.* |& tee test/report/integration.report.log'

test-acceptance-sync: 
	${bash} 'SUT_URL="${url-server}" go test -v -cover ./internal/app/... -run TestAcceptance.* |& tee test/report/acceptance.report.log'

test-acceptance: test-clean
	TEST_SYNC="$(MAKE) test-acceptance-sync" $(MAKE) test-sync

tests-sync:
	${bash} 'SUT_URL="${url-server}" go test -v -cover ./... |& tee test/report/report.log'

tests: test-clean
	TEST_SYNC="$(MAKE) tests-sync" $(MAKE) test-sync

#*______Build______

build: lint
	rm -rf build/
	go build -ldflags="-s -w" -o build/codexgo ./cmd/codexgo

#*______Release______

release:
	${release-it}

release-alpha:
	${release-it} --preRelease=alpha
	
release-beta:
	${release-it} --preRelease=beta

release-ci:
	${release-it} --ci --no-git.requireCleanWorkingDir $(OPTIONS)

release-dry:
	${release-it-dry}
 
release-dry-version:
	${release-it-dry} --release-version

release-dry-changelog:
	${release-it-dry} --changelog

#*______Git______

commit:
	${npx} cz

WARNING-git-forget:
	git rm -r --cached .
	git add .

WARNING-git-genesis:
	git clean -e .env* -fdx
	${git-reset-hard}
	$(MAKE) init

#*______Docker______

docker-usage:
	docker system df

docker-it:
	docker exec -it $(ID) bash

compose-up:
	${compose-env} $(ENV) up $(SERVICE)

compose-down:
	${compose-env} $(ENV) down $(SERVICE)

compose-dev-down:
	${compose-env} .env.dev down
	${docker-rm-vol} codexgo-database-mongodb-dev

compose-dev: compose-dev-down
	${compose-env} .env.dev up

compose-test-down:
	${compose-env} .env.test down
	${docker-rm-vol} codexgo-database-mongodb-test

compose-test-integration: compose-test-down
	${compose-env} .env.test --env-file .env.test.integration up --exit-code-from codexgo

compose-test-acceptance: compose-test-down
	${compose-env} .env.test --env-file .env.test.acceptance up --exit-code-from codexgo

compose-tests: compose-test-down
	${compose-env} .env.test up --exit-code-from codexgo

compose-prod-down:
	${compose-env} .env.prod down
	${docker-rm-img} codexgo

compose-prod: compose-prod-down
	${compose-env} .env.prod up --exit-code-from codexgo

demo-down:
	${compose-env} .env.demo down

demo: demo-down
	${compose-env} .env.demo up --exit-code-from codexgo

compose-down: compose-dev-down compose-test-down compose-prod-down demo-down

WARNING-docker-prune-soft:
	docker system prune
	$(MAKE) compose-down
	$(MAKE) docker-usage

WARNING-docker-prune-hard:
	docker system prune --volumes -a
	$(MAKE) compose-down
	$(MAKE) docker-usage

#*______Devcontainer______

devcontainer:
	${bash} 'echo -e "$(USER_PASSWORD)\n$(USER_PASSWORD)" | sudo passwd vscode'

connect:
	ssh -p $(PORT) -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -o GlobalKnownHostsFile=/dev/null vscode@localhost

#*______Fix______

fix-dev: upgrade-go install-tools-dev

fix-test: upgrade-go install-tools-test
