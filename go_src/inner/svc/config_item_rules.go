package svc

// configItemRule 定义已知配置项的展示和默认选择规则。
type configItemRule struct {
	Description     string
	DefaultSelected bool
	IsPlugin        bool
}

var configItemRules = map[string]configItemRule{
	// 默认同步
	"app.json": {
		Description:     "编辑器的通用配置",
		DefaultSelected: true,
	},
	"core-plugins.json": {
		Description:     "Obsidian 核心插件开关",
		DefaultSelected: true,
	},
	"hotkeys.json": {
		Description:     "自定义快捷键",
		DefaultSelected: true,
	},
	"appearance.json": {
		Description:     "外观配置和启用的 CSS 片段，应与 snippets/ 一起同步",
		DefaultSelected: true,
	},
	"snippets/": {
		Description:     "自定义编辑器样式",
		DefaultSelected: true,
	},

	// 默认跳过
	"community-plugins.json": {
		Description:     "已启用的社区插件列表，需要与对应插件目录一起同步",
		DefaultSelected: false,
		IsPlugin:        true,
	},
	"graph.json": {
		Description:     "关系图谱的显示和布局设置，通常不需要跨库同步",
		DefaultSelected: false,
	},
	"bookmarks.json": {
		Description:     "当前库的书签列表，通常包含仅在本库存在的文件",
		DefaultSelected: false,
	},
	"types.json": {
		Description:     "当前库的属性类型定义，可能依赖本库的笔记结构",
		DefaultSelected: false,
	},
	"workspace.json": {
		Description:     "桌面端当前打开文件和界面布局，变化频繁且仅适用于当前库",
		DefaultSelected: false,
	},
	"workspace-mobile.json": {
		Description:     "移动端当前打开文件和界面布局，变化频繁且仅适用于当前库",
		DefaultSelected: false,
	},
	"workspaces.json": {
		Description:     "保存的工作区布局，可能引用仅在当前库存在的文件",
		DefaultSelected: false,
	},
}

// newConfigItem 创建配置项并应用已知规则。
func newConfigItem(path string, name string, isDir bool, isPlugin bool) ConfigItem {
	rule := configItemRules[path]
	if rule.IsPlugin {
		isPlugin = true
	}
	description := rule.Description
	if isPlugin && description == "" {
		description = "社区插件程序和设置"
	}
	return ConfigItem{
		Path:            path,
		Name:            name,
		IsDir:           isDir,
		Description:     description,
		DefaultSelected: rule.DefaultSelected,
		IsPlugin:        isPlugin,
	}
}
