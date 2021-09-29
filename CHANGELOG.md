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

## [0.2.0] Unreleased

- `Changed` template.Create() renamed in template.CreateRow() for readability.
- `Fixed` row.Raw() now return the correct value (a map[string]interface{} of raw values).
- `Fixed` conversions error no longer ignored by unmarshaling.

## [0.1.0] 2021-09-21

- `Added` First official version of `jl` command line.
