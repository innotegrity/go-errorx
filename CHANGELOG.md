# Changelog

## Unreleased

No unreleased changes

## v1.1.0 (Released 2023-10-05)

* Changed `BaseError` to use internal variables to force creation via `NewBaseError()`
* Fixed `BaseError` class so that `nil` errors are never returned by `InternalError()`
  
## v1.0.15 (Released 2023-08-11)

* Initial stable release of the module
