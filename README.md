Humanize BOSH Manifest
======================

Convert a BOSH deployment manifest created by Spiff (keys are alphabetically sorted) into a readable format.

Usage
-----

```
humanize-manifest manifest.yml > readable_manifest.yml
```

Without an argument the it will use the currently targeted deployment manifest (from `~/.bosh_config`\)

```
bosh deployment path/to/manifest.yml
humanize-manifest
```

Install
-------

```
go get cloudfoundry-community/humanize-manifest
```
