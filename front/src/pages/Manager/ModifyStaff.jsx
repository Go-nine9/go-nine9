import React, {useState} from 'react'
import { getCookie, setCookie } from '../../AuthContext/AuthContext';
import { useNavigate, useLocation } from 'react-router-dom';


const ModifyMySalon = () => {
    const navigate = useNavigate();
    let { state  } = useLocation();
    console.log(state.state)

    const [ModifiedSalon, setModifiedSalon] = useState(state.state);
    const [errors, setErrors] = useState("")

    const token = getCookie("authToken")



    // Change les valeurs du salon
    const handleInputChange = (event) => {
        const { name, value } = event.target;
        setModifiedSalon({ ...ModifiedSalon, [name]: value });
      };

    const handleModify = async() =>{
        const {Name, Address, Phone, user} = ModifiedSalon;
        const requestData = {Name, Address, Phone, user};
        try {
            const response = await fetch(`http://localhost:8097/api/management/salons/${state.state.ID}`, {
              method: 'PATCH',
              headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,
              },
              mode: 'cors',
              body: JSON.stringify(requestData),
            });
        
            if (!response.ok) {
              const errorResponse = await response.json();
              console.log(errorResponse)
              setErrors(errorResponse.message || 'Échec de la requête d\'inscription');
            }
            if(response.ok){
                setErrors("")
                navigate("/admin/")
            }
           
            

          } catch (err) {
            setErrors(err.message || 'Une erreur inattendue s\'est produite');
          }

    }

    // // Change les valeurs d'une personne du staff
    // const handleStaffInputChange = (index, event) => {
    //     const { name, value } = event.target;
    //     const updatedStaff = [...ModifiedSalon.User];
    //     updatedStaff[index][name] = value;
    //     setModifiedSalon({ ...ModifiedSalon, user: updatedStaff });
    //   };

    //   //Ajoute une nouvelle entrée de input
    //   const addStaff = () => {
    //     setModifiedSalon({
    //       ...ModifiedSalon,
    //       user: [
    //         ...ModifiedSalon.user,
    //         { Lastname: '', Firstname: '', Email: '', Password: '' , Roles : "staff"},
    //       ],
    //     });
    //   };

    //   //retire un staff
    //   const removeStaff = (index) => {
    //     const updatedStaff = [...ModifiedSalon.user];
    //     updatedStaff.splice(index, 1);
    //     setModifiedSalon({ ...ModifiedSalon, user: updatedStaff });
    //   };
    

      
  return (
    <div>
        <h1> Modifier un Salon </h1>
        {errors.length > 0 && <h3>{errors}</h3>}
        <div>
        {ModifiedSalon && 
        
            <>
            <button onClick={handleModify}>Modifies ton salon</button>
            <label>Nom du salon</label>
            <input name="Name" type="text" placeholder='Ex: Mon salon' value={ModifiedSalon.Name} onChange={handleInputChange}/>
            <label>Adresse</label>
            <input name="Address" type="text" placeholder='Ex: 10 rue de Louvois 75002 Paris' value={ModifiedSalon.Address} onChange={handleInputChange}/>
            <label>Numéro de téléphone</label>
            <input name="Phone" type="tel" placeholder="Ex: 667" value={ModifiedSalon.Phone} onChange={handleInputChange}/>
            {/* <button onClick={addStaff}>Ajouter des collaborateurs</button> */}
            {/* {ModifiedSalon.User.map((staffMember, index) => (
          <div key={index}>
          
          <h3>Staff {index+1}</h3>
            <label>Nom</label>
            <input
              name="Lastname"
              type="text"
              value={staffMember.Lastname}
              onChange={(e) => handleStaffInputChange(index, e)}
            />
            <label>Prénom</label>
            <input
              name="Firstname"
              type="text"
              value={staffMember.Firstname}
              onChange={(e) => handleStaffInputChange(index, e)}
            />
            <label>Email</label>
            <input
              name="Email"
              type="email"
              value={staffMember.Email}
              onChange={(e) => handleStaffInputChange(index, e)}
            />
            <label>Mot de passe</label>
            <input
              name="Password"
              type="password"
              value={staffMember.Password}
              onChange={(e) => handleStaffInputChange(index, e)}
            />
              <button onClick={() => removeStaff(index)}>Supprimer</button>
          </div>
        ))} */}
            
          
            </>
            }

            </div>
      
    </div>
  )
}

export default ModifyMySalon
