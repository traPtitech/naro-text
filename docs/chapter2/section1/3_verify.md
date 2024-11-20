# 検証

## 完成形

まずは完成形です。

:::details 完成形
::: code-group
<<<@/chapter2/section1/src/final/main.rs{rs:line-numbers}[main.rs]
<<<@/chapter2/section1/src/final/handler.rs{rs:line-numbers}[handler.rs]
<<<@/chapter2/section1/src/final/repository.rs{rs:line-numbers}[repository.rs]
<<<@/chapter2/section1/src/final/handler/country.rs{rs:line-numbers}[handler/country.rs]
<<<@/chapter2/section1/src/final/repository/country.rs{rs:line-numbers}[repository/country.rs]
<<<@/chapter2/section1/src/final/handler/auth.rs{rs:line-numbers}[handler/auth.rs]
<<<@/chapter2/section1/src/final/repository/users.rs{rs:line-numbers}[repository/users.rs]
<<<@/chapter2/section1/src/final/repository/users_session.rs{rs:line-numbers}[repository/users_session.rs]
:::

## 検証

自分の実装が正しく動くか検証しましょう。

:::warning
全て Postman での検証です。  
`go run main.go`でサーバーを起動した状態で行ってください。

また、`GET`と`POST`を間違えないようにして下さい。
:::

<a href="http://localhost:8080/cities/Tokyo">localhost:8080/cities/Tokyo</a>にアクセスすると、ログインしていないため`401 Unauthorized`が返ってきます。そのため、情報を入手できません。
![](images/3/postman1-unauthorized.png)

ユーザーを作成します。
上手く作成できれば Status 201 が返ってくるはずです。  
(注意:`POST`です)
![](images/3/postman2-signup.png)

そのままパスを変えてログインリクエストを送ります。
![](images/3/postman3-login.png)

ログインに成功したら、レスポンスの方の Cookies を開いて value の中身をコピーします
![](images/3/postman4-cookie-place.png)

![](images/3/postman5-cookie-value.png)

リクエストの方の Headers で Cookie をセットします。

Key に`Cookie`を
Value に`sessions={コピーした値};`をセットします(既に自動で入っている場合もあります、その場合は追加しなくて大丈夫です)。

もう一度 <a href="http://localhost:8080/cities/Tokyo">localhost:8080/cities/Tokyo</a> にアクセスすると正常に API が取れるようになりました。  
(注意:`GET`です)
![](images/3/postman6-authorized.png)

ここで、作成されたユーザーがデータベースに保存されていることを確認してみましょう。
`mysql > SELECT * FROM users;`
![](images/3/database1-user.png)

ユーザー名とハッシュ化されたパスワードが確認できますね。
![](images/3/database2-session.png)

ちょっと分かりにくい表示ですが、セッションもしっかり確認できます。
