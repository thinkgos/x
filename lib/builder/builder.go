package builder

var (
	// BuildTime 编译日期 由外部ldflags指定
	BuildTime = "unknown"
	// GitCommit git提交版本(短) 由外部ldflags指定
	GitCommit = "unknown"
	// GitFullCommit git提交版本(完整) 由外部ldflags指定
	GitFullCommit = "unknown"
	// Version 版本 由外部ldflags指定
	Version = "unknown"
	// APIVersion api版本 由外部ldflags指定
	APIVersion = "unknown"
	// Model 型号 由外部ldflags指定
	Model = "unknown"
	// Name 应用名称 由外部ldflags指定
	Name = "unknown"
)
