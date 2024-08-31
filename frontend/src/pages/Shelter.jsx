import { useLocation } from 'react-router-dom';

const Abrigos = () => {
  const location = useLocation();
  const queryParams = new URLSearchParams(location.search);
  const value = queryParams.get('value');

  return (
    <div>
      <h1>Abrigos</h1>
      <p>Query Parameter Value: {value}</p>

      {/* Listas os abrigos que contem esse valor na lista de coisas que o abrigo precisa
        
      */}
    </div>
  );
};

export default Abrigos;