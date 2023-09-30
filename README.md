<div align="center">
	<h1>Go-CTFd</h1>
	<a href="https://pkg.go.dev/github.com/ctfer-io/go-ctfd"><img src="https://shields.io/badge/-reference-blue?logo=go&style=for-the-badge" alt="reference"></a>
	<a href="https://goreportcard.com/report/github.com/ctfer-io/go-ctfd"><img src="https://goreportcard.com/badge/github.com/ctfer-io/go-ctfd?style=for-the-badge" alt="go report"></a>
	<a href="https://coveralls.io/github/ctfer-io/go-ctfd?branch=main"><img src="https://img.shields.io/coverallsCoverage/github/ctfer-io/go-ctfd?style=for-the-badge" alt="Coverage Status"></a>
	<a href=""><img src="https://img.shields.io/github/license/ctfer-io/go-ctfd?style=for-the-badge" alt="License"></a>
	<br>
	<a href="https://github.com/ctfer-io/go-ctfd/actions/workflows/ci.yaml"><img src="https://img.shields.io/github/actions/workflow/status/ctfer-io/go-ctfd/ci.yaml?style=for-the-badge" alt="CI"></a>
	<a href="https://github.com/ctfer-io/go-ctfd/actions/workflows/codeql-analysis.yaml"><img src="https://img.shields.io/github/actions/workflow/status/ctfer-io/go-ctfd/codeql-analysis.yaml?style=for-the-badge" alt="CodeQL"></a>
</div>

Golang client for interacting with CTFd.

Last version tested on: [3.5.2](https://github.com/CTFd/CTFd/releases/tag/3.5.2).

## Feedbacks

 - The `Authorization` header is not properly documented. It should be stated that it must have the format `Authorization: Token <api_key>`. To find it, you must either check the CTFd API or an existing tool such as ctfcli.
 - The HEAD request to `/api/v1/notifications` does not required authentication, so the number of notifications could be extracted from a closed CTFd. This has low impact on confidentiality and causes no other impact.
 - The Swagger is highly incomplete and contains errors.
