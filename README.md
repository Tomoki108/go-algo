# go-algo

solve Atcoder Beginer Contests

## pre-requisists

- install oj, acc, goenv and atcoder's go version

```sh
pip3 install online-judge-tools
npm install -g atcoder-cli

brew update
brew install goenv

# actual atcoder's go version is 1.20.6, but using 1.20.14 to use statick check
# https://img.atcoder.jp/file/language-update/language-list.html
goenv install 1.20.14
```

- add alias to a file like `~/.bash_profile`

```sh
# AtCoder
alias acct='oj t -c "go run main.go"'
alias accs='acct && acc s -s -- -y'
alias accss='acc s -s -- -y'
```

## commands

- how to solve a contest

```sh
cd ABC

acc new {contest_id} # init dir for the contest

cd a # move problem dir

acct # test by samples

accs # test by samples && submit

accss # just submit
```

- update Go template file (sync file registered atocoder-cli with `./template.go`)

```sh
make update-tpl
```
