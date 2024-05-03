package log

import (
	"time"

	"github.com/huashusu/rotate"
)

// createWriter 创建一个日志文件写入器，支持文件轮转和大小、年龄限制。
//
// 参数:
// cfg: 指向Log结构体的指针，包含日志配置信息，如日志文件目录、最大年龄和最大大小。
// layout: 日志文件的格式布局。
//
// 返回值:
// 返回一个已初始化的rotate.Rotate指针，用于向日志文件写入日志。
// 如果在创建过程中遇到错误，将返回nil和相应的错误信息。
func createWriter(cfg *Log, layout string) (*rotate.Rotate, error) {
	// 使用提供的配置信息和选项创建一个新的日志文件轮转器
	file, err := rotate.New(cfg.Directory, layout, "log",
		rotate.WithMaxAge(time.Duration(cfg.MaxAge)*rotate.Day), // 设置日志文件的最大年龄
		rotate.WithMaxSize(rotate.MB*int64(cfg.MaxSize)),        // 设置日志文件的最大大小
	)
	if err != nil {
		return nil, err // 如果创建失败，返回错误信息
	}
	return file, err
}
