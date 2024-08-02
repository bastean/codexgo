<h1 align="center">

<!-- [![README Logo](https://raw.githubusercontent.com/bastean/codexgo/main/assets/readme/logo.png)](https://github.com/bastean) -->

[![README Logo](assets/readme/logo.png)](https://github.com/bastean/codexgo)

</h1>

<div align="center">

> Example CRUD project applying Hexagonal Architecture, Domain-Driven Design (DDD), Event-Driven Architecture (EDA), Command Query Responsibility Segregation (CQRS), Behavior-Driven Development (BDD), Continuous Integration (CI), and more... in Go.

</div>

<br />

<div align="center">

[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/bastean/codexgo/v4)](https://goreportcard.com/report/github.com/bastean/codexgo/v4)
[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](https://github.com/commitizen/cz-cli)
[![Release It!](https://img.shields.io/badge/%F0%9F%93%A6%F0%9F%9A%80-release--it-orange.svg)](https://github.com/release-it/release-it)

</div>

<div align="center">

[![Upgrade workflow](https://github.com/bastean/codexgo/actions/workflows/upgrade.yml/badge.svg)](https://github.com/bastean/codexgo/actions/workflows/upgrade.yml)
[![CI workflow](https://github.com/bastean/codexgo/actions/workflows/ci.yml/badge.svg)](https://github.com/bastean/codexgo/actions/workflows/ci.yml)
[![Release workflow](https://github.com/bastean/codexgo/actions/workflows/release.yml/badge.svg)](https://github.com/bastean/codexgo/actions/workflows/release.yml)

</div>

<div align="center">

[![Go Reference](https://pkg.go.dev/badge/github.com/bastean/codexgo/v4.svg)](https://pkg.go.dev/github.com/bastean/codexgo/v4)
[![GitHub Releases](https://img.shields.io/github/v/release/bastean/codexgo.svg)](https://github.com/bastean/codexgo/releases)

</div>

## Showcase

<div align="center">

<img src="assets/readme/desktop-home.png" />

<img src="assets/readme/desktop-dashboard.png" />

<img width="49%" src="assets/readme/mobile-home.png" />

<img width="49%" src="assets/readme/mobile-dashboard.png" />

<img src="assets/readme/mail-confirm-account.png" />

</div>

## Usage

### Docker (Demo)

> [!NOTE]
>
> - [System Requirements](#locally)
> - In the Demo version, the link to confirm the account is sent through the Terminal.
>   - _"Hi \<username\>, please confirm your account through this link: \<link\>"_
> - We can define our own **SMTP** configuration in the [.env.demo](deployments/.env.demo) file by simply modifying the `CODEXGO_SMTP_*` variables, then we will be able to receive the links by mail.

```bash
make demo
```

## Features

### Project Layout

- Based on [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

### Git

- Hooks managed by [husky](https://github.com/typicode/husky):
  - Pre-Push:
    - Scanning Repository for leaks using [TruffleHog CLI](https://github.com/trufflesecurity/trufflehog) and [Trivy](https://github.com/aquasecurity/trivy)
  - Pre-Commit: [lint-staged](https://github.com/lint-staged/lint-staged)
    - Scanning files for leaks using [TruffleHog CLI](https://github.com/trufflesecurity/trufflehog?tab=readme-ov-file#8-scan-individual-files-or-directories)
    - Formatting
  - Commit-Msg: [commitlint](https://github.com/conventional-changelog/commitlint)
    - Check [Conventional Commits](https://www.conventionalcommits.org) rules
- Commit message helper using [Commitizen](https://github.com/commitizen/cz-cli).
  - Interactive prompt that allows you to write commits following the [Conventional Commits](https://www.conventionalcommits.org) rules:
    ```bash
    make commit
    ```

### Linting/Formatting Tools

- Go: **staticcheck** and **gofmt**.
- templ: **templ fmt**.
- Gherkin: **Cucumber extension**.
- Others: **Prettier cli/extension**.

### Scanners

- [TruffleHog CLI](https://github.com/trufflesecurity/trufflehog): Secrets.
- [Trivy](https://github.com/aquasecurity/trivy): Secrets, Vulnerabilities and Misconfigurations.
- [OSV-Scanner](https://github.com/google/osv-scanner): Vulnerabilities.

### Testing Packages

- Random data generator: [Gofakeit](https://github.com/brianvoe/gofakeit).
- Unit/Integration: [Testify](https://github.com/stretchr/testify).
- Acceptance: [Testify](https://github.com/stretchr/testify), [Godog (Cucumber)](https://github.com/cucumber/godog) and [Playwright](https://github.com/playwright-community/playwright-go).

### Releases

- Automatically managed by [Release It!](https://github.com/release-it/release-it):
  - Before/After Hooks for:
    - Linting
    - Testing
  - Bump version based on [Conventional Commits](https://www.conventionalcommits.org) and [SemVer](https://semver.org/):
    - CHANGELOG generator
    - Commits and Tags generator
    - GitHub Releases

### GitHub

- Actions for:
  - Setup Languages and Dependencies
- Workflows running:
  - Automatically (Triggered by **Push** or **Pull requests**):
    - Secrets Scanning ([TruffleHog Action](https://github.com/trufflesecurity/trufflehog?tab=readme-ov-file#octocat-trufflehog-github-action))
    - Linting
    - Testing
  - Manually (Using the **Actions tab** on GitHub):
    - Upgrade Dependencies
    - Automate Release
- Issue Templates **(Defaults)**.

### Devcontainer

- Multiple Features already pre-configured:
  - Go
  - Node
  - Docker in Docker
- Extensions and their respective settings to work with:
  - Go
  - templ
  - Cucumber
    - Gherkin
  - Prettier
  - Better Comments
  - Todo Tree
  - cSpell

### Docker

- Dockerfile
  - **Multi-stage builds**:
    - Development
    - Testing
    - Build
    - Production
- Compose
  - Switched by ENVs.

### Message Broker

- Routing Key based on [AsyncAPI Topic Definition](https://github.com/fmvilas/topic-definition).

### Security

- Form validation at the client using [Fomantic - Form Validation](https://fomantic-ui.com/behaviors/form.html).
  - On the server, the validations are performed using the **Value Objects** defined in the **Context**.
- Data **authentication** via **JWT** managed by **Session Cookies**.
- Account confirmation via **Mail** or **Terminal**.
- Password hashing using [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt).
- Requests **Rate Limiting**.
- Server log files.

### Scripts

- [syncenv](scripts/syncenv/syncenv.go)
  - Synchronize all **.env\*** files in the directory using an **.env** model.
- [copydeps](scripts/copydeps/copydeps.go)
  - Copies the files required by the browser dependencies from the **node_modules** folder and places them inside the **static** folder on the server.
- [upgrade](scripts/upgrade/upgrade.go)
  - Perform the following steps to upgrade the project:
    - Upgrade Go, Node and Tools.
    - Linting and Testing.
    - Commit upgrades.
- [run](deployments/run.sh)
  - Display the logs and redirect them to a file whose name depends on the time at which the service was run.
  - Used in Production Image.

## Basic Workflow (Domain > (Infrastructure | Application) > Presentation)

### Bounded Context (App/Business/Department) > Modules (Troubleshooting) > Layers (Domain, Infrastructure & Application)

- **Domain (Logic Core)**
  - Value Objects (Entities)
    - Mother Creators
    - Unit Tests
  - Messages (Event/Command)
    - Mother Creators
  - Aggregates (Sets of Entities)
    - Aggregate Root (Core Set)
    - Mother Creators
  - Role Interfaces (Ports)
    - Repository
    - Broker
  - Model Interfaces
    - Use Cases
    - Handlers/Consumers
  - Services (Abstract Logic)
  - Errors (Management)
- **Infrastructure (Port Adapters)**
  - Persistence
    - Repository Mocks
    - Implementations (Adapters)
    - Integration Tests
  - Communication
    - Broker Mocks
    - Implementations (Adapters)
    - Integration Tests
- **Application (Orchestration of Domain Logic)**
  - Use Cases
    - Implementations
  - Commands
    - Mother Creators
  - Queries/Responses
    - Mother Creators
  - Handlers/Consumers
    - Implementations
    - Unit Tests

### App > Server > (Presentation)

- **Presentation (Consumers of Bounded Context Modules)**
  - Services (Mapping)
    - Centralize Imports
  - Server
    - Templates
    - Handlers
    - Routes
    - Features (Gherkin)
      - Acceptance Tests

### Idiomatic

- **Domain**
  - `errors.New*()`, `errors.BubbleUp()` & `errors.Panic()`
    - Only in the "Domain" can we throw a `panic()`.
- **Infrastructure**
  - `Open()` & `Close()`
    - `session`
  - `errors.New*()` & `errors.BubbleUp()`
- **Application**
  - `Run()`, `Handle()` & `On()`
  - `errors.New*()` & `errors.BubbleUp()`
- **Modules**
  - `Start()` & `Stop()`
  - `errors.BubbleUp()`
- **Services / Apps**
  - `Up()` & `Down()`
    - `log.[Wrap]()`
  - `errors.New*()` & `errors.BubbleUp()`
- **Main**
  - `log.Fatal()` & `log.[Wrap]()`
    - Only `main()` can use `log.Fatal()`.
- **Logs**
  - `[embed]`
    - We use `[]` to "embed" external values such as error messages, fields, etc... inside our messages.

## First Steps

### Clone

#### HTTPS

```bash
git clone https://github.com/bastean/codexgo.git && cd codexgo
```

#### SSH

```bash
git clone git@github.com:bastean/codexgo.git && cd codexgo
```

### Initialize

#### Dev Container (recommended)

1. System Requirements

   - [Docker](https://docs.docker.com/get-docker)

     - [Dev Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

2. Start VS Code

   ```bash
   code .
   ```

3. Open Command Palette

   - Ctrl+Shift+P

4. Run

   ```txt
   Dev Containers: Reopen in Container
   ```

#### Locally

1. System Requirements

   - [Go](https://go.dev/doc/install)
   - [Node](https://nodejs.org/en/download)
   - [Make](https://www.gnu.org/software/make)
   - [Docker](https://docs.docker.com/get-docker)

2. Run

   ```bash
   make init
   ```

### ZIP

> [!NOTE]
>
> - [System Requirements](#locally)
> - We need to change `<user>` and `<repository>` with our own values.

```bash
curl -sSfLO https://github.com/bastean/codexgo/archive/refs/heads/main.zip \
&& unzip main.zip \
&& mv codexgo-main <repository> \
&& rm main.zip \
&& cd <repository> \
&& make genesis \
&& git commit -m "chore: add codexgo" \
&& git branch -M main \
&& git remote add github https://github.com/<user>/<repository>.git \
&& git push -u github main \
&& git status
```

### GitHub Repository

> [!IMPORTANT]
> These settings are necessary to be able to execute the Actions Workflows.

#### Settings tab

##### Actions

- General

  - Workflow permissions

    - [x] Read and write permissions

##### Secrets and variables

- Actions

  - New repository secret

    - `BOT_GPG_PRIVATE_KEY`

      ```bash
      gpg --armor --export-secret-key [Pub_Key_ID (*-BOT)]
      ```

    - `BOT_GPG_PASSPHRASE`

### Run

#### ENVs

> [!IMPORTANT]
> Before running it, we must initialize the following environment variable files:
>
> - [.env.example](deployments/.env.example)
>   - We will have to create a `.env.(dev|test|prod)` for each runtime environment.
>   - In the [.env.example.demo](deployments/.env.example.demo) file, we can see the values that can be used.
>
> In case we only want to run the **Integration** or **Acceptance** tests, in addition to having the `.env.test` file, we must have the following files created:
>
> - [.env.example.test.integration](deployments/.env.example.test.integration)
>   - Rename the file to `.env.test.integration`.
> - [.env.example.test.acceptance](deployments/.env.example.test.acceptance)
>   - Rename the file to `.env.test.acceptance`.

#### Development

```bash
make compose-dev
```

#### Tests

##### Unit

```bash
make test-unit
```

##### Integration

```bash
make compose-test-integration
```

##### Acceptance

```bash
make compose-test-acceptance
```

##### Unit/Integration/Acceptance

```bash
make compose-tests
```

#### Production

```bash
make compose-prod
```

## Tech Stack

#### Base

- [Go](https://go.dev)
- [Gin](https://gin-gonic.com)
- [templ](https://templ.guide)
  - [Fomantic-UI](https://fomantic-ui.com)
- [RabbitMQ](https://www.rabbitmq.com/tutorials/tutorial-one-go)
- [MongoDB](https://www.mongodb.com/docs/drivers/go)

#### Please see

- [go.mod](go.mod)
- [package.json](package.json)

## Contributing

- Contributions and Feedback are always welcome!
  - [Open a new issue](https://github.com/bastean/codexgo/issues/new/choose)

## License

- [MIT](LICENSE)
