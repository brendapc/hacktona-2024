import InputApp from "../components/input/input";

export const LoginPage = () => {
  return (
    <div className="flex items-center justify-center min-h-screen bg-custom-gradient">
      <div className="w-full max-w-md p-8 space-y-6 bg-[#F5F9E9] rounded-2xl shadow-lg">
        <h2 className="text-2xl font-bold text-center text-black">Login</h2>
        
        <form className="space-y-6">
          <div>
            <InputApp placeHolder="Digite seu e-mail" />
          </div>

          <div>
            <InputApp placeHolder="Digite sua senha" />
          </div>

          <div>
            <button
              type="submit"
              className="w-full px-4 py-2 text-white rounded-md shadow-lg bg-[#6060D4] hover:bg-[#4E4EB8] focus:outline-none focus:ring-2 focus:ring-[#6060D4] focus:ring-opacity-50"
            >
              Entrar
            </button>
          </div>
        </form>

        <div className="text-center">
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
