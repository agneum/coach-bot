VERSION = $(shell v=$$(git tag -l --points-at HEAD) && [ -n "$$v" ] && echo $$v || TZ=UTC git show -s --abbrev=12 --date=format-local:%Y%m%d%H%M%S --pretty=format:v0.0.0-%cd-%h HEAD)
LDFLAGS = -ldflags '-X github.com/agneum/scheduler-bot/pkg/version.Version=$(VERSION)'
GOBUILD = go build $(LDFLAGS) -o


version:
	@echo $(VERSION)

reform:
	reform pkg/models

build:
	$(GOBUILD) bin/scheduler-bot cmd/bot/main.go

run:
	$(GOBUILD) bin/scheduler-bot cmd/bot/main.go
	bin/scheduler-bot

imports:
	goimports -w .
