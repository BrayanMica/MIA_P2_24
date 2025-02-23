import diskIMG from "../../assets/disk.png";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

export default function DiskCreen({ip="localhost"}) {
  const [data, setData] = useState([]) 
  const navigate = useNavigate()
  
  // execute the fetch command only once and when the component is loaded
  useState(() => {
 
    const rawData = {
      "rutas":["A.dsk","B.dsk","C.dsk","D.dsk","A.dsk","B.dsk","C.dsk","D.dsk","A.dsk","B.dsk","C.dsk","D.dsk","E.dsk","F.dsk","G.dsk","H.dsk","I.dsk","J.dsk","K.dsk","L.dsk","M.dsk","N.dsk","O.dsk","P.dsk","Q.dsk","R.dsk","S.dsk","T.dsk","U.dsk","V.dsk","W.dsk","X.dsk","Y.dsk","Z.dsk"]
     }
    setData(rawData.rutas)

    var dataF = {
      User: 'root',
      Password: 'admin'
    }
    console.log(`fech to http://${ip}:4000/`)
    fetch(`http://${ip}:4000/tasks`, {
      method: 'POST', 
      headers: {
        'Content-Type': 'application/json' 
      },
      body: JSON.stringify(dataF)
    })
    .then(response => response.json())
    .then(data => {
      console.log(data); // Do something with the response
      setData(data.List)
    })
    .catch(error => {
      console.error('There was an error with the fetch operation:', error);
    });
  }, [])

  const onClick = (objIterable) => {
    //e.preventDefault()
    console.log("click",objIterable)
    navigate(`/discos/${objIterable}`)
  }

  return (
    <>
      <h1>Discos Encontrados</h1>
      <br/>

      <div style={{border:"red 1px solid",display: "flex", flexDirection: "row", flexWrap: "wrap"}}>

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
                <img src={diskIMG} alt="disk" style={{width: "100px"}} />
                <p>{objIterable}</p>
              </div>
            )
          })
        }
      
      </div>
    </>
   )
 }
 