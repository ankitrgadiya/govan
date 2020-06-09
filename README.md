<div align="center">
    <img width="200" height="200" alt="Govan"
    src="https://user-images.githubusercontent.com/18071765/84123804-f7884100-aa57-11ea-9db9-87831104f0f5.png">
</div>

# Overview

Golang package management comes built-in with standard distribution. There is no
central package repository (like NPM, RubyGems) in Go. Instead code repository
in itself is a package. Go uses specific meta tags or query paramters to figure
out which VCS to use and where is the source code located.

All popular code hosting provides have configured the repository paths to work
out of the box with Go and so packages like
`[github.com/ankitrgadiya/govan](http://github.com/ankitrgadiya/govan)` is
commonly used in projects. However the problem with such imports is that its
tightly coupled with hostring provider and changing that will break all the
dependencies ([Not Exactly](https://proxy.golang.org/)).

Go supports using [custom vanity
import](https://golang.org/cmd/go/#hdr-Remote_import_paths) paths which can be
used to mitigate the issue. A custom domain name can be used like
`argc.in/govan` instead of code hosting's domain which gives the developers
flexibility to move around different hosting providers. This also simplifies the
imports for the consumers of packages.

<blockquote class="twitter-tweet">
  <p lang="en" dir="ltr">
    Non vanity import paths are considered harmful. If you are a serious
    project, enforce a vanity import path. <a
    href="https://twitter.com/hashtag/golang?src=hash&amp;ref_src=twsrc%5Etfw">#golang</a>
  </p>
  &mdash; Jaana Dogan (@rakyll) <a
  href="https://twitter.com/rakyll/status/892805962867683328?ref_src=twsrc%5Etfw">August
  2, 2017</a>
</blockquote> 

Govan is a static HTML generator for vanity imports. It generates HTML files for
the packages which implements the necesarry meta tags for Go distribution to
process and also redirects the user (from browser) to package documentation.

# Installation

Govan can be installed in the `GOPATH` using the following command.

```go
go get -u argc.in/govan
```

- [ ]  Setup Releases with Pre-built binaries for various platforms

# Usage

Govan requires initial domain setup which can be done easily using the `init`
sub-command. This will create a configuration file `govan.yaml` in the current
directory and create HTML for root redirect in the output directory. By default,
the output directory is `./site` which can be overriden by `--output` flag.

```bash
govan init argc.in
OR
govan init argc.in --output ./output
```

After the initial setup, packages can be added using `add` sub-command. This
takes package path (without domain) and source url. This will update the
configuration file and generate new HTML file for the package under output
directory. For updating the source of an already existing package, same command
can be re-used.

```bash
govan add govan github.com/ankitrgadiya/govan
```

If you need to remove an existing package, `remove` sub-command can be used.
This just requires package name. This is an no-op command and it won't complain
if there is no such package in the configuration.

```bash
govan remove govan
```

Finally, if you need to manually modify the configuration file for any reason or
if there is a requirement to run govan in non-interactive environment then
`generate` sub-command can be used to (re-)generate HTML files for all the
packages.

```bash
govan generate
```

# License

This project uses [BSD 3-Clause](./LICENSE) License.
