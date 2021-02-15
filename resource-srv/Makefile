APP_NAME:=resource-srv
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
	@go get git.yitum.com/mygomod/yitea@master
	@go get git.yitum.com/mygomod/yitea-contrib@master
	@go get git.yitum.com/mygomod/proto@master

local.job.syncvideo:
	@cd $(APP_PATH) && go run main.go --config=config/syncprod.toml --job=syncvideo

local.job.syncpdf:export EGO_DEBUG=true
local.job.syncpdf:
	@cd $(APP_PATH) && go run main.go --config=config/syncprod.toml --job=syncpdf

debug: clear install mock

macbuild:
	@echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>making build app<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
	@chmod +x $(SCRIPT_PATH)/build/*.sh
	@cd $(APP_PATH) && $(SCRIPT_PATH)/build/macgobuild.sh $(APP_NAME) $(COMPILE_OUT)
	@echo "\n"

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

prod.data:
	@cd $(APP_PATH) && $(APP_PATH)/scripts/deploy/deploy_data.sh $(APP_NAME) root@prod.your.machine

prod.deploy: build deployprod


job.videolist:export EGO_DEBUG=true
job.videolist:
	@cd $(APP_PATH) && go run main.go --config=config/local.toml --job=videolist

job.syncvideo:export EGO_DEBUG=true
job.syncvideo:
	@cd $(APP_PATH) && go run main.go --config=config/local.toml --job=syncvideo

job.genvideoplay:export EGO_DEBUG=true
job.genvideoplay:
	@cd $(APP_PATH) && go run main.go --config=config/local.toml --job=genvideoplay
job.mock:export EGO_DEBUG=true
job.mock:
	@cd $(APP_PATH) && go run main.go --config=config/local.toml --job=mock
job.csv:export EGO_DEBUG=true
job.csv:
	@cd $(APP_PATH) && go run main.go --config=config/local.toml --job=csv
job.pdf:export EGO_DEBUG=true
job.pdf:
	@cd $(APP_PATH) && go run main.go --config=config/local.toml --job=pdf
job.syncpdf:export EGO_DEBUG=true
job.syncpdf:
	@cd $(APP_PATH) && go run main.go --config=config/local.toml --job=syncpdf

prod.install:export EGO_DEBUG=true
prod.install:
	@cd $(APP_PATH) && ./bin/$(APP_NAME) --config=config/prod.toml --job=install

prod.job.syncvideo:export EGO_DEBUG=true
prod.job.syncvideo:
	@cd $(APP_PATH) && ./bin/$(APP_NAME) --config=config/prod.toml --job=syncvideo

prod.job.syncpdf:export EGO_DEBUG=true
prod.job.syncpdf:
	@cd $(APP_PATH) && ./bin/$(APP_NAME) --config=config/prod.toml --job=syncpdf

prod.job.synccsv:export EGO_DEBUG=true
prod.job.synccsv:
	@cd $(APP_PATH) && ./bin/$(APP_NAME) --config=config/prod.toml --job=csv


dev.ssh.mysql:
	@ssh root@dev.your.machine -L 3306:127.0.0.1:3306 -N

prod.ssh.mysql:
	@ssh root@prod.your.machine -L 3306:127.0.0.1:3306 -N