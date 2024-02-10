import React, { useContext, useEffect } from 'react';
import { Routes, Route, BrowserRouter, useNavigate } from 'react-router-dom';
import { AuthContext, AuthProvider } from './AuthContext/AuthContext';
import Login from './pages/login/Login';
import Home from './pages/home/Home';
import Dashboard from './pages/Manager/Dashboard';
import CreateSalons from './pages/Manager/CreateSalons';
import Register from './pages/register/Register';
import AddStaff from './pages/Manager/AddStaff';
import ModifyMySalon from './pages/Manager/ModifyStaff';

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
      <Route path="/addStaff" element={<AddStaff />} />
      <Route path="/modify" element={<ModifyMySalon />} />
    </Routes>
  );
}

export default App;
