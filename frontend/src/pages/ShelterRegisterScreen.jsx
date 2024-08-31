import { useState } from "react";
import InputApp from "../components/input/input";
import ButtonComponent from "../components/ButtonComponent";
import axios from 'axios';

export default function ShelterRegister() {
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [address, setAddress] = useState('');
  const [phone, setPhone] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [capacity, setCapacity] = useState('');
  const [capacityPets, setCapacityPets] = useState('');
  const [accessibility, setAccessibility] = useState(false);

  const handleTextareaChange = (event) => {
    setDescription(event.target.value);
    event.target.style.height = 'auto';
    event.target.style.height = event.target.scrollHeight + 'px';
  };

  const handleRegister = async () => {
    const shelterData = {
      name: name,
      description: description,
      address: address,
      phone: phone,
      email: email,
      password: password,
      capacity: parseInt(capacity),
      capacity_pets: parseInt(capacityPets),
      acessibility: accessibility
    };

    try {
      const response = await axios.post('http://localhost:8080/api/v1/shelters', shelterData);
      if (response.status === 201) {
        alert("Abrigo cadastrado com sucesso!");
      }
    } catch (error) {
      console.error("Erro ao cadastrar abrigo:", error);
      alert("Ocorreu um erro ao cadastrar o abrigo.");
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-custom-gradient">
      <div className="w-5/6 flex-row p-16 space-y-14 bg-[#F5F9E9] rounded-2xl shadow-lg">
        <h2 className="text-3xl font-bold text-center text-black">Cadastre-se como Abrigo</h2>
        <div className="w-full flex">
          <div className="w-1/2 px-6">
            <div className="w-full mb-3 space-y-1">
              <label>Nome do abrigo:</label>
              <InputApp placeHolder="" onChangeText={setName} />
            </div>
            <div className="w-full mb-2 space-y-1">
              <label>Descrição:</label>
              <textarea
                placeholder=""
                className="w-full px-3 py-2 placeholder-gray-400 bg-white border border-none rounded-md shadow-md drop-shadow-md focus:outline-none focus:ring-2 focus:ring-[#74D480] resize-none"
                rows="1"
                onChange={handleTextareaChange}
              ></textarea>
            </div>
            <div className="w-full mb-3 space-y-1">
              <label>Endereço:</label>
              <InputApp placeHolder="" onChangeText={setAddress} />
            </div>
            <div className="w-full mb-3 space-y-1">
              <label>Telefone:</label>
              <InputApp placeHolder="" onChangeText={setPhone} />
            </div>
          </div>

          <div className="w-1/2 px-6">
            <div className="w-full mb-3 space-y-1">
              <label>Email:</label>
              <InputApp placeHolder="" onChangeText={setEmail} />
            </div>
            <div className="w-full mb-3 space-y-1">
              <label>Senha:</label>
              <InputApp placeHolder="" onChangeText={setPassword} type="password" />
            </div>
            <div className="w-full mb-3 space-y-1">
              <label>Capacidade:</label>
              <InputApp placeHolder="" onChangeText={setCapacity} />
            </div>
            <div className="w-full mb-3 space-y-1">
              <label>Capacidade para Pets:</label>
              <InputApp placeHolder="" onChangeText={setCapacityPets} />
            </div>
            <div className="w-full mb-3 space-y-1">
              <label>Acessibilidade:</label>
              <input 
                type="checkbox" 
                className="h-6 w-6"
                checked={accessibility} 
                onChange={() => setAccessibility(!accessibility)} 
              />
            </div>
          </div>
        </div>
        <div className="flex justify-center">
          <ButtonComponent label={"Cadastrar"} onClick={handleRegister} />
        </div>
      </div>
    </div>
  );
}
