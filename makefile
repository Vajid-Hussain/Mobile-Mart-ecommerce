

GO=go

run:
	${GO} run ./cmd/api

# build:
# 	${GO} build ./cmd

air:
	air

buildDeployment: 
	${GO} build -o ./cmd/api/tmp/deploy ./cmd/api/main.go

build:
	${GO} build -o ./cmd/api/tmp/api.exe ./cmd/api/main.go 

buildrun:
	./cmd/api/build/mobileMart

swaggo:
	swag init -g ./cmd/api/main.go

swaggoformat:
	swag fmt