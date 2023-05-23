.DEFAULT_GOAL := publish-rsocket-ping

IMAGE-PREFIX ?= ssingh3339

export DOCKER_CLI_EXPERIMENTAL=enabled

.PHONY: build # Build the container image
build-rsocket-ping:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--output "type=docker,push=false" \
		--tag $(IMAGE-PREFIX)/rsocket-ping:latest \
		rsocket/.

.PHONY: publish # Push the image to the remote registry
publish-rsocket-ping:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=image,push=true" \
		--tag $(IMAGE-PREFIX)/rsocket-ping:latest \
		rsocket/.