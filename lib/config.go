package lib

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// 创建一个导出变量，以便其它包中可以使用，而不用每次都读取文件
var ConfigSetting *Config

const sub_key_sep string = ":"

type Config struct {
	Configs map[string]string
	sub_key string
}

func (c *Config) InitConfig(path string) {
	c.Configs = make(map[string]string)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("文件不存在")
	}

	// 读取配置文件.
	// 打开文件
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// 准备一个ioReader
	r := bufio.NewReader(f)

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		// 去除左右空格
		line = bytes.TrimSpace(line)

		// 判断是否是注释
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		// 如果是[开头的表示是一个配置数组
		sub_start_index := bytes.IndexByte(line, '[')
		sub_end_index := bytes.IndexByte(line, ']')
		if sub_start_index > -1 && sub_end_index > -1 && sub_end_index > sub_start_index+1 {
			c.sub_key = string(line[sub_start_index+1 : sub_end_index])
			continue
		}

		// 使用=分割字符串
		index := bytes.IndexByte(line, '=')
		if index < 0 {
			continue
		}

		// 获取配置项键值
		value := string(line[index+1:])

		// 获取配置项键名
		key := string(line[:index])

		// 可能会出现bug，比如配置项[default]下的name和default_name会出现冲突
		if c.sub_key != "" {
			key = c.sub_key + sub_key_sep + key
		}

		c.Configs[key] = value
	}
	ConfigSetting = c
}

// 读取配置
func (c *Config) Read(sub_key string, key string) string {
	if c == nil {
		panic("配置文件未初始化")
	}
	real_key := sub_key + sub_key_sep + key
	if v, existes := c.Configs[real_key]; existes {
		return v
	}
	return ""
}
