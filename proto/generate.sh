#!/bin/sh

dirs=`find gocn -type d`
for dir in $dirs; do
    # 只处理存在proto文件的目录
    if [ "$(ls ${dir}/*.proto 2>/dev/null)" ]; then
        # 生成go桩代码
        echo "protoc --go_out=plugins=grpc:gen/go $dir/*.proto"
#        protoc --go_out=plugins=grpc:../../../ -I . $dir/*.proto;
        protoc --go_out=plugins=grpc:. -I . $dir/*.proto;

        # 生成python桩代码
        # echo "python -m grpc_tools.protoc -I$dir/ --python_out=$dir/ --grpc_python_out=$dir/ $dir/*.proto"
	      # python -m grpc_tools.protoc -I./ --python_out=./ --grpc_python_out=./  $dir/*.proto;
    fi
done


### 生成go Mock代码
#dirs=`find gen/go/fun -type d | grep -v mock`
#for dir in $dirs; do
#    if [ "$(ls ${dir}/*.pb.go 2>/dev/null)" ]; then
#      files=`ls ${dir}/*.pb.go`
#      mkdir -p $dir/mock
#      for file in $files; do
#        file_name=${file##*/}
#        echo "mock gen for $file ..."
#        mockgen --source=$file > ${dir}/mock/${file_name%%.pb.go}.mock.go
#      done
#    fi
# done

