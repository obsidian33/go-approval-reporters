# go-approval-reporters
Custom Reporters for [ApprovalTests.go](https://github.com/approvals/go-approval-tests)

## Delta
This reporter simply displays the diff using [delta](https://github.com/dandavison/delta), a CLI. It displays the diff side-by-side and then uses the clipboard reporter so you can simply paste in a move command to accept the received result.

By default, `delta` will use a width of 200 columns, if you'd like to adjust it you can use the `NewDeltaDiffReporterWithWidth()` instead of the default `NewDeltaDiffReporter()`.
