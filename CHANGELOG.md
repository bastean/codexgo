# Changelog

## [4.6.1](https://github.com/bastean/codexgo/compare/v4.6.0...v4.6.1) (2024-08-03)

### Bug Fixes

- **server:** relocate unreachable logs ([1f2df6a](https://github.com/bastean/codexgo/commit/1f2df6ab02ec362a27233344a9f2035e374bfa00))

## [4.6.0](https://github.com/bastean/codexgo/compare/v4.5.0...v4.6.0) (2024-08-03)

### Chores

- **deps:** upgrade ([240d8d4](https://github.com/bastean/codexgo/commit/240d8d4ed5334cb8b6561773752c9a83184101b4))
- **docker:** add ignore file ([9bbaa59](https://github.com/bastean/codexgo/commit/9bbaa598143dfaabb69988a0c503c4550abe004e))

### Documentation

- **readme:** add script to initialize repository from zip file ([78c843b](https://github.com/bastean/codexgo/commit/78c843b12e9bde9870bc758eef2f47783d592171))

### New Features

- **devcontainer:** add ssh server ([5492305](https://github.com/bastean/codexgo/commit/54923055d13b4bde924b65dd10f7d8f95bfb6d49))

### Refactors

- **docker:** rename default network ([494b3b4](https://github.com/bastean/codexgo/commit/494b3b4ad5796f3757ea03bf943c592772e837bc))
- **makefile:** rename targets ([9093e2d](https://github.com/bastean/codexgo/commit/9093e2db7ba1f0f4d1e3b4ceb5b2167fc8665906))
- **server:** separate views from api endpoints ([5fdd1ea](https://github.com/bastean/codexgo/commit/5fdd1ea6897493c06bf62f9199ca35cb6454adff))

## [4.5.0](https://github.com/bastean/codexgo/compare/v4.4.0...v4.5.0) (2024-07-30)

### ⚠ BREAKING CHANGES

- **deployments:** rename envs

### Chores

- change git ignore list ([f93dea2](https://github.com/bastean/codexgo/commit/f93dea2bb31c4a1d1dda834a832533a18ba44612))
- **deps:** upgrade ([df92fda](https://github.com/bastean/codexgo/commit/df92fda92b054363249a9ef86440ad5910fff812))
- **deps:** upgrade ([c994a97](https://github.com/bastean/codexgo/commit/c994a97ee5e635b37aa7181a22d97010ae22aaa0))
- **makefile:** add set of rules to install tools ([94f0b8f](https://github.com/bastean/codexgo/commit/94f0b8fb8dcc7241bdd6c94013503bb9a9eed74a))
- **mod:** upgrade ([31e32d2](https://github.com/bastean/codexgo/commit/31e32d2e250e58d5b1c5c93507ba33fe55bf33ab))

### Documentation

- **readme:** add basic idiomatic ([1af2146](https://github.com/bastean/codexgo/commit/1af2146739b07d3f4682ea56bbefb914b1415d7b))

### New Features

- **logger:** add colored outputs ([9e9d041](https://github.com/bastean/codexgo/commit/9e9d0415dfe982294405f5d5323ddc1fa8d7c22b))
- **server:** add initial health check endpoint ([95f2117](https://github.com/bastean/codexgo/commit/95f2117797593780f2c77e6cd5900beb1be650b3))
- **server:** add proxy check ([367773c](https://github.com/bastean/codexgo/commit/367773ce8ea859fe62c81b05e148be2dcb99101d))

### Refactors

- change format of messages in errors ([dda8c9a](https://github.com/bastean/codexgo/commit/dda8c9ae2dfd919a15cfce597356e8026a20b8ce))
- **deployments:** rename envs ([a3186bd](https://github.com/bastean/codexgo/commit/a3186bd899904d237a20e347940115796e22702f))
- **dockerfile:** use makefile targets to install the tools ([c77f092](https://github.com/bastean/codexgo/commit/c77f0926af64ec0dc9b6240c61f1a05d27dbdb9e))
- rename functions according to their layer ([e37b64c](https://github.com/bastean/codexgo/commit/e37b64c6a7744c81f995f821dbbed91b87085dc1))
- **scripts:** reuse command execution in upgrade ([c48eec6](https://github.com/bastean/codexgo/commit/c48eec66c1143917cfc750fc628407d416246479))
- **server:** change error handling ([a789393](https://github.com/bastean/codexgo/commit/a789393f07ce92914c27096cdaee7a745de2aba3))
- **service:** organize envs ([71f4f4c](https://github.com/bastean/codexgo/commit/71f4f4ca29b9c6449612be18c990d8a6592cdece))
- **service:** use aliases of context types ([51d7701](https://github.com/bastean/codexgo/commit/51d770172f757a4283bed9434d3c13cb8f89dabb))

### Tests

- **acceptance:** assert errors directly ([6471fe2](https://github.com/bastean/codexgo/commit/6471fe2a71f6ec9c235c53fb880bb422e9724b20))

## [4.4.0](https://github.com/bastean/codexgo/compare/v4.3.1...v4.4.0) (2024-07-14)

### ⚠ BREAKING CHANGES

- **context:** rename packages
- **infrastructure:** rename packages
- rename envs
- add internal folder to manage apps and services

### Chores

- **deps:** upgrade dependencies ([dd29c51](https://github.com/bastean/codexgo/commit/dd29c5141e6f43b8fc9c4cb52806074d64cd6c5a))

### Documentation

- **readme:** add more details about the interaction between layers ([99476ba](https://github.com/bastean/codexgo/commit/99476ba7c1d9f01bc8d42d59f5a417c1c43b59c2))

### New Features

- **server:** show proxy port when running ([42c84ed](https://github.com/bastean/codexgo/commit/42c84edb2a2d4f37245b31da512e8e4f924a9fd8))

### Refactors

- add internal folder to manage apps and services ([6081521](https://github.com/bastean/codexgo/commit/6081521f2f4496fcc4b3e162c31e5f9bb8b8c265))
- **cmd:** stop apps before services ([0b52af7](https://github.com/bastean/codexgo/commit/0b52af71e25e7d0c2e98c0fb7c42cdd9f0148b66))
- **context:** rename packages ([07f39eb](https://github.com/bastean/codexgo/commit/07f39eb8b477460c3ba15ee99fd893e8e9384232))
- **context:** replace use of generic models with specific ones ([6ee85b2](https://github.com/bastean/codexgo/commit/6ee85b2aeb12d037e639cdd20c6696b38ae72349))
- **domain:** add err prefix to bubble errors ([814a73b](https://github.com/bastean/codexgo/commit/814a73b55118751a5d3c7db77a3aa3fd364c57cb))
- **infrastructure:** rename packages ([aa6b04d](https://github.com/bastean/codexgo/commit/aa6b04d42348cec42425eea5375b63328854eac1))
- rename envs ([12ab31d](https://github.com/bastean/codexgo/commit/12ab31d7aeadb6129df463f558add5896b519a8b))

### Tests

- **infrastructure:** add assertions for missing mongo errors ([d42e73e](https://github.com/bastean/codexgo/commit/d42e73e45d8b781b9b88c52edf6d7528e4effe4e))

## [4.3.1](https://github.com/bastean/codexgo/compare/v4.3.0...v4.3.1) (2024-07-01)

### Chores

- **deps:** upgrade dependencies ([1011a81](https://github.com/bastean/codexgo/commit/1011a814c472f1d00b75aa125e05f8b21721db8d))

### Bug Fixes

- **makefile:** remove previous production docker image ([bcd39a0](https://github.com/bastean/codexgo/commit/bcd39a0d1b51eb6314e59c6239454d98dfeaa350))
- **templ:** resolve imported and not used error ([84cbd69](https://github.com/bastean/codexgo/commit/84cbd69f90eccc3909e0ae620876bc3a1729acab))

## [4.3.0](https://github.com/bastean/codexgo/compare/v4.2.1...v4.3.0) (2024-06-26)

### Chores

- **deps:** upgrade dependencies ([bce72d8](https://github.com/bastean/codexgo/commit/bce72d872897ad527dd3e435fe45fe5f93300d81))
- **deps:** upgrade jwt to v5 ([18054b7](https://github.com/bastean/codexgo/commit/18054b747b8250e7a169359dc1f54287328d0edc))

### Documentation

- add scanners ([a34ea61](https://github.com/bastean/codexgo/commit/a34ea616ff1425c3053115df604e64ef9609e283))

### New Features

- add trivy and osv scanners ([6fe3c53](https://github.com/bastean/codexgo/commit/6fe3c53255d3fb68dbbda985a0c15e7b9b68ae5c))

### Refactors

- **context:** change domain message components ([157ef5b](https://github.com/bastean/codexgo/commit/157ef5bcd3a54c9784ddcdd6339e9223da735a38))
- **context:** rename variables in value objects ([c8f46a1](https://github.com/bastean/codexgo/commit/c8f46a1e857c25d38a72da5a9f30273c8dc64e3d))
- **scripts:** change panic on error ([4577790](https://github.com/bastean/codexgo/commit/4577790af035df66e18aa96ee8289c30409a1687))
- **server:** rename service logs ([65bd07c](https://github.com/bastean/codexgo/commit/65bd07c900686326c91d5fba163c501348a8589e))
- **server:** reorganize services ([f559715](https://github.com/bastean/codexgo/commit/f5597153bd121cce16d77aecba2cf0e268bcf1ce))

### Tests

- **context:** add assertion for duplication error in mongo ([201fd53](https://github.com/bastean/codexgo/commit/201fd53f5970546dddaf62c2b21d50a35841052c))
- **context:** add assertion for omitted json errors ([a08ea38](https://github.com/bastean/codexgo/commit/a08ea38efd3fa0d191505f48de571a234c9b2217))

## [4.2.1](https://github.com/bastean/codexgo/compare/v4.2.0...v4.2.1) (2024-06-19)

### Bug Fixes

- add boolean format verb in strings ([e317911](https://github.com/bastean/codexgo/commit/e317911d22e1798a4bc0a1c917089b06d8228a35))
- remove default format verb in strings ([57beb5e](https://github.com/bastean/codexgo/commit/57beb5e3339a6ac9685e47bd2c798fdcccf619b2))

## [4.2.0](https://github.com/bastean/codexgo/compare/v4.1.1...v4.2.0) (2024-06-17)

### Chores

- **deps:** upgrade dependencies ([7b19ee7](https://github.com/bastean/codexgo/commit/7b19ee7dee15d14a7fad328977d090c031ef2964))

### New Features

- **context:** add handling of omitted errors ([f70c276](https://github.com/bastean/codexgo/commit/f70c2767a368859560c1c2986411384ee36ae92e))

### Refactors

- add default format verb to strings ([c459caa](https://github.com/bastean/codexgo/commit/c459caaf4c6fb7c5aede730caad10ab0093e3a3e))
- change panic on error ([1fc83fa](https://github.com/bastean/codexgo/commit/1fc83fa7b6433059e098bbe8a21762b7eb95000a))
- **context:** remove notify module ([064815f](https://github.com/bastean/codexgo/commit/064815f9efb7b75638c4f2d88efe1a27cacecc80))
- **context:** remove redundant details from type names ([27a666a](https://github.com/bastean/codexgo/commit/27a666a11bf7232df9d9ba34c20f1d15e03abade))

### Tests

- **context:** add handling of unexpected errors in mothers to avoid flaky tests ([2fbea22](https://github.com/bastean/codexgo/commit/2fbea223fada03944b1deb6a1b83f2f7df879e93))

## [4.1.1](https://github.com/bastean/codexgo/compare/v4.1.0...v4.1.1) (2024-06-12)

### Bug Fixes

- **makefile:** add pipefail to return an error when a test fails ([5b4c26e](https://github.com/bastean/codexgo/commit/5b4c26e4621fec259eb1b6cfa0bd263534b50588))

## [4.1.0](https://github.com/bastean/codexgo/compare/v4.0.0...v4.1.0) (2024-06-10)

### Chores

- **deps:** upgrade dependencies ([93fc426](https://github.com/bastean/codexgo/commit/93fc4264ee2f2937e07e9a8f9c96156df4dfb4f1))

### Documentation

- **readme:** add basic layers workflow ([b6f6d5d](https://github.com/bastean/codexgo/commit/b6f6d5d0f8b6bb75759a8a0744f86bd5abc44893))

### New Features

- **makefile:** add tee in test rules ([1d21d7a](https://github.com/bastean/codexgo/commit/1d21d7a31fac96db750c9266e181f93765ca1089))

### Bug Fixes

- **dockerfile:** update air module name ([71bc376](https://github.com/bastean/codexgo/commit/71bc376c7a6e834d16eaf10684806ca652bd3e51))

### Refactors

- add type alias ([f55bb9d](https://github.com/bastean/codexgo/commit/f55bb9d1fbe3c933c1bb48e6885270c92e69beee))
- **context:** add pointer to search criteria type ([8648a18](https://github.com/bastean/codexgo/commit/8648a184464c355d5b6216a4cb6a2f58c4bc1b95))
- **context:** change empty type from struct to interface ([4e5dcf0](https://github.com/bastean/codexgo/commit/4e5dcf0152a5eec35818edf90030ae449037e0a2))
- **context:** change errors in shared module ([47fe621](https://github.com/bastean/codexgo/commit/47fe62172d93792694beef6df3145502690e8d6a))
- **context:** change parameters to use primitive type in user module ([db8fc5b](https://github.com/bastean/codexgo/commit/db8fc5b12b24968677f2a452e703a5af3192020f))
- **context:** change updates in user module ([ba294af](https://github.com/bastean/codexgo/commit/ba294aff99ac5a6d35b2701531d97bf16f6191fd))
- squash struct fields ([8ccc22c](https://github.com/bastean/codexgo/commit/8ccc22c0ed57d02e1717d723462673ee1ebc55d3))

### Tests

- **context:** add more explicit test case names ([5027a31](https://github.com/bastean/codexgo/commit/5027a31a0fafcdbcac17ad09a04b557e145c5a55))

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
