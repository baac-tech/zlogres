# Changelog

## v0.2.4

- [Edited] Updated dependencies
- [Edited] Updated module version

## v0.2.3

- [Edited] Updated dependencies
- [Edited] Updated module version

## v0.2.2

- [Edited] Updated dependencies

## v0.2.1

- [Edited] Updated dependencies

## v0.2.0

- [Edited] Updated dependencies
- [Edited] Renamed module name for github public repo

## v0.1.7

- [Added] `ResponseStatusTextTag` in log message, for showing http status message (key_name: "status_text")
- [Added] `TimeUsageUnitTag` in log message, for showing time unit usage (key_name: "elapsed_time_unit")

## v0.1.6

- [Edited] Changed log template by encapsulated all thing into `EventTag` (key_name: "event")
- [Edited] New configurable of `ContextMessageKey` (default: "message")
- [Edited] Moved `response_time` into `ResponseTimeTag`

## v0.1.5

- [Edited] New configurable of `LogLevel` (default: "info") and `ElapsedTimeUnit` (default: "micro")
- [Edited] Fixed `TimeFieldFormat` is `time.RFC3339Nano`

## v0.1.4

- [Edited] - New module and package name for ipanda

- ## v0.1.3

- [Edited] `README.md` for supporting local repository
- [Edited] converted `took` into `elasped_time`

## v0.1.2

- [Edited] changed repository to ipanda


## v0.1.1

- [Edited] removed unused go mod package
- [Edited] beautify code

## v0.1.0

- [Initial] zlogres, the middleware for fiber. Supported with `requestid` middleware and can be custom context key with config `RequestIDContextKey`
