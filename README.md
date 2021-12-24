# ProtoBuf 安装与使用（Golang）

## 安装Protoc

1. 安装protobuf的编译器：从https://github.com/google/protobuf/releases找编译好的二进制包发行版，下载安装，在cmd输入protoc，能找到该可执行文件即为安装成功

2. 安装golang的protobuf编译插件：

   ```
   go get -u github.com/golang/protobuf/proto // golang protobuf 库
   go get -u github.com/golang/protobuf/protoc-gen-go //protoc --go_out 编译插件工具
   ```

3. 安装golang grpc库：

   ```
   go get google.golang.org/grpc	//代码中调用grpc
   ```

## 使用

1. 参考命令：

   ```
   protoc -I ./idl/ --go_out=./ --go-grpc_out=./ ./idl/*.proto
   ```

2. -I 参数：编译protobuf代码时，.pb文件中import引入的文件的位置；

3. --go_out参数：使用protoc的go的编译插件（记得把$GOPATH/bin放在环境变量中,protoc-gen-go工具在这个路径下）为.proto中定义的message生成golang对应的struct代码, =后面为生成golang代码的存储路径；

4. --go-grpc_out参数：使用protoc的go的编译插件（记得把$GOPATH/bin放在环境变量中,protoc-gen-go工具在这个路径下）为.proto文件中定义的service生成client和service相关代码，=后面为生成的golang代码的存储路径；

5. 最后的./idl/*.proto为本次编译的pb文件，表示编译./idl路径下的全部pb文件；

6. 记得将goland中设置/language/protobuf/中的location添加一下，否则.proto文件的import会有红色下划线错误（找不到import路径）；

7. 最后通过该命令成功生成了对应的golang代码（开启gomod后import可能会有报错，需要简单修改一下生成代码的import路径）

## .proto文件

```protobuf
syntax = "proto3";		//规定protobuf语法的版本 proto2、proro3
package rpc.baas.bit;	//定义profobuf文件的包，将多个.proto文件归集到一个package下；在某个包的文件中使用rpc.laptop包下的Laptop message时，需要使用rpc.laptop.Laptop（package_name.message_name）的方式
option go_package = "pb/chaincodeCallService";	//生成go代码放在哪个路径下

message ContractCallRequest {
  enum CallType {
    UNKNOWN_TYPE = 0;
    INVOKE = 1;
    QUERY = 2;
  };
  int64 contractId = 1; //stub to call chaincode, use chaincode_name or chaincode_id
  string contractName = 2;  //stub to call chaincode, use chaincode_name or chaincode_id
  int64 blockchainId = 3; //blockchain id
  string contractVersion = 4; //contract version
  string args = 5; //parameters for querying or invoking contract
  CallType callType = 6; // invoke or query
}

message ContractCallResponse {
  int64 statusCode = 2; //response status code
  string msg = 3; //response message
  string package = 4; //response
}

service ContractCallService {
  rpc CallContract(ContractCallRequest) returns(ContractCallResponse) {};
}
```

## git的使用
1. git remote -v //显示所有的分支信息
2. git remote add [nickname] url //添加远程仓库的url，并将该仓库起名为**nickname** e.g. git remote add origin  git@github.com:tianqixin/runoob-git-test.git
3. git push [远程仓库名] [本地分支名]:[远程分支名] //将本地仓库的[本地分支名]对应的分支的更新提交至[远程仓库]对应的[远程分支]上 e.g. git push origin master 将本地仓库的master分支的更新提交到远程仓库origin的master分支上