package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "command_client"
	app.Usage = "github.com/codegangsta/cli のテンプレートです"
	app.Version = "0.0.1"

	// グローバルオプション設定
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "dryrun, d", // 省略指定 => d
			Usage: "グローバルオプション dryrunです。",
		},
	}

	app.Commands = []cli.Command{
		// コマンド設定
		{
			Name:    "hello",
			Aliases: []string{"h"},
			Usage:   "hello world を表示します",
			Action:  helloAction,
		},
	}

	app.Before = func(c *cli.Context) error {
		// 開始前の処理をここに書く
		fmt.Println("開始")
		return nil // error を返すと処理全体が終了
	}

	app.After = func(c *cli.Context) error {
		// 終了時の処理をここに書く
		fmt.Println("終了")
		return nil
	}

	app.Run(os.Args)
}

func helloAction(c *cli.Context) {

	// グローバルオプション
	var isDry = c.GlobalBool("dryrun")
	if isDry {
		fmt.Println("this is dry-run")
	}

	// パラメータ
	var paramFirst = ""
	if len(c.Args()) > 0 {
		paramFirst = c.Args().First() // c.Args()[0] と同じ意味
	}

	fmt.Printf("Hello world! %s\n", paramFirst)
}
