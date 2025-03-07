# How to contribute

## Communication

- Issues: [GitHub](https://github.com/authzed/spicedb-operator/issues)
- Email: [Google Groups](https://groups.google.com/g/authzed-oss)
- Discord: [Zanzibar Discord](https://authzed/discord)

All communication must follow our [Code of Conduct].

[Code of Conduct]: CODE-OF-CONDUCT.md

## Creating issues

If any part of the project has a bug or documentation mistakes, please let us know by opening an issue.
All bugs and mistakes are considered very seriously, regardless of complexity.

Before creating an issue, please check that an issue reporting the same problem does not already exist.
To make the issue accurate and easy to understand, please try to create issues that are:

- Unique -- do not duplicate existing bug report.
  Deuplicate bug reports will be closed.
- Specific -- include as much details as possible: which version, what environment, what configuration, etc.
- Reproducible -- include the steps to reproduce the problem.
  Some issues might be hard to reproduce, so please do your best to include the steps that might lead to the problem.
- Isolated -- try to isolate and reproduce the bug with minimum dependencies.
  It would significantly slow down the speed to fix a bug if too many dependencies are involved in a bug report.
  Debugging external systems that rely on this project is out of scope, but guidance or help using the project itself is fine.
- Scoped -- one bug per report.
  Do not follow up with another bug inside one report.

It may be worthwhile to read [Elika Etemad’s article on filing good bug reports][filing-good-bugs] before creating a bug report.

Maintainers might ask for further information to resolve an issue.

[filing-good-bugs]: http://fantasai.inkedblade.net/style/talks/filing-good-bugs/

## Contribution flow

This is a rough outline of what a contributor's workflow looks like:

- Create an issue
- Fork the project
- Create a branch from where to base the contribution -- this is almost always `main`
- Push changes into a branch of your fork
- Submit a pull request
- Respond to feedback from project maintainers

Creating new issues is one of the best ways to contribute.
You have no obligation to offer a solution or code to fix an issue that you open.
If you do decide to try and contribute something, please submit an issue first so that a discussion can occur to avoid any wasted efforts.

## Legal requirements

In order to protect both you and ourselves, all commits will require an explicit sign-off that acknowledges the [DCO].

Sign-off commits end with the following line:

```git
Signed-off-by: Random J Developer <random@developer.example.org>
```

This can be done by using the `--signoff` (or `-s` for short) git flag to append this automatically to your commit message.
If you have already authored a commit that is missing the signed-off, you can amend or rebase your commits and force push them to GitHub.

[DCO]: /DCO

## Common tasks

### Testing & building a binary

In order to build and test the project, the [latest stable version of Go] and knowledge of a [working Go environment] are required.

[latest stable version of Go]: https://golang.org/dl
[working Go environment]: https://golang.org/doc/code.html

Install [mage](https://magefile.org/#installation):

```sh
# homebrew, see link for other options
brew install mage
```

Run e2e tests:

```sh
mage test:e2e
```

Run unit tests:

```sh
mage test:unit
```

### Adding dependencies

This project does not use anything other than the standard [Go modules] toolchain for managing dependencies.

[Go modules]: https://golang.org/ref/mod

```sh
go get github.com/org/newdependency@version
```

Continuous integration enforces that `go mod tidy` has been run.

### Regenerating `proposed-update-graph.yaml`

The update graph can be regenerated whenever there is a new spicedb release.
CI will validate all new edges when there are changes to `proposed-update-graph.yaml` and will copy them into `validated-update-graph.yaml` if successful.

```go
mage gen:graph
```
