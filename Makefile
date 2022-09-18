GO_FILES:=$(shell find . -type f -name '*.go' -print)

i3blocksgo: $(GO_FILES)
	go build -o i3blocksgo .
