package main

import (
	li "demo/logwrapper"
)

func main() {
	var standardLogger = li.NewLogger()
	standardLogger.InvalidArgValue("client", "nil")
}
