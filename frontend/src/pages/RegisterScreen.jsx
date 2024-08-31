import { CardButtonComponent } from '../components/CardButtonComponent';
import { useNavigate } from 'react-router-dom';

export const RegisterScreen = () => {
  const navigate = useNavigate();
  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-custom-gradient">
      <h1 className="text-3xl font-semibold text-center mb-8">Cadastrar</h1>
      
      <div className="flex flex-wrap items-center justify-center gap-8 mt-8">
        <CardButtonComponent 
          onClick={() => navigate('/registerVolunteer')} 
          title='Voluntário' 
          image='https://images03.brasildefato.com.br/ba2b3042aea1c53aeb6429a7e63a0719.webp'
        />
        <CardButtonComponent 
          onClick={() => navigate('/CompanyRegister')} 
          title='Empresa' 
          image='https://your-image-url.com/company.png'
        />
        <CardButtonComponent 
          onClick={() => navigate('/ShelterRegister')} 
          title='Abrigo' 
          image='https://your-image-url.com/shelter.png'
        />
      </div>
    </div>
  );
};
