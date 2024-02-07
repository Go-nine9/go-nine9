import React, {useEffect, useState} from 'react'
import { getCookie, getJWT } from '../../AuthContext/AuthContext';
import { useNavigate } from "react-router-dom";


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
      setSalon(response.data)
  
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
    Hello
      
    </div>
  )
}

export default Dashboard
