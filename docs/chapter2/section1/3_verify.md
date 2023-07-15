# 検証

## 完成形

まずは完成形です。

<details>

main.go 

<<<@/chapter2/section1/src/0/final/main.go{go:line-numbers}

handler.go

<<<@/chapter2/section1/src/0/final/handler.go{go:line-numbers}


</details>

## 検証

自分の実装が正しく動くか検証しましょう。

:::warning
全て Postman での検証です
`go run main.go`でサーバーを起動した状態で行ってください
:::

<http://localhost:8080/cities/Tokyo> へ
初めに普通にアクセスするとダメです
![](https://md.trapti.tech/uploads/upload_96a03d609e761150a2136963dd34006a.png)

ユーザーを作成します。
上手く作成できれば Status 201 が返ってくるはずです。
![](https://md.trapti.tech/uploads/upload_4d891187b392debc9732aeff7ecaca08.png)

そのままパスを変えてログインリクエストを送ります。
![](https://md.trapti.tech/uploads/upload_7b21cf42397801806bab12f5180ce888.png)

ログインに成功したら、レスポンスの方の Cookies を開いて value の中身をコピーします
![](https://md.trapti.tech/uploads/upload_13756985fc7e93bd6d032083d340ea6b.png)

リクエストの方の Headers で Cookie をセットします。

Key に`Cookie`を
Value に`sessions=(コピーした値);`をセットします(既に自動で入っている場合もあります、その場合は追加しなくて大丈夫です)。

もう一度 <http://localhost:8080/cities/Tokyo> にアクセスすると正常に API が取れるようになりました。
![](https://md.trapti.tech/uploads/upload_59c6c86e127d982f511946d2a183d0a6.png)

ここで、作成されたユーザーがデータベースに保存されていることを確認してみましょう。
`mysql > SELECT * FROM users;`
![](https://md.trap.jp/uploads/upload_f713b7da16df6729729a25ca2b5a6816.png)
ユーザー名とハッシュ化されたパスワードが確認できますね。
![](https://md.trap.jp/uploads/upload_7f007d73bd0ff508dff12246546b1a5b.png)
ちょっと分かりにくい表示ですが、セッションもしっかり確認できます。

