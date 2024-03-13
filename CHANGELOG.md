# Changelog

## [2.0.1](https://github.com/bastean/codexgo/compare/v2.0.0...v2.0.1) (2024-03-13)

### Bug Fixes

- **deps:** upgrade dependencies ([4e3f621](https://github.com/bastean/codexgo/commit/4e3f621bf8b3833ef2cd4d7bbe877cf5d38a81ac))

### Refactors

- change domain models ([f80911a](https://github.com/bastean/codexgo/commit/f80911acf48a9bfb115d3328a7834babfa123b02))

## [2.0.0](https://github.com/bastean/codexgo/compare/v1.5.0...v2.0.0) (2024-03-02)

### ⚠ BREAKING CHANGES

- add standard project layout

### Bug Fixes

- **deps:** upgrade dependencies ([f11b15f](https://github.com/bastean/codexgo/commit/f11b15f77899ab50ae0ac744dd84346cb71a7760))

### Refactors

- add standard project layout ([307089c](https://github.com/bastean/codexgo/commit/307089c56975716fb6788e6fafd06ffa8b42f620))

## [1.5.0](https://github.com/bastean/codexgo/compare/v1.4.0...v1.5.0) (2024-02-18)

### New Features

- add script to sync .env\* files ([e7fcc0b](https://github.com/bastean/codexgo/commit/e7fcc0b6355e5abf00a97526e7becb111cdf2dda))

### Bug Fixes

- **deps:** upgrade dependencies ([fecaafa](https://github.com/bastean/codexgo/commit/fecaafa9bf35e6a5fa71ae0468845bd32bef26ea))

## [1.4.0](https://github.com/bastean/codexgo/compare/v1.3.1...v1.4.0) (2024-02-15)

### New Features

- add commit message types to include in the changelog ([db06cf9](https://github.com/bastean/codexgo/commit/db06cf95d6d637f097a6745d04302b8f272a50a6))

### Bug Fixes

- **deps:** upgrade dependencies ([80c2256](https://github.com/bastean/codexgo/commit/80c22563516b5da15ea07475fbc94c4fcbffd5c6))

## [1.3.1](https://github.com/bastean/codexgo/compare/v1.3.0...v1.3.1) (2024-02-14)

### Bug Fixes

- **actions:** upgrade go setup action ([da7bc21](https://github.com/bastean/codexgo/commit/da7bc213a052d088efaac6b20c5ec5ad92f4d037))
- change live reload ([2f97bdb](https://github.com/bastean/codexgo/commit/2f97bdbe0675a747ac4eddcfe99632dcf0803b0f))

## [1.3.0](https://github.com/bastean/codexgo/compare/v1.2.0...v1.3.0) (2024-02-06)

### Features

- **actions:** add upgrade workflow ([e2d62d4](https://github.com/bastean/codexgo/commit/e2d62d4d76e56e0dfaabe0cbd474ef23ee1e5687))
- add script to upgrade dependencies ([a7cd088](https://github.com/bastean/codexgo/commit/a7cd088099d336526e00c6187a835f9938e48a55))
- **backend:** add secure middleware ([370db08](https://github.com/bastean/codexgo/commit/370db087b6df9ce1aaa3fa2e5589abc9756ec9b2))

### Bug Fixes

- **actions:** add commit push to upgrade workflow ([9ea06db](https://github.com/bastean/codexgo/commit/9ea06dbfdb1f2af22f187734d757fe2ad8b0e88a))
- **deps:** upgrade dependencies ([811345c](https://github.com/bastean/codexgo/commit/811345c603d2b07ed76f83620fc6386ea90d1861))
- **deps:** upgrade dependencies ([c99be30](https://github.com/bastean/codexgo/commit/c99be30ae77f766ca09dece90a627f657f8458c3))

## [1.2.0](https://github.com/bastean/codexgo/compare/v1.1.0...v1.2.0) (2024-01-28)

### Features

- **backend:** add rate limiter middleware ([a6c1b2b](https://github.com/bastean/codexgo/commit/a6c1b2b2a484d0b4ac63364b76c1ba18f8c3e4b3))

### Bug Fixes

- **deps:** upgrade modules dependencies ([d9851aa](https://github.com/bastean/codexgo/commit/d9851aaeb9ff510148935043ab446bea52e3dc26))
- remove go vet from lint-staged ([3869cbf](https://github.com/bastean/codexgo/commit/3869cbf84fb83bc105b16be0fb6f1a03ab830e9f))

## [1.1.0](https://github.com/bastean/codexgo/compare/v1.0.0...v1.1.0) (2024-01-22)

### Features

- **actions:** add brew setup ([ef7a00d](https://github.com/bastean/codexgo/commit/ef7a00de57e7cf524223f6e4ced5f7bf2ad71e55))
- add go vet on lint-staged ([8c52de4](https://github.com/bastean/codexgo/commit/8c52de4ace6d34c2174fe8f03c35e84b6a4040a5))
- add upx to compress binaries ([9d4e926](https://github.com/bastean/codexgo/commit/9d4e926a3b764f6fe2e49009fb69adc127acb7ea))
- **devcontainer:** add brew to simplify installation of tools ([8c77ed4](https://github.com/bastean/codexgo/commit/8c77ed45692b6303fcb6235b4ddb612d4e175505))
- **makefile:** add go mod tidy on lint rule ([d203639](https://github.com/bastean/codexgo/commit/d203639765560e1e375b77b9759bed581c2176ab))

### Bug Fixes

- **docker:** add optimization to compose ([0730183](https://github.com/bastean/codexgo/commit/0730183bfbc522f0c5278a733e63b346fbe41044))

## [1.0.0](https://github.com/bastean/codexgo/compare/v0.1.1...v1.0.0) (2024-01-17)

### ⚠ BREAKING CHANGES

- **readme:** Ready for v1

### Features

- add codexgo logos ([7ff0641](https://github.com/bastean/codexgo/commit/7ff0641a5db2df5f180242e3d05d93c1ba0cfc92))
- add trufflehog scan on lint-staged ([bdc473c](https://github.com/bastean/codexgo/commit/bdc473c56a446e0268c1ed02222af2a37185c244))
- **ci:** add tests job to workflow ([c033f54](https://github.com/bastean/codexgo/commit/c033f5429ddbe07f55281a921b706e6820537ffa))
- **devcontainer:** add cucumber extension ([fabb6d8](https://github.com/bastean/codexgo/commit/fabb6d8e990b87a71f0152cdeb43b6c28f3cd878))
- **docker:** add production compose ([4963296](https://github.com/bastean/codexgo/commit/49632964e8a238fd676204cd4eb0bff03a959ac7))

### Bug Fixes

- **backend:** add responsive to alerts ([fc1e4c8](https://github.com/bastean/codexgo/commit/fc1e4c80ba071edc7bfaf39a43bae4aaad8f5b1d))

### Documentation

- **readme:** add contributing section ([54e95f6](https://github.com/bastean/codexgo/commit/54e95f65a5deddd73e1021ad520c848a75ca29cc))

## [0.1.1](https://github.com/bastean/codexgo/compare/v0.1.0...v0.1.1) (2024-01-07)

## 0.1.0 (2024-01-07)

### Features

- **backend:** add basis to use htmx with tailwindcss ([5f260b5](https://github.com/bastean/codexgo/commit/5f260b5a594c0eaa50324d04f715f614145f7adc))
- **backend:** add crud endpoints ([08957ba](https://github.com/bastean/codexgo/commit/08957ba38446d9c3d52e75d225f5b77c1541c7f3))
- **backend:** add development dockerfile ([a8cef51](https://github.com/bastean/codexgo/commit/a8cef51c8f158fac22c71a567441b3efb49abfc8))
- **backend:** add pwa ([a370906](https://github.com/bastean/codexgo/commit/a3709064ba027b9c2e60cd157de55c95e41802a3))
- **context|backend:** add authentication to protected endpoints ([b582c11](https://github.com/bastean/codexgo/commit/b582c112f11e9a206fd037a2ecd7ea2bafa252ce))
- **context|backend:** add password hashing ([8264127](https://github.com/bastean/codexgo/commit/82641276144048800e1a99ac0806f3a1402f98a7))
- **context:** add basis to run use cases ([fa28f1f](https://github.com/bastean/codexgo/commit/fa28f1f87471e6c84dccdcebd35fc198bb46b96d))
- **context:** add crud use cases ([d503539](https://github.com/bastean/codexgo/commit/d5035397c7485e2826e4ed4cf594db2f32ef7145))
- **context:** add mongo repository adapter ([f25a793](https://github.com/bastean/codexgo/commit/f25a7931fab225edebc6e03a4f6f0a124a8ab05d))
- **devcontainer:** add prettier extension ([d833c1b](https://github.com/bastean/codexgo/commit/d833c1b62defdc6407534662358c89396948e7dc))

### Bug Fixes

- **ci:** upgrade actions ([3054b85](https://github.com/bastean/codexgo/commit/3054b85e405293668d0e2647b584e7cb0f815710))
- **release:** change manifest path ([05918f0](https://github.com/bastean/codexgo/commit/05918f0961fed086665c3e8572efcee5cdf9a025))
