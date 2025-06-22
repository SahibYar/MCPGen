package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("mcpgen CLI - Modular Code Pipeline Generator")
	fmt.Println("Usage: mcpgen <openapi-file> <arazzo-file> [options]")
	// TODO: Wire up OpenAPILoader, ArazzoParser, FlowCompiler, CodeGenerator, PluginManager
	os.Exit(0)
}
