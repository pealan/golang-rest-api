all: test
test: unit-test

.PHONY: unit-test
unit-test:
	@docker build . --target unit-test