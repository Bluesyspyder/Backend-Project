import { useEffect, useState } from 'react'

type Todo = {
  id: number
  body: string
  completed: boolean
}

const App = () => {

  const [todos,setTodos] = useState<Todo[]>([]);

  useEffect(() => {
    fetch('/api/todos')
    .then(res => res.json())
    .then(setTodos)
  },[])


  return (
    <>
      <div className='flex flex-col items-center justify-center text-3xl'>
        <span >
          Todos
        </span>
        <h3>
          {todos.map(t => (
            <div key={t.id}>
              hi
            </div>
          ))}
        </h3>
      </div>
    </>
  )
}

export default App