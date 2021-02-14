#!/bin/sh

dirs=`find fun -type d`
for dir in $dirs; do
    # 只处理存在proto文件的目录
    if [ "$(ls ${dir}/*.proto 2>/dev/null)" ]; then
        echo "python -m grpc_tools.protoc -I$dir/ --python_out=$dir/ --grpc_python_out=$dir/ $dir/*.proto"
	    python -m grpc_tools.protoc -I./ --python_out=./ --grpc_python_out=./  $dir/*.proto;
    fi
done