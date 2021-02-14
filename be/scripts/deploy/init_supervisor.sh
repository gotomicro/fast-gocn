#!/bin/bash
rm -rf bin.tar.gz
echo "应用名：$1"
echo "配置路径：$2"
echo "服务器：$3"

tar zcvf etc.tar.gz -C ${2} .

ssh ${3} "mkdir -p /home/www/server/${1}"
ssh ${3} "mkdir /tmp/etc"
ssh ${3} "cd /tmp/etc && rm -rf etc.tar.gz"
scp etc.tar.gz ${3}:/tmp/etc
ssh ${3} "cd /tmp/etc && \
	    tar xvf etc.tar.gz && \
	    cd /tmp/etc/supervisor/ && \
	    cp ${1}.conf /etc/supervisor/conf.d/ && \
	    supervisorctl update"
