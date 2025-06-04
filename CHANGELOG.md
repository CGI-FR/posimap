# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Types of changes

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [0.3.0]

- `Added` comp-3 numeric encoding support

## [0.1.0]

- `Added` fold to jsonline and unfold back to fixed-width file.
- `Added` schema definition file in YAML format.
- `Added` ability to reference an external schema file.
- `Added` support for nested records, occurrences, and redefines.
- `Added` charset conversion, supporting most mono-byte charsets.
- `Added` fold command automatically trim spaces (can be disabled via the `--notrim` flag).
- `Added` unfold command automatically cut overflow runes.
- `Added` graph and charsets commands.
