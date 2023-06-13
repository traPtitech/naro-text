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
  const [newItemName, setNewItemName] = useState('')
  const [newItemPrice, setNewItemPrice] = useState(0)

  const addItem = () => {
    setItems([...items, { name: newItemName, price: newItemPrice }])
  }

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
      <div>
        <label>
          名前
          <input onChange={(e) => setNewItemName(e.target.value)} type="text" value={newItemName} />
        </label>
        <label>
          価格
          <input
            onChange={(e) => setNewItemPrice(parseInt(e.target.value))}
            type="number"
            value={newItemPrice}
          />
        </label>
        <button onClick={addItem}>add</button>
      </div>
    </div>
  )
}
