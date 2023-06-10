# 演習問題

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
    "answer": Number
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

---

自分のサーバーが正しく動作しているか確認したい方は以下の Makefile をご利用ください。

:::details Makefile
```bash
# ----- Makefile by ptr -----
# このファイルをMakefileという名前でカレントディレクトリに保存する。
# 使い方:
#    "make test"で課題をすべてテストする。
#     SSH鍵をサーバーに登録しておくと、"make"でローカルのmain.goをサーバーにコピーできる。
PORT=8080
# ！ここを変更する！
# 自分のtraQ ID
ID=ptr
# !ここまで!


# ！ここから（分からない人は）変更しない！
all:
	scp main.go naro:~/go/src/hello-server/main.go

test:
	@echo "\n====================\n"
	@echo "[TEST] /$(ID)"
	curl -X GET "http://localhost:$(PORT)/$(ID)"
	@echo "\n====================\n"
	@echo "[TEST] /ping"
	curl -X GET "http://localhost:$(PORT)/ping"
	@echo "\n====================\n"
	@echo "[TEST] /fizzbuzz 1of2"
	curl -X GET "http://localhost:$(PORT)/fizzbuzz?count=20"
	@echo "\n====================\n"
	@echo "[TEST] /fizzbuzz 2of2"
	curl -X GET "http://localhost:$(PORT)/fizzbuzz"
	@echo "\n====================\n"
	@echo "[TEST] /add"
	curl -X POST "http://localhost:$(PORT)/add" -d "left=18781&right=18783"
	@echo "\n====================\n"
	@echo "[TEST] /students"
	curl -X GET "http://localhost:$(PORT)/students/3/1"
```

使用例は以下の通りです。
```
$ make test

====================

[TEST] /ptr
curl -X GET "http://localhost:8080/ptr"
id: ptr
twitter: @ptrYudai
pwn wannabe
====================

[TEST] /ping
curl -X GET "http://localhost:8080/ping"
pong
====================

[TEST] /fizzbuzz 1of2
curl -X GET "http://localhost:8080/fizzbuzz?count=20"
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

[TEST] /fizzbuzz 2of2
curl -X GET "http://localhost:8080/fizzbuzz"
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

[TEST] /add
curl -X POST "http://localhost:8080/add" -d "left=18781&right=18783"
37564
====================

[TEST] /students
curl -X GET "http://localhost:8080/students/3/1"
{"student_number":1,"name":"Hikaru"}
```
:::

---

:::tip
最後の課題のデータは次のような構造体を用意して、json.Unmarshal すると定義しやすいです。
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
