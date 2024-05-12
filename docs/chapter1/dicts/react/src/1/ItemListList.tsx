import { useState } from 'react'

interface Item {
  name: string
  price: number
}

export default function ItemList() {
  const [items, setItems] = useState<Item[]>([
    { name: 'たまご', price: 100 },
    { name: 'りんご', price: 160 }
  ])

  return (
    <div>
      <div>ItemList</div>
      <div>
        {items.map((item) => (
          <div key={item.name}>
            <div className="name">名前：{item.name}</div>
            <div className="price">{item.price}円</div>
          </div>
        ))}
      </div>
    </div>
  )
}
