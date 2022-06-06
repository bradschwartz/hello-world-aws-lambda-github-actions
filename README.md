# Hello World AWS Lambda + GitHub Actions

Exploring GitHub actions by deploying an AWS Lamba function.

General Steps:

1. Lint
1. Build/Compile
1. Unit Tests
1. Code Quality
1. Deploy


Diving in, I'm using Go as the language/tooling for the lambda. We'll note here
every step we said we'd do above for the specific project steps, and try to document
the corresponding GitHub Actions setup.

1. ~~Lint~~
1. Build/Compile
  - `go build -o bootstrap -v ./...`
1. Unit Tests
  - `go test -cover -v ./...`
1. Code Quality
  - Implicitly using SonarCloud, no specific Actions
1. Deploy
  - `appleboy/lambda-action@master`
