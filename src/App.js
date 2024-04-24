import React from 'react';
import { BrowserRouter as Router, Route, Link, Routes } from 'react-router-dom';
import Primercomponente from './componentes/Primercomponente';
import Segundocomponente from './componentes/Segundocomponente';
import Tercercomponente from './componentes/Tercercomponente';
import './Styles/Estilos.css'; // Asegúrate de que esta ruta es correcta

function App() {
  return (
    <Router>
      <div className="App">
        <div className="sidebar">
          <ul>
            <li><Link className="link" to="/primercomponente">Primer Componente</Link></li>
            <li><Link className="link" to="/segundocomponente">Segundo Componente</Link></li>
            <li><Link className="link" to="/tercercomponente">Tercer Componente</Link></li>
          </ul>
        </div>
        <div className="main">
          <Routes>
            <Route path="/primercomponente" element={<Primercomponente />} />
            <Route path="/segundocomponente" element={<Segundocomponente />} />
            <Route path="/tercercomponente" element={<Tercercomponente />} />
          </Routes>
        </div>
      </div>
    </Router>
  );
}

export default App;