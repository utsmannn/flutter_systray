package main

// DO NOT EDIT, this file is generated by hover at compile-time for the flutter_systray plugin.

import (
	flutter_systray "github.com/JanezStupar/flutter_systray/go"
	flutter "github.com/go-flutter-desktop/go-flutter"
)

func init() {
	// Only the init function can be tweaked by plugin maker.
	options = append(options, flutter.AddPlugin(&flutter_systray.FlutterSystrayPlugin{}))
}
