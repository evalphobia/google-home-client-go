Google Home client for golang
----

[![GoDoc][1]][2] [![License: MIT][3]][4] [![Release][5]][6] [![Build Status][7]][8] [![Codecov Coverage][11]][12] [![Go Report Card][13]][14] [![Downloads][15]][16]

[1]: https://godoc.org/github.com/evalphobia/google-home-client-go?status.svg
[2]: https://godoc.org/github.com/evalphobia/google-home-client-go
[3]: https://img.shields.io/badge/License-MIT-blue.svg
[4]: LICENSE.md
[5]: https://img.shields.io/github/release/evalphobia/google-home-client-go.svg
[6]: https://github.com/evalphobia/google-home-client-go/releases/latest
[7]: https://travis-ci.org/evalphobia/google-home-client-go.svg?branch=master
[8]: https://travis-ci.org/evalphobia/google-home-client-go
[9]: https://coveralls.io/repos/evalphobia/google-home-client-go/badge.svg?branch=master&service=github
[10]: https://coveralls.io/github/evalphobia/google-home-client-go?branch=master
[11]: https://codecov.io/github/evalphobia/google-home-client-go/coverage.svg?branch=master
[12]: https://codecov.io/github/evalphobia/google-home-client-go?branch=master
[13]: https://goreportcard.com/badge/github.com/evalphobia/google-home-client-go
[14]: https://goreportcard.com/report/github.com/evalphobia/google-home-client-go
[15]: https://img.shields.io/github/downloads/evalphobia/google-home-client-go/total.svg?maxAge=1800
[16]: https://github.com/evalphobia/google-home-client-go/releases
[17]: https://img.shields.io/github/stars/evalphobia/google-home-client-go.svg
[18]: https://github.com/evalphobia/google-home-client-go/stargazers

# Quick Usage

```go
package main

import (
	"github.com/evalphobia/google-home-client-go/googlehome"
)

func main() {
	cli, err := googlehome.NewClientWithConfig(googlehome.Config{
		Hostname: "192.168.0.1",
		Lang:     "en",
		Accent:   "GB",
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// Speak text on Google Home.
	cli.Notify("Hello")

	// Change language
	cli.SetLang("ja")
	cli.Notify("こんにちは、グーグル。")

	// Or set language in Notify()
	cli.Notify("你好、Google。", "zh")
}
```

## Environment variables

|Name|Description|
|:--|:--|
| `GOOGLE_HOME_HOST` | Hostname or IP address of Google Home for speech feature. |
| `GOOGLE_HOME_PORT` | Port number of Google Home. Default is `8009`. |
| `GOOGLE_HOME_LANG` | Speaking language of Google Home. Default is `en`. |
| `GOOGLE_HOME_ACCENT` | Speaking accent of Google Home. Default is `us`. |


# Credit

This library is based on [github.com/kunihiko-t/google-home-notifier-go](https://github.com/kunihiko-t/google-home-notifier-go) by [kunihiko-t](https://github.com/kunihiko-t/) and heavily depends on [github.com/barnybug/go-cast](https://github.com/barnybug/go-cast).

This port version supports environment value config and some feature.
