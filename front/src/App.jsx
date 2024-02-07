// eslint-disable-next-line no-unused-vars
import React from 'react';
import { useContext, useEffect } from 'react';
import { AuthContext, AuthProvider } from './AuthContext/AuthContext';
import Login from './pages/login/Login';
import { Routes, Route, BrowserRouter, useNavigate } from 'react-router-dom';
import Home from './pages/home/Home';
import Register from './pages/register/Register';
import Dashboard from './pages/Manager/Dashboard';
import CreateSalons from './pages/Manager/CreateSalons';

function Layout() {
  const { isAuthenticated, isManager, logout} = useContext(AuthContext);
  const navigate = useNavigate();
  console.log(isManager)

  useEffect(() => {
    if (!isAuthenticated) {
      navigate('/login');
    }
  }, [isAuthenticated, navigate]);

  if (!isAuthenticated) {
    return null;
  }

  return (
    <div>
      <section>
      <header>
      <h1> Planity </h1>
        <button onClick={logout}> Se déconnecter </button>
      </header>
        <Routes>
          <Route path="/" element={isManager ? <Staff />: <Home />} />
        </Routes>
      </section>
    </div>
  );
}

function Staff(){

  return(
    <Routes>
    <Route path="/admin" element={Dashboard} />
          <Route path="/admin/create" element={CreateSalons} />
          
    </Routes>
  )

}

function App() {
  return (
    <div>
      <BrowserRouter>
        <AuthProvider>
          <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="/*" element={<Layout />} />
          </Routes>
        </AuthProvider>
      </BrowserRouter>
    </div>
  );
}

export default App;