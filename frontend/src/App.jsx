/* App.jsx */
import { BrowserRouter as Router, Link, Routes, Route } from 'react-router-dom';
import { useState } from "react";
import './App.css';
import Comandos from './pages/Comandos/Comandos';
import Informes from './pages/Informes/Informes';
import Disk from './pages/DiskCreen/DiskCreen';
import Partition from './pages/Partition/Partition';
import SingIn from './pages/SingIn/SingIn'; 

function App() {
  const [ip, setIP] = useState("localhost") 

    const handleChage = (e) => {
      console.log(e.target.value)
      setIP(e.target.value)
    }
  return (
    <>
    IP: <input type="text" onChange={handleChage}/> -- {ip}
    <Router>
      <div style={{ display: 'flex', height: '100vh' }}>
        <div style={{ width: '20%', background: 'black', padding: '10px' }}>
          <ul>
            <li className="link"><Link to="/comandos"> Comandos</Link></li>
            <li className="link"><Link to="/discos"> Disk</Link></li>
            <li className="link"><Link to="/informes"> Informes</Link></li>
          </ul>
        </div>
        <div style={{ flex: 1, padding: '10px', overflow: 'auto', display: 'flex', flexDirection: 'column' }}>
          <Routes>
            <Route path="/comandos" element={<Comandos ip={ip}/>} />
            <Route path="/discos" element={<Disk ip={ip}/>} />
            <Route path="/informes" element={<Informes ip={ip}/>} />
            <Route path="/discos/:id/" element={<Partition ip={ip}/>} />"
            <Route path="/login/:disk/:part" element={<SingIn ip={ip}/>} />
          </Routes>
        </div>
      </div>
    </Router>
    </>
  );
}

export default App;