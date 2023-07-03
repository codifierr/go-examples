.DEFAULT_GOAL := publishall

IMAGE-PREFIX ?= ssingh3339

export DOCKER_CLI_EXPERIMENTAL=enabled

.PHONY: build-rsocket-ping
build-rsocket-ping:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=docker,push=false" \
		--tag $(IMAGE-PREFIX)/rsocket-ping:latest \
		rsocket/.

.PHONY: publish-rsocket-ping
publish-rsocket-ping:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=image,push=true" \
		--tag $(IMAGE-PREFIX)/rsocket-ping:latest \
		rsocket/.

.PHONY: build-udp-server
build-udp-server:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=docker,push=false" \
		--tag $(IMAGE-PREFIX)/udp-server:latest \
		udp-server/.

.PHONY: publish-udp-server
publish-udp-server:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=image,push=true" \
		--tag $(IMAGE-PREFIX)/udp-server:latest \
		udp-server/.

.PHONY: build-echo-graphql
build-echo-graphql:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=docker,push=false" \
		--tag $(IMAGE-PREFIX)/echo-graphql:latest \
		echo-graphql/.

.PHONY: publish-echo-graphql
publish-echo-graphql:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=image,push=true" \
		--tag $(IMAGE-PREFIX)/echo-graphql:latest \
		echo-graphql/.

.PHONY: buildall
buildall: build-rsocket-ping build-udp-server build-echo-graphql

.PHONY: publishall
publishall: publish-rsocket-ping publish-udp-server publish-echo-graphql
