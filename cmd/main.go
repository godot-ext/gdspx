package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"godot-ext/gdspx/cmd/gdextensionparser"
	"godot-ext/gdspx/cmd/gdextensionparser/clang"
	"godot-ext/gdspx/cmd/generate/ffi"
	"godot-ext/gdspx/cmd/generate/gdext"

	"github.com/spf13/cobra"
)

var (
	verbose          bool
	cleanAll         bool
	cleanGdextension bool
	cleanTypes       bool
	cleanClasses     bool
	genClangAPI      bool
	genExtensionAPI  bool
	packagePath      string
	godotPath        string
	parsedASTPath    string
	buildConfig      string
)

func init() {
	absPath, _ := filepath.Abs(".")
	var (
		defaultBuildConfig string
	)
	if strings.Contains(runtime.GOARCH, "32") {
		defaultBuildConfig = "float_32"
	} else {
		defaultBuildConfig = "float_64"
	}
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", true, "Toggle extra debug output")
	rootCmd.PersistentFlags().BoolVarP(&genClangAPI, "clang-api", "", true, "Generate GDExtension C wrapper")
	rootCmd.PersistentFlags().BoolVarP(&genExtensionAPI, "extension-api", "", false, "Generate Extension API")
	rootCmd.PersistentFlags().StringVarP(&packagePath, "package-path", "p", absPath, "Specified package path")
	rootCmd.PersistentFlags().StringVarP(&godotPath, "godot-path", "", "godot", "Specified path where the Godot executable is located")
	rootCmd.PersistentFlags().StringVarP(&parsedASTPath, "parsed-ast-path", "", "_debug_parsed_ast.json", "Specified path where the AST structure should be written to")
	rootCmd.PersistentFlags().StringVarP(&buildConfig, "build-config", "", defaultBuildConfig, "Specified build configuration for built-in class sizes")
}

var rootCmd = &cobra.Command{
	Use:        "godot-go",
	Aliases:    []string{},
	SuggestFor: []string{},
	Short:      "Godot Go",
	Long:       "",
	Example:    "",
	ValidArgs:  []string{},
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	ArgAliases:             []string{},
	BashCompletionFunction: "",
	Deprecated:             "",
	Hidden:                 false,
	Annotations:            map[string]string{},
	Version:                "",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	PreRun: func(cmd *cobra.Command, args []string) {
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			ast clang.CHeaderFileAST
			err error
		)
		if verbose {
			println(fmt.Sprintf(`build configuration "%s" selected`, buildConfig))
		}
		// generte c++ ext header file
		if genClangAPI {
			if verbose {
				println("Generating gdextension godot ext functions...")
			}
			gdext.GenerateHeader(packagePath)
		}

		// generate go wrap code
		if genClangAPI {
			ast, err = gdextensionparser.GenerateGDExtensionInterfaceAST(packagePath, parsedASTPath)
			if err != nil {
				panic(err)
			}
		}
		if genClangAPI {
			if verbose {
				println("Generating gdextension C wrapper functions...")
			}
			ffi.Generate(packagePath, ast)
			gdext.Generate(packagePath, ast)
		}

		if verbose {
			println("cli tool done")
		}
		return nil
	},
	PostRun: func(cmd *cobra.Command, args []string) {
	},
	PostRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	SilenceErrors:              false,
	SilenceUsage:               false,
	DisableFlagParsing:         false,
	DisableAutoGenTag:          false,
	DisableFlagsInUseLine:      false,
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 0,
	TraverseChildren:           false,
	FParseErrWhitelist:         cobra.FParseErrWhitelist{},
}

func execGoFmt(filePath string) {
	cmd := exec.Command("gofmt", "-w", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Panic(fmt.Errorf("error running gofmt: \n%s\n%w", output, err))
	}
}

func execGoImports(filePath string) {
	cmd := exec.Command("goimports", "-w", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Print(fmt.Errorf("error running goimports: \n%s\n%w", output, err))
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
