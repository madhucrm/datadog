TEST?=$$(go list ./...)
RECORD?=false
GOIMPORTS_FILES?=$$(find . -name '*.go')
PKG_NAME=datadog
DIR=~/.terraform.d/plugins
ZORKIAN_VERSION?=master
API_CLIENT_VERSION?=master
LOCAL_PACKAGE="github.com/terraform-providers/terraform-provider-datadog"

default: build

build: fmtcheck
	go install

install: fmtcheck
	mkdir -vp $(DIR)
	go build -o $(DIR)/terraform-provider-datadog

uninstall:
	@rm -vf $(DIR)/terraform-provider-datadog

# Run unit tests; these tests don't interact with the API and don't support/need RECORD
test: get-test-deps fmtcheck
	gotestsum --hide-summary skipped --format testname --debug --packages $(TEST) -- $(TESTARGS) -timeout=30s

# Run acceptance tests (this runs integration CRUD tests through the terraform test framework)
testacc: get-test-deps
	RECORD=$(RECORD) TF_ACC=1 gotestsum --format testname --debug --rerun-fails --packages ./... -- -v $(TESTARGS) -timeout 120m

# Run both unit and acceptance tests
testall: test testacc

cassettes: get-test-deps fmtcheck
	RECORD=true TF_ACC=1 gotestsum --format testname --packages ./... -- -v $(TESTARGS) -timeout 120m

vet:
	@echo "go vet ."
	@go vet $$(go list ./...) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	goimports -format-only -local $(LOCAL_PACKAGE) -w $(GOIMPORTS_FILES)
	terraform fmt -recursive examples

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/fmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

test-compile: get-test-deps
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	gotestsum --format testname -- -c $(TEST) $(TESTARGS)

update-go-client:
	echo "Updating the Zorkian client to ${ZORKIAN_VERSION} and the API Client to ${API_CLIENT_VERSION}"
	go get github.com/zorkian/go-datadog-api@$(ZORKIAN_VERSION)
	go get github.com/DataDog/datadog-api-client-go@${API_CLIENT_VERSION}
	go mod tidy

get-test-deps:
	gotestsum --version || go install gotest.tools/gotestsum@latest
	which goimports || go install golang.org/x/tools/cmd/goimports@latest

license-check:
	@sh -c "'$(CURDIR)/scripts/license-check.sh'"

tools:
	go generate -tags tools tools/tools.go

docs: tools
	@sh -c "'$(CURDIR)/scripts/generate-docs.sh'"

check-docs: docs
	@if [ "`git status --porcelain docs/`" ]; then \
	    git diff; \
		echo "Uncommitted changes were detected in the autogenerated docs folder. Please run 'make docs' to autogenerate the docs, and commit the changes" && echo `git status --porcelain docs/` && exit 1; \
	else \
		echo "Success: No generated documentation changes detected"; \
	fi

.PHONY: build check-docs docs test testall testacc cassettes vet fmt fmtcheck errcheck test-compile tools get-test-deps license-check
