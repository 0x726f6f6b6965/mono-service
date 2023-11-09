PROJECTNAME := $(shell basename "$(PWD)")
include .env
export $(shell sed 's/=.*//' .env)

# Bazel

# gazelle-run - 
.PHONY: gazelle-run
gazelle-run:
	@bazel run //:gazelle

# baz-clean - 
.PHONY: baz-clean
baz-clean:
	@bazel clean

.PHONY: baz-build
baz-build:
	@bazel build //services/graph-service/app
	@bazel build //services/graph-service/app:app_layer
	@bazel build //services/graph-service/app:image
	@bazel build //services/graph-service/app:tarball
	@bazel build //services/user-service/app
	@bazel build //services/user-service/app:app_layer
	@bazel build //services/user-service/app:image
	@bazel build //services/user-service/app:tarball

# Protobuf

# proto-lint - 
.PHONY: proto-lint
proto-lint:
	@buf lint

# proto-gen - 
.PHONY: proto-gen
proto-gen:
	@buf generate

# proto-check-breaking - 
.PHONY: proto-check-breaking
proto-check-breaking:
	@buf breaking --against '.git#branch=main' --error-format=json | jq .

# proto-all - 
.PHONY: proto-all
proto-all: proto-gen proto-lint proto-check-breaking

# proto-clean
.PHONY: proto-clean
proto-clean: 
	@find protos -type f -name "*.go" -delete

## DB

.PHONY: set-psql
set-psql:
	@docker run --name MyPostgres -d -p 5432:5432 \
		-e POSTGRES_DB=$(POSTGRES_DB) \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-v ./psql-data:/var/lib/postgresql/data \
		--rm postgres:latest

# GraphQL

.PHONY: graph-gen
graph-gen:
	@cd services/graph-service && go run github.com/99designs/gqlgen generate
	@go generate services/graph-service/graph/resolver.go

.PHONY: graph-init
graph-init:
	@cd services/graph-service && go run github.com/99designs/gqlgen init 

.PHONY: k8s-run
k8s-run:
	@envsubst < k8s/mono-service.yaml | kubectl apply -f -
	@nohup kubectl port-forward deployment/mono-service 8080:8080 > /dev/null1 2>&1 &
	@nohup kubectl port-forward deployment/mono-service 5432:5432 > /dev/null2 2>&1 &

k8s-clean:
	@kubectl delete -f k8s/mono-service.yaml
	# @kill $(shell ps | grep "kubectl port-forward deployment/mono-service 8080:8080" | awk '{print $1}')
	# @kill $(shell ps | grep "kubectl port-forward deployment/mono-service 5432:5432" | awk '{print $1}')




