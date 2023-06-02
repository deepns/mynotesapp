ENTRY_POINT = httpHandler
RUNTIME = go120
FUNCTION_NAME = mynotesapp

local:
	FUNCTION_TARGET=$(ENTRY_POINT) go run cmd/main/main.go
deploy:
	gcloud functions deploy $(FUNCTION_NAME) \
		--gen2 \
		--runtime=$(RUNTIME) \
		--source=. \
		--entry-point=$(ENTRY_POINT) \
		--trigger-http \
		--set-env-vars SOURCE_DIR=./serverless_function_source_code \
		--allow-unauthenticated