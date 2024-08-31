import InputApp from "../components/input/input";
import ButtonComponent from "../components/ButtonComponent";

export const RegisterVolunteer = () => {
  const handleSelectChange = (event) => {
    event.target.blur();
  };
  return (
    <div className="flex items-center justify-center min-h-screen bg-custom-gradient">
      <div className="w-full max-w-lg p-16 space-y-14 bg-[#F5F9E9] rounded-2xl shadow-lg">
        <h2 className="text-3xl font-bold text-center text-black">Cadastra-se como voluntário</h2>
        
        <form className="flex flex-col w-full">
          <div className="w-full mb-6 mt-4">
            <InputApp placeHolder="Digite seu nome" />
          </div>

          <div className="w-full mb-6">
            <InputApp placeHolder="Digite seu email" />
          </div>

          <div className="w-full mb-6">
            <InputApp placeHolder="Digite sua senha" />
          </div>

          <div className="w-full mb-32 relative">
            <select 
              className="w-full px-4 py-2 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-300 rounded-md shadow-lg appearance-none focus:outline-none focus:ring-2 focus:ring-[#74D480]"
              style={{ 
                backgroundImage: 'url("data:image/svg+xml,%3Csvg xmlns=\'http://www.w3.org/2000/svg\' width=\'20\' height=\'20\' viewBox=\'0 0 24 24\' fill=\'none\' stroke=\'%23000\' stroke-width=\'2\' stroke-linecap=\'round\' stroke-linejoin=\'round\' class=\'feather feather-chevron-down\'%3E%3Cpolyline points=\'6 9 12 15 18 9\'/%3E%3C/svg%3E")', 
                backgroundRepeat: 'no-repeat', 
                backgroundPosition: 'right 12px center',
                backgroundSize: '16px 16px'
              }}
              onChange={handleSelectChange}
            >
              <option value="" disabled selected hidden>
                Selecione uma opção
              </option>
              <option value="volunteer1">Geral</option>
              <option value="volunteer2">Terapeuta</option>
              <option value="volunteer3">Doutor</option>
              <option value="volunteer4">Enfermeiro</option>
              <option value="volunteer5">Assistente Social</option>
              <option value="volunteer6">Veterinário</option>
              <option value="volunteer7">Outro</option>
            </select>
          </div>

          <div className="flex justify-center">
            <ButtonComponent label="Entrar" />
          </div>
        </form>

        <div className="flex flex-col items-center text-center">
          <p className="text-sm mb-2 text-gray-600">
            Já possui uma conta?
          </p>
          <a
            href="/register"
            className="text-sm font-medium text-[#6060D4] hover:underline"
          >
            Fazer Login
          </a>
        </div>
      </div>
    </div>
  );
};

export default RegisterVolunteer;