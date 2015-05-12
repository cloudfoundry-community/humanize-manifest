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

The output will vary based on the manifest, but here is an example from a bosh-lite CF manifest:

```
$ humanize-manifest | head -n 20
meta:
  default_env:
    bosh:
      password: $6$4gDD3aV0rdqlrKC$2axHCxGKIObs6tAmMTqYCspcdvQXh3JJcvWOY2WGb4SrdXtnCyNaWlrf3WEqvYR2MYizEGp3kMmbpwBC6jsHt0
  environment: cf-warden
  releases:
  - name: cf
    version: latest
  stemcell:
    name: bosh-warden-boshlite-ubuntu-trusty-go_agent
    version: latest
name: cf-warden
director_uuid: c6f166bd-ddac-4f7d-9c57-d11c6ad5133b
releases:
- name: cf
  version: "200"
compilation:
  workers: 6
  network: cf1
  reuse_compilation_vms: true
```

Install
-------

```
go get github.com/cloudfoundry-community/humanize-manifest
```
