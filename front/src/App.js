import logo from './logo.svg';
import './App.css';
import {useState} from "react";

function App() {
  const [data, setData] = useState(null)

  const load = () => {
    fetch("/api/user").then(x => x.json()).then(x => setData(x))
  }

  return (
    <div className="App">
      <button onClick={() => load()} >load</button>
      <pre>{JSON.stringify(data, null, 2)}</pre>
    </div>
  );
}

export default App;
