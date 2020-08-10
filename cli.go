package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/rakyll/statik/fs"
	_ "github.com/sago35/umedago-cli-init/statik"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	appName        = "umedago-cli-init"
	appDescription = ""
)

type cli struct {
	outStream io.Writer
	errStream io.Writer
}

// コマンドライン引数
var (
	app    = kingpin.New(appName, appDescription)
	target = app.Arg("target", "target app name").Required().String()
)

// Run : main 関数にあたる本体
func (c *cli) Run(args []string) error {
	app.UsageWriter(c.errStream)

	// コマンドラインオプションの処理
	if VERSION != "" {
		app.Version(fmt.Sprintf("%s version %s build %s", appName, VERSION, BUILDDATE))
	} else {
		app.Version(fmt.Sprintf("%s version - build -", appName))
	}
	app.HelpFlag.Short('h')

	k, err := app.Parse(args[1:])
	if err != nil {
		return err
	}

	switch k {
	default:
		err := umedagoCliInit(*target)
		if err != nil {
			return err
		}
	}

	return nil
}

func umedagoCliInit(appName string) error {
	fmt.Printf("umedagoCliInit(%q)\n", appName)
	sfs, err := fs.New()
	if err != nil {
		return err
	}

	err = fs.Walk(sfs, `/`, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			target := filepath.Join(appName, path)
			fmt.Println(path, `->`, target)

			err := os.Mkdir(filepath.Dir(target), 0777)
			if err != nil && !os.IsExist(err) {
				return err
			}

			w, err := os.Create(target)
			if err != nil {
				return err
			}
			defer w.Close()

			b, err := fs.ReadFile(sfs, path)
			if err != nil {
				return err
			}

			//_, err = w.Write(b)
			//if err != nil {
			//	return err
			//}
			tmpl, err := template.New(`xxx`).Parse(string(b))
			if err != nil {
				return err
			}

			err = tmpl.Execute(w, struct{ App, User, Email string }{
				App:   appName,
				User:  `sago35`,
				Email: `sago35@gmail.com`,
			})
			if err != nil {
				return err
			}

		}
		return nil
	})
	if err != nil {
		return err
	}

	{
		cmd := exec.Command(`git`, `init`, `.`)
		cmd.Dir = appName
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	{
		cmd := exec.Command(`git`, `add`, `.`)
		cmd.Dir = appName
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	{
		cmd := exec.Command(`git`, `status`)
		cmd.Dir = appName
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
