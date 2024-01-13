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
- Copy & modify the file `./config/config.yaml` with your own parameters config

## Command Runner
- `./scripts/run.sh` for running app 
- `./scripts/run-worker.sh` running worker producer with kafka
- `./scripts/run-lint.sh` linters runner 
- `./scripts/run-container.sh` run with docker
- `./scripts/generate-coverage-report.sh` generate test coverage result report
- `./scripts/run-test.sh` running unit testing


## Next Feature Soon
- Kafka
- GRPC 
- Elasticsearch
- Swagger API Documentation
- Kubernetes
- Many More

