import { useContext } from 'react';
import { Routes, Route, BrowserRouter } from 'react-router-dom';
import { AuthContext, AuthProvider } from './AuthContext/AuthContext';
import Login from './pages/login/Login';
import Home from './pages/home/Home';
import Dashboard from './pages/Manager/Dashboard';
import CreateSalons from './pages/Manager/CreateSalons';
import Register from './pages/register/Register';
import Salon from './pages/salons/Salon';
import AddStaff from './pages/Manager/AddStaff';
import ModifyMySalon from './pages/Manager/ModifyStaff';
import Reservation from './pages/reservation/Reservation';
import Confirmation from './pages/confirmation/Confirmation';
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
      <Route path="/" element={<Home />} />
      <Route path="/salons/:id" element={<Salon />} />
      <Route path="salons/:id/reservation" element={<Reservation />} />
      <Route path="/confirmation" element={<Confirmation />} />
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
      <Route path="/dashboard" element={<Dashboard />} />
      <Route path="/create" element={<CreateSalons />} />
      <Route path="/addStaff" element={<AddStaff />} />
      <Route path="/modify" element={<ModifyMySalon />} />
    </Routes>
  );
}

export default App;
