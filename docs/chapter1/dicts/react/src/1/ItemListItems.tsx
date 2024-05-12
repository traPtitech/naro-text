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

  return <div>ItemList</div>
}
