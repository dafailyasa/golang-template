# Golang Template (Fiber & Mongo) 🚀

All the project is based in interfaces that mean you can implement your own logic and use it in the project. And this project structure references to [Golang Standard Layout.](https://github.com/golang-standards/project-layout) 

## Stack
- Router: [Fiber 🚀](https://gofiber.io)
- Env: [Viper 🔐](https://github.com/spf13/viper)
- Database: [Mongo 💾](https://www.mongodb.com/docs/drivers/go/current/) 
- Logger: [Zap ⚡](https://github.com/uber-go/zap)
- Deploy: [Docker 🐳](https://www.docker.com)
- CI: [Github Actions 🐙](https://docs.github.com/en/actions)

## Before The Execution
- Copy & modify the file `./config/config.example.yaml` with your own parameters config

## Command Runner
- `./scripts/run.sh` for running app 
- `./scripts/run-worker.sh` running worker producer with kafka
- `./scripts/run-lint.sh` linters runner 
- `./scripts/run-container.sh` run with docker
- `./scripts/generate-coverage-report.sh` generate test coverage result report
- `./scripts/run-test.sh` running unit testing

## Commit type

| Type | Description |
| --- | --- |
| feat | A new feature |
| fix | A bug fix |
| docs | Documentation only changes |
| style | Changes that do not affect the meaning of the code (white-space, formatting etc) |
| refactor | A code change that neither fixes a bug nor adds a feature |
| perf | A code change that improves performance |
| test | Adding missing tests or correcting existing tests |
| build | Changes that affect the build system or external dependencies |
| ci | Changes to our CI configuration files and scripts |
| chore | Other changes that don't modify src or test files |
| revert | Reverts a previous commit |

## Next Feature Soon
- Kafka
- GRPC 
- Elasticsearch
- Swagger API Documentation
- Kubernetes
- Many More

