import { CardButtonComponent } from '../components/CardButtonComponent';
import { useNavigate } from 'react-router-dom';
import image1 from '../assets/image_1.png';
import image2 from '../assets/image_2.png';
import card_abrigo from '../assets/card-abrigo.png';

export const RegisterScreen = () => {
  const navigate = useNavigate();
  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-custom-gradient">
      <h1 className="text-3xl font-semibold text-center mb-8">Cadastrar</h1>
      
      <div className="flex flex-wrap items-center justify-center gap-8 mt-8">
        <CardButtonComponent 
          onClick={() => navigate('/registerVolunteer')} 
          title='Voluntário' 
          image={image1}
        />
        <CardButtonComponent 
          onClick={() => navigate('/CompanyRegister')} 
          title='Empresa' 
          image={card_abrigo}
        />
        <CardButtonComponent 
          onClick={() => navigate('/ShelterRegister')} 
          title='Abrigo' 
          image={image2}
        />
      </div>
    </div>
  );
};
