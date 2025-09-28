LOCAL_BIN:=$(CURDIR)/bin

deps-up:
	docker-compose -f test-docker-compose.yml up --force-recreate --build
.PHONY: deps-up