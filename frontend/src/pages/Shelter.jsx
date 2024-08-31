import { useLocation } from 'react-router-dom';
import { ShelterCard } from '../components/ShelterCard';

const Abrigos = () => {
  const location = useLocation();
  const queryParams = new URLSearchParams(location.search);
  const value = queryParams.get('value')?.toLowerCase();

  const shelterData = [
    {
      name: 'Abrigo Esperança',
      address: 'Avenida Ipiranga, 7090, Bairro Ipiranga',
      remainingSpots: 60,
      needs: ['Roupa de cama', 'Alimento', 'Tar Ner'],
      description: 'Descrição detalhada do abrigo, suas necessidades e outras informações importantes.'
    },
    {
      name: 'Abrigo Solidário',
      address: 'Rua das Flores, 123, Centro',
      remainingSpots: 45,
      needs: ['Agasalhos', 'Produtos de higiene', 'Alimentos não perecíveis'],
      description: 'Abrigo com foco em fornecer suporte emergencial durante o inverno e necessidades básicas para os menos favorecidos.'
    },
    {
      name: 'Abrigo Novo Horizonte',
      address: 'Rua São Pedro, 456, Vila Nova',
      remainingSpots: 30,
      needs: ['Alimentos', 'Produtos de limpeza', 'Fraldas'],
      description: 'Abrigo que oferece assistência a famílias em situação de vulnerabilidade com uma abordagem centrada na comunidade.'
    },
    {
      name: 'Abrigo Esperança do Futuro',
      address: 'Avenida Paulista, 987, Bela Vista',
      remainingSpots: 25,
      needs: ['Roupas', 'Medicamentos', 'Materiais escolares'],
      description: 'Abrigo que apoia pessoas em situação de rua e suas famílias, oferecendo uma variedade de serviços e suporte educacional.'
    }
  ];

  return (
    <div className='bg-custom-gradient h-screen w-screen overflow-scroll'>
      <h1 className='text-3xl font-bold ml-2 my-4 text-black'>Abrigos</h1>
      

      {
        shelterData.filter(shelter =>
          shelter.needs.some(need => need.toLowerCase().includes(value))
        ).map((shelter, index) => (
          <ShelterCard
            key={index}
            name={shelter.name}
            address={shelter.address}
            remainingSpots={shelter.remainingSpots}
            needs={shelter.needs}
            description={shelter.description}
            value={value}
          />
        ))
      }
      {shelterData
        .filter(shelter =>
          shelter.needs.some(need => need.toLowerCase().includes(value))
        ).length === 0 && value == !undefined && (
          <p className='font-bold'>Nenhum abrigo encontrado para a necessidade especificada.</p>
        )}

      {value == undefined && shelterData.map((shelter, index) =>
      (<ShelterCard
        key={index}
        name={shelter.name}
        address={shelter.address}
        remainingSpots={shelter.remainingSpots}
        needs={shelter.needs}
        description={shelter.description}
        value={value}
      />))}
    </div>
  );
};

export default Abrigos;
