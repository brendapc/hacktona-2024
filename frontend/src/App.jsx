import { BrowserRouter as Router, Route, Routes, } from 'react-router-dom';
import Home from './pages/Home';
import { LoginPage } from './pages/LoginPage';
import Abrigos from './pages/Shelter';
import { ShelterProfile } from './pages/ShelterProfile/ShelterProfile';
import { ShelteredForm } from './pages/ShelterProfile/ShelteredForm';
import { VolunteerList } from './pages/ShelterProfile/VolunteerList';

const App = () => {
  return (
    <Router>
      <div>
        <Routes>
          <Route path="/" exact element={<Home />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/abrigos" element={<Abrigos />} />
          <Route path="/perfil-abrigo" element={<ShelterProfile />} />
          <Route path="/cadastro-abrigados" element={<ShelteredForm />} />
          <Route path="/banco-de-voluntarios" element={<VolunteerList />} />
        </Routes>
      </div>
    </Router>
  );
};

export default App;