all: isuda isutar

deps:

isuda: deps
	go build -o isuda isuda.go type.go util.go user.go

isutar: deps
	go build -o isutar isutar.go type.go util.go

.PHONY: all deps
