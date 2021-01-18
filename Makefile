DOCKER_REGISTRY=docker.io
OWNER_ACCOUNT=goforbroke1006
SERVICE_NAME=lock-free-research

IMAGE_TAG=$(shell git describe --tags --abbrev=0 2>/dev/null | git rev-parse --abbrev-ref HEAD )

image:
	@docker build -t ${DOCKER_REGISTRY}/${OWNER_ACCOUNT}/${SERVICE_NAME}:${IMAGE_TAG} ./
	@docker build -t ${DOCKER_REGISTRY}/${OWNER_ACCOUNT}/${SERVICE_NAME}:latest ./
	@docker push ${DOCKER_REGISTRY}/${OWNER_ACCOUNT}/${SERVICE_NAME}:${IMAGE_TAG}
	@docker push ${DOCKER_REGISTRY}/${OWNER_ACCOUNT}/${SERVICE_NAME}:latest


include .stand
export $(shell sed 's/=.*//' .stand)

stand: stand/deploy stand/help

stand/deploy:
	@ssh ${PERFORMANCE_STAND_USER}@${PERFORMANCE_STAND_IP} "mkdir -p ${PERFORMANCE_STAND_DIR}"
	@scp ./docker-compose.yaml ${PERFORMANCE_STAND_USER}@${PERFORMANCE_STAND_IP}:${PERFORMANCE_STAND_DIR}/
	@scp -r ./.compose         ${PERFORMANCE_STAND_USER}@${PERFORMANCE_STAND_IP}:${PERFORMANCE_STAND_DIR}/
	@ssh ${PERFORMANCE_STAND_USER}@${PERFORMANCE_STAND_IP} "cd ${PERFORMANCE_STAND_DIR} && docker-compose pull"
	@ssh ${PERFORMANCE_STAND_USER}@${PERFORMANCE_STAND_IP} "cd ${PERFORMANCE_STAND_DIR} && bash ./.compose/rolling-update.sh api-tcp-counter-mutex-based"
	@ssh ${PERFORMANCE_STAND_USER}@${PERFORMANCE_STAND_IP} "cd ${PERFORMANCE_STAND_DIR} && bash ./.compose/rolling-update.sh api-tcp-counter-lock-free"
	@ssh ${PERFORMANCE_STAND_USER}@${PERFORMANCE_STAND_IP} "cd ${PERFORMANCE_STAND_DIR} && docker-compose up -d"

stand/lb:
	@scp ./docker-compose.yaml ${PERFORMANCE_STAND_USER}@${PERFORMANCE_STAND_IP}:${PERFORMANCE_STAND_DIR}/
	@scp -r ./.compose         ${PERFORMANCE_STAND_USER}@${PERFORMANCE_STAND_IP}:${PERFORMANCE_STAND_DIR}/
	@ssh ${PERFORMANCE_STAND_USER}@${PERFORMANCE_STAND_IP} "cd ${PERFORMANCE_STAND_DIR} && docker-compose restart api-tcp-counter-mutex-based-lb api-tcp-counter-lock-free-lb"

stand/help:
	@echo "Grafana             http://${PERFORMANCE_STAND_IP}:3000/login"
	@echo "Prometheus targets  http://${PERFORMANCE_STAND_IP}:9090/targets"
	@echo "Mutex-based counter http://${PERFORMANCE_STAND_IP}:8010/metrics | http://${PERFORMANCE_STAND_IP}:8011/metrics | http://${PERFORMANCE_STAND_IP}:8012/metrics"
	@echo "Lock-free counter   http://${PERFORMANCE_STAND_IP}:8020/metrics | http://${PERFORMANCE_STAND_IP}:8021/metrics | http://${PERFORMANCE_STAND_IP}:8022/metrics"

