DOCKER_REGISTRY=docker.io
OWNER_ACCOUNT=goforbroke1006
SERVICE_NAME=lock-free-research

IMAGE_TAG=$(shell git describe --tags --abbrev=0 2>/dev/null | git rev-parse --abbrev-ref HEAD )

image:
	docker build -t ${DOCKER_REGISTRY}/${OWNER_ACCOUNT}/${SERVICE_NAME}:${IMAGE_TAG} ./
	docker build -t ${DOCKER_REGISTRY}/${OWNER_ACCOUNT}/${SERVICE_NAME}:latest ./
	docker push ${DOCKER_REGISTRY}/${OWNER_ACCOUNT}/${SERVICE_NAME}:${IMAGE_TAG}
	docker push ${DOCKER_REGISTRY}/${OWNER_ACCOUNT}/${SERVICE_NAME}:latest
