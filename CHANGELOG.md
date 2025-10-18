# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.3.3

- Update Go version to 1.25.3

## v1.3.2

- Update Go version to 1.25.2
- Add golangci-lint configuration file (.golangci.yml)
- Enhance Makefile with additional quality checks (lint, gosec, trivy, osv-scanner)
- Add Trivy security scanner to GitHub Actions CI workflow
- Update development dependencies (golangci-lint, gosec, osv-scanner)
- Improve format target to handle vendor directories correctly
- Update copyright year to 2025

## v1.3.1

- Add comprehensive Go doc comments to all exported functions and types
- Create package documentation in doc.go with usage examples
- Enhance README.md with detailed API documentation and usage examples
- Improve .gitignore with comprehensive Go-specific entries
- Apply Go doc best practices throughout the codebase

## v1.3.0

- remove vendor
- go mod update

## v1.2.0

- add NotEmptySlice validation
- add NotEmptyString validation 

## v1.1.0

- add slice length validations
- go mod update

## v1.0.0

- Initial Version
