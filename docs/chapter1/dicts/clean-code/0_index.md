# きれいなコードを書くために

この章では、きれいなコードを書くためにはどうすればよいか、具体的な例とともに説明します。

[[toc]]

## きれいなコードとは何か
そもそもきれいなコードとは何でしょうか。
きれいなコードとは、「他の人が短時間で理解できる」コードです。

ここでいう「他の人」とは、自分以外の人だけではなく、将来の自分も含みます。
例えば、これから SysAd で扱っていくコードは、長い間様々な人の手によってメンテナンスされます。
そこには当然自分が書いたコードを別の人がメンテナンスすることもあるし、逆に他の人が書いたコードを自分がメンテナンスすることもあるでしょう。

そんな時、次のようなコードが出てきたらあなたはすっと理解できるでしょうか。

```go
var (	
	a = 4 // X1-X3
	b = 7
	c = 2
	d = 5 // Y1-Y3
	e = 1
	f = 9
) 
func main() {
	tmp_var := 252521.0
	if tmp_var < math.Sqrt(float64((0-a)*(0-a)+(0-d)*(0-d))) {
		tmp_var = math.Sqrt(float64((a-0)*(a-0) + (d-0)*(d-0)))
	} 
	if tmp_var > math.Sqrt(float64((0-b)*(0-b)+(0-e)*(0-e))) {
	tmp_var = math.Sqrt(float64((-b)*(-b) + (-e)*(-e)))}
	}
	if tmp_var < math.Sqrt(float64((0-c)*(0-c)+(0-f)*(0-f))) { 
	fmt.Println(math.Sqrt(float64((c)*(c) + (f)*(f))))} else { 
	fmt.Println(tmp_var) }} 
```

まず`tmp_var`は何を指しているのだろうかという疑問が湧くでしょう。
他にも`if`文の条件で何が制限されているか、`252521.0`という数字が何を意味するか、などと疑問が尽きず、理解に時間がかかります。

では、次のコードはどうでしょうか。

<<< @/chapter1/dicts/clean-code/src/example.go

このコードであれば、原点からの距離を計算して、最も近い距離を出力するということがすぐに理解できます。

では、後者のコードを書けるようになるためにはどのようなことに気を付けるべきでしょうか。この章では、そのための具体的な方法を説明します。

## フォーマットをそろえよう

まず、インデントがズレていたり、括弧の位置がそろっていなかったりするといま自分がどの関数にいるのか、どの`if`文の中にいるのかがわかりにくくなります。
そのため、フォーマットをそろえることが重要です。

基本どの言語にもフォーマッタというものがあります。これは、コードを自動でフォーマットしてくれるものです。
さっきのコードにフォーマッタをかけてみましょう。

```go

var (
	a = 4 // X1-X3
	b = 7
	c = 2
	d = 5 // Y1-Y3
	e = 1
	f = 9
)

func main() {
	tmp_var := 252521.0
	if tmp_var < math.Sqrt(float64((0-a)*(0-a)+(0-d)*(0-d))) {
		tmp_var = math.Sqrt(float64((a-0)*(a-0) + (d-0)*(d-0)))
	}
	if tmp_var > math.Sqrt(float64((0-b)*(0-b)+(0-e)*(0-e))) {
		tmp_var = math.Sqrt(float64((-b)*(-b) + (-e)*(-e)))
	}
	if tmp_var < math.Sqrt(float64((0-c)*(0-c)+(0-f)*(0-f))) {
		fmt.Println(math.Sqrt(float64((c)*(c) + (f)*(f))))
	} else {
		fmt.Println(tmp_var)
	}
}
```
これだけでもだいぶ見やすくなりましたね。

## 意味のある名前を付けよう

前者の汚いコードが出てきたとき、`tmp_var`って何だろうという疑問が出てきました。これは、`tmp_var`から分かる情報が、「一時変数」という情報しかないからです。

実際には、`tmp_var`は「原点からの距離の最小値」を表しています。であるならば、そのような内容を表す名前をつけるべきです。
先ほどのコードに意味のある名前を付けてみましょう。
```go
var (	
	a = 4 // X1-X3
	b = 7
	c = 2
	d = 5 // Y1-Y3
	e = 1
	f = 9
) 
func main() {
	minDistance := 252521.0
	if minDistance < math.Sqrt(float64((0-a)*(0-a)+(0-d)*(0-d))) {
		minDistance = math.Sqrt(float64((a-0)*(a-0) + (d-0)*(d-0)))
	}
	if minDistance > math.Sqrt(float64((0-b)*(0-b)+(0-e)*(0-e))) {
		minDistance = math.Sqrt(float64((-b)*(-b) + (-e)*(-e)))
	}
	if minDistance < math.Sqrt(float64((0-c)*(0-c)+(0-f)*(0-f))) {
		fmt.Println(math.Sqrt(float64((c)*(c) + (f)*(f))))
	} else {
		fmt.Println(minDistance)
	}
}
```

変数が何を指しているか少しわかりやすくなりましたね。

これだけを見ると、「具体的な単語をつければいいんでしょ、簡単じゃん」と感じることもあるでしょう。しかし、実際にはそう簡単ではなく、プログラマが日々頭を悩ませる重要な問題の 1 つです。
なぜなら、名前は変数以外にも関数、型等にもつける必要があり、それらの名前が適切な意味を持ちつつ、重複しないようにすることは容易でないからです。

そこで、名前をつけるときには、次のようなことを意識すると良いでしょう。
### 命名規則に従おう
コードの質が管理されたプロジェクトには、「命名規則」というものが存在します。

例えば、次のように定めることが出来ます。
- 「型」を表すものは、先頭を大文字にして、単語のつなぎ目も大文字にする
- 「定数」を表すものは、すべて大文字にして、単語のつなぎ目はアンダースコアにする
- 「関数」を表すものは、先頭を小文字にして、単語のつなぎ目はアンダースコアにする
- 「変数」を表すものは、先頭を小文字にして、単語のつなぎ目は大文字にする

このように定めることで、`TEST_NAME`は定数、`test_name`は関数、`testName`は変数ということがわかります。
今、「先頭を大文字にして、単語のつなぎ目も大文字にする」という形式が出てきましたが、それには名前が存在して、以下の表のようになっています。

| 名前 | 形式 | 例 |
| --- | --- | --- |
| キャメルケース | 先頭を小文字にして、単語のつなぎ目を大文字にする | `testName` |
| パスカルケース | 先頭を大文字にして、単語のつなぎ目を大文字にする | `TestName` |
| スネークケース | 単語のつなぎ目をアンダースコアにする | `test_name` |
| ケバブケース | 単語のつなぎ目をハイフンにする | `test-name` |

このほかにも、プライベートメンバーは`_`で終わりにするなど、プロジェクトごとに決まっていることがあります。

### 明確な単語を選ぼう
名前をつけるときには、出来るだけより明確な単語を選ぶようにしましょう。

例えば、つぎのような`Movie`というクラスがあったとします。
```cpp
class Movie {
	Stop();
}
```

`Stop()`というメソッド名は特段悪くないです。しかし、後から再開（Resume）出来るのであれば、`Pause()`の方が適切でしょう。逆に出来ないのであれば、`Terminate()`などより強い言葉を使ってもよいでしょう。

::: tip
**単語の探し方**

とはいえど、このように明確な単語を選ぶには語彙力が要求されます。そこで、いくつかの方法を紹介します。

- `{単語} synonyms`で検索する
  - `synonyms`は類義語という意味です。ペアとなる単語を調べたいときは`{単語} antonyms`とするとよいでしょう。
- 変数名を考えてくれるサイトを利用する
	- [codic](https://codic.jp/engine)等が有名
::: 

## 適切に分割・再利用しよう

1 ファイル、1 関数内に書かれるコードが長くなれば長くなるほど、そのコードを理解するのは難しくなります。
また、同じ処理を行うコードをコピー&ペーストしてしまうと、修正の際に複数の箇所を修正する必要が出てきます。

それを避けるために、適切なサイズで関数として切り出したりして、コードを分割、共通化することが重要です。

先ほどのコードをもう一度見てみましょう。
```go
var (	
	a = 4 // X1-X3
	b = 7
	c = 2
	d = 5 // Y1-Y3
	e = 1
	f = 9
) 
func main() {
	minDistance := 252521.0
	if minDistance < math.Sqrt(float64((0-a)*(0-a)+(0-d)*(0-d))) {
		minDistance = math.Sqrt(float64((a-0)*(a-0) + (d-0)*(d-0)))
	}
	if minDistance > math.Sqrt(float64((0-b)*(0-b)+(0-e)*(0-e))) {
		minDistance = math.Sqrt(float64((-b)*(-b) + (-e)*(-e)))
	}
	if minDistance < math.Sqrt(float64((0-c)*(0-c)+(0-f)*(0-f))) {
		fmt.Println(math.Sqrt(float64((c)*(c) + (f)*(f))))
	} else {
		fmt.Println(minDistance)
	}
}
```
まだまだ読みにくいですね。ここで、`if`文の中に距離を求める処理が何度も出てきていることがわかります。これを関数として切り出してみましょう。

```go
var (	
	a = 4 // X1-X3
	b = 7
	c = 2
	d = 5 // Y1-Y3
	e = 1
	f = 9
) 
func distance(x1, y1, x2, y2 int) float64 {
	dx := float64(x1 - x2)
	dy := float64(y1 - y2)
	return math.Sqrt(dx*dx + dy*dy)
}
func main() {
	minDistance := 252521.0
	if minDistance < distance(0,0,a,d) {
		minDistance = distance(0,0,a,d)
	} 
	if minDistance > distance(0,0,b,e) {
		minDistance = distance(0,0,b,e)
	}
	if minDistance < distance(0,0,c,f) {
		fmt.Println(distance(0,0,c,f))
	} else {
		fmt.Println(minDistance)
	}
```

計算部分がなくなって少しわかりやすくなったでしょうか。しかし、今度は`distance()`の引数が何を表しているのか使う側からだと分かりにくいという問題があります。`distance(x1,x2,y1,y2)`なのか、`distance(x1,y1,x2,y2)`なのか見た目からは分かりません。
そこで、「座標」としてまとめてしまいましょう。

```go

type Point struct {
	x int
	y int
}
var points = []Point{{4, 5}, {7, 1}, {2, 9}}
func distance(a, b Point) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	return math.Sqrt(dx*dx + dy*dy)
}
func main() {
	minDistance := 252521.0
	origin := Point{0, 0}
	if minDistance < distance(origin, points[0]) {
		minDistance = distance(origin, points[0])
	} 
	if minDistance > distance(origin, points[1]) {
		minDistance = distance(origin, points[1])
	}
	if minDistance < distance(origin, points[2]) {
		fmt.Println(distance(origin, points[2]))
	} else {
		fmt.Println(minDistance)
	}
```
なんか同じ計算をしているのが見えてきましたね。`for`文でまとめちゃいましょう。

```go
type Point struct {
	x int
	y int
}
var points = []Point{{4, 5}, {7, 1}, {2, 9}}
func distance(a, b Point) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	return math.Sqrt(dx*dx + dy*dy)
}
func main() {
	minDistance := 252521.0
	origin := Point{0, 0}
	for _, p := range points {
		if minDistance < distance(origin, p) {
			minDistance = distance(origin, p)
		}
	}
	fmt.Println(minDistance)
}
```

## 名前を付けよう

最後に`252521.0`とは何の数字でしょうか。今回の処理を考えると、`minDistance`と比較したときに、最初は必ず値を更新する必要があります。なので、十分大きな値を設定したのでしょう。
しかし、それであれば、`float64`の最大値を設定するのがいいでしょう。
```go
type Point struct {
	x int
	y int
}
var points = []Point{{4, 5}, {7, 1}, {2, 9}}
func distance(a, b Point) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	return math.Sqrt(dx*dx + dy*dy)
}
func main() {
	minDistance := 1.79769313486231570814527423731704356798070e+308
	origin := Point{0, 0}
	for _, p := range points {
		if minDistance < distance(origin, p) {
			minDistance = distance(origin, p)
		}
	}
	fmt.Println(minDistance)
}
```

さて、数値を置き換えましたが、結局この数字が何か一目でわからないという問題は解決していません。そこで、この数字に名前を付けましょう。
:::info
このように、名前がついていなくて、意味が分かりにくい数字を「マジックナンバー」と呼びます。
:::

```go
type Point struct {
	x int
	y int
}
var points = []Point{{4, 5}, {7, 1}, {2, 9}}
const maxFloat64 = 1.79769313486231570814527423731704356798070e+308
func distance(a, b Point) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	return math.Sqrt(dx*dx + dy*dy)
}
func main() {
	minDistance := maxFloat64
	origin := Point{0, 0}
	for _, p := range points {
		if minDistance < distance(origin, p) {
			minDistance = distance(origin, p)
		}
	}
	fmt.Println(minDistance)
}
```
最初のコードと比べると圧倒的に読みやすくなりましたね。

## コメントをつけよう
ここまで、がんばってコードをきれいにしてきました。しかし、それでも理解するのが難しい部分はあります。例えば、数学的に難解な暗号処理をしていたり、複雑なアルゴリズムを使うときです。
そういう時は、必ずコメントをつけましょう。
```go
type Point struct {
	x int
	y int
}
var points = []Point{{4, 5}, {7, 1}, {2, 9}}
const maxFloat64 = 1.79769313486231570814527423731704356798070e+308
func distance(a, b Point) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	return math.Sqrt(dx*dx + dy*dy)
}
func main() {
	minDistance := maxFloat64
	origin := Point{0, 0}
	// 入力した点のうち最も原点に近い点を探し、その距離を求める
	for _, p := range points {
		if minDistance < distance(origin, p) {
			minDistance = distance(origin, p)
		}
	}
	fmt.Println(minDistance)
}
```


## おわりに

ここで述べたルールはほんの一例です。実際は言語ごとにルールが存在したり、プロジェクトごとにルールが存在したりします。参考文献には、きれいなコードを書くためのエッセンスが詰まっています。ぜひ読んでみてください。

## 参考文献
[リーダブルコード](https://www.amazon.co.jp/%E3%83%AA%E3%83%BC%E3%83%80%E3%83%96%E3%83%AB%E3%82%B3%E3%83%BC%E3%83%89-%E2%80%95%E3%82%88%E3%82%8A%E8%89%AF%E3%81%84%E3%82%B3%E3%83%BC%E3%83%89%E3%82%92%E6%9B%B8%E3%81%8F%E3%81%9F%E3%82%81%E3%81%AE%E3%82%B7%E3%83%B3%E3%83%97%E3%83%AB%E3%81%A7%E5%AE%9F%E8%B7%B5%E7%9A%84%E3%81%AA%E3%83%86%E3%82%AF%E3%83%8B%E3%83%83%E3%82%AF-Theory-practice-Boswell/dp/4873115655)