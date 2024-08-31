import { BrowserRouter as Router, Route, Routes, } from 'react-router-dom';
import Home from './pages/Home';
import { LoginPage } from './pages/LoginPage';
import CompanyRegister from './pages/CompanyRegisterScreen';
import ShelterRegister from './pages/ShelterRegisterScreen';
import RegisterVolunteer from './pages/RegisterVolunteer';
import { RegisterScreen } from './pages/RegisterScreen';
import Abrigos from './pages/Shelter';

const App = () => {
  return (
    <Router>
      <div>
        <Routes>
          <Route path="/" exact element={<Home/>} />
          <Route path="/login" element={<LoginPage/>} />
          <Route path="/register" element={<RegisterScreen/>} />
          <Route path="/registerVolunteer" element={<RegisterVolunteer/>} />
          <Route path="/CompanyRegister" element={<CompanyRegister/>} />
          <Route path="/ShelterRegister" element={<ShelterRegister/>} />
          <Route path="/abrigos" element={<Abrigos />} />
        </Routes>
      </div>
    </Router>
  );
};

export default App;