package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/mod/modfile"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var projectName string

var host string

func Init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

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

	// 读取环境变量文件，如果不存在则忽略
	_ = godotenv.Load("version", ".env")

	address := os.Getenv("IP_QUERY_ADDRESS")

	urls := strings.Split(address, ",")
	for _, url := range urls {
		var err error
		host, err = getHost(fmt.Sprintf("http://%v", url))
		if err != nil {
			log.Println(err, url)
			continue
		}
		return
	}
}

func ProjectName() string {
	return projectName
}

func Debug() bool {
	return os.Getenv("DEBUG") == "true"
}

func MQTTDebug() bool {
	return os.Getenv("MQTT_DEBUG") == "true"
}

func Host() string {
	return host
}

func getHost(url string) (host string, err error) {
	// 获取项目运行环境的IP
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	host = string(buf)
	return
}
