// Copyright [2020] [thinkgos] thinkgo@aliyun.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package textcolor

import (
	"fmt"
	"runtime"
)

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

// 颜色定义
const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

// Black 黑色
func Black(msg string) string { return String(msg, 0, 0, TextBlack) }

// Red 红色
func Red(msg string) string { return String(msg, 0, 0, TextRed) }

// Green 绿色
func Green(msg string) string { return String(msg, 0, 0, TextGreen) }

// Yellow 黄色
func Yellow(msg string) string { return String(msg, 0, 0, TextYellow) }

// Blue 蓝色
func Blue(msg string) string { return String(msg, 0, 0, TextBlue) }

// Magenta 紫红色
func Magenta(msg string) string { return String(msg, 0, 0, TextMagenta) }

// Cyan 青蓝色
func Cyan(msg string) string { return String(msg, 0, 0, TextCyan) }

// White 白色
func White(msg string) string { return String(msg, 0, 0, TextWhite) }

// String 自定义文件
func String(msg string, conf, bg, text int) string {
	if runtime.GOOS == "windows" {
		return msg
	}
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
}
