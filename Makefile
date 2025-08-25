MAIN=cmd/main.go

run:
	go run ${MAIN}

run-with-env:
	@set -a && source .env.example && set +a && go run ${MAIN}
