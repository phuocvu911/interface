package main

import "flag"

var (
	addr  = flag.String("addr", ":8080", "server listen address")
	paint = flag.Bool("paint", false, "colorize decoded output (bonus)")
)
