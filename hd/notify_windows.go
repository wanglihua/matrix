package main

import (
	"github.com/lxn/walk"
	"log"
)

var notify_icon *walk.NotifyIcon

func notify_setup() {
	var err error
	//mw, err := walk.NewMainWindow()
	//if err != nil {
	//	log.Fatal(err)
	//}

	icon := walk.IconApplication()

	notify_icon, err = walk.NewNotifyIcon()
	if err != nil {
		log.Fatal(err)
	}

	if err := notify_icon.SetIcon(icon); err != nil {
		log.Fatal(err)
	}

	if err := notify_icon.SetToolTip("hd build notify"); err != nil {
		log.Fatal(err)
	}

	if err := notify_icon.SetVisible(true); err != nil {
		log.Fatal(err)
	}

	//mw.Run()
}

func notify_dispose() {
	if notify_icon == nil {
		return
	}

	err := notify_icon.Dispose()
	if err != nil {
		log.Fatal(err)
	}
}

func notify_show_success() {
	icon := walk.IconWinLogo()
	if err := notify_icon.SetIcon(icon); err != nil {
		log.Fatal(err)
	}

	if err := notify_icon.ShowInfo("Success", "Build Success and Running!"); err != nil {
		log.Fatal(err)
	}
}

func notify_show_error() {
	icon := walk.IconError()
	if err := notify_icon.SetIcon(icon); err != nil {
		log.Fatal(err)
	}

	if err := notify_icon.ShowError("Error", "Stopped and Build Error!"); err != nil {
		log.Fatal(err)
	}
}
