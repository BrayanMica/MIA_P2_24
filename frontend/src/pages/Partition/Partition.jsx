import partitionIMG from "../../assets/partition.png";
import { useState } from "react";
import { Link, useParams, useNavigate } from "react-router-dom";

export default function Partition({ip="localhost"}) {
  const { id } = useParams()
  const [data, setData] = useState([])
  const navigate = useNavigate()
  const [data2, setData2] = useState([])

  // execute the fetch command only once and when the component is loaded
  useState(() => {
    console.log(`fech to http://${ip}:4000/`)
    fetch(`http://${ip}:4000/tasks`)
      .then(response => response.json())
      .then(data => {console.log(data); setData2(data);})
    
    const rawData = {
      "rutas": ["Part1", "Part2", "Part3", "Part4", "Part5", "Part6", "Part7", "Part8", "Part9", "Part10", "Part11", "Part12", "Part13", "Part14", "Part15", "Part16", "Part17", "Part18", "Part19", "Part20"]
    }
    setData(rawData.rutas)

  }, [])

  const onClick = (objIterable) => {
    console.log("click", objIterable)
    navigate(`/login/${id}/${objIterable}`)
  }

  return (
    <>
      <h1>Partition {id}</h1>
      <br />
      <Link to="/discos">Regresar a los discos</Link>
    
      <h1>{data2.Status}</h1>
      <h2>{data2.Value}</h2>

      <div style={{ border: "red 1px solid", display: "flex", flexDirection: "row" , flexWrap: "wrap"}}>

        {
          data.map((objIterable, index) => {
            return (
              <div key={index} style={{
                border: "green 1px solid",
                display: "flex",
                flexDirection: "column", // Alinea los elementos en columnas
                alignItems: "center", // Centra verticalmente los elementos
                maxWidth: "100px",
              }}
                onClick={() => onClick(objIterable)}
              >
                <img src={partitionIMG} alt="disk" style={{ width: "100px" }} />
                <p>{objIterable}</p>
              </div>
            )
          })
        }

      </div>
    </>
  )
}