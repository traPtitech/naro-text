# 演習問題

基本問題は解答を置いてありますが、できるだけ自力で頑張ってみてください。

最後に全問題をまとめて動作確認できるシェルスクリプトが置いてあるので、作ったエンドポイントは消さないで残しておくと良いです。

## 基本問題 GET /ping

pong と返すエンドポイントを作成してください。

### スキーマ
#### リクエスト
`/ping`


#### レスポンス
content-type: `text/plain`
status code: `200`
```
pong
```

#### 完成したら、以下のコマンドをターミナルで実行して上手く機能しているか確認しましょう。
```bash
$ curl -X GET "http://localhost:8080/ping" # pong
```
:::details 解答
<<<@/chapter1/section3/src/4-1_ping.go
:::

## 基本問題 GET /fizzbuzz
クエリパラメータ`count`で渡された数までの FizzBuzz を返してください。  
`count`がない場合は 30 として扱い、`count`が整数として解釈できない場合はステータスコード`400`を返してください。

### スキーマ
#### リクエスト
`/fizzbuzz?count=10`

#### レスポンス
content-type: `text/plain`
status code: `200`
```
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
```

##### 適切でないリクエストの場合
content-type: `text/plain`
status code: `400`
```
Bad Request
```

#### 完成したら、以下のコマンドをターミナルで実行して上手く機能しているか確認しましょう。
```bash
$ curl -X GET "http://localhost:8080/fizzbuzz?count=20"
$ curl -X GET "http://localhost:8080/fizzbuzz" # count=30 と同じ
$ curl -X GET "http://localhost:8080/fizzbuzz?count=a" # Bad Request
```

**/fizzbuzzが上手く動いたら、講習会の実況用チャンネルに↑の実行結果を投稿しましょう！**

:::details 解答
<<<@/chapter1/section3/src/4-2_fizzbuzz.go
:::

## 基本問題 POST /add
送信される値を足した数値を返してください。

### スキーマ
#### リクエスト
content-type: `application/json`
```json
{
    "left": 27,
    "right": 57
}
```

#### レスポンス
content-type: `application/json`
status code: `200`
```json
{
    "answer": 84
}
```

##### 適切でないリクエストの場合
content-type: `application/json`
status code: `400`
```json
{
    "error": "Bad Request"
}
```
を返してください。

#### 完成したら、以下のコマンドをターミナルで実行して上手く機能しているか確認しましょう。
```bash
$ curl -X POST "http://localhost:8080/add" -H "Content-Type: application/json" -d '{"left": 18781, "right": 18783}' # 37564
$ curl -X POST "http://localhost:8080/add" -H "Content-Type: application/json" -d '{"left": 0, "right": -0}' # 0
$ curl -X POST "http://localhost:8080/add" -H "Content-Type: application/json" -d '{"left": a, "right": b}' # Bad Request
$ curl -X POST "http://localhost:8080/add" -H "Content-Type: application/json" -d '{"left": 100}' # Bad Request
```
:::details 解答
<<<@/chapter1/section3/src/4-3_add.go
:::

## 発展問題 GET /students/:class/:studentNumber

前提:以下のデータをサーバー側で持っておく。
``` json
[
  {"class_number": 1, "students": [
    {"student_number": 1, "name": "pikachu"},
    {"student_number": 2, "name": "ikura-hamu"},
    {"student_number": 3, "name": "noc7t"}
  ]},
  {"class_number": 2, "students": [
    {"student_number": 1, "name": "Sora"},
    {"student_number": 2, "name": "Kaito"},
    {"student_number": 3, "name": "Haruka"},
    {"student_number": 4, "name": "Shingo"}
  ]},
  {"class_number": 3, "students": [
    {"student_number": 1, "name": "Hikaru"},
    {"student_number": 2, "name": "Eri"},
    {"student_number": 3, "name": "Ryo"}
  ]},
  {"class_number": 4, "students": [
    {"student_number": 1, "name": "Marina"},
    {"student_number": 2, "name": "Takumi"}
  ]}
]
```

classNumber，studentNumber に対応する学生の情報を JSON で返してください。  
学生が存在しない場合、`404`を返してください。

### スキーマ
#### リクエスト
`/students/2/3`

#### レスポンス
content-type: `application/json`
status code: `200`
```json
{
    "student_number": 3,
    "name": "Haruka"
}
```

##### 学生が存在しない場合
content-type: `application/json`
status code: `404`
```json
{
    "error": "Student Not Found"
}
```

:::tip
ヒント: 最後の課題のデータは次のような構造体を用意して、json.Unmarshal すると定義しやすいです。
```go
type Student struct {
	Number int    `json:"student_number"`
	Name   string `json:"name"`
}
type Class struct {
	Number   int       `json:"class_number"`
	Students []Student `json:"students"`
}
```
:::

#### 完成したら、以下のコマンドをターミナルで実行して上手く機能しているか確認しましょう。
```bash
$ curl -X GET "http://localhost:8080/students/1/1" # pikachu
$ curl -X GET "http://localhost:8080/students/3/4" # Student Not Found
```


## 自分のサーバーが正しく動作しているか確認しよう

カレントディレクトリ(`main.go`があるディレクトリ)に、新しく`test.sh`というファイルを作成し、以下をコピー&ペーストしてください。
```bash
#!/bin/bash

# ↓これを自分のIDに変更してください！
ID=pikachu
# ↑これを自分のIDに変更してください！

echo ""
echo "===================="
echo "[TEST] /${ID}"
echo 'curl -X GET http://localhost:8080/'${ID}
curl -X GET "http://localhost:8080/${ID}"
echo ""
echo "===================="
echo "[TEST] /ping"
echo 'curl -X GET http://localhost:8080/ping'
curl -X GET "http://localhost:8080/ping"
echo ""
echo "===================="
echo "[TEST] /fizzbuzz 1of3"
echo '-X GET http://localhost:8080/fizzbuzz?count=20'
curl -X GET "http://localhost:8080/fizzbuzz?count=20"
echo ""
echo "===================="
echo "[TEST] /fizzbuzz 2of3"
echo 'curl -X GET http://localhost:8080/fizzbuzz'
curl -X GET "http://localhost:8080/fizzbuzz"
echo ""
echo "===================="
echo "[TEST] /fizzbuzz 3of3"
echo 'curl -X GET http://localhost:8080/fizzbuzz?count=a'
curl -X GET "http://localhost:8080/fizzbuzz?count=a"
echo ""
echo "===================="
echo "[TEST] /add 1of4"
echo 'curl -X POST http://localhost:8080/add -H "Content-Type: application/json" -d "{\"left\": 18781, \"right\": 18783}"'
curl -X POST "http://localhost:8080/add" -H "Content-Type: application/json" -d '{"left": 18781, "right": 18783}'
echo ""
echo "===================="
echo "[TEST] /add 2of4"
echo 'curl -X POST http://localhost:8080/add -H "Content-Type: application/json" -d "{\"left\": 0, \"right\": -0}"'
curl -X POST "http://localhost:8080/add" -H "Content-Type: application/json" -d '{"left": 0, "right": -0}'
echo ""
echo "===================="
echo "[TEST] /add 3of4"
echo 'curl -X POST http://localhost:8080/add -H "Content-Type: application/json" -d "{\"left\": a, \"right\": b}"'
curl -X POST "http://localhost:8080/add" -H "Content-Type: application/json" -d '{"left": a, "right": b}'
echo ""
echo "===================="
echo "[TEST] /add 4of4"
echo 'curl -X POST http://localhost:8080/add -H "Content-Type: application/json" -d "{\"left\": 100}"'
curl -X POST "http://localhost:8080/add" -H "Content-Type: application/json" -d '{"left": 100}'
echo ""
echo "===================="
echo "[TEST] /students"
echo 'curl -X GET http://localhost:8080/students/1/1'
curl -X GET "http://localhost:8080/students/1/1"
echo ""
echo "===================="
echo "[TEST] /students"
echo 'curl -X GET http://localhost:8080/students/3/4'
curl -X GET "http://localhost:8080/students/3/4"
echo ""
```

ペーストした後、ファイル内の以下の部分を自分の ID に書き換えてください。
```
# ↓これを自分のIDに変更してください！
ID=pikachu
# ↑これを自分のIDに変更してください！
```

最後に、ターミナルを開き、以下を実行してください。
```bash
$ chmod +x test.sh # 実行権限を付与
$ ./test.sh # シェルスクリプトtest.shを実行
```

使用例は以下の通りです。
:::details 使用例
```
$ ./test.sh 

====================
[TEST] /pikachu
curl -X GET http://localhost:8080/pikachu
始めまして、@pikachuです。
ケモノ(特に四足歩行)や、低頭身デフォルメマスコット(TDM)が大好きです。
普段はVRChatに生息しています。twitter: @pikachu0310VRC
====================
[TEST] /ping
curl -X GET http://localhost:8080/ping
pong

====================
[TEST] /fizzbuzz 1of3
-X GET http://localhost:8080/fizzbuzz?count=20
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16
17
Fizz
19
Buzz

====================
[TEST] /fizzbuzz 2of3
curl -X GET http://localhost:8080/fizzbuzz
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16
17
Fizz
19
Buzz
Fizz
22
23
Fizz
Buzz
26
Fizz
28
29
FizzBuzz

====================
[TEST] /fizzbuzz 3of3
curl -X GET http://localhost:8080/fizzbuzz?count=a
Bad Request

====================
[TEST] /add 1of4
curl -X POST http://localhost:8080/add -H "Content-Type: application/json" -d "{\"left\": 18781, \"right\": 18783}"
{"answer":37564}

====================
[TEST] /add 2of4
curl -X POST http://localhost:8080/add -H "Content-Type: application/json" -d "{\"left\": 0, \"right\": -0}"
{"answer":0}

====================
[TEST] /add 3of4
curl -X POST http://localhost:8080/add -H "Content-Type: application/json" -d "{\"left\": a, \"right\": b}"
{"error":"Bad Request"}

====================
[TEST] /add 4of4
curl -X POST http://localhost:8080/add -H "Content-Type: application/json" -d "{\"left\": 100}"
{"error":"Bad Request"}

====================
[TEST] /students 1of2
curl -X GET http://localhost:8080/students/1/1
{"student_number":1,"name":"pikachu"}

====================
[TEST] /students 2of2
curl -X GET http://localhost:8080/students/3/4
{"error":"Student Not Found"}
```
:::

**ここまで出来たら、講習会の実況用チャンネルに`./test.sh`の出力結果を貼りましょう！**  
(発展問題は出来てなくても大丈夫ですが、チャレンジしてみてください。)  
(出力結果は長いので、\`\`\`ここに内容\`\`\`のように内容を\`\`\`で囲い、コードブロックにして送信すると良いです。)

### お疲れ様でした！
