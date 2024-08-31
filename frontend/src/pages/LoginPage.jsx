import InputApp from "../components/input/input";
import ButtonComponent from "../components/ButtonComponent";

export const LoginPage = () => {
  return (
    <div className="flex items-center justify-center min-h-screen bg-custom-gradient">
      <div className="w-full max-w-lg p-16 space-y-14 bg-[#F5F9E9] rounded-2xl shadow-lg">
        <h2 className="text-3xl font-bold text-center text-black">Login</h2>
        
        <form className="flex flex-col items-center w-full">
          <div className="w-full mb-6 mt-4">
            <InputApp placeHolder="Digite seu e-mail" />
          </div>

          <div className="w-full mb-32">
            <InputApp placeHolder="Digite sua senha" />
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
