# umedago-cli-init

CLI ツールの雛形生成ツール

## Description

以下の機能を持った雛形を生成します

* テスト機能
* リリース機能

## Requirement

Go 1.13 or later

## Usage

    usage: umedago-cli-init [<flags>] [<target>]

    Flags:
      -h, --help     Show context-sensitive help (also try --help-long and
                     --help-man).
          --version  Show application version.

    Args:
      [<target>]  target app name

## Installation

    go get github.com/sago35/umedago-cli-init

## Build

ソースコードは、 `%GOPATH%/src/github.com/sago35/umedago-cli-init/main.go` となるように配置してください。  

`go run dist/make_dist.go VERSION` を実行すると、 ./dist/release 以下にファイルが作成されます。  
release用にBuildする際は、Changes.mdに該当Versionの記載が必要です。  

    (例) release用にv1.2.3でbuild
    $ go run dist/make_dist.go 1.2.3

    (例) 開発用にv1.2.3でbuild
    $ go run dist/make_dist.go 1.2.3 --dev

public 以下を更新したときは以下を実行してください。  
実行には [statik](https://github.com/rakyll/statik) が必要です。  

    $ go generate

### Environment

* go
* kingpin.v2

## Notice

なし

## License

Copyright (c) 2020 sago35  
Released under the MIT license  
https://opensource.org/licenses/mit-license.php  

## Author

sago35 - <sago35@gmail.com>
