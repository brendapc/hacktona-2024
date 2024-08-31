export const HomeHeader = () => {
  return (
    <header className="flex justify-between items-center h-16 p-4 backdrop-blur-md bg-white/70 shadow-md fixed w-full z-10">
      <div className="flex items-center">
        <span className="rounded-full p-2 bg-blue-500 text-white shadow-md">
          icon
        </span>
        <div className="relative">
          <input
            type="text"
            placeholder="Pesquisar..."
            className="rounded-full px-4 py-2 pl-10 ml-2 bg-gray-100 border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-400 transition duration-300 ease-in-out"
          />
          <span className="absolute left-3 top-2 text-gray-500">🔍</span>
        </div>
      </div>
      <div>
        <a
          href="/login"
          className="rounded-full px-6 py-2 bg-blue-500 text-white shadow-lg hover:bg-blue-600 transition duration-300 ease-in-out"
        >
          Entrar
        </a>
      </div>
    </header>
  );
};