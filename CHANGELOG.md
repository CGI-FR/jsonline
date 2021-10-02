# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Types of changes

- **`Added`** for new features.
- **`Changed`** for changes in existing functionality.
- **`Deprecated`** for soon-to-be removed features.
- **`Removed`** for now removed features.
- **`Fixed`** for any bug fixes.
- **`Security`** in case of vulnerabilities.

## [0.2.0] Unreleased

- **`Added`** cast values to specific raw type on row import with `template.WithMapped*` or `value.NewValue(v interface{}, f Format, rawtype interface{})`.
- **`Added`** `value.GetRawType()` return the raw type that will be used on import.
- **`Added`** new package `cast` to expose all type cast-related functions.
- **`Changed`** `template.Create()` renamed in `template.CreateRow()` for readability.
- **`Changed`** default raw type and export value for `Number` format is `json.Number`.
- **`Changed`** module name to `github.com/cgi-fr/jsonline`.
- **`Changed`** exporting `int32` or `rune` to `String` format will no longer convert the value to the corresponding unicode character, but rather print the numeric value.
- **`Changed`** exporting to `Binary` format will no longer convert the value to a `string` before base64 encoding, but will be encoded to in-memory byte representation.
- **`Changed`** exporting or importing to `Timestamp` format will always produce an `int64` (if not specified otherwise).
- **`Fixed`** `Format` type now correctly exported by jsonline package.
- **`Fixed`** `row.Raw()` now return the correct value (a `map[string]interface{}` of raw values).
- **`Fixed`** conversions error no longer ignored by unmarshaling.

## [0.1.0] 2021-09-21

- **`Added`** First official version of `jl` command line.
