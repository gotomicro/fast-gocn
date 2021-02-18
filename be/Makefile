APP_NAME:=gocn-wechat-be
APP_PATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
SCRIPT_PATH:=$(APP_PATH)/scripts
COMPILE_OUT:=$(APP_PATH)/bin/$(APP_NAME)

go:export EGO_DEBUG=true
go:
	@cd $(APP_PATH) && bee run
install:export EGO_DEBUG=true
install:
	@cd $(APP_PATH) && go run main.go --config=config/local.toml --job=install
update:
	@go get github.com/gotomicro/ego@main

debug: clear install mock

build:
	@echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>making build app<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
	@chmod +x $(SCRIPT_PATH)/build/*.sh
	@cd $(APP_PATH) && $(SCRIPT_PATH)/build/gobuild.sh $(APP_NAME) $(COMPILE_OUT)

deploydev:
	@cd $(APP_PATH) && $(APP_PATH)/scripts/deploy/deploy_go.sh $(APP_NAME) root@dev.your.machine

dev.init_supervisor:
	@cd $(APP_PATH) && $(APP_PATH)/scripts/deploy/init_supervisor.sh $(APP_NAME) data/etcdev root@dev.your.machine

dev.init_nginx:
	@cd $(APP_PATH) && $(APP_PATH)/scripts/deploy/init_nginx.sh $(APP_NAME) data/etcdev root@dev.your.machine

dev.deploy: build deploydev


deployprod:
	@cd $(APP_PATH) && $(APP_PATH)/scripts/deploy/deploy_go.sh $(APP_NAME) root@prod.your.machine

prod.init_supervisor:
	@cd $(APP_PATH) && $(APP_PATH)/scripts/deploy/init_supervisor.sh $(APP_NAME) data/etcprod root@prod.your.machine

prod.init_nginx:
	@cd $(APP_PATH) && $(APP_PATH)/scripts/deploy/init_nginx.sh $(APP_NAME) data/etcprod root@prod.your.machine

prod.deploy: build deployprod





