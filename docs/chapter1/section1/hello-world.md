# Go で Hello World

`~/develop/go/hello`ディレクトリの中にプログラムを作成します。

## VSCodeで`~/develop/go/hello`ディレクトリを開く
- ディレクトリ作成
`mkdir -p ~/develop/go/hello`

- 移動
`cd ~/develop/go/hello`

- vscode を開く
`code .`


## `main.go`の作成

- `main.go`を作成する
- ファイル > 新規ファイル (または``Ctrl + N``)  で空のファイルを開く
- ファイル > 名前をつけて保存 (または``Ctrl + Shift + S``) で`main.go`と命名して保存する

以下の内容を入力して保存(`Ctrl + S`)します。

<<<@/chapter1/section1/src/main.go

## 実行

`go run main.go`

うまくできれば結果は次のようになります。
```bash
hijiki51@DESKTOP-JF9KJFE:~/develop/go/hello$ go run main.go
Hello, World!
```
