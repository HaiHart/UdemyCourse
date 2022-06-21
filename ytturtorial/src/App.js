import { useState } from 'react';
import { ethers } from 'ethers';
import './App.css';
import Greeter from './artifacts/contracts/Greeter.sol/Greeter.json'
const GreeterAddress ='0x5fbdb2315678afecb367f032d93f642f64180aa3'

function App() {
  const [greeting, setGreeting] = useState('')
  const [newData,setNewData]=useState()
  async function requestAccounts() {
    await window.ethereum.request({method:'eth_requestAccounts'})
  }
  
  async function fetchGreeting() {
    if (typeof window.ethereum !== 'undefined') {
      const provider = new ethers.providers.Web3Provider(window.ethereum)
      const contract = new ethers.Contract(GreeterAddress, Greeter.abi, provider)
      try {
        const data = await contract.greet()
        setNewData(data)
        setGreeting(data)
        console.log('data: ',data)  
      } catch (err) {
        console.log(err)
      }
    }
  }
  
  async function setGreets() {
    if (!greeting) return
    if (typeof window.ethereum !== 'undefined') {
      await requestAccounts()
      const provider = new ethers.providers.Web3Provider(window.ethereum)
      const signer = provider.getSigner()
      const contract = new ethers.Contract(GreeterAddress, Greeter.abi, signer)
      const transaction = await contract.setGreeting(greeting)
      await transaction.wait();
      fetchGreeting()
    }
  }
  
  return (
    <div className="App">
      <header className="App-header">
        <button onClick={fetchGreeting}>fetchGreeting</button>
        <button onClick={setGreets}>set greeting</button>
        <input
          onChange={e => { setGreeting(e.target.value) }}
          placeholder='set greeting'
          value={greeting}
        />
      </header>
      <p>
        {newData}
      </p>
    </div>
  );
}

export default App;
