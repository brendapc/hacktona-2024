import { BrowserRouter as Router, Route, Routes, } from 'react-router-dom';
import Home from './pages/Home';
import { LoginPage } from './pages/LoginPage';
// import CompanyRegister from './pages/CompanyRegisterScreen';
// import ShelterRegister from './pages/ShelterRegisterScreen';
import Abrigos from './pages/Shelter';

const App = () => {
  return (
    <Router>
      <div>
        <Routes>
          <Route path="/" exact element={<Home/>} />
          <Route path="/login" element={<LoginPage/>} />
          <Route path="/abrigos" element={<Abrigos />} />
        </Routes>
      </div>
    </Router>
  );
};

export default App;