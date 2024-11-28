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

## commands

- コンテストの流れ（一部 `~/.bash_profile` の alias を利用）

```sh
cd ABC

acc new {contest_id} # ディレクトリ初期化

cd a # 解きたい問題のディレクトリに移動

acct # サンプルでテスト

accs # サンプルでテスト && 提出
```

- go の template ファイルを更新（atocoder-cli に登録しているファイルを `./template.go` と同期）

```sh
make update-tpl
```
