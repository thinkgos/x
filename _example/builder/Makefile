# 应用名称
name=builder
# 型号
model = ${name}
# 固件版本
# git describe --tags `git rev-list --tags --max-count=1`
version = `git describe --tags`
# api版本
APIversion = v0.0.1
# 设置固件名称
firmwareName = ${name}

execveFile := ${firmwareName}

# 编译flags
path=github.com/thinkgos/go-core-package/lib/builder
# TIPS: 如果值内容中含有空格，可以用单引号！
flags = -ldflags "-X '${path}.BuildTime=`date "+%F %T %z"`' \
	-X '${path}.GitCommit=`git rev-parse --short=8 HEAD`' \
	-X '${path}.GitFullCommit=`git rev-parse HEAD`' \
	-X '${path}.Name=${name}' \
	-X '${path}.Model=${model}' \
	-X '${path}.Version=${version}' \
	-X '${path}.APIVersion=${APIversion}' -w" # -s 引起gops无法识别go版本号,upx压缩也同样

system:
	@go build ${flags} -o ${execveFile} .

clean:
	@echo "----> cleanning"
	@rm ${execveFile}
	@echo "----> clean successful"

.PHONY: system clean