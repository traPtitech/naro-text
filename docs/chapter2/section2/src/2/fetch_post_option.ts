const data = {
  /* ここにサーバーに渡すデータを入れる。 */
}
fetch('送信したい先のURL', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json'
  },
  body: JSON.stringify(data)
})
