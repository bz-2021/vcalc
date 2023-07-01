package utils

import "strings"

// FileTreeNode 存储文件树的节点信息
type FileTreeNode struct {
	Level    int    // 节点深度
	Size     int64  // 文件大小
	Duration int64  // 总时长
	Filename string // 文件名
	IsDir    bool   // 是否为文件夹
	IsVideo  bool   // 是否为视频类型

	Children []*FileTreeNode // 子节点
	Parent   *FileTreeNode   // 父节点
	Left     *FileTreeNode   // 左兄弟
	Right    *FileTreeNode   // 右兄弟

	Prefix strings.Builder // 用于绘制树状结构的前缀
	Attr   strings.Builder //
	Path   strings.Builder
}
