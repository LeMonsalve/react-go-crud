import {FC} from "react"

const App: FC = () => {

    const getData = async () => {
        const data = await fetch('http://localhost:4000/users').then(response => response.json())
        console.log( data )
    }

    return (
        <div>
            <h1>React Go CRUD</h1>
            <button onClick={ getData }>Get Data</button>
        </div>
    )
}

export default App