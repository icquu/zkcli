package main

import (
    "github.com/urfave/cli"
    "fmt"
    "path"
)

func init() {
    command := cli.Command{
        Name: "get",
        Usage: "Get path info",
        ArgsUsage: "path",
        Action: func(ctx *cli.Context) error {
            if !ctx.Args().Present() {
                cli.ShowSubcommandHelp(ctx)
                return nil
            }
            p := ctx.Args().First()
            data, stat, err := conn.Get(p)
            if err != nil {
                fmt.Println(err)
                return nil
            }
            fmt.Printf("DATA:\n%s\n", string(data))
            if len(data) > 0 {
                println("")
            }
            fmt.Println("STAT:")
            printObj(stat)

            children, _, _ := conn.Children(p)
            if len(children) > 0 {
                fmt.Println("\nCHILDREN:")
                for _, item := range children {
                    fmt.Println(path.Join(p, item))
                }
            }
            return nil
        },
    }
    app.Commands = append(app.Commands, &command)
    app.BashComplete = nodeCompletion
}


