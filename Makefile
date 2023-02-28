build-local:
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o handler ./cmd/main.go

sls-local:
	sls deploy --stage local --force

deploy-infra-local:
	docker compose up -d

deploy: deploy-infra-local build-local sls-local