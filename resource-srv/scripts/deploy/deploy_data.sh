#!/bin/bash
rm -rf data.tar.gz
tar zcvf data.tar.gz data
echo "应用名：$1"
echo "服务器：$2"
ssh ${2} "cd /home/www/server/${1} && mv data.tar.gz databak.tar.gz"
scp data.tar.gz ${2}:/home/www/server/${1}/

ssh ${2} "cd /home/www/server/${1} && \
	    tar xvf data.tar.gz && \
	    chown www:www -R /home/www/server/${1}"
