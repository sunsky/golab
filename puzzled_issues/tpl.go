package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
	"time"
)

type VolumeTpl struct {
	Source        string
	ContainerName string
	JobStartTime  *time.Time
	Buffer        *bytes.Buffer
}

func (v *VolumeTpl) WriteOutput() string {
	e := ioutil.WriteFile("__stdout_stderr__.log", v.Buffer.Bytes(), 0644)
	return fmt.Sprintf("%v", e)
}
func main() {
	//
	now := time.Now()
	fmt.Println(now.Format("RFC3339"))
	bf2 := &bytes.Buffer{}
	bf2.WriteString("Buffer..in tpl")
	var e error
	tpl, e := template.New("xx").Funcs(template.FuncMap{
		"writeFile": func(filePath string, bf *bytes.Buffer) error {
			return ioutil.WriteFile(filePath, bf.Bytes(), 0644)
		},
		"hasPermission": func(feature string) bool {
			if feature == "A" {
				return true
			}
			return false
		},
	}).Parse("{{.WriteOutput}} {{ writeFile `/labs/go/puzzled_issues/test.log` .Buffer}} {{if hasPermission `A`}}YES{{end}} {{print `-` (.JobStartTime.Format `RFC3339`) \"和\" .JobStartTime.Year `+` .JobStartTime.Month    }}  DIR={{print `VOLUME/` .JobStartTime.Year `/` .JobStartTime.Month `/` .JobStartTime.Day `/` .JobStartTime.Hour `/` .JobStartTime.Minute `/` .JobStartTime.Second}}")

	fmt.Println("222", e)
	e = tpl.Execute(os.Stdout, &VolumeTpl{
		Source:        "Sourcex",
		ContainerName: "Containerx",
		JobStartTime:  &now,
		Buffer:        bf2,
	})
	fmt.Println(e)
}
