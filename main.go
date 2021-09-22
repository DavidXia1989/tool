package main

import (
	servertpl "code.zm.shzhanmeng.com/go-common/zmtool/microserver"
	tpl "code.zm.shzhanmeng.com/go-common/zmtool/template"
	"flag"
	"fmt"
	"github.com/xlab/treeprint"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type config struct {
	// foo
	Name string
	//go build   /usr/local/go1.14/bin
	GoRoot string
	// micro new example -type
	FQDN string
	// github.com/micro/foo
	Dir string
	// $GOPATH/src/github.com/micro/foo
	GoDir string
	// $GOPATH
	GoPath string
	// UseGoPath
	UseGoPath bool
	// Files
	Files []file
	// Comments
	Comments []string
	// Plugins registry=etcd:broker=nats
	Plugins []string
}
type file struct {
	Path string
	Tmpl string
}

func create(c config) error {
	// check if dir exists
	if _, err := os.Stat(c.GoDir); !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists", c.GoDir)
	}
	// just wait
	<-time.After(time.Millisecond * 250)

	fmt.Printf("Creating service %s in %s\n\n", c.FQDN, c.GoDir)

	t := treeprint.New()

	nodes := map[string]treeprint.Tree{}
	nodes[c.GoDir] = t
	// write the files
	for _, file := range c.Files {
		f := filepath.Join(c.GoDir, file.Path)
		dir := filepath.Dir(f)

		b, ok := nodes[dir]
		if !ok {
			d, _ := filepath.Rel(c.GoDir, dir)
			b = t.AddBranch(d)
			nodes[dir] = b
		}

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}
		p := filepath.Base(f)

		b.AddNode(p)
		if err := write(c, f, file.Tmpl); err != nil {
			return err
		}
	}

	// print tree
	fmt.Println(t.String())

	for _, comment := range c.Comments {
		fmt.Println(comment)
	}

	// just wait
	<-time.After(time.Millisecond * 250)

	return nil
}

func main(){
	useGoModule := os.Getenv("GO111MODULE")
	var dir,name,goRoot,projType string
	flag.StringVar(&goRoot, "goroot", "/usr/local/go1.14/bin", "构建时 go目录 /usr/local/go1.14/bin")
	flag.StringVar(&dir, "dir", "", "创建目录")

	if len(os.Args) < 2 {
		fmt.Println("项目名不能为空")
	}
	if len(os.Args) < 3 {
		projType = "proj"
	} else {
		projType = os.Args[2]
	}
	name = os.Args[1]

	flag.Parse()


	if name=="" {
		fmt.Println("项目名不能为空")
		return
	}

	var c config

	switch projType {
		case "api":
			c = config{
				Name:      name,
				Dir:       dir,
				GoRoot:    goRoot,
				Files: []file{
					{name+"/main.go", tpl.MainFunc},
					{name+"/go-build.sh", tpl.Sh_go_build},
					{name+"/shell/check_monitor.sh", tpl.Sh_check_monitor},
					{name+"/shell/monitor_exec.sh", tpl.Sh_monitor_exec},
					{name+"/shell/publish-script.sh", tpl.Sh_publish_script},
					{name+"/conf/app.test.yaml", tpl.Yaml},
					{name+"/conf/app.yaml", tpl.Yaml},
					{name+"/routers/route.go", tpl.Route},
					{name+"/controllers/exampleController/example.go", tpl.ExampleController},
					{name+"/kernel/kernel.go", tpl.Kernel},
					{name+"/common/cmd_run.go", tpl.Cmd_run},
					{name+"/common/common.go", tpl.Common},
					{name+"/common/response.go", tpl.Response},
					{name+"/conf/config.go", tpl.Conf},
					{name+"/.gitignore", tpl.Gitignore},
				},
			}
	    case "server":
			c = config{
				Name:      name,
				Dir:       dir,
				GoRoot:    goRoot,
				Files: []file{
					{name+"/main.go", servertpl.MainFunc},
					{name+"/readme.md", servertpl.Readme},
					{name+"/go-build.sh", servertpl.Sh_go_build},
					{name+"/shell/check_monitor.sh", servertpl.Sh_check_monitor},
					{name+"/shell/monitor_exec.sh", servertpl.Sh_monitor_exec},
					{name+"/shell/publish-script.sh", servertpl.Sh_publish_script},
					{name+"/conf/app.test.yaml", servertpl.Yaml},
					{name+"/conf/app.yaml", servertpl.Yaml},
					{name+"/kernel/kernel.go", servertpl.Kernel},
					{name+"/common/cmd_run.go", servertpl.Cmd_run},
					{name+"/common/common.go", servertpl.Common},
					{name+"/common/response.go", servertpl.Response},
					{name+"/conf/config.go", servertpl.Conf},
					{name+"/.gitignore", servertpl.Gitignore},
					{name+"/proto/example.proto", servertpl.Proto},
					{name+"/domain/service/exampleService.go", servertpl.Service},
					{name+"/handler/example.go", servertpl.Handler},
					{name+"/handler/RegistryHandle.go", servertpl.RegisterHandler},
					{name+"/domain/model/example.go", servertpl.Model},
					{name+"/common/jargre.go", servertpl.Jaegre},
				},
			}
		case "proj":
			c = config{
				Name:      name,
				Dir:       dir,
				GoRoot:    goRoot,
				Files: []file{
					{name+"/main.go", tpl.MainFunc},
					{name+"/go-build.sh", tpl.Sh_go_build},
					{name+"/shell/check_monitor.sh", tpl.Sh_check_monitor},
					{name+"/shell/monitor_exec.sh", tpl.Sh_monitor_exec},
					{name+"/shell/publish-script.sh", tpl.Sh_publish_script},
					{name+"/conf/app.test.yaml", tpl.Yaml},
					{name+"/conf/app.yaml", tpl.Yaml},
					{name+"/routers/route.go", tpl.Route},
					{name+"/controllers/exampleController/example.go", tpl.ExampleController},
					{name+"/kernel/kernel.go", tpl.Kernel},
					{name+"/common/cmd_run.go", tpl.Cmd_run},
					{name+"/common/common.go", tpl.Common},
					{name+"/common/response.go", tpl.Response},
					{name+"/conf/config.go", tpl.Conf},
					{name+"/.gitignore", tpl.Gitignore},
				},
			}
		default:
	}

	if path.IsAbs(dir) {
		fmt.Println("require relative path as service will be installed in GOPATH")
		return
	}

	if useGoModule != "off" {
		c.Files = append(c.Files, file{name+"/go.mod", tpl.Module})
	}

	if err := create(c); err != nil {
		fmt.Println(err)
		return
	}
}
func write(c config, file, tmpl string) error {
	fn := template.FuncMap{
		"title": strings.Title,
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	t, err := template.New("f").Funcs(fn).Parse(tmpl)
	if err != nil {
		return err
	}

	return t.Execute(f, c)
}