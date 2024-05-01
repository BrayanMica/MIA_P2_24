import { Link } from "react-router-dom";

export default function Home() {
  return (
    <>
      <p>Administrador de discos</p>
      <br/>
      <Link to="/AppNavigator/DiskCreen">Commands</Link>
   </>
  )
}