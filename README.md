# go-algo

challenge Atcoder Beginer Contests

## pre-requisists

- install oj, acc, goenv, atcoder's go version, etc.

```sh
pip3 install online-judge-tools
npm install -g atcoder-cli

brew update
brew install goenv

# NOTE: I was usign go 1.20.14 till abc397
# current version: https://img.atcoder.jp/file/language-update/2025-10/language-list.html?_gl=1*2e4isk*_ga*Mjc0MTk1OTUxLjE3NjAyNDc5ODA.*_ga_RC512FD18N*czE3NjcxNTkwNjQkbzYkZzEkdDE3NjcxNTkxNDgkajYwJGwwJGgw
goenv install 1.25.1

# install compatible version of gopls(language server) and dlv(debug tool).
go install golang.org/x/tools/gopls@v0.15.3
go install github.com/go-delve/delve/cmd/dlv@v1.20.2
```

- setup atcoder-cli

```sh
# to download all tasks for acc new
acc config default-task-choice all

# language setting
cd `acc config-dir`
mkdir go && cd go
vi template.json # then copy and paste ./template.json
vi main.go # then copy and paste ./template.go
```

- add goenv setting and alias to `~/.bashrc` or `~/.zshrc`

```sh
# goenv
eval "$(goenv init -)"
export PATH="$HOME/.goenv/shims:$PATH"

# AtCoder
alias acct='oj t -c "go run main.go"'
alias accs='acct && acc s -s -- -y'
alias accss='acc s -s -- -y'
```

## commands

- how to challenge a contest

```sh
cd ABC

acc new {contest_id} # init dir for the contest

cd a # move problem dir

acct # test by samples

# currently unable to submit by cli due to AtCoder's update (ref: https://github.com/Tatamo/atcoder-cli/issues/68)
# accs # test by samples && submit
# accss # just submit
```

- update Go template file (sync file registered atocoder-cli with `./template.go`)

```sh
make update-tpl
```
