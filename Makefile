GO ?= go
PROJECT ?= github.com/mheon/testvol
BUILDTAGS ?=
BUILDFLAGS ?=

all: testvol

testvol:
	$(GO) build $(BUILDFLAGS) -tags "$(BUILDTAGS)" -o bin/$@ . $(PROJECT)

clean:
	$(RM) *.sock
	$(RM) -r bin/
