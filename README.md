Humanize BOSH Manifest
======================

Convert a BOSH deployment manifest created by `bosh interpolate` (where keys
are sorted in alphabetical order) into a human-friendly readable format.

The latests BOSH v2.0 deployment manifest schema is primarily supported, but
v1 is also supported, and `bosh create-env` (formely known as `bosh-init`)
specificities are also supported.

This also basically works for old v1 manifests created by `spiff` or `spruce`
because v1 schema is supported.


Usage
-----

The command just takes a manifest as argument and writes the result to the
standard output.

```
humanize-manifest machine-created-manifest.yml > human-readable-manifest.yml
```

The output will vary based on the manifest, but here is an example with a
[SHIELD v8](https://github.com/starkandwayne/shield-boshrelease) deployment
manifests from the [Easy Foundry distribution](https://github.com/gstackio/gstack-bosh-environment):

```yaml
name: easyfoundry-shield-v8  # name comes first
instance_groups:
- name: shield               # name comes first
  instances: 1
  azs: [z1]                  # AZs are folded
  jobs:
  - name: core
    release: shield
    provides:
      shield:
        as: shield
        shared: true
    properties:
      # ...
  - name: shield-agent
    release: shield
    consumes:
      shield:
        instances:
        - address: shield-v8.easyfoundry.internal
        properties:
          domain: shield-v8.easyfoundry.internal
          port: 10443
    properties:
      core:
        ca: ((shield-tls.ca))
  stemcell: default
  vm_type: default
  vm_extensions:
  - shield-v8-loadbalancer
  persistent_disk_type: 1GB
  networks:
  - name: shield-v8-network
update:
  serial: true
  canary_watch_time: 1000-120000
  max_in_flight: 1
  update_watch_time: 1000-120000
variables:
- name: shield-agent-key
  type: ssh
- name: shield-ca
  type: certificate
  options:
    is_ca: true
    common_name: shieldca
- name: shield-tls
  type: certificate
  options:
    ca: shield-ca
    common_name: shield
    alternative_names:
    - 127.0.0.1
    - '*.shield.default.shield.bosh'
    - '*.shield.shield-v8-network.easyfoundry-shield-v8.bosh'
    - shield-v8.easyfoundry.prototyp.it
    extended_key_usage: [client_auth, server_auth]
- name: vault-ca
  type: certificate
  options:
    is_ca: true
    common_name: vaultca
- name: vault-tls
  type: certificate
  options:
    ca: vault-ca
    common_name: vault
    alternative_names:
    - 127.0.0.1
    - '*.vault.default.shield.bosh'
    - '*.vault.shield-v8-network.easyfoundry-shield-v8.bosh'
    extended_key_usage: [client_auth, server_auth]
releases:
- name: shield
  version: 8.0.8
  url: https://github.com/starkandwayne/shield-boshrelease/releases/download/v8.0.8/shield-8.0.8.tgz
  sha1: 55d1d6d8557f9b185fef7b5c6d73017b4c654f03
stemcells:
- alias: default
  os: ubuntu-trusty
  version: "3541.9"
```


Install
-------

```
go get github.com/cloudfoundry-community/humanize-manifest
```


Contributing
------------

If you modify this code, build it with:

```
go mod download
go build
```

And then test it with the `-d` debug flag:

```
humanize-manifest -d machine-created-manifest.yml 2>&1 | less
```
