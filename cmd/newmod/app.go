package main

import (
	"fmt"
	"github.com/leaanthony/clir"
	"github.com/leaanthony/newmod/assets"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

const version = "v1.0.0"

func Run(args []string, output io.Writer) error {
	cli := clir.NewCli("newmod", "A Go Module scaffolder", version)

	modulepath := ""
	cli.StringFlag("name", " The name of the module, EG: github.com/myusername/mymod", &modulepath)
	cmd := false
	cli.BoolFlag("cmd", "Generate the cmd structure", &cmd)
	license := ""
	cli.StringFlag("license", "Generate license file (mit,apache,gpl2,gpl3)", &license)
	verbose := false
	cli.BoolFlag("verbose", "Verbose output", &verbose)

	cli.Action(func() error {
		if modulepath == "" {
			return fmt.Errorf("please provide a module name")
		}
		if license != "" && !assets.LicenseIsValid(license) {
			return fmt.Errorf("invalid license key '%s'", license)
		}
		// create the directory
		moduleName := filepath.Base(modulepath)
		err := os.Mkdir(moduleName, 0755)
		if err != nil {
			return err
		}

		// cd to the directory
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		moduleDir := filepath.Join(cwd, moduleName)
		err = os.Chdir(moduleDir)
		if err != nil {
			return err
		}

		// Create cmd directory structure
		if cmd {
			cliDir := filepath.Join(moduleDir, "cmd", moduleName)
			err = os.MkdirAll(cliDir, 0755)
			if err != nil {
				return err
			}

			err := os.WriteFile(filepath.Join(cliDir, "main.go"), assets.MainGoFile, 0755)
			if err != nil {
				return err
			}
		}

		// Generate README
		err = os.WriteFile(filepath.Join(moduleDir, "README.md"), assets.Readme, 0755)
		if err != nil {
			return err
		}

		// Generate License
		if license != "" {
			assets.CopyLicense(license, moduleDir)
		}

		// Run `go mod init`
		cmd := exec.Command("go", "mod", "init", modulepath)
		if !verbose {
			cmd.Stdout = nil
		}
		err = cmd.Run()
		if err != nil {
			return err
		}

		// Run git init
		cmd = exec.Command("git", "init")
		if !verbose {
			cmd.Stdout = nil
		}
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			return err
		}

		return nil
	})

	return cli.Run(args...)
}
