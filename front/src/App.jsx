import React, { useContext, useEffect } from 'react';
import { Routes, Route, BrowserRouter, useNavigate } from 'react-router-dom';
import { AuthContext, AuthProvider } from './AuthContext/AuthContext';
import Login from './pages/login/Login';
import Home from './pages/home/Home';
import Register from './pages/register/Register';
import Dashboard from './pages/Manager/Dashboard';
import CreateSalons from './pages/Manager/CreateSalons';

function App() {
  const { isAuthenticated, isManager } = useContext(AuthContext);

  const ManagerRoutes = (
    <Routes>
        <Route path="/admin/*" element={<StaffRoutes />} />
      <Route path="/*" element={<Home />} />
    </Routes>
  );

  const publicRoutes = (
    <Routes>
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
      <Route path="/*" element={<Login />} />
    </Routes>
  );

  const Layout = () => {
    return (
      <div>
        <BrowserRouter>
          <AuthProvider>
            {isManager ? ManagerRoutes : isAuthenticated ? publicRoutes : publicRoutes}
          </AuthProvider>
        </BrowserRouter>
      </div>
    );
  };

  return <Layout />;
}

function StaffRoutes() {
  return (
    <Routes>
      <Route path="/" element={<Dashboard />} />
      <Route path="/create" element={<CreateSalons />} />
    </Routes>
  );
}

export default App;
