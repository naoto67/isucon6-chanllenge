all: isuda

deps:
	go get "github.com/patrickmn/go-cache"

isuda: deps
	go build -o isuda isuda.go type.go util.go user.go keyword.go entryCache.go

.PHONY: all deps
