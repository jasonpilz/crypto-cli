# Crypto CLI

> Command line interface for cryptocurrency prices and data

## Usage

> [WIP]

## Releasing

> This project follows [semantic versioning](semver.org).

When it's time to create a new release version:

1. `make git/changes` to review changes since the last release.
   Based on the changes, decide what kind of release is required (bugfix, feature or breaking).

2. `BUMP=(major|minor|patch|bugfix|feature|breaking) make git/tag` to create a new git tag.
   (bugfix, feature and breaking are aliases for semver's patch, minor and major).
   BUMP defaults to patch. The command assumes you have a remote repository named
   `origin` pointing to this repository.  If you'd prefer to specify a different remote
   repository, you can do so by setting `ORIGIN=(preferred remote name)`.

3. `make github/release` to create a new release based on the most recent git tag.
   This step will build application binaries for each os and architecture, docker images
   based on the binary, upload the images to a docker registry and the binaries to the github release.

### Prerequisites

* [goreleaser](https://goreleaser.com/install/)
* [docker](https://docs.docker.com/install/)
* `GITHUB_TOKEN` defined in `.env.goreleaser` with `repo` access.
