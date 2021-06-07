# newmod

A utility to quickly scaffold new Go modules.

## Install

`go get github.com/leaanthony/newmod`

## Usage

`newmod -name "your/module/path" [-cmd] [-license <license key>]`

## What it does

  - Creates a directory using the last part of the module path
  - Runs `go mod init` using your module path
  - Runs `git init`
  - Add a README.md
  - If `-cmd` is given: 
    - Create the `cmd` directory structure
    - Create a basic `main.go` file
  - If `-license` is given:
    - Create a `LICENSE.md` file with the corresponding license text
    
### Licenses

Licenses that may be generated are:

| License key | License | Notes |
| ----------- | ------- | ----- |
| mit         | MIT License | A short and simple permissive license with conditions only requiring preservation of copyright and license notices. Licensed works, modifications, and larger works may be distributed under different terms and without source code. |
| apache2     | Apache License 2.0 | A permissive license whose main conditions require preservation of copyright and license notices. Contributors provide an express grant of patent rights. Licensed works, modifications, and larger works may be distributed under different terms and without source code. |
| agpl3       | GNU Affero General Public License v3.0 | Permissions of this strongest copyleft license are conditioned on making available complete source code of licensed works and modifications, which include larger works using a licensed work, under the same license. Copyright and license notices must be preserved. Contributors provide an express grant of patent rights. When a modified version is used to provide a service over a network, the complete source code of the modified version must be made available. |
| gpl3        | GNU General Public License v3.0 | Permissions of this strong copyleft license are conditioned on making available complete source code of licensed works and modifications, which include larger works using a licensed work, under the same license. Copyright and license notices must be preserved. Contributors provide an express grant of patent rights. |
| lgpl3       | GNU Lesser General Public License v3.0 | Permissions of this copyleft license are conditioned on making available complete source code of licensed works and modifications under the same license or the GNU GPLv3. Copyright and license notices must be preserved. Contributors provide an express grant of patent rights. However, a larger work using the licensed work through interfaces provided by the licensed work may be distributed under different terms and without source code for the larger work. |
| mpl2        | Mozilla Public License 2.0 | Permissions of this weak copyleft license are conditioned on making available source code of licensed files and modifications of those files under the same license (or in certain cases, one of the GNU licenses). Copyright and license notices must be preserved. Contributors provide an express grant of patent rights. However, a larger work using the licensed work may be distributed under different terms and without source code for files added in the larger work. |
| boost       | Boost Software License 1.0 | A simple permissive license only requiring preservation of copyright and license notices for source (and not binary) distribution. Licensed works, modifications, and larger works may be distributed under different terms and without source code. |
| unlicense   | The Unlicense | A license with no conditions whatsoever which dedicates works to the public domain. Unlicensed works, modifications, and larger works may be distributed under different terms and without source code. |

Credit: Above text from [Choose a license](https://choosealicense.com).