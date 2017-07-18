package go_crud

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func Start() {
	config := NewConfig()

	fp, err := os.Open(config.Tables)
	if err != nil {
		panic(err)
	}
	defer func() {
		if fp != nil {
			fp.Close()
		}
	}()

	result.Headers = append(result.Headers, "FilePath")

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		result.Headers = append(result.Headers, scanner.Text())
	}

	err = filepath.Walk(config.Target, Apply)
	if err != nil {
		panic(err)
	}

	breakProj := ""
	breakOxAry := make([]string, len(result.Headers)-1)

	editBodies := [][]string{}
	proj := ""
	for idx, body := range result.Bodies {
		paths := strings.Split(body[0], "\\")
		proj = paths[1]
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

	tmpl := template.Must(template.ParseFiles("../tmpl.csv"))
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
	Headers []string
	Bodies  [][]string
}

var result = &Result{
	Headers: []string{},
	Bodies:  [][]string{},
}
