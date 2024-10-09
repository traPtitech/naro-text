---
prev: 
  text: 'はじめに'
  link: '/chapter1/index'
next:
  text: '環境構築'
  link: '/chapter1/section1/0_setup-windows'
---


# きれいなコードを書くために

この章では、きれいなコードを書くためにはどうすればよいか、具体的な例とともに説明します。

[[toc]]

## きれいなコードとは何か
そもそもきれいなコードとは何でしょうか。
きれいなコードとは、「他の人が短時間で理解できる」コードです。

ここでいう「他の人」とは、自分以外の人だけではなく、将来の自分も含みます。
書いた直後は何を目的にして書いたか覚えていられます。しかし、1 か月後、1 年後に見直すとなった時に覚えていられるでしょうか。
もはや初めて見るコードと大差ないでしょう。

これから SysAd で扱っていくコードは、長い間様々な人の手によってメンテナンスされます。
そこには当然自分が書いたコードを別の人がメンテナンスすることもあるし、逆に他の人が書いたコードを自分がメンテナンスすることもあります。

そんな時、次のようなコードが出てきたらあなたはすっと理解できるでしょうか。

```rs
fn main() {
    let a = 4; // X1-X3
    let b = 7;
    let c = 2;
    let d = 5; // Y1-Y3
    let e = 1;
    let f = 9;
    let mut tmp_var = 252521.0;
    if tmp_var > (((0-a)*(0-a)+(0-d)*(0-d)) as f64).sqrt() {
		tmp_var = (((a-0)*(a-0) + (d-0)*(d-0)) as f64).sqrt();
	} 
	if tmp_var > (((0-b)*(0-b)+(0-e)*(0-e)) as f64).sqrt() {
	tmp_var = (((-b)*(-b) + (-e)*(-e)) as f64).sqrt();
	}
	if tmp_var > (((0-c)*(0-c)+(0-f)*(0-f)) as f64).sqrt() { 
	println!("{}", ((((c)*(c) + (f)*(f) )) as f64).sqrt());} else { 
	println!("{}", tmp_var); }}

```

まず`tmp_var`は何を指しているのだろうかという疑問が湧くでしょう。
他にも`if`文の条件で何が制限されているか、`252521.0`という数字が何を意味するか、などと疑問が尽きず、理解に時間がかかります。

では、次のコードはどうでしょうか。

<<< @/chapter1/dicts/clean-code/src/example.rs

このコードであれば、原点からの距離を計算して、最も近い距離を出力するということがすぐに理解できます。

では、後者のコードを書けるようになるためにはどのようなことに気を付けるべきでしょうか。この章では、そのための具体的な方法を説明します。

## フォーマットをそろえよう

まず、インデントがズレていたり、括弧の位置がそろっていなかったりするといま自分がどの関数にいるのか、どの`if`文の中にいるのかがわかりにくくなります。
そのため、フォーマットをそろえることが重要です。

基本どの言語にもフォーマッタというものがあります。これは、コードを自動でフォーマットしてくれるものです。
例えば、次のようなものがあげられます。
- `rustfmt` (rust)
- `gofmt`(go)
- `prettier`(JavaScript)
- `black`(Python)

さっきのコードにフォーマッタをかけてみましょう。

```rs

fn main() {
    let a = 4; // X1-X3
    let b = 7;
    let c = 2;
    let d = 5; // Y1-Y3
    let e = 1;
    let f = 9;
    let mut tmp_var = 252521.0;
    if tmp_var > (((0 - a) * (0 - a) + (0 - d) * (0 - d)) as f64).sqrt() {
        tmp_var = (((a - 0) * (a - 0) + (d - 0) * (d - 0)) as f64).sqrt();
    }
    if tmp_var > (((0 - b) * (0 - b) + (0 - e) * (0 - e)) as f64).sqrt() {
        tmp_var = (((-b) * (-b) + (-e) * (-e)) as f64).sqrt();
    }
    if tmp_var > (((0 - c) * (0 - c) + (0 - f) * (0 - f)) as f64).sqrt() {
        println!("{}", (((c) * (c) + (f) * (f)) as f64).sqrt());
    } else {
        println!("{}", tmp_var);
    }
}
```
これだけでもだいぶ見やすくなりましたね。

## 意味のある名前を付けよう

前者の汚いコードが出てきたとき、`tmp_var`って何だろうという疑問が出てきました。これは、`tmp_var`から分かる情報が、「一時変数」という情報しかないからです。

実際には、`tmp_var`は「原点からの距離の最小値」を表しています。であるならば、そのような内容を表す名前をつけるべきです。
先ほどのコードに意味のある名前を付けてみましょう。
```rs
fn main() {
    let a = 4; // X1-X3
    let b = 7;
    let c = 2;
    let d = 5; // Y1-Y3
    let e = 1;
    let f = 9;
    let mut min_distance = 252521.0;
    if min_distance > (((0 - a) * (0 - a) + (0 - d) * (0 - d)) as f64).sqrt() {
        min_distance = (((a - 0) * (a - 0) + (d - 0) * (d - 0)) as f64).sqrt();
    }
    if min_distance > (((0 - b) * (0 - b) + (0 - e) * (0 - e)) as f64).sqrt() {
        min_distance = (((-b) * (-b) + (-e) * (-e)) as f64).sqrt();
    }
    if min_distance > (((0 - c) * (0 - c) + (0 - f) * (0 - f)) as f64).sqrt() {
        println!("{}", (((c) * (c) + (f) * (f)) as f64).sqrt());
    } else {
        println!("{}", min_distance);
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
- 「変数」を表すものは、先頭を小文字にして、単語のつなぎ目はアンダースコアにする

このように定めることで、`TEST_NAME`は定数、`test_name`は関数や変数ということがわかります。

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
```rs
fn main() {
    let a = 4; // X1-X3
    let b = 7;
    let c = 2;
    let d = 5; // Y1-Y3
    let e = 1;
    let f = 9;
    let mut min_distance = 252521.0;
    if min_distance > (((0 - a) * (0 - a) + (0 - d) * (0 - d)) as f64).sqrt() {
        min_distance = (((a - 0) * (a - 0) + (d - 0) * (d - 0)) as f64).sqrt();
    }
    if min_distance > (((0 - b) * (0 - b) + (0 - e) * (0 - e)) as f64).sqrt() {
        min_distance = (((-b) * (-b) + (-e) * (-e)) as f64).sqrt();
    }
    if min_distance > (((0 - c) * (0 - c) + (0 - f) * (0 - f)) as f64).sqrt() {
        println!("{}", (((c) * (c) + (f) * (f)) as f64).sqrt());
    } else {
        println!("{}", min_distance);
    }
}
```
まだまだ読みにくいですね。ここで、`if`文の中に距離を求める処理が何度も出てきていることがわかります。これを関数として切り出してみましょう。

```rs
fn distance(x1: i32, y1: i32, x2: i32, y2: i32) -> f64 {
    let dx = (x1 - x2) as f64;
    let dy = (y1 - y2) as f64;
    (dx * dx + dy * dy).sqrt()
}

fn main() {
    let a = 4; // X1-X3
    let b = 7;
    let c = 2;
    let d = 5; // Y1-Y3
    let e = 1;
    let f = 9;
    let mut min_distance = 252521.0;
    if min_distance > distance(0, 0, a, d) {
        min_distance = distance(0, 0, a, d);
    }
    if min_distance > distance(0, 0, b, e) {
        min_distance = distance(0, 0, b, e);
    }
    if min_distance > distance(0, 0, c, f) {
        println!("{}", distance(0, 0, c, f));
    } else {
        println!("{}", min_distance);
    }
}

```

計算部分がなくなって少しわかりやすくなったでしょうか。しかし、今度は`distance()`の引数が何を表しているのか使う側からだと分かりにくいという問題があります。`distance(x1,x2,y1,y2)`なのか、`distance(x1,y1,x2,y2)`なのか見た目からは分かりません。
そこで、「座標」としてまとめてしまいましょう。

```rs
#[derive(Clone, Copy)]
struct Point {
    x: i32,
    y: i32,
}

fn distance(a: Point, b: Point) -> f64 {
    let dx = (a.x - b.x) as f64;
    let dy = (a.y - b.y) as f64;
    (dx * dx + dy * dy).sqrt()
}

fn main() {
    let points = [
        Point { x: 4, y: 5 },
        Point { x: 7, y: 1 },
        Point { x: 2, y: 9 },
    ];
    let mut min_distance = 252521.0;
    let origin = Point { x: 0, y: 0 };
    if min_distance > distance(origin, points[0]) {
        min_distance = distance(origin, points[0]);
    }
    if min_distance > distance(origin, points[1]) {
        min_distance = distance(origin, points[1]);
    }
    if min_distance > distance(origin, points[2]) {
        println!("{}", distance(origin, points[2]));
    } else {
        println!("{}", min_distance);
    }
}

```
なんか同じ計算をしているのが見えてきましたね。`for`文でまとめちゃいましょう。

```rs
#[derive(Clone, Copy)]
struct Point {
    x: i32,
    y: i32,
}

fn distance(a: Point, b: Point) -> f64 {
    let dx = (a.x - b.x) as f64;
    let dy = (a.y - b.y) as f64;
    (dx * dx + dy * dy).sqrt()
}

fn main() {
    let points = [
        Point { x: 4, y: 5 },
        Point { x: 7, y: 1 },
        Point { x: 2, y: 9 },
    ];
    let mut min_distance = 252521.0;
    let origin = Point { x: 0, y: 0 };
    for &p in &points {
        if min_distance > distance(origin, p) {
            min_distance = distance(origin, p);
        }
    }
    println!("{}", min_distance);
}
```

## 数字の意味を明確にしよう

最後に`252521.0`とは何の数字でしょうか。今回の処理を考えると、`minDistance`と比較したときに、最初は必ず値を更新する必要があります。なので、十分大きな値を設定したのでしょう。
しかし、それであれば、`f64`の最大値を設定するのがいいでしょう。
Rust では、``f64::MAX`` や ``f64::INFINITY`` といった定数が用意されているため、これらのいずれかを用いましょう。
```rs
#[derive(Clone, Copy)]
struct Point {
    x: i32,
    y: i32,
}

fn distance(a: Point, b: Point) -> f64 {
    let dx = (a.x - b.x) as f64;
    let dy = (a.y - b.y) as f64;
    (dx * dx + dy * dy).sqrt()
}

fn main() {
    let points = [
        Point { x: 4, y: 5 },
        Point { x: 7, y: 1 },
        Point { x: 2, y: 9 },
    ];
    let mut min_distance = f64::INFINITY;
    let origin = Point { x: 0, y: 0 };
    for &p in &points {
        if min_distance > distance(origin, p) {
            min_distance = distance(origin, p);
        }
    }
    println!("{}", min_distance);
}
```

最初のコードと比べると圧倒的に読みやすくなりましたね。

## コメントをつけよう
ここまで、がんばってコードをきれいにしてきました。しかし、それでも理解するのが難しい部分はあります。例えば、数学的に難解な暗号処理をしていたり、複雑なアルゴリズムを使うときです。
そういう時は、必ずコメントをつけましょう。
```rs
#[derive(Clone, Copy)]
struct Point {
    x: i32,
    y: i32,
}

fn distance(a: Point, b: Point) -> f64 {
    let dx = (a.x - b.x) as f64;
    let dy = (a.y - b.y) as f64;
    (dx * dx + dy * dy).sqrt()
}

fn main() {
    let points = [
        Point { x: 4, y: 5 },
        Point { x: 7, y: 1 },
        Point { x: 2, y: 9 },
    ];
    let mut min_distance = f64::INFINITY;
    let origin = Point { x: 0, y: 0 };
    // 入力した点のうち最も原点に近い点を探し、その距離を求める
    for &p in &points {
        if min_distance > distance(origin, p) {
            min_distance = distance(origin, p);
        }
    }
    println!("{}", min_distance);
}
```

## おまけ
Rust のイテレータを使用すると、``min_distance`` を求める操作をメソッドチェイン(メソッドの連鎖)で書くことができます。

```rs
#[derive(Clone, Copy)]
struct Point {
    x: i32,
    y: i32,
}

fn distance(a: Point, b: Point) -> f64 {
    let dx = (a.x - b.x) as f64;
    let dy = (a.y - b.y) as f64;
    (dx * dx + dy * dy).sqrt()
}

fn main() {
    let points = [
        Point { x: 4, y: 5 },
        Point { x: 7, y: 1 },
        Point { x: 2, y: 9 },
    ];
    let origin = Point { x: 0, y: 0 };
    // 入力した点のうち最も原点に近い点を探し、その距離を求める
    let min_distance = points
        .iter()
        .map(|&p| distance(origin, p))
        .fold(f64::INFINITY, f64::min);
    println!("{}", min_distance);
}
```

## おわりに

ここで述べたルールはほんの一例です。実際は言語ごとにルールが存在したり、プロジェクトごとにルールが存在したりします。参考文献には、きれいなコードを書くためのエッセンスが詰まっています。ぜひ読んでみてください。

## 参考文献
[リーダブルコード](https://www.amazon.co.jp/%E3%83%AA%E3%83%BC%E3%83%80%E3%83%96%E3%83%AB%E3%82%B3%E3%83%BC%E3%83%89-%E2%80%95%E3%82%88%E3%82%8A%E8%89%AF%E3%81%84%E3%82%B3%E3%83%BC%E3%83%89%E3%82%92%E6%9B%B8%E3%81%8F%E3%81%9F%E3%82%81%E3%81%AE%E3%82%B7%E3%83%B3%E3%83%97%E3%83%AB%E3%81%A7%E5%AE%9F%E8%B7%B5%E7%9A%84%E3%81%AA%E3%83%86%E3%82%AF%E3%83%8B%E3%83%83%E3%82%AF-Theory-practice-Boswell/dp/4873115655)
