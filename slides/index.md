---
marp: true
theme: SysAd
---

<style>
img:not(.emoji) {
  display: block;
  margin: 0 auto;
}
</style>

<!--
_class: title
-->

# なろう講習会　シラバス

---

<!--
_class: section-head
-->

# 第一部

---

# Day 1

- 講習会の目的
- 講習会の心構え
- ファイル・ディレクトリ
- ts の紹介？
  - 簡単な ts 説明
- go の紹介？
  - ポインタについて
  - インタプリター・コンパイラ
- （実習） 環境構築
- （実習）言語の練習をしてもいいかも

---

# Day 2

- おさらい・フロントエンド
  - HTML / CSS / JS
- Web サービスの仕組み
  - traQ を題材にして知る
  - バックエンド・フロント・DB を知ってもらう
- Vue について　（座学編）
  - なぜ HTML JS じゃないのか
- Vite npm node
- Alt 〇〇
- Vue について（実習編）
- フロントエンドのデバッグ　（実習編）

---

# Day 3

- HTTP - Part 1
  - URL URI
  - Path
  - リクエストとレスポンス
  - ステータスコード
  - Cookie
- 実習 - フロントエンド Day 2 - ページ分け, CSS
  - 軽いデザインとか？
- Cookie 使ってみる？

---

# Day 4

- バックエンドのデバッグ
  - エラーの扱い方
  - いろんなエラー (コンパイル・ネットワーク・ランタイム 500 400 などなど panic とかも)
- HTTP - Part 2
  - Header
  - JSON
  - Query
  - IP と ドメイン
  - Protocol GET 以外
- WebSocket? むずいかも　座学で触れるくらいかな

---

# Day 5

- DB
  - データの保存について
  - データの正規化
  - RDB
- 環境変数と Git
  - 認証情報の扱い方
  - SQL の構文とか？（実習でもいいかも）
- フロントエンドからの貫通

---

<!--
_class: section-head
-->

# 第二部

---

# Day 6

- デプロイ
- Docker
- サーバー接続
- 実習： FE から BE DB まで繋いで compose　を作ってみる

---

# Day 7

- アカウント（認証認可の基礎）
- セキュリティ（基礎）
- 実習
  - FE と BE 共同作業で、アカウントを作ってみる
  - CTF をしてみる

---

# Day 8

- テスト
- CI CD
- WebSocket、相互通信
- 実習：それぞれの CI/CD をやってみる

---

# Day 9

- 設計をしてみる
  - 機能要件と非機能要件
- コードのきれいさを知る
- 生成 AI との共同作業
- 実習：FE - API ページ設計 / BE - API - DB スキーマ設計
- 実習：それぞれで AI を触る