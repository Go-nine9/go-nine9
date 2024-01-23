import { useContext, useEffect, useState } from 'react';
import './login.css';
import { AuthContext } from '../../AuthContext/AuthContext';
import { useNavigate } from 'react-router-dom';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState(null);
  const { login, isAuthenticated } = useContext(AuthContext);
  const navigate = useNavigate();

  useEffect(() => {
    const checkAuthentication = async () => {
      if (isAuthenticated) {
        navigate('/');
      }
    };

    checkAuthentication();
  }, [isAuthenticated, navigate]);

  async function handleSubmit(event) {
    event.preventDefault();
    if (!email || !password) {
      setError('Veuillez remplir tout les champs.');
      return;
    }
    try {
      await login(email, password);
      navigate('/');
    } catch (err) {
      setError(err.message);
    }
  }

  return (
    <div className="container-page">
      <div className="login-container">
        <h1 className="login-title">Bienvenue chez Gonine</h1>
        <form action="" className="container">
          <h2>se connecter</h2>
          <label htmlFor="">
            <input
              type="text"
              placeholder="mon@gmail.com"
              name="email"
              className="form-control"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </label>
          <label htmlFor="">
            <input
              type="password"
              placeholder="Password"
              className="form-control"
              name="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </label>
          <button onClick={(event) => handleSubmit(event)}>se connecter</button>
          {error && <p className="error">{error}</p>}
        </form>
      </div>
    </div>
  );
}

export default Login;
