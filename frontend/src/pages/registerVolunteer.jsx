import { useState } from "react";
import InputApp from "../components/input/input";
import ButtonComponent from "../components/ButtonComponent";

export const RegisterVolunteer = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [selectedOption, setSelectedOption] = useState("");

  const [errors, setErrors] = useState({}); // Estado para armazenar os erros

  const handleSelectChange = (event) => {
    setSelectedOption(event.target.value);
  };

  const validateForm = () => {
    const newErrors = {};
    if (!name) newErrors.name = "Nome é obrigatório.";
    if (!email) {
      newErrors.email = "Email é obrigatório.";
    } else if (!/\S+@\S+\.\S+/.test(email)) {
      newErrors.email = "Email inválido.";
    }
    if (!password) newErrors.password = "Senha é obrigatória.";
    if (!selectedOption) newErrors.selectedOption = "Por favor, selecione uma opção.";

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0; // Retorna true se não houver erros
  };

  const handleSubmit = (event) => {
    event.preventDefault(); // Prevenir o comportamento padrão de envio do formulário
    if (validateForm()) {
      console.log("Formulário enviado com sucesso!");
      // Adicione aqui o código para enviar os dados ao backend
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-custom-gradient">
      <div className="w-full max-w-lg p-16 space-y-8 bg-[#F5F9E9] rounded-2xl shadow-lg">
        <h2 className="text-3xl font-bold text-center text-black">Cadastre-se como voluntário</h2>
        
        <form className="flex flex-col w-full" onSubmit={handleSubmit}>
          {/* Campo de Nome */}
          <div className="w-full mb-4">
            <InputApp 
              placeHolder="Digite seu nome" 
              value={name} 
              onChangeText={setName} 
            />
            {errors.name && <p className="text-red-500 text-sm mt-1">{errors.name}</p>}
          </div>

          {/* Campo de Email */}
          <div className="w-full mb-4">
            <InputApp 
              type="email" 
              placeHolder="Digite seu email" 
              value={email} 
              onChangeText={setEmail} 
            />
            {errors.email && <p className="text-red-500 text-sm mt-1">{errors.email}</p>}
          </div>

          {/* Campo de Senha */}
          <div className="w-full mb-4">
            <InputApp 
              type="password" // Define o tipo como password para esconder o texto digitado
              placeHolder="Digite sua senha" 
              value={password} 
              onChangeText={setPassword} 
            />
            {errors.password && <p className="text-red-500 text-sm mt-1">{errors.password}</p>}
          </div>

          {/* Dropdown de Seleção */}
          <div className="w-full mb-16">
            <select 
              className="w-full px-4 py-2 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-300 rounded-md shadow-lg appearance-none focus:outline-none focus:ring-2 focus:ring-[#74D480]"
              style={{ 
                backgroundImage: 'url("data:image/svg+xml,%3Csvg xmlns=\'http://www.w3.org/2000/svg\' width=\'20\' height=\'20\' viewBox=\'0 0 24 24\' fill=\'none\' stroke=\'%23000\' stroke-width=\'2\' stroke-linecap=\'round\' stroke-linejoin=\'round\' class=\'feather feather-chevron-down\'%3E%3Cpolyline points=\'6 9 12 15 18 9\'/%3E%3C/svg%3E")', 
                backgroundRepeat: 'no-repeat', 
                backgroundPosition: 'right 12px center',
                backgroundSize: '16px 16px'
              }}
              value={selectedOption}
              onChange={handleSelectChange}
            >
              <option value="" disabled hidden>
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
            {errors.selectedOption && <p className="text-red-500 text-sm mt-1">{errors.selectedOption}</p>}
          </div>

          {/* Botão de Envio */}
          <div className="flex justify-center">
            <ButtonComponent label="Cadastrar" type="submit" />
          </div>
        </form>

        {/* Link para Login */}
        <div className="flex flex-col items-center text-center">
          <p className="text-sm mb-2 text-gray-600">
            Já possui uma conta?
          </p>
          <a
            href="/login"
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
