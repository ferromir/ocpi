# ocpi - Go types for the OCPI protocol

Module: `github.com/ferromir/ocpi`

This repository provides typed Go structs for OCPI protocol versions and the versions endpoint.
Each protocol version lives in its own sub-package, and `versions` contains the types for the /versions endpoint.

Packages:
- `versions`    : types for the versions endpoint (Version, Endpoint, VersionsResponse)
- `v2_1_1`      : example types for OCPI v2.1.1 (Response envelope + a sample Location)
- `v2_2`        : example types for OCPI v2.2 (Response envelope + a sample Tariff)
- `v2_3`        : example types for OCPI v2.3 (Response envelope + a sample CDR)

Struct fields include `json` and `validate` tags (using github.com/go-playground/validator/v10 style).

Generated files are examples to help you start modelling the full spec.
