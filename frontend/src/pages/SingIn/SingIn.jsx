import { Link, useParams } from "react-router-dom";


export default function SingIn({ip="localhost"}) {
   const { disk, part } = useParams()

   
   const handleSubmit = (e) => {
      e.preventDefault()
      console.log("submit", disk, part)

      const user = e.target.uname.value
      const pass = e.target.psw.value

      console.log("user", user, pass)
   }

  return (
    <>
      <h1>Login</h1>
      <br/>
      <Link to="/discos">Regresar a los Discos</Link>
      <br />
      <br />
      <br />
      <br />

      <form onSubmit={handleSubmit}>
         

         <div className="container">
            <label htmlFor="uname"><b>Username: </b></label>
            <input type="text" placeholder="Enter Username" name="uname" required/>

            <br />
            <br />   

            <label htmlFor="psw"><b>Password: </b></label>
            <input type="password" placeholder="Enter Password" name="psw" required/>
            <br />
            <br />
            <button type="submit">Login</button>
           
         </div>
        
      </form>


   </>
  )
}