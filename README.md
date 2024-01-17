<h1 align="center">

<!-- [![logo readme](https://raw.githubusercontent.com/bastean/codexgo/main/docs/readme/logo-readme.png)](https://github.com/bastean) -->

[![logo readme](docs/readme/logo-readme.png)](https://github.com/bastean/codexgo)

</h1>

<div align="center">

> Example CRUD project applying Hexagonal Architecture, Domain-Driven Design (DDD), Command Query Responsibility Segregation (CQRS), Behavior-Driven Development (BDD), Continuous Integration (CI), and more... in Go

</div>

<br />

<div align="center">

[![license MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](https://github.com/commitizen/cz-cli)
[![release it](https://img.shields.io/badge/%F0%9F%93%A6%F0%9F%9A%80-release--it-orange.svg)](https://github.com/release-it/release-it)

</div>

<div align="center">

[![ci workflow](https://github.com/bastean/codexgo/actions/workflows/ci.yml/badge.svg)](https://github.com/bastean/codexgo/actions/workflows/ci.yml)
[![release workflow](https://github.com/bastean/codexgo/actions/workflows/release.yml/badge.svg)](https://github.com/bastean/codexgo/actions/workflows/release.yml)

</div>

<div align="center">

[![github release](https://img.shields.io/github/v/release/bastean/codexgo.svg)](https://github.com/bastean/codexgo/releases)

</div>

## Features

- Devcontainer

  - Features
  - Extensions & Settings

- Docker

  - Dockerfile
    - Multistage
  - Compose
    - Setup by ENV

- GitHub

  - Actions & Workflows
    - Setup Languages and Dependencies
    - Secrets Scanning, Linting & Test Checks
    - Automate Release
  - Issue Templates (Defaults)

- Git

  - Hooks
    - Pre-Commit
      - Secrets Scanning & Formatting
    - Commit-Msg
      - Check [Conventional Commits](https://www.conventionalcommits.org) rules

- Releases

  - Automatically
    - Hooks
      - Linting & Test Checks
    - Bump Version (based on [Conventional Commits](https://www.conventionalcommits.org) & [SemVer](https://semver.org/))
    - CHANGELOG
    - Commit & Tag
    - GitHub Release

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

1. Install required

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

1. Install required

   - [Go](https://go.dev/doc/install)
   - [Node](https://nodejs.org/en/download)
   - [Make](https://www.gnu.org/software/make)
   - [Docker](https://docs.docker.com/get-docker)

2. Run

   ```bash
   make init
   ```

#### ZIP

1. [Install required](#locally)

2. Run

   ```bash
   make from-zero
   ```

### Run

#### Development

```bash
make compose-dev
```

#### Test

```bash
make compose-test
```

#### Production

```bash
make compose-prod
```

## Screenshots

<div align="center">

<img src="docs/readme/codexgo-desktop-welcome.png" />

<img src="docs/readme/codexgo-desktop-dashboard.png" />

<img width="49%" src="docs/readme/codexgo-mobile-welcome.png" />

<img width="49%" src="docs/readme/codexgo-mobile-dashboard.png" />

</div>

## Tech Stack

#### Base

- [Go](https://go.dev)
- [HTMX](https://htmx.org)
- [Tailwind CSS](https://tailwindcss.com)
  - [daisyUI](https://daisyui.com)

#### Please see

- [go.mod](go.work) (from Workspaces)
- [package.json](package.json)

## Contributing

- Contributions and Feedback are always welcome!

## License

- [MIT](LICENSE)
