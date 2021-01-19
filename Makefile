help: ## show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sed -e 's/^GNUmakefile://' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

secret: ## make secret
	if [ -z "${LARGE_SECRET_PASSPHRASE}" ]; then echo need $$LARGE_SECRET_PASSPHRASE; exit 1; fi
	echo ${LARGE_SECRET_PASSPHRASE} | gpg --batch --passphrase-fd 0 --symmetric --cipher-algo AES256 debug_client_secret.json
	echo ${LARGE_SECRET_PASSPHRASE} | gpg --batch --passphrase-fd 0 --symmetric --cipher-algo AES256 debug.token

build:
	go build -v -tags=prod
