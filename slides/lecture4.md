---
marp: true
theme: SysAd
---

<!--
_class: title
-->

# データベース入門

Webエンジニアになろう講習会 第6回

---

# 自己紹介

<div class="columns">
  <div>
    <img src="assets/lecture4/icon.png"/>
  </div>
  <div>
    <h2>renkon</h2>
    <div>情報通信系</div>
    <div>おすすめのエディター募集中です</div>
    <div>インフラやってます</div>
  </div>
</div>

---

<!--
_class: section-head
-->
# 前回のおさらい

---

# HTTPって何だったっけ？

- <span class="underlined">お願い (リクエスト)</span> と <span class="underlined">お返事 (レスポンス)</span> で<br>やりとりをする仕組み
- リソースを取得・編集・削除するやりとりができる

---

<!--
_class: section-head
-->
# リクエスト

これください！

---

# リクエストの構造

<div class="center">
  <img src="assets/lecture4/http-request.svg" width="80%"/>
</div>

---

<!--
_class: section-head
-->
# レスポンス

リクエストに対する返事

---

# レスポンスの構造

<div class="center">
  <img src="assets/lecture4/http-response.svg" width="70%"/>
</div>

---

<!--
_class: section-head
-->
# データの伝送
HTTPが依存するプロトコル

---

# TCP

HTTPが依存する通信プロトコル．
  - 通信が確実に/正しい順序で届くことを保証する．
  - あらゆるデジタルなデータの伝送をすることができる．
    - HTTPのリクエストやメールの送信，コンピュータの遠隔操作(ssh)まで
- 通信をパケット(小包)に分けて送信する

---

# IP (Internet Protocol)

- 世界中に何億台も機械があるが、そのうち届いてほしい機械にデータを届けたい
- 機械に番号を振る
- ものすごく良い感じに機械に番号を振ることで、機械同士が直接つながっていなくてもデータが届くようになっている

---

# 目次

- 座学
  - データベース ◀️
  - 環境変数
- 実習
  - SQLで遊ぶ
  - Goでデータベースを扱う
  - サーバーからデータベースを扱う
---

<!--
_class: section-head
-->
# データベース

---

# Webサービス概観

<div class="center">
  <img src="assets/lecture4/traq-concept-3.png" width="100%"/>
</div>

---

# Webサービス概観

<div class="center">
  <img src="assets/lecture4/traq-concept-4.png" width="100%"/>
</div>

---

# DB  (DataBase) とは

- DataBase → データ基地 → データを蓄積し、管理する場所
- 例: Excelファイル
  - 大量のデータをまとめられる
  - 効率よく検索、追加、削除ができる
- <span class="underlined">SQL</span>と呼ばれる言語がよく使われる

---

# DBMS  (**D**ata**B**ase **M**anagement **S**ystem)

- データベースの情報管理を行うアプリ
- <span class="underlined">データの整合性</span>を保つ役割
- 整合性を保つ: **ACID特性**
  - **A**tomicity（原子性）、**C**onsistency（整合性）、**I**solation（独立性）、**D**urability（永続性）
  - 関係が壊れたりデータが勝手に無くなったりしない

<div class="columns-3">
  <img src="assets/lecture4/mysql.png"/>
  <img src="assets/lecture4/postgresql.png"/>
  <img src="assets/lecture4/mongodb.png"/>
</div>
<div>

---

# RDBMS（Relational DBMS）（1/2）

- 表形式でデータを格納 （テーブル）
- 管理する情報の種類を列（カラム）にする
- 1つのデータを1行（レコード）で管理
- **ACID**特性を持つ

<div class="center">
  <img src="assets/lecture4/table.png"/>
</div>

---

# RDBMS（2/2）

- テーブル同士に**関係**（**Relation**）を作れる
- 例: 特定のカラムにあるテーブルから制限をかける
  - 部員テーブルの“所属班”は“班”テーブルの”班名”しか挿入できない

<div class="center">
  <img src="assets/lecture4/relation.png" width="80%"/>
</div>


---

# SQL

**S**tructured **Q**uery **L**anguage

例: `member` テーブルから<br>`Id`, `Name`, `Team` の3カラムを選択する

```sql
SELECT Id, Name, Team FROM member;
```



---

# NoSQL

- Not Only SQLの略
  - RDBMS以外のDBMS
- 高パフォーマンス、サーバー分散
- SQLを<span class="underlined">使えない</span>
- 整合性が<span class="underlined">保たれない・弱い</span>場合も

---

# 様々なDBMS

![](assets/lecture4/db.png)

---

# 目次

- 座学
  - データベース
  - 環境変数　◀️
- 実習
  - SQLで遊ぶ
  - Goでデータベースを扱う
  - サーバーからデータベースを扱う

---

# 環境変数って何？ (1/2)

- **定義**: プロセスに紐づく名前付きの設定（Key=Valueの形で書く）
- **特徴**:
  - アプリは通常「ひとつの実行プロセス」として動く。環境変数はそのプロセスの設定であり、プロセスごとに値が決まる
  - 実行時に読み取って挙動を変える（起動前に設定を用意する）

---
# 環境変数って何？ (2/2)

- 例えば
  - **環境によって変わる設定**: DBのホスト/ポート、サービスのURL
  - **機密情報**: DBのユーザー名・パスワード、APIキー、シークレット

<div class="center">
  <img src="assets/lecture4/password.png" width="80%"/>
</div>

---

# 注意点・運用の基本

- どこに書く？
  - ❌️ ソースコード
    - Gitでコミットしたら漏洩
    - 手元のパソコンとサーバーで設定が異なる
  - ⭕️ 環境変数（よく`.env`ファイルに記述）
    - `.gitignore`で、認証情報を含むファイルのみGitから除外
    - 環境ごとに設定を切り替えられる

---

# Webエンジニアになるにあたって

<div class="center">
  <img src="assets/lecture4/web-engineer-growth.png" width="90%"/>
</div>

---

# Webエンジニアになるにあたって

<div class="center">
  <img src="assets/lecture4/web-engineer-growth-2.png" width="90%"/>
</div>

---

<!--
_class: section-head
-->
# なろう講習会 第一部 完！

お疲れ様でした！

本日の実習範囲:
データベースを扱う準備～サーバーからデータベースを扱う
