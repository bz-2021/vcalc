package utils

import (
	"path/filepath"
	"strings"
)

// 对于 IO 密集型任务，适合开多个协程进行并发处理

var DirectoryPrefix string = "└── "
var MidChildPrefix string = "├── "
var LastChildPrefix string = "│   "
var SpacePrefix string = "    "

var VisitedPath map[string]bool

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
	Path   strings.Builder // 用于保存路径信息
}

func New(path string) *FileTreeNode {
	builder := &strings.Builder{}
	prefix := &strings.Builder{}
	builder.WriteString(path)
	prefix.WriteString(DirectoryPrefix)
	return &FileTreeNode{Level: 0, Prefix: *prefix, Path: *builder}
}

// Visit 遍历节点下所有的子目录和子文件
func (node *FileTreeNode) Visit(Level int8) (dirs, files int) {
	if path, err := filepath.Abs(node.Path.String()); err == nil {
		path = filepath.Clean(path)
		VisitedPath[path] = true
	}
	return 0, 0
}
