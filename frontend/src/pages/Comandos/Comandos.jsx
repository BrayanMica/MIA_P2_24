import { Link } from "react-router-dom";

export default function Home() {
  return (
    <>
      <h1>Ingreso de comandos</h1>
      <div style={{height: "90vh", padding: "10px"}}>
      <textarea readOnly style={{width: "90%", height: "calc(100% - 100px)", resize: "none", padding: "5x", marginBottom: "5px"}} placeholder="Lista de comandos ejecutados..."></textarea>
      <input style={{width: "90%", padding: "10px", border: "none", borderBottom: "2px solid #000"}} type="text" placeholder="Escribe un comando..." />
      <button style={{width: "8%", backgroundColor: "#4CAF50", color: "white", border: "none", padding: "10px", cursor: "pointer"}}
              onMouseOver={e => e.target.style.backgroundColor = "#45a049"}
              onMouseOut={e => e.target.style.backgroundColor = "#4CAF50"}>
        Enviar
      </button>
    </div>
   </>
  )
}