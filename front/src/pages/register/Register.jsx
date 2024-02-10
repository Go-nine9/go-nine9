// eslint-disable-next-line no-unused-vars
import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from '../../AuthContext/AuthContext';
import { useNavigate } from 'react-router-dom';
import { Link } from 'react-router-dom';
import './register.css'

function Register() {
  const [firstname, setFirstname] = useState('');
  const [lastname, setLastname] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [role, setRole] = useState("users");
  const [error, setError] = useState(null);
  const navigate = useNavigate();
  const { register, isAuthenticated } = useContext(AuthContext);

  // Redirigez l'utilisateur vers une autre page si déjà authentifié
  useEffect(() => {
    if (isAuthenticated) {
      navigate('/');
    }
  }, [isAuthenticated, navigate]);

  async function handleSubmit(event) {
    event.preventDefault();
    if (!firstname || !lastname || !email || !password || !role) {
      setError('Veuillez remplir tous les champs.');
      return;
    }
    try {
      await register(firstname, lastname, email, password, role);
      navigate('/');
    } catch (err) {
      setError(err.message);
    }
  }


  return (
    <div className="container-page">
      <div className="left">
        <div className="register-container">
          <h1 className="register-title">Nouveau sur Gonine ?</h1>
          <form action="" className="container">
            <div className="top-info">
              <label htmlFor="lastname">
                Nom *
                <input
                  type="text"
                  placeholder="Nom de famille"
                  name="lastname"
                  className="form-control"
                  value={lastname}
                  onChange={(e) => setLastname(e.target.value)}
                />
              </label>
              <label htmlFor="firstname">
                Prénom *
                <input
                  type="text"
                  placeholder="Prénom"
                  name="firstname"
                  className="form-control"
                  value={firstname}
                  onChange={(e) => setFirstname(e.target.value)}
                />
              </label>
            </div>
            <label htmlFor="email">
              Email *
              <input
                type="text"
                placeholder="Email"
                name="email"
                className="form-control"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
            </label>
            <label htmlFor="password">
              Mot de passe *
              <input
                type="password"
                placeholder="Mot de passe"
                className="form-control"
                name="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </label>
            
            <div className="role-container">
              <h2>Êtes-vous un professionnel</h2>
              <div className="role-content">
                <label htmlFor="dewey" className='roles'>
                  <input 
                    type="radio" 
                    id="user"
                    name="user"
                    value="manager"
                    onChange={(e)=>{setRole(e.target.value)}}
                  />
                  Oui
                </label>
                <label htmlFor="huey" className='roles'>
                  <input 
                    type="radio"
                    id="user"
                    name="user"
                    value="users"
                    checked
                    onChange={(e)=>{setRole(e.target.value)}}
                  />
                  Non
                </label>
              </div>
            </div>
            <button className="btn-primary"onClick={(event) => handleSubmit(event)}>Créer mon compte</button>
            {error && <p className="error">{error}</p>}
          </form>
          <span className="separation">ou</span>
            <div className="register">
              <h2>Vous avez déjà utilisé Gonine?</h2>
              <Link to="/login" className='btn-secondary'>Créer mon compte</Link>
          </div>
        </div>
      </div>
      <div className="right">
        <img src="https://res.cloudinary.com/planity/image/upload/q_auto,f_auto/v1701340648/portail/illustrations/LOGIN/2023/3.jpg" alt="" />
      </div>
    </div>
  );
}

export default Register;