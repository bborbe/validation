# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## Unreleased

- Update Go to 1.26.0

## v1.4.3

- Update Go from 1.25.5 to 1.25.7
- Update github.com/bborbe/errors from v1.5.1 to v1.5.2
- Update github.com/onsi/ginkgo/v2 from v2.27.5 to v2.28.1
- Update github.com/onsi/gomega from v1.39.0 to v1.39.1
- Update numerous indirect dependencies

## v1.4.1

- Update Go to 1.25.5
- Update golang.org/x/crypto to v0.47.0
- Update dependencies

## v1.4.0

- update go and deps

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
