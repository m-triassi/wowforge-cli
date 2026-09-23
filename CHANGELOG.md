# [1.5.0](https://github.com/m-triassi/wowforge-cli/compare/v1.4.2...v1.5.0) (2026-09-23)


### Bug Fixes

* [#31](https://github.com/m-triassi/wowforge-cli/issues/31) delete all recorded addon folders on remove to prevent orphaned files ([ccb0b44](https://github.com/m-triassi/wowforge-cli/commit/ccb0b442fc56878a17495d9d6d55b88fd8e07b66))
* [#31](https://github.com/m-triassi/wowforge-cli/issues/31) retain addons with empty folder lists to prevent dropping migrated entries ([e835b2f](https://github.com/m-triassi/wowforge-cli/commit/e835b2fbe95f3814ab55a85795c829688364013a))
* [#31](https://github.com/m-triassi/wowforge-cli/issues/31) warn on remove when addon folders are unknown to prevent silent orphaned files ([ccdae72](https://github.com/m-triassi/wowforge-cli/commit/ccdae72af0eaefd40a249559b8b3123e2d8acee5))


### Features

* [#31](https://github.com/m-triassi/wowforge-cli/issues/31) return installed folder names from InstallAddon to enable accurate tracking ([496ac5c](https://github.com/m-triassi/wowforge-cli/commit/496ac5cd224ce7065e9ac0e1afa5d61041819368))
* [#31](https://github.com/m-triassi/wowforge-cli/issues/31) store installed addon folders in config to improve readability ([31db9d8](https://github.com/m-triassi/wowforge-cli/commit/31db9d8940898c019433dda5fa4aac14e5b7ac8f))



## [1.4.2](https://github.com/m-triassi/wowforge-cli/compare/v1.4.1...v1.4.2) (2026-09-21)


### Bug Fixes

* create parent directory for zip entries lacking dir entries (closes [#25](https://github.com/m-triassi/wowforge-cli/issues/25)) ([5d209e9](https://github.com/m-triassi/wowforge-cli/commit/5d209e914ed24636b81e08ede31f4c1d63716fe9))



## [1.4.1](https://github.com/m-triassi/wowforge-cli/compare/v1.4.0...v1.4.1) (2026-07-14)


### Bug Fixes

* **build:** add automated version and hash calculattion for nix on new releases ([7d96094](https://github.com/m-triassi/wowforge-cli/commit/7d960945bdd0a30ba21120dc16080f09e5adff1c))



# [1.4.0](https://github.com/m-triassi/wowforge-cli/compare/v1.3.0...v1.4.0) (2026-07-14)


### Bug Fixes

* [#28](https://github.com/m-triassi/wowforge-cli/issues/28) accomodate flavor when removing unneeded files ([5c12530](https://github.com/m-triassi/wowforge-cli/commit/5c12530e1d62e2cdfa38084db4598c983060c5d0))
* [#28](https://github.com/m-triassi/wowforge-cli/issues/28) account for flavor in download flow ([81ebc91](https://github.com/m-triassi/wowforge-cli/commit/81ebc91c4a883c1bf86ed4d400c2e5bf5290ab89))
* [#28](https://github.com/m-triassi/wowforge-cli/issues/28) add flavor specifier and return files by latest addon patch version ([e7c2487](https://github.com/m-triassi/wowforge-cli/commit/e7c248754c38545d4225b66ab2b2dc1d5e48f7a3))


### Features

* [#28](https://github.com/m-triassi/wowforge-cli/issues/28) allow users to set their game flavor via the game version id map ([581e1ca](https://github.com/m-triassi/wowforge-cli/commit/581e1ca3e59601179581f830e93db0823e98b5f3))



# [1.3.0](https://github.com/m-triassi/wowforge-cli/compare/v1.2.0...v1.3.0) (2026-05-23)


### Features

* **build:** [#26](https://github.com/m-triassi/wowforge-cli/issues/26) add nix build support for local dev and installation via flakes ([70ec29e](https://github.com/m-triassi/wowforge-cli/commit/70ec29ed38ab4595dd51a7696c20cf7329fcb03c))



# [1.2.0](https://github.com/m-triassi/wowforge-cli/compare/v1.1.0...v1.2.0) (2024-02-01)


### Features

* [#21](https://github.com/m-triassi/wowforge-cli/issues/21) add prints that shows what the application is doing during runtime ([a1b76fa](https://github.com/m-triassi/wowforge-cli/commit/a1b76faf300c5b1695bfb507bd8774a3b7f81cd0))



# [1.1.0](https://github.com/m-triassi/wowforge-cli/compare/v1.0.0...v1.1.0) (2023-12-23)


### Features

* [#19](https://github.com/m-triassi/wowforge-cli/issues/19) add automatic incrementing version flag ([68e14de](https://github.com/m-triassi/wowforge-cli/commit/68e14deeedac5ede115120af2bf90dee8e6d4074))



# [1.0.0](https://github.com/m-triassi/wowforge-cli/compare/v0.11.0...v1.0.0) (2023-12-20)


### Features

* add install instruction ([e310614](https://github.com/m-triassi/wowforge-cli/commit/e310614aa15a560bab29f2a1077852293bdd72f7))


### BREAKING CHANGES

* Finalize release, even though this is a docs change.



# [0.11.0](https://github.com/m-triassi/wowforge-cli/compare/v0.10.0...v0.11.0) (2023-12-20)


### Bug Fixes

* [#17](https://github.com/m-triassi/wowforge-cli/issues/17) add error handling to prevent trying to delete bad directories ([ffb784a](https://github.com/m-triassi/wowforge-cli/commit/ffb784a61aa67030880380c134fd22946bac3bff))
* [#17](https://github.com/m-triassi/wowforge-cli/issues/17) reverse equality sign to get correct result ([58c59b8](https://github.com/m-triassi/wowforge-cli/commit/58c59b84433274986026511fbcf0ed67a14045d2))


### Features

* [#17](https://github.com/m-triassi/wowforge-cli/issues/17) add basic filesystem package for managing folders ([d4d287b](https://github.com/m-triassi/wowforge-cli/commit/d4d287b3a9b535a454a2298be92a331fb9cfdb2b))
* [#17](https://github.com/m-triassi/wowforge-cli/issues/17) add command to remove installed addons ([0ff7045](https://github.com/m-triassi/wowforge-cli/commit/0ff7045383be7c02e7f454147cd3c72271bceebe))
* [#17](https://github.com/m-triassi/wowforge-cli/issues/17) add find command and deduplicate functionality ([e3f1194](https://github.com/m-triassi/wowforge-cli/commit/e3f1194a819e533d0e11bfa80c3dbe397868dba2))



# [0.10.0](https://github.com/m-triassi/wowforge-cli/compare/v0.9.0...v0.10.0) (2023-12-11)


### Features

* [#13](https://github.com/m-triassi/wowforge-cli/issues/13) add wider platform support ([b834dc6](https://github.com/m-triassi/wowforge-cli/commit/b834dc639c36f1d7d2aabcae6f1f85e9ba58cc54))



# Changelog

All notable changes to `wowforge-cli` will be documented in this file.
