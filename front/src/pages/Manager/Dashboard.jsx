import React, {useEffect, useState} from 'react'
import { getCookie, getJWT } from '../../AuthContext/AuthContext';
import { Link, useNavigate } from "react-router-dom";


const Dashboard = () => {
  const [salon, setSalon] = useState();
  const navigate = useNavigate();

  const getSalon = async (token) =>{
    try {
      const response = await fetch('http://localhost:8097/api/management/salons', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
        mode: 'cors',
      });
  
      if (!response.ok) {
        const errorResponse = await response.json();
        throw new Error(errorResponse.message || 'Échec de la requête d\'inscription');
      }

      const goodResponse = await response.json();
      console.log(goodResponse[0])
      setSalon(goodResponse[0])
  
    } catch (err) {
      throw new Error(err.message || 'Une erreur inattendue s\'est produite');
    }
  }

  useEffect(() => {
    const cookie = getCookie('authToken');
    const token = getJWT(cookie)
    const salonID = token.salonID
    if(salonID !== null){
       getSalon(cookie)
    }else{
      navigate('/admin/create')
    }
  }, []);


  return (
    <div>
   <h1> Dashboard de mon salon </h1>
   <Link to="/admin/modify"><button> Modifier mon salon </button></Link>
   <Link to="/admin/addStaff"><button> Ajouter des salariés </button></Link>
   <button> Supprimer mon salon </button>
   {salon ?
   <>
    <h2>{salon.Name}</h2>
    <p>adresse de mon salon : {salon.Address}</p>
    <p>Téléphone: {salon.Phone}</p>
    <h2>L'équipe</h2>
    {salon.User.map((user)=>(
      <>
      <h2>{user.Firstname}  {user.Lastname} {user.Roles === "manager" && "(moi)"}</h2>
      <p>{user.Email} </p>
      <p>{user.Roles} </p>
      </>
    ))}

    </>  
    
   : <h1> Vous n'avez pas encore de salon snif <Link to="/admin/create"> Cliquez ici pour en créer un </Link></h1> }
      
    </div>
  )
}

export default Dashboard
