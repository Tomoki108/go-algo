# go-algo

solve Atcoder Beginer Contests

## pre-requisists

- install oj, acc, goenv and atcoder's go version

```sh
pip3 install online-judge-tools
npm install -g atcoder-cli

brew update
brew install goenv

goenv install 1.20.6
```

- add alias to `~/.bash_profile` or some file like that

```sh
# AtCoder
alias acct='oj t -c "go run main.go"'
alias accs='actest && acc s'
```

## commands

- how to solve a contest

```sh
cd ABC

acc new {contest_id} # init dir for the contest

cd a # move problem dir

acct # test by samples

accs # test by samples && submit
```

- update Go template file (sync file registered atocoder-cli with `./template.go`)

```sh
make update-tpl
```
