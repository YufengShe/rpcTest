# Go Module文档阅读笔记

## modules、packages  和versions

1. module：一个packages的集合，一个module中包含多个package

2. versions：module的版本，下载module时会直接或者间接的选择该module的某一个版本

3. packages：实现某一个功能的放在同一个文件夹下的源文件；

4. e.g.一个第三方库（或者自己写的一个工具包）（例如：golang.org/x/net）可以看做一个module。他下面有一个包含html功能的package html，则调用这个package的import路径为golang.org/x/net/html。

5. e.g. module path：module的标准名称。**在module内部引用本module的某一个pacakge的时候，需要将module path作为引用前缀**。这是go mod包管理下引用自己写的package的方式！！！ e.g. (创建module path为)go mod init github.com/yufeng/blockchain； （在该module中创建两个package为）service、tools；在service package中调用tools package时，import路径为github.com/yufeng/blockchain/tools

   module path需要描述这个module的功能以及在哪里能找到这个包，因此一个module path通常包含仓库的根路径和版本（一般从2开始），例如：

   github.com/yufeng/blockchain同时为该module在github的保存路径；

   当该包升级到2.0.0版本之后（大版本的改变），为防止版本更替导致的不兼容问题，修改2.0.0版本之后的module path为github/yufeng/blockchain/v2，注意，在实践中发现有些开源第三方库不遵守这个规范，这时go.mod文件中对应包的后面会被加上+incompatible的后缀

6. e.g. version标识一个module的不可变版本，版本号由v开头，例如v1.0.0；一个版本号中包含三个非负数，由.分隔，从左至右代表了大版本、小版本和补丁版本（划分的粒度从左至右递减）；当对module中定义的某些API进行**不兼容的更新**之后，大版本必须增加，同时将小版本号和补丁版本号置零；当对module进行了兼容旧版本的更新后，例如增加了新的功能，此时大版本不变，小版本增加，同时补丁版本置零；当对module进行兼容的一些小修改，例如fix a bug时，大版本和小版本号不变，补丁版本增加即可；补丁版本之后可以跟一个可选的连字符开头的预发布字符串（例如v1.0.0-beta，预发布串表示这是在v1.0.0版本之前的一个版本，预发版本可能与正式版本不兼容）；版本号还有可能有这种后缀example.com/m v4.1.2+incompatible，这表示该module的升级和维护没有按照moudule path的标准来（在module path中添加example.com/m/v4）,因此+incompatible表示可能因为该module的不规范导致接口不兼容问题，在对这个包进行版本升级时需要谨慎（注意是否跨越了大版本导致兼容性问题）

7. go命令可以使用tag、分支、哈希来拉取想要的版本：

    - go get -d golang.org/x/net@daa7c041 (hash)
    - go get -d golang.org/x/net@v1.0.0 (tag)
    - go get -d golang.org/x/net@master (branch)
    - go get 命令会将其转换成伪版本形式：golang.org/x/net/v1.0.0-yyyymmddhhmmss-daa7c041

## GOPROXY

1. goproxy是为main module获取所依赖module设置的代理url列表，url列表由逗号分隔开； e.g. https://corp.example.com,https://proxy.golang.org,direct 意思是在获取依赖时，先去第一个url对应的代理找，找不到的话再去第二个找，direct表示在设置的代理中找不到的话回源查找（如github、golang.org）等版本仓库查找，还可以设置为off，表示不回源查找，找不到直接报错；
2. 在proxy中根据import的一个package查找包含这个package的module：前缀匹配法；e.g. 查找提供golang.org/x/net/html这个包的module：在proxy中并行请求golang.org/x/net/html、golang.org/x/net、golang.org/x、golang.org包的最新版本，如果对应的包不存在，则返回404或者410错误；通过前缀匹配的方式找到适合的module后，将module path和version添加到go.mod文件中；

## go.mod

1. go.mod 为utf-8编码的文本文件，主要定义了一个module的依赖关系（其依赖的module path和version）；
2. go.mod文件中包含的关键字：module、go、require、replace、exclude；
    - module：用于定义当前项目的模块路径（module path）；
    - go：用于设置预期的 Go 版本；
    - require：用于设置一个特定的模块版本；
    - exclude：用于从使用中排除一个特定的模块版本；
    - replace：用于将一个模块版本替换为另外一个模块版本；
3. go.mod中的//indirect注释的含义：
    - //indirect表示这个依赖库没有被main module中的package直接引用（被main module直接引用的包引用）
    - 不是所有的间接依赖都会被填入go.mod文件
    - main module依赖A，仅当A没有启用go module包管理机制或者A的go.mod文件不完整时，才会将缺失的依赖写入main module的go.mod文件中并添加//indirect注释；
4. 













