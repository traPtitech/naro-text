import { useState } from 'react'

export default function ClickCounter() {
  const [count, setCount] = useState<number>(0)
  return (
    <div>
      <div>回数： {count}</div>
      <button onClick={() => setCount(count + 1)}>クリック！</button>
      <button onClick={() => setCount(0)}>リセット！</button>
    </div>
  )
}
