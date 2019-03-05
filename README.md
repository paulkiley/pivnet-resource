# Pivnet Resource

Interact with [Pivotal Network](https://network.pivotal.io) from concourse.

## Installing

The recommended method to use this resource is with
[resource_types](http://concourse.ci/configuring-resource-types.html) in the
pipeline config as follows:

```yaml
---
resource_types:
- name: pivnet
  type: docker-image
  source:
    repository: pivotalcf/pivnet-resource
    tag: latest-final
```

Using `tag: latest-final` will automatically pull the latest final release,
which can be found on the
[releases page](https://github.com/pivotal-cf/pivnet-resource/releases).

**To avoid automatically upgrading, use a fixed tag instead e.g. `tag: v0.6.3`**

Releases are semantically versioned; these correspond to the git tags in this
repository.

## Source Configuration

```yaml
resources:
- name: p-mysql
  type: pivnet
  source:
    api_token: {{api-token}}
    product_slug: p-mysql
```

* `api_token`: *Required.*
  Token from your Pivotal Network profile. Accepts either your Legacy API Token or UAA Refresh Token.

* `product_slug`: *Required.*
  Name of product on Pivotal Network.

* `release_type`: *Optional.*
  Lock to a specific release type.

* `copy_metadata`: *Optional.*
  Set to `true` to copy metadata from the latest All Users release within the minor. Defaults to `false`.
  The following metadata is copied:

  * Release Notes URL
  * End of General Support
  * End of Technical Guidance
  * End of Availability
  * EULA
  * License Exception
  * ECCN
  * Controlled
  * Dependency Specifiers
  * Upgrade Path Specifiers

* `endpoint`: *Optional.*
  Endpoint of Pivotal Network.

  Defaults to `https://network.pivotal.io`.

* `product_version`: *Optional.*
  Regex to match product version e.g. `1\.2\..*`.

  Empty values match all product versions.

* `sort_by`: *Optional.*
  Mechanism for sorting releases.

  Defaults to `none` which returns them in the order they come back from Pivotal Network.

  Other permissible values for `sort_by` include:
  - `semver` - this will order the releases by semantic version,
    returning the release with the highest-valued version.

## Example Pipeline Configuration

See [example pipeline configurations](https://github.com/pivotal-cf/pivnet-resource/blob/master/examples).

## Behavior

### `check`: Check for new product versions on Pivotal Network.

Discovers all versions of the provided product.
Returned versions are optionally filtered and ordered by the `source` configuration.

### `in`: Download the product from Pivotal Network.

Downloads the provided product from Pivotal Network. You will be required to accept a EULA for any product you're downloading for the first time.

The metadata for the product is written to both `metadata.json` and
`metadata.yaml` in the working directory (typically `/tmp/build/get`).
Use this to programmatically determine metadata of the release.

See [metadata](https://github.com/pivotal-cf/pivnet-resource/blob/master/metadata)
for more details on the structure of the metadata file.

#### Parameters

* `globs`: *Optional.* Array of globs matching files to download.

  If multiple files are matched, they are all downloaded.
  - The globs match on the actual *file names*, not the display names in Pivotal
  Network. This is to provide a more consistent experience between uploading and
  downloading files.
  - If the globs fail to match any files the release download fails
  with error.
  - If one or more globs fails to match any files, only the matched files will be downloaded.
  - If `globs` is not provided (or is nil), **all files will be downloaded**.
  - Setting `globs` to the empty array (i.e. `globs: []`) will not attempt to
  download any files.
  - Files are downloaded to the working directory (e.g. `/tmp/build/get`) and the
  file names will be the same as they are on Pivotal Network - e.g. a file with
  name `some-file.txt` will be downloaded to `/tmp/build/get/some-file.txt`.

* `unpack`: *Optional.* Whether to unpack the downloaded file.  
  This can be used to use a root filesystem that is packaged as a archive file on network.pivotal.io as the image to run a given concourse task

  Example of how to unpack with `get` and pass as image to task definition

  ```yaml
  resource:
  - name: image
    type: pivnet
    source:
      api_token: {{pivnet_token}}
      product_slug: {{image-slug}}
      product_version: 0\.0\..*

  jobs:
  - name: sample
    serial: true
    plan:
    - get: tasks
    - get: image
      resource: pcf-automation
      params:
        globs: ["image-*.tar"]
        unpack: true

    - task: say hello
      image: image
      file: tasks/say-hello.yml
  ```

### `out`: Upload a product to Pivotal Network.

Creates a new release on Pivotal Network with the provided version and metadata.

It can also upload one or more files to Pivotal Network bucket and calculate the MD5 checksum locally for each file in order to add MD5 checksum to the file metadata in Pivotal Network.

**Existing product files with the same AWS key will no longer be deleted and recreated.**

**If you want to associate an existing product file with a new release, you can do so by specifying the existing AWS key when creating the release. This will no longer break past release associations.**

**If the AWS key matches an existing file, but the SHA does not match, you will now receive an error and need to rename the file.**

**Existing releases with the same version will _not_ be deleted and recreated by
default, and will instead result in an error.**

See [metadata](https://github.com/pivotal-cf/pivnet-resource/blob/master/metadata)
for more details on the structure of the metadata file.

#### Parameters

* `file_glob`: *Optional.* Glob for matching files to upload.

  If multiple files are matched by the glob, they are all uploaded. If no files are matched, release creation fails with an error.

* `metadata_file`: *Required.*
  File containing metadata for releases and product files.

  See [metadata](https://github.com/pivotal-cf/pivnet-resource/blob/master/metadata)
  for more details on the structure of the metadata file.

* `skip_product_file_polling`: *Optional.*
  Boolean. Skip product file validation checks after upload. Pivotal Network still validates the files asynchronously
  but _waiting_ for the results will not happen as part of the `put:` process. **Note:** All associated product files
  in a release must still clear validation before the release can be promoted from _Admins Only_ visibility.

* `override`: *Optional.*
  Boolean. Forces re-upload of releases of releases and versions that are
  already present on the Pivotal Network.

### Some Common Gotchas

#### Using Glob Patterns Instead of Regex Patterns

We commonly see `product_version` patterns that look something like these:

```yaml
product_version: Go*          # Go Buildpack
#....
product_version: 1\.12\.*       # ERT
```

These superficially resemble Globs, not Regexes. They will generally work, but not because they
are a glob. They work because the regex will also match.

For example, the first pattern, `Go*` will match "Go Buildpack 1.1.1". But it would also match
"Goooooooo" or "Go Tell It On A Mountain". The second pattern, `1\.12\.*`, will match "1.12.0".
But it will also match "1.12........." and "1.12.notanumber"

Instead, try patterns like:

```yaml
product_version: Go.*\d+\.\d+\.\d+  # Go Buildpack
#....
product_version: 1\.12\.\d+         # ERT
```

Note that [the regex syntax is Go's, which is slightly limited](https://github.com/google/re2/wiki/Syntax)
compared to PCRE and other popular syntaxes.

#### Using `check-resource` for sorted but non-sequential releases (eg. Buildpacks, Stemcells)

When doing a `check`, pivnet-resource defaults to using the server-provided order. This works
fine for simple cases where the response from the server is already in semver order. For example,
imagine this order from a product:

```
1.12.3
1.12.2
1.12.1
1.12.0
1.11.4
1.11.3
1.11.2
1.11.1
1.11.0
```

This list is in descending semver order. All the 1.12 patch releases are together, followed
by all the 1.11 patch releases and so on.

Some products do not group into major or major.minor groups in their responses. This is usually
because a product has multiple concurrent version releases. For example, Stemcells typically
have multiple major versions available. When a CVE is announced affected them, multiple releases
occur at once, giving a order like:

```
9999.21
7777.19
9999.20
7777.18
```

In this example, the Stemcell versions 9999 and 7777 are _sorted_ but not _sequential_.

To fix, use `sort_by: semver` in your resource definition.

Note: Buildpack "versions" are actually a name and a version combined. You'll need to escape spaces
in your `check-resource` command for it to work properly. Eg:

```
fly -t pivnet check-resource \
  --resource pivnet-resource-bug-152616708/binary-buildpack \
  --from product_version:Binary\ 1.0.11#2017-03-23T13:57:51.214Z
```

In this example we escaped the space between "Binary" and "1.0.11".

## Integration Environment

The Pivotal Network team maintain an integration environment at
`https://pivnet-integration.cfapps.io/`.

This environment is useful for teams to develop against, as changes to products
in this account are separated from the live account.

Example configuration for the integration environment:

```yaml
resources:
- name: p-mysql
  type: pivnet
  source:
    api_token: {{api-token}}
    product_slug: p-mysql
    endpoint: https://pivnet-integration.cfapps.io
```

## Developing

### Prerequisites

A valid install of golang is required - version 1.7.x is tested; earlier
versions may also work.

### Dependencies

We use [go modules](https://github.com/golang/go/wiki/Modules) for dependencies, so you will have to make sure to turn them on with `GO111MODULE=on`.


### Running the tests

Install the ginkgo executable with:

```
go get -u github.com/onsi/ginkgo/ginkgo
```

The tests require a valid Pivotal Network API token and valid AWS S3 configuration.

Refer to the
[official docs](https://network.pivotal.io/docs/api#how-to-authenticate)
for more details on obtaining a Pivotal Network API token.

It is advised to run the acceptance tests against the Pivotal Network integration
environment endpoint `https://pivnet-integration.cfapps.io`.

Run the tests with the following command (dummy values must be replaced by user):

```
PRODUCT_SLUG=my-product-slug-eg-pivotal-diego-pcf \
API_TOKEN=my-token \
PIVNET_ENDPOINT=some-pivnet-endpoint \
PIVNET_RESOURCE_REFRESH_TOKEN=some-pivnet-resource-token \
./bin/test
```

### Contributing

Please make all pull requests to the `master` branch, and
[ensure the tests pass locally](https://github.com/pivotal-cf/pivnet-resource#running-the-tests).
