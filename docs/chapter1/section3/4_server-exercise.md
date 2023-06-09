# 演習問題

## 基本問題 GET /ping

pong と返すエンドポイントを作成してください。
(作成済みのはず)

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
`count`がない場合は 30 として，`count`が整数として解釈できない場合はステータスコードとして`400`を返してください。

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
ポストされる値の足し算を返してください。

### スキーマ
#### リクエスト
content-type: `application/json`
```javascript=
{
    "right": Number,
    "left": Number
}
```

#### レスポンス
content-type: `application/json`
status code: `200`
```javascript=
{
    "answer": Number
}
```

##### 適切でないリクエストの場合
content-type: `application/json`
status code: `400`
```javascript=
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
    {"student_number": 1, "name": "hijiki51"},
    {"student_number": 2, "name": "logica"},
    {"student_number": 3, "name": "Ras"}
  ]},
  {"class_number": 2, "students": [
    {"student_number": 1, "name": "asari"},
    {"student_number": 2, "name": "irori"},
    {"student_number": 3, "name": "itt"},
    {"student_number": 4, "name": "mehm8128"}
  ]},
  {"class_number": 3, "students": [
    {"student_number": 1, "name": "reyu"},
    {"student_number": 2, "name": "yukikurage"},
    {"student_number": 3, "name": "anko"}
  ]},
  {"class_number": 4, "students": [
    {"student_number": 1, "name": "Uzaki"},
    {"student_number": 2, "name": "yashu"}
  ]}
]
```

class，studentNumber に対応する学生の情報を JSON で返してください。
学生が存在しない場合，`404`を返してください。

### スキーマ
#### リクエスト
`/students/2/3`

#### レスポンス
content-type: `application/json`
status code: `200`
```javascript=
{
    "student_number": 3,
    "name": "itt"
}
```

##### 学生が存在しない場合
content-type: `application/json`
status code: `404`
```javascript=
{
    "error": "Student Not Found"
}
```

---

自分のサーバーが正しく動作しているか確認したい方は以下の Makefile をご利用ください。

:::details Makefile
```shell=bash
# ----- Makefile by ptr -----
# このファイルをMakefileという名前でカレントディレクトリに保存する。
# 使い方:
#    "make test"で課題をすべてテストする。
#     SSH鍵をサーバーに登録しておくと、"make"でローカルのmain.goをサーバーにコピーできる。

# ！ここを変更する！
# 自分に割り当てられたポート番号
PORT=31337
# 自分のtraQ ID
ID=ptr
# !ここまで!


# ！ここから（分からない人は）変更しない！
all:
	scp main.go naro:~/go/src/hello-server/main.go

test:
	@echo "\n====================\n"
	@echo "[TEST] /$(ID)"
	curl -X GET "http://133.130.109.224:$(PORT)/$(ID)"
	@echo "\n====================\n"
	@echo "[TEST] /ping"
	curl -X GET "http://133.130.109.224:$(PORT)/ping"
	@echo "\n====================\n"
	@echo "[TEST] /fizzbuzz 1of2"
	curl -X GET "http://133.130.109.224:$(PORT)/fizzbuzz?count=20"
	@echo "\n====================\n"
	@echo "[TEST] /fizzbuzz 2of2"
	curl -X GET "http://133.130.109.224:$(PORT)/fizzbuzz"
	@echo "\n====================\n"
	@echo "[TEST] /add"
	curl -X POST "http://133.130.109.224:$(PORT)/add" -d "left=18781&right=18783"
	@echo "\n====================\n"
	@echo "[TEST] /students"
	curl -X GET "http://133.130.109.224:$(PORT)/students/3/1"
```

使用例は以下の通りです。
```
$ make test

====================

[TEST] /ptr
curl -X GET "http://133.130.109.224:31337/ptr"
id: ptr
twitter: @ptrYudai
pwn wannabe
====================

[TEST] /ping
curl -X GET "http://133.130.109.224:31337/ping"
pong
====================

[TEST] /fizzbuzz 1of2
curl -X GET "http://133.130.109.224:31337/fizzbuzz?count=20"
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
curl -X GET "http://133.130.109.224:31337/fizzbuzz"
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
curl -X POST "http://133.130.109.224:31337/add" -d "left=18781&right=18783"
37564
====================

[TEST] /students
curl -X GET "http://133.130.109.224:31337/students/3/1"
{"student_number":1,"name":"reyu"}
```
:::

---

最後の課題のデータは次のような構造体を用意して、json.Unmarshal すると定義しやすいかも？
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