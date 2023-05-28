# naro-text

## 記法
https://vitepress.dev/guide/markdown

## 記入の注意点
- 内輪ネタは避けましょう
- section以下は基本フラットな構造を保って下さい
- HackMD独自記法をすべてカバーしているわけではないので注意しましょう
- 節タイトルは内容を簡潔に表したものにしてください
  - その１，その２みたいなのは出来るだけ避ける
- container（`:::info` など）は以下の方針を目安に利用してください
  - info: 注釈（git/cliの操作法など）
  - tips: 発展的な内容、コラム
  - warning: 注意事項（気をつけないと正常に実行できないことがあるもの）
  - danger: 破壊的な事象（サーバーに負荷を与える操作など）
  - 練習問題はheadingで区切って、containerは使用しない
    - 答えを載せる場合はdetailsを使用して隠す