package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/mod/modfile"
	"log"
	"os"
	"strings"
)

var projectName string

func Init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 读取环境变量文件，如果不存在则忽略
	_ = godotenv.Load()

	projectName = os.Getenv("PROJECT_NAME")

	if projectName == "" {
		// 获取模块名作为项目名称
		mod, err := os.ReadFile("go.mod")
		if err != nil {
			log.Fatal(err)
		}

		path := modfile.ModulePath(mod)

		if path == "" || strings.Contains(path, "/") {
			log.Fatal(fmt.Sprintf("invalid module path: %v", path))
		}
		projectName = path

	}

}

func ProjectName() string {
	return projectName
}

func Debug() bool {
	return os.Getenv("debug") == "true"
}
