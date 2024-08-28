# Developing on this project

## Setting up a local development environment

The goal here is to have a setup that is consistent with both your colleagues' and this repositories automation,
allowing for coherence best practices and a fast feedback loop.

### Minimal setup

Apart from a working Golang setupo, the `pre-commit` tools is mandatory. It will guard you from creating inconsistencies
between you and other team members, lower CI costs by limiting the necessary number of runs and even do some basic
security checks.

To install `pre-commit` and its dependencies, please refer the following docs:

1. [Pre-commit framework installation](https://pre-commit.com/#install)
2. [golangci-lint installation](https://golangci-lint.run/usage/install/)
3. [helm-docs installation](https://github.com/norwoodj/helm-docs#installation)

Once you have the dependencies installed, you have to register `pre-commit` as your local hooks for the git repo
(you need to do this only once):

```nohighlight
pre-commit install --install-hooks
```

To run all the `pre-commit` checks without creating a commit, run:

```nohighlight
pre-commit run -a
```

### Setup for full build workflow

Install these additional tools, if you don't have them installed yet:

- [Docker](https://www.docker.com/products/docker-desktop/)
- [GoReleaser](https://goreleaser.com/)
- [syft](https://github.com/anchore/syft)
- [Cosign](https://github.com/sigstore/cosign)
- [app-build-suite](https://github.com/giantswarm/app-build-suite/releases) - download the latest script (dockerized)

Now, to build the Go binary and Docker image locally, run:

```nohighlight
goreleaser release --verbose --snapshot
```

To build the Helm chart, run:

```nohighlight
dabs.sh -c helm
```

## Creating a release

To create a release a new version (create a GitHub release, container image and Helm chart), create and push a tag
to trigger the process, ie.

```nohighlight
git tag v0.1.0
git push origin v0.1.0
```

Make sure to stick to the exact semver format, and always prepend the `v` for release tags.

## Automation

This sections provides a better understanding of what happens in GitHub workflows and pre-commit actions.

### Running validation and linting tools

This repo uses the [pre-commit](https://pre-commit.com/) framework to run a variety of static code quality analysis tools,
both in CI and local runs. Please [consult the config file](../.pre-commit-config.yaml) to check what is run by default.

The checks configured include:

- General good practices for code files (e.g. consistent line endings).
- Golang validation based on [golangci-lint](https://golangci-lint.run/).
- Syntax linting for Dockerfile, Markdown, JSON, YAML and shell scripts.
- Protection from committing secrets to the repository, using [gitleaks](https://github.com/gitleaks/gitleaks).
- Ensure up-to-date Helm chart documentation in the chart README via [helm-docs](https://github.com/norwoodj/helm-docs).

### Building binaries, images, and charts

For this purpose, the repo uses [goreleaser](https://goreleaser.com/). Please consult [its config file](../.goreleaser.yaml)
to check the included configuration and tune it to your needs.

By default, the `goreleaser` configuration includes:

- Building a Go binary for these Linux architectures: `amd64`, `arm` and `arm64`.
- Building multi-architecture container image with the binary built in the previous step.
- Generating a SBoM manifest.
- Signing all the artifacts with `cosign`.
- Creating a GitHub release with an automatically generated changelog based on git commits.

For building a Helm chart, some code linting and validation tools are executed via
[app-build-suite](https://github.com/giantswarm/app-build-suite/). The tools applied here include
[ct](https://github.com/helm/chart-testing) and [kube-linter](https://github.com/stackrox/kube-linter).

### GitHub workflows

There following GitHub workflows are configured:

- [Release project](../.github/workflows/release.yaml): triggered on a tag push to create a full project release.
- [Validate and test code](../.github/workflows/validate-test.yaml): triggered on each commit, runs basic
   code validation.
- [Pre-commit auto-update](.github/workflows/update-pre-commit.yaml): ran on schedule to detect
  automatic updates to `pre-commit` actions.

Additionally, Renovate is [configured](renovate.json) to automatically handle dependency updates.
