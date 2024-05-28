# Changelog

## [4.0.0](https://github.com/bastean/codexgo/compare/v3.0.1...v4.0.0) (2024-05-28)

### ⚠ BREAKING CHANGES

- **server:** decouple service initializations
- **context:** change notification system workflow
- **server:** change acceptance tests to work with the new ui
- **server:** add fomantic-ui
- **server:** change error handling from panic to wrapped errors
- **context:** change package names in shared module
- **context:** change integration tests to check for wrapped errors instead of panic
- **context:** change unit tests to check for wrapped errors instead of panic
- **context:** change error handling from panic to wrapped errors

### Chores

- change air config ([a3b2f94](https://github.com/bastean/codexgo/commit/a3b2f94cb9f1f5b69ab719ae09001bac8382e7b2))
- change git ignore list ([122f7b2](https://github.com/bastean/codexgo/commit/122f7b2ba448cf79e741e5c7c3e3b7283a2dcaaf))
- change go version in mod file ([95ac107](https://github.com/bastean/codexgo/commit/95ac107f7b6dee5618ec459c911ca07440521c04))
- change makefile rules ([285b30b](https://github.com/bastean/codexgo/commit/285b30b7e1c88503681279e097d4a6e0abf8154c))
- **deps:** upgrade dependencies ([f804fc8](https://github.com/bastean/codexgo/commit/f804fc8c0e22f660407e80c826811455d9d13303))

### Documentation

- **readme:** add updated screenshots ([ffd6b17](https://github.com/bastean/codexgo/commit/ffd6b175b19292111ba5bff3611322c7ee8cdbb6))
- **readme:** change description ([122b14c](https://github.com/bastean/codexgo/commit/122b14c69174c5de0f9cc0cac54e4b55ea3a25d2))

### New Features

- **air:** enable live-reloading on the browser ([7714c38](https://github.com/bastean/codexgo/commit/7714c38cf0a74ef07dff07d232b98f462070a71a))
- **context:** add json marshal error handler to error bubble ([68819fe](https://github.com/bastean/codexgo/commit/68819fe9521117adb43504105b90f91abecebce4))
- **context:** add new terminal transport port adapter to notify module ([28fd1fe](https://github.com/bastean/codexgo/commit/28fd1fe865666b65f25d93e1f7b1895f2b40a998))
- **scripts:** add copy-deps script ([79e2d73](https://github.com/bastean/codexgo/commit/79e2d73989674b9a2d62841db87da89ef4bd8564))
- **server:** add accepts cookies nag ([682c370](https://github.com/bastean/codexgo/commit/682c370366139911fa145defa9127310503d86e9))
- **server:** add cookies cleaning ([30d4b9a](https://github.com/bastean/codexgo/commit/30d4b9aca620098e1ed4b9c0c0501003570f8007))
- **server:** add fomantic-ui ([738bf51](https://github.com/bastean/codexgo/commit/738bf5140c0dcb310c4effcbd3659720ba2c20a5))
- **server:** add log files ([141001a](https://github.com/bastean/codexgo/commit/141001a2d0bf0269a25030597c45d3a2b7c2f891))
- **server:** add missing error handlers ([99938f6](https://github.com/bastean/codexgo/commit/99938f6a5c00a47a5d4471d07354fe565a6c50f5))
- **server:** add popup to inform about account status ([77eb4a9](https://github.com/bastean/codexgo/commit/77eb4a917e1487a3736da346064befa0ff3d35c8))

### Bug Fixes

- add missing pointers ([b6b9343](https://github.com/bastean/codexgo/commit/b6b934305caecf51eea9c058e84d75cfaa2d353f))
- **deps:** upgrade dependencies ([4574f31](https://github.com/bastean/codexgo/commit/4574f3147dff657eb1672a97d0acea236bf8d5c4))
- **server:** add json unmarshal type error handler ([eb9eeb5](https://github.com/bastean/codexgo/commit/eb9eeb5e4442caf2abd07c294cd68a0c9db8f9b9))

### Refactors

- add field names at struct initialization ([55c5de3](https://github.com/bastean/codexgo/commit/55c5de3902be83a9066b9867e9493ad4ebab6f87))
- **context:** change error handling from panic to wrapped errors ([ec3245c](https://github.com/bastean/codexgo/commit/ec3245c9caf81562cfb8c2ae61aa16ee34b4d5e6))
- **context:** change exchange to router in broker model ([be19870](https://github.com/bastean/codexgo/commit/be198705206d0f0370b73dc5ed8bb9cf1c3d3572))
- **context:** change notification system workflow ([f7ec73c](https://github.com/bastean/codexgo/commit/f7ec73cd038337f63b792f1f2623285b7e0eb854))
- **context:** change package names in shared module ([e26da1e](https://github.com/bastean/codexgo/commit/e26da1e123692559f9a84a2ad3f0e53c4a1b1743))
- **context:** change time format in errors ([23362e4](https://github.com/bastean/codexgo/commit/23362e40d7ea9fafe2b83a120f3bbc6652644057))
- **context:** change type name of shared errors ([61c9b93](https://github.com/bastean/codexgo/commit/61c9b93fcec544e6d95c4e3ecd76a0f49ba38e43))
- **context:** rename folders using plural names instead of the prefix s in shared module ([fe5aabf](https://github.com/bastean/codexgo/commit/fe5aabf7b3815b60a48382076f7f453e065f15f2))
- **context:** rename packages using plural names in shared module ([db66e1d](https://github.com/bastean/codexgo/commit/db66e1d7e9c62ed8600f048512c37ec30bcba4ef))
- **makefile:** add MAKE variable to rules with a recursive recipe ([75e31a8](https://github.com/bastean/codexgo/commit/75e31a84e9b58b1723270bce738162240628e21d))
- **makefile:** change target names of test rules ([f078581](https://github.com/bastean/codexgo/commit/f0785811deabe3ab8343185cb14ca97933c6341c))
- rename files using flatcase ([28d3e5f](https://github.com/bastean/codexgo/commit/28d3e5fafad4cdbeed23a8c6bb51724e6e6f746e))
- **scripts:** change commit message on upgrade script ([9b257a2](https://github.com/bastean/codexgo/commit/9b257a223f1c2cc22dfe9c7dc5b1764389f3f8f1))
- **server:** add ui class in jquery component selectors ([3c1743e](https://github.com/bastean/codexgo/commit/3c1743ed59d0e35fe5b4d2e2e7d9a4e62066a4a5))
- **server:** change broker service components to individual files ([93d3c29](https://github.com/bastean/codexgo/commit/93d3c29e4669be36e5ed2f196fdb4e1183314a72))
- **server:** change error handling from panic to wrapped errors ([1e3d766](https://github.com/bastean/codexgo/commit/1e3d766be194b64960e302b83318a9929f104e5c))
- **server:** change error messages in services ([0f6a21e](https://github.com/bastean/codexgo/commit/0f6a21e98a89347d721f32b611fdacd6b405430d))
- **server:** decouple service initializations ([61961d2](https://github.com/bastean/codexgo/commit/61961d2327948df5b9611ff75034c28d9f34a859))

### Tests

- **context:** add spaces between definitions in setup test ([4318ea2](https://github.com/bastean/codexgo/commit/4318ea27f6cac20154489c55917b9d639bd54fe0))
- **context:** change integration tests to check for wrapped errors instead of panic ([6bb93ac](https://github.com/bastean/codexgo/commit/6bb93ac315e0049a2407952406e8349886274b2a))
- **context:** change time on expected error messages ([56137d6](https://github.com/bastean/codexgo/commit/56137d695a6663e18272db70c1d98dcf97c03140))
- **context:** change unit tests to check for wrapped errors instead of panic ([971b9de](https://github.com/bastean/codexgo/commit/971b9de188f998f9b6ecdec5435375be46789c9f))
- **server:** change acceptance tests to work with the new ui ([8df4c59](https://github.com/bastean/codexgo/commit/8df4c59c9c280d7c7a55285fa662b682b2df9469))

## [3.0.1](https://github.com/bastean/codexgo/compare/v3.0.0...v3.0.1) (2024-04-08)

### Bug Fixes

- **deps:** upgrade dependencies ([bd92cf7](https://github.com/bastean/codexgo/commit/bd92cf74fced77cb9011171e60f15d687ddc94f7))
- **makefile:** add phony target ([3c33a90](https://github.com/bastean/codexgo/commit/3c33a9005396067e0a2c130d78648a41bc677f73))
- **makefile:** remove init-ci rule ([3311e14](https://github.com/bastean/codexgo/commit/3311e144798aaff6dfe34df1c2c8aa9751f3ca68))

### Refactors

- **makefile:** change rules order ([95c6170](https://github.com/bastean/codexgo/commit/95c6170e4719d0702efe2ed75197c42cd3103494))

## [3.0.0](https://github.com/bastean/codexgo/compare/v2.0.1...v3.0.0) (2024-04-04)

### Documentation

- **readme:** add features ([6d36f5d](https://github.com/bastean/codexgo/commit/6d36f5d75dfc6e1e3cf9cb50ca01d6f7ab1a4b7a))

### New Features

- add account confirmation via email ([66f7b6e](https://github.com/bastean/codexgo/commit/66f7b6eda53e2f3ea897603c032e85f51fe6cf83))
- add event-driven architecture using rabbitmq ([1fd11cb](https://github.com/bastean/codexgo/commit/1fd11cb1b1b9096dc2aafccd2da8982d6d041279))
- add example env demo file ([c288d3c](https://github.com/bastean/codexgo/commit/c288d3ccdd7e4d348b915f22c9fa0236df7de247))
- add gracefully close infrastructure connections ([fb91c9a](https://github.com/bastean/codexgo/commit/fb91c9a569d7f2d6be67e983e19bb53ec5cb5191))

### Bug Fixes

- **deps:** upgrade dependencies ([a27389f](https://github.com/bastean/codexgo/commit/a27389f51cfc9b37d6b11dcc1013c2e02be84ea4))
- remove files generated by templ ([0adc9c6](https://github.com/bastean/codexgo/commit/0adc9c6cd570b672f8f9719ee0f3691b895804fd))

### Refactors

- change env handling in context to app ([f395933](https://github.com/bastean/codexgo/commit/f395933288ba2ad6fbb95eb683faa7195ebad890))
- change templ components ([e54c18a](https://github.com/bastean/codexgo/commit/e54c18a8421a1f5a62bb7ef0e2e58f90f2b50f4b))

### Tests

- add individual execution of unit, integration and acceptance tests ([4dc646f](https://github.com/bastean/codexgo/commit/4dc646f52794c0ad80803ce95900c0bf402029fd))
- remove shared value objects ([618ab5c](https://github.com/bastean/codexgo/commit/618ab5c990f2c54ec574380f54778fb64757ea2a))

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
