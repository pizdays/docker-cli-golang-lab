unit-test:
	go clean -testcache && go test -v -cover ./...

rundev:
	cp env_dev .env
	swag init
	goreload run main.go

runstage:
	cp env_stage .env
	swag init
	goreload run main.go

# Credentials (NOTE: Avoid storing sensitive details directly. Consider using environment variables)
DOCKER_REGISTRY_USER = 
DOCKER_REGISTRY_PASSWORD = 
PASS_SSH_API_STAGING = 
SERVER_API_STAGING = 

# Constants for Docker
REPO_NAME_STAGE = komgrip/docker-cli-golang-lab
REPO_NAME_PROD = komgrip/docker-cli-golang-lab
STAGE_TAG = latest
PROD_TAG = latest
STAGE_DOCKERFILE = docker/Dockerfile.stage
PROD_DOCKERFILE = docker/Dockerfile.prod

# Constants for Notifications
LINE_NOTIFY_API = https://notify-api.line.me/api/notify
LINE_HEADERS = 'Authorization: Bearer OTWjsy0nWfIWoNW16Vg8lF7FeojSBIuuHwSOhfh8JAZ'
LINE_MESSAGE_STAGE = "message=Deploy docker-cli-golang-lab | Staging | Status: Successful"
LINE_MESSAGE_PROD = "Deploy docker-cli-golang-lab | Production | Status: Successful"
STICKER_DETAILS = 'stickerPackageId=1070' 'stickerId=17843'

# Commands
SSH_COMMAND = sshpass -p "$(PASS_SSH_API_STAGING)" ssh -o "StrictHostKeyChecking=no" $(SERVER_API_STAGING)
DOCKER_LOGIN_COMMAND = docker login --username "$(DOCKER_REGISTRY_USER)" --password "$(DOCKER_REGISTRY_PASSWORD)"
DOCKER_CLEANUP_COMMAND = docker-compose down && docker system prune -af && docker-compose pull && docker-compose up -d

# Function for building and pushing docker images
define build_and_push
	docker build --cache-from $(1):$(STAGE_TAG) -t $(1):$(2) . -f $(3)
	docker push $(1):$(2)
endef

# Function for sending notifications
define send_notification_staging
	curl -X POST $(LINE_NOTIFY_API) \
	-H $(LINE_HEADERS) \
	-F "message=$(LINE_MESSAGE_STAGE)" \
	-F $(STICKER_DETAILS)
endef

# Function for sending notifications
define send_notification_prod
	curl -X POST $(LINE_NOTIFY_API) \
	-H $(LINE_HEADERS) \
	-F "message=$(LINE_MESSAGE_PROD)" \
	-F $(STICKER_DETAILS)
endef

define git_actions
	git checkout $(1)
	git pull
endef

buildstage:
	git pull
	$(DOCKER_LOGIN_COMMAND)
	$(call build_and_push,$(REPO_NAME_STAGE),$(STAGE_TAG),$(STAGE_DOCKERFILE))
	docker rmi $(REPO_NAME_STAGE):$(STAGE_TAG)
	$(SSH_COMMAND) "cd /home/exat/docker && ls -la && $(DOCKER_LOGIN_COMMAND) && $(DOCKER_CLEANUP_COMMAND)"
	$(call send_notification_staging)

buildprod:
	git pull
	$(DOCKER_LOGIN_COMMAND)
	$(call build_and_push,$(REPO_NAME_PROD),$(PROD_TAG),$(PROD_DOCKERFILE))
	docker rmi $(REPO_NAME_PROD):$(PROD_TAG)
	$(call send_notification_prod)
