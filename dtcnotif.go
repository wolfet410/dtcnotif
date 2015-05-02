package main

import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
    "strings"
)

func main() {
    var inTE, outTE *walk.TextEdit

    MainWindow{
        Title:   "User Generated Title",
        MinSize: Size{600, 400},
        Layout:  VBox{},
        Children: []Widget{
            TextEdit{AssignTo: &outTE, ReadOnly: true},
            PushButton{
                Text: "OK",
                OnClicked: func() {
                    outTE.SetText(strings.ToUpper(inTE.Text()))
                },
            },
        },
    }.Run()
}