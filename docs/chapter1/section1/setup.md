# 環境構築
::: tip

質問をするときにはできるだけスクリーンショットを貼るようにしましょう。テキストだけで説明しても解決に必要な情報を全て伝えるのは難しいです。

Mac: `Control + Shift + Command + 4`を押すと、矩形選択でスクリーンショットが撮れます。 traQ のメッセージ入力欄に`Command + V`で貼り付けられます。  
Windows: `Winキー + Shift + S`を押すと、矩形選択でスクリーンショットが撮れます。 traQ のメッセージ入力欄に`Ctrl + V`で貼り付けられます。
:::


## 拡張機能の導入

### 共通

- Go
- Volar (第 3 回で補足)
- ESLint (第 3 回で補足)

::: info
コマンドは手打ちではなくてコピー&ペーストで打ってください。
手打ちだと写し間違いをする可能性があります。
:::

## goのinstall

https://golang.org/doc/install
インストールが終わった後に`go version`してみて`go version go1.20.4`と出れば成功です。

通常のインストールと asdf を使ったインストールの 2 種類がありますが、asdf を使った方が後からバージョンを上げるのが簡単になるので、長期的にオススメです。
どちらか好みのほうを選択しましょう。
## 通常のインストール

::: info
コマンドは一行づつ実行してください。
:::

### Mac

Mac のタブを選択し、pkg をダウンロード->インストーラ起動で設定完了です。
もしくは`brew install go@1.20`を実行してください。

### Windows (WSL2)
``` bash
sudo apt install tar git
wget https://golang.org/dl/go1.17.10.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.10.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bash_profile
source ~/.bash_profile
```

## with asdf(バージョン管理ツール)

asdf を導入してから、その asdf を使って golang を導入します。

### asdf導入

[公式資料](https://asdf-vm.com/#/core-manage-asdf)

#### Mac
bash
``` bash
brew install asdf
echo -e "\n. $(brew --prefix asdf)/libexec/asdf.sh" >> ~/.bash_profile
echo -e "\n. $(brew --prefix asdf)/etc/bash_completion.d/asdf.bash" >> ~/.bash_profile
source ~/.bash_profile
```
zsh
``` zsh
brew install asdf
echo -e "\n. $(brew --prefix asdf)/libexec/asdf.sh" >> ${ZDOTDIR:-~}/.zshrc
source ~/.zshrc
```
#### Windows(WSL2)

``` bash
sudo apt install git
git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.11.3
echo '. $HOME/.asdf/asdf.sh' >> ~/.bashrc
echo '. $HOME/.asdf/completions/asdf.bash' >> ~/.bashrc
source ~/.bashrc
```

### Golang 導入

``` bash
asdf plugin-add golang
asdf install golang 1.20.4
asdf global golang 1.20.4
```
