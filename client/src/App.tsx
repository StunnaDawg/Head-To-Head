import React from "react"
import useSWR from "swr"
import "./index.css"

export const ENDPOINT = "http://localhost:4000"

const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((r) => r.json())

function App() {
  const { data, mutate } = useSWR("dashboard", fetcher)

  return (
    <>
      <div className="flex-1 flex flex-row justify-center">
        <h1 className="text-3xl font-bold underline">
          data:{JSON.stringify(data)}
        </h1>
      </div>
    </>
  )
}

export default App
