all: isuda

deps:

isuda: deps
	go build -o isuda isuda.go type.go util.go user.go keyword.go entryCache.go keyCache.go starCache.go

.PHONY: all deps
