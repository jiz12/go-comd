package template

import (
	_ "embed"
	"os"
	"runtime"
	"strings"
)

//1,选择模板
//2.从模板中读取内容，并写入到新的文件中
//3.返回成功

type TemplateFile struct {
	Content     string
	DefaultName string
}

//go:embed mysql/migrate.txt
var migrateStr string

//go:embed config/config.yml
var configStr string

//go:embed shortcut/shortcut.txt
var shortcutStr string

const (
	Migrate  = "migrate"
	Config   = "config"
	ShortCut = "shortcut"
)

var Templates = map[string]TemplateFile{
	Migrate:  {Content: migrateStr, DefaultName: "template.go"},
	Config:   {Content: configStr, DefaultName: "config.yml"},
	ShortCut: {Content: shortcutStr, DefaultName: "shortcut.txt"},
}

// LoadContentFromTemplate load content from template and write into the file
func LoadContentFromTemplate(f *os.File, dir string, content string) error {

	content = replacePackageName(dir, f.Name(), content)

	_, err := f.WriteString(content)

	return err
}

// replacePackageName replace package name in generated go file.
func replacePackageName(dir string, fileName string, content string) string {

	if getFileSuffix(fileName) != "go" {
		return content
	}
	packageName := dir
	if dir == "." {
		packageName = handleCurrentDirPackageName()
	}

	content = strings.Replace(content, "%s", packageName, 1)

	return content
}

// handleCurrentDirPackageName  use current dir name as package name
func handleCurrentDirPackageName() string {
	var pkgName string
	var res []string
	pwd, _ := os.Getwd()
	sysType := runtime.GOOS
	if sysType == "windows" {
		res = strings.Split(pwd, "\\")
	} else {
		res = strings.Split(pwd, "/")
	}
	if len(res) != 0 {
		pkgName = res[len(res)-1]
	}
	return pkgName
}

func getFileSuffix(fileName string) string {
	res := strings.Split(fileName, ".")
	if len(res) == 0 {
		return ""
	}
	return res[len(res)-1]
}
