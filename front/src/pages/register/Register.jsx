// eslint-disable-next-line no-unused-vars
import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from '../../AuthContext/AuthContext';
import { useNavigate } from 'react-router-dom';
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
      <div className="register-container">
        <h1 className="register-title">Inscription chez Gonine</h1>
        <form action="" className="container">
          <h2>S'inscrire</h2>
          <label htmlFor="firstname">
            <input
              type="text"
              placeholder="Prénom"
              name="firstname"
              className="form-control"
              value={firstname}
              onChange={(e) => setFirstname(e.target.value)}
            />
          </label>
          <label htmlFor="lastname">
            <input
              type="text"
              placeholder="Nom de famille"
              name="lastname"
              className="form-control"
              value={lastname}
              onChange={(e) => setLastname(e.target.value)}
            />
          </label>
          <label htmlFor="email">
            <input
              type="text"
              placeholder="mon@gmail.com"
              name="email"
              className="form-control"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </label>
          <label htmlFor="password">
            <input
              type="password"
              placeholder="Mot de passe"
              className="form-control"
              name="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </label>
          <fieldset>
  <legend>S'inscrire en tant que</legend>

  <div>
    <input type="radio" id="user" name="user" value="users" onChange={(e)=>{setRole(e.target.value)}} />
    <label for="huey">Utilisateur</label>
  </div>

  <div>
    <input type="radio" id="user" name="user" value="manager" onChange={(e)=>{setRole(e.target.value)}} />
    <label for="dewey">Professionnel</label>
  </div>

</fieldset>

          <button onClick={(event) => handleSubmit(event)}>S'inscrire</button>
          {error && <p className="error">{error}</p>}
        </form>
      </div>
    </div>
  );
}

export default Register;