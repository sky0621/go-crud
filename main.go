package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

// TODO 機能実現スピード最優先での実装なので要リファクタ
func main() {
	cfg := flag.String("config", "config.toml", "Config File")
	flag.Parse()

	ReadConfig(*cfg)

	config := NewConfig()

	//fp, err := os.Open(config.Tables)
	//if err != nil {
	//	panic(err)
	//}
	//defer func() {
	//	if fp != nil {
	//		fp.Close()
	//	}
	//}()
	dsn := fmt.Sprintf(config.DB, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASS"), os.Getenv("MYSQL_SCHEMA"))
	//fmt.Println(dsn)
	db, err := gorm.Open("mysql", dsn)
	//db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/sakila?charset=utf8&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASS")))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var res []InformationSchema
	db.Raw(fmt.Sprintf("SELECT table_name FROM `information_schema`.`TABLES` WHERE `TABLE_SCHEMA` = '%s' ORDER BY table_name", os.Getenv("MYSQL_SCHEMA"))).Scan(&res)
	//db.Raw("SELECT table_name FROM `information_schema`.`TABLES` WHERE `TABLE_SCHEMA` = 'sakila' ORDER BY table_name").Scan(&res)
	//fmt.Println(res)

	result.Headers = append(result.Headers, "FilePath")

	//scanner := bufio.NewScanner(fp)
	//for scanner.Scan() {
	//	result.Headers = append(result.Headers, scanner.Text())
	//}
	for _, t := range res {
		result.Headers = append(result.Headers, t.TableName)
	}

	err = filepath.Walk(config.Target, Apply)
	if err != nil {
		panic(err)
	}

	breakProj := ""
	breakOxAry := make([]string, len(result.Headers)-1)

	editBodies := [][]string{}
	proj := ""
	//fmt.Println("##################################################")
	//fmt.Println(result.Bodies)
	//fmt.Println("##################################################")
	for idx, body := range result.Bodies {
		paths := strings.Split(body[0], string(os.PathSeparator))
		projIdx := -1
		for idx, p := range paths {
			if p == config.Topdir {
				projIdx = idx + 1
			}
		}
		if projIdx == -1 {
			panic(errors.New("No Target Directory"))
		}
		proj = paths[projIdx]
		oxAry := body[1:]

		if proj != breakProj {
			if idx != 0 {
				oneAry := []string{breakProj}
				for _, bo := range breakOxAry {
					oneAry = append(oneAry, bo)
				}
				editBodies = append(editBodies, oneAry)
			}
			breakProj = proj
			breakOxAry = oxAry
		} else {
			for idx, ox := range oxAry {
				if ox == "-" {
					continue
				}
				breakOxAry[idx] = ox
			}
		}
	}
	oneAry := []string{breakProj}
	for _, bo := range breakOxAry {
		oneAry = append(oneAry, bo)
	}
	editBodies = append(editBodies, oneAry)

	result.Bodies = editBodies

	tmpl := template.Must(template.ParseFiles(config.Template))
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, result)
	if err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}

func Apply(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	fp, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if fp != nil {
			fp.Close()
		}
	}()

	fmngr := &FilterManager{Filter: NewFilterConfig()}
	if !fmngr.IsTarget(path, info) {
		return nil
	}

	strs := ""
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		strs = fmt.Sprintf("%s%s", strs, scanner.Text())
	}

	body := []string{}
	body = append(body, path)

	for idx, tblName := range result.Headers {
		if idx == 0 {
			continue
		}
		if strings.Contains(strs, tblName) {
			body = append(body, "o")
		} else {
			body = append(body, "-")
		}
	}

	result.Bodies = append(result.Bodies, body)
	return nil
}

type Result struct {
	Datetime string
	Headers  []string
	Bodies   [][]string
}

var result = &Result{
	Datetime: time.Now().Format("2006-01-02 15:04"),
	Headers:  []string{},
	Bodies:   [][]string{},
}

// Config ...
type Config struct {
	Target   string
	Topdir   string
	Tables   string
	Template string
	DB       string
	Filter   *FilterConfig
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Target:   viper.GetString("target"),
		Topdir:   viper.GetString("topdir"),
		Tables:   viper.GetString("tables"),
		Template: viper.GetString("template"),
		DB:       viper.GetString("db"),
		Filter:   NewFilterConfig(),
	}
}

// FilterConfig ...
type FilterConfig struct {
	Out []string
	In  []string
}

// NewFilterConfig ...
func NewFilterConfig() *FilterConfig {
	return &FilterConfig{
		Out: viper.GetStringSlice("filter.out"),
		In:  viper.GetStringSlice("filter.in"),
	}
}

// ReadConfig ...
func ReadConfig(configFilePath string) error {
	viper.SetConfigFile(configFilePath)
	return viper.ReadInConfig()
}

// FilterManager ...
type FilterManager struct {
	Filter *FilterConfig
}

// IsTarget ...
func (m *FilterManager) IsTarget(path string, info os.FileInfo) bool {
	if info.IsDir() {
		return false
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	outs := m.Filter.Out
	for _, out := range outs {
		outExp, err := regexp.Compile(out)
		if err != nil {
			panic(err)
		}
		if outExp.MatchString(absPath) {
			return false
		}
	}

	ins := m.Filter.In
	for _, in := range ins {
		inExp, err := regexp.Compile(in)
		if err != nil {
			panic(err)
		}
		if inExp.MatchString(absPath) {
			return true
		}
	}

	return false
}

type InformationSchema struct {
	TableName string `gorm:"column:table_name"`
}
