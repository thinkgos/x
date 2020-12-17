package habit

// 运行工作模式
const (
	ModeDev   = "dev"   // 开发模式
	ModeDebug = "debug" // 测试模式
	ModeProd  = "prod"  // 生产模式
)

// IsModeDev 是否是开发模式
func IsModeDev(mode string) bool {
	return mode == ModeDev
}

// IsModeDebug 是否是测试模式
func IsModeDebug(mode string) bool {
	return mode == ModeDebug
}

// IsModeProd 是否是生产模式
func IsModeProd(mode string) bool {
	return mode == ModeProd
}
