package main

import (
	"fmt"
	"log"
	"matrix/core/lic"
	"os"
	"path/filepath"
	"unicode/utf8"
)

var cmdMeta = &Command{
	UsageLine: "meta [module_to_generate_meta]",
	Short:     "Generate models meta",
	Long: `
    generate meta for application
    module_to_generate_meta can be "all" for all models, "app" for app root modles, or other app's modules names.
    sample: hd meta auth
`,
}

func init() {
	cmdMeta.Run = generate_meta
}

func generate_meta(cmd *Command, args []string) int {
	if len(args) < 2 {
		fmt.Fprint(os.Stderr, cmdMeta.Long)
		return 1
	}

	lic_app_name := args[0]
	lic_date := args[1]

	if utf8.RuneCountInString(lic_date) != 8 {
		fmt.Fprintf(os.Stderr, "lic date invalid: %s", lic_date)
	}

	lic_text := lic.EncryptLic(lic_app_name, lic_date)

	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	lic_file_path := currentDir + string(os.PathSeparator) + "app.lic"

	dstFile, err := os.Create(lic_file_path)
	if err != nil {
		log.Fatal(err)
		return 2
	}
	dstFile.WriteString(lic_text)
	dstFile.Close()

	fmt.Println("Your lic is ready:", lic_file_path)
	return 0
}
