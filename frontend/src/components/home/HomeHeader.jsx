import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Dropdown from '../Dropdown';

export const HomeHeader = () => {
  const [searchValue, setSearchValue] = useState('');
  const navigate = useNavigate();
  const [selectedA11yOption, setSelectedA11yOption] = useState("Select an option");
  const [selectedHelpOption, setSelectedHelpOption] = useState("Select an option");
  const [selectedSpotsOption, setSelectedSpotsOption] = useState("Select an option");

  const handleKeyDown = (event) => {
    if (event.key === 'Enter') {
      navigate(`/abrigos?value=${searchValue}`);
    }
  };

  return (
    <>
    <header className="flex justify-between items-center h-16 p-4 backdrop-blur-md bg-white/70 shadow-md w-full z-10">
      <div className="flex items-center">
        <span className="rounded-full p-2 bg-blue-500 text-white shadow-md">
          icon
        </span>
        <div className="relative">
          <input
            type="text"
            placeholder="Pesquisar..."
            value={searchValue}
            onChange={(e) => setSearchValue(e.target.value)}
            onKeyDown={handleKeyDown}
            className="rounded-full px-4 py-2 pl-10 ml-2 bg-gray-100 border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-400 transition duration-300 ease-in-out"
          />
          <span className="absolute left-5 top-2 text-gray-500">🔍</span>
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
    <div className='flex items-center bg-white h-10' >
      <div className='flex items-center mr-2'>
        <span> Acessibilidade: </span>
        <Dropdown 
          options={["tanto faz", "é acessivel"]} 
          selectedOption={selectedA11yOption} 
          setSelectedOption={setSelectedA11yOption} 
        />
      </div>
       <div className='flex items-center mr-2'>
        <span> Precisa de voluntarios de: </span>
        <Dropdown 
          options={["ajuda geral", "psicologia", "medicina", "enfermagem", "nutrição", "transporte", "veterinario", "outros"]} 
          selectedOption={selectedHelpOption} 
          setSelectedOption={setSelectedHelpOption} 
        />
      </div>
       <div className='flex items-center mr-2'>
        <span> Capacidade: </span>
        <Dropdown 
          options={["todos", "com vagas"]} 
          selectedOption={selectedSpotsOption} 
          setSelectedOption={setSelectedSpotsOption} 
        />
      </div>
    </div>
    </>
  );
};