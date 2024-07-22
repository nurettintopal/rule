project_name = rule

help: ## This help dialog.
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -F -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

run-local: ## Run the app locally
	go run main.go

requirements: ## Generate go.mod & go.sum files
	go mod tidy

clean-packages: ## Clean packages
	go clean -modcache

test: ## Run the tests
	go test

test-v: ## Run the tests
	go test -v

test-cover: ## Run the tests
	go test -cover

test-cover-detail: ## Run the tests with coverage details
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out