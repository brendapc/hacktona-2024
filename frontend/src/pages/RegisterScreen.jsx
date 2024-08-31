import { CardButtonComponent } from '../components/CardButtonComponent';

export const RegisterScreen = () => {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-custom-gradient">
      <h1 className="text-3xl font-semibold text-center mb-20">Cadastrar</h1>
      
      <div className="flex flex-wrap items-center justify-center gap-8 mt-8">
        <CardButtonComponent 
          onClick={() => console.log('Voluntário')} 
          title='Voluntário' 
          image='https://images03.brasildefato.com.br/ba2b3042aea1c53aeb6429a7e63a0719.webp'
        />
        <CardButtonComponent 
          onClick={() => console.log('Empresa')} 
          title='Empresa' 
          image='https://your-image-url.com/company.png'
        />
        <CardButtonComponent 
          onClick={() => console.log('Abrigo')} 
          title='Abrigo' 
          image='https://your-image-url.com/shelter.png'
        />
      </div>
    </div>
  );
};
