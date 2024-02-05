// eslint-disable-next-line no-unused-vars
import React from 'react';
import { useContext, useEffect } from 'react';
import { AuthContext, AuthProvider } from './AuthContext/AuthContext';
import Login from './pages/login/Login';
import { Routes, Route, BrowserRouter, useNavigate } from 'react-router-dom';
import Home from './pages/home/Home';
import Register from './pages/register/Register';
import Dashboard from './pages/Manager/Dashboard';

function Layout() {
  const { isAuthenticated, isManager} = useContext(AuthContext);
  const navigate = useNavigate();

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
        <Routes>
          <Route path="/" element={isManager ? <Dashboard />: <Home />} />
        </Routes>
      </section>
    </div>
  );
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