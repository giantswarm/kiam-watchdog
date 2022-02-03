# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Added --probe-mode flag to allow using either 'route53' or 'sts' to probe AWS API.

## [0.4.0] - 2021-10-06

### Added

- Add node-selector and tolerations.

## [0.3.0] - 2021-07-30

### Changed

- Use `.Values.kiam.region` as field to get AWS region.

## [0.2.0] - 2021-07-30

### Changed

- Remove default value for AWS region.

## [0.1.1] - 2021-07-29

### Added

- Add push of docker image to aliyun.

## [0.1.0] - 2021-07-26

[Unreleased]: https://github.com/giantswarm/kiam-watchdog/compare/v0.4.0...HEAD
[0.4.0]: https://github.com/giantswarm/kiam-watchdog/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/giantswarm/kiam-watchdog/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/kiam-watchdog/compare/v0.1.1...v0.2.0
[0.1.1]: https://github.com/giantswarm/kiam-watchdog/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/giantswarm/kiam-watchdog/releases/tag/v0.1.0
