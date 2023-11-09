

GO=go

run:
	${GO} run ./cmd/api

# build:
# 	${GO} build ./cmd

air:
	air

build:
	${GO} build -o ./cmd/api/tmp/api ./cmd/api/main.go 

buildrun:
	./cmd/api/build/mobileMart



