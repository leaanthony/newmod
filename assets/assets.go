package assets

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/leaanthony/debme"
	"path/filepath"
	"strings"
)

//go:embed main/main.go
var MainGoFile []byte

//go:embed main/README.md
var Readme []byte

//go:embed licenses
var licenses embed.FS

func LicenseIsValid(licenseKey string) bool {
	l, err := debme.FS(licenses, "licenses")
	if err != nil {
		return false
	}

	_, err = l.Open(licenseKey + ".txt")
	return err == nil
}

func CopyLicense(licenseKey string, path string) error {
	key := strings.ToLower(licenseKey)
	l, err := debme.FS(licenses, "licenses")
	if err != nil {
		return err
	}
	err = l.CopyFile(key+".txt", filepath.Join(path, "LICENSE"), 0755)

	if err != nil {
		return fmt.Errorf("license key '%s' is invalid", licenseKey)
	}
	return nil
}