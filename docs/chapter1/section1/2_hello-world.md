# Rust で Hello World

`~/develop/rust-hello`ディレクトリの中にプログラムを作成します。

## VSCodeで`~/develop/rust-hello`ディレクトリを開く

- ディレクトリ作成
`mkdir -p ~/develop/rust-hello`

- 移動
`cd ~/develop/rust-hello`

- vscode を開く
`code .`

## Rust プロジェクトの初期化

`~/develop/rust-hello`ディレクトリで以下のコマンドを実行します。

- Rust プロジェクトの初期化
`cargo init`

実行すると、`src/main.rs`を含むいくつかのファイルが生成されます。`src/main.rs`には以下の内容が記述されています。

<<<@/chapter1/section1/src/main.rs

## 実行

`cargo run`

うまくできれば結果は次のようになります。

```bash
kenken@kenken:~/develop/rust-hello$ cargo run
   Compiling rust_playground v0.1.0 (/home/kenken/main/temp/rust_playground)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.31s
     Running `target/debug/rust_playground`
Hello, world!
```
