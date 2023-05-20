# Go-CTFd

Golang client for interacting with CTFd.

## Feedbacks

 - The `Authorization` header is not properly documented. It should be stated that it must have the format `Authorization: Token <api_key>`. To find it, you must either check the CTFd API or an existing tool such as ctfcli.
 - The HEAD request to `/api/v1/notifications` does not required authentication, so the number of notifications could be extracted from a closed CTFd. This has low impact on confidentiality and causes no other impact.
 - The Swagger is highly incomplete and contains errors.
