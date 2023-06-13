import { useState } from 'react'

interface Item {
  name: string
  price: number
}

function ItemList() {
  const [items, setItems] = useState<Item[]>([
    { name: 'たまご', price: 100 },
    { name: 'りんご', price: 160 }
  ])
  return <div>ItemList</div>
}

export default ItemList
