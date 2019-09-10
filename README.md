# tty

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Sponsor][sponsor-image]][sponsor-url]

Terminal detection for Go.

## API

```go
isTerminal := tty.IsTerminal(os.Stdout.Fd())
```

## Authors

* Yasuhiro Matsumoto (a.k.a mattn)
* Eduard Urbach (minor modifications)

## Style

Please take a look at the [style guidelines](https://github.com/akyoto/quality/blob/master/STYLE.md) if you'd like to make a pull request.

## Sponsors

| [![Cedric Fung](https://avatars3.githubusercontent.com/u/2269238?s=70&v=4)](https://github.com/cedricfung) | [![Scott Rayapoullé](https://avatars3.githubusercontent.com/u/11772084?s=70&v=4)](https://github.com/soulcramer) | [![Eduard Urbach](https://avatars3.githubusercontent.com/u/438936?s=70&v=4)](https://eduardurbach.com) |
| --- | --- | --- |
| [Cedric Fung](https://github.com/cedricfung) | [Scott Rayapoullé](https://github.com/soulcramer) | [Eduard Urbach](https://eduardurbach.com) |

Want to see [your own name here?](https://github.com/users/akyoto/sponsorship)

[godoc-image]: https://godoc.org/github.com/akyoto/tty?status.svg
[godoc-url]: https://godoc.org/github.com/akyoto/tty
[report-image]: https://goreportcard.com/badge/github.com/akyoto/tty
[report-url]: https://goreportcard.com/report/github.com/akyoto/tty
[tests-image]: https://cloud.drone.io/api/badges/akyoto/tty/status.svg
[tests-url]: https://cloud.drone.io/akyoto/tty
[coverage-image]: https://codecov.io/gh/akyoto/tty/graph/badge.svg
[coverage-url]: https://codecov.io/gh/akyoto/tty
[sponsor-image]: https://img.shields.io/badge/github-donate-green.svg
[sponsor-url]: https://github.com/users/akyoto/sponsorship
