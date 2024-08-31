import { useState } from "react";
import InputApp from "../components/input/input";
import ButtonComponent from "../components/ButtonComponent";
import { FaEye, FaEyeSlash } from "react-icons/fa";

export const LoginPage = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [showPassword, setShowPassword] = useState(false); // Estado para controle da visibilidade da senha

  return (
    <div className="flex items-center justify-center min-h-screen bg-custom-gradient">
      <div className="w-full max-w-lg p-16 space-y-14 bg-[#F5F9E9] rounded-2xl shadow-lg">
        <h2 className="text-3xl font-bold text-center text-black">Login</h2>
        
        <form className="flex flex-col w-full" onSubmit={handleSubmit}>
          <div className="w-full mb-6 mt-4">
            <InputApp placeHolder="Email" onChangeText={setEmail} />
            {errors.email && <p className="text-red-500 text-sm mt-1">{errors.email}</p>}
          </div>

          {/* Campo de Senha com Funcionalidade de Olhinho */}
          <div className="w-full mb-32 relative">
            <InputApp 
              placeHolder="Senha" 
              type={showPassword ? "text" : "password"} // Alterna entre texto e senha
              onChangeText={setPassword} 
            />
            <button
              type="button"
              className="absolute right-4 top-1/2 transform -translate-y-1/2"
              onClick={() => setShowPassword(!showPassword)} // Alterna o estado da visibilidade
            >
              {showPassword ? <FaEyeSlash /> : <FaEye />} {/* Ícones para mostrar/ocultar senha */}
            </button>
          </div>

          <div className="flex justify-center">
            <ButtonComponent label="Entrar" />
          </div>
        </form>

        <div className="flex flex-col items-center text-center">
          <p className="text-sm mb-2 text-gray-600">
            Não possui conta?
          </p>
          <a
            href="/register"
            className="text-sm font-medium text-[#6060D4] hover:underline"
          >
            Cadastre-se
          </a>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
