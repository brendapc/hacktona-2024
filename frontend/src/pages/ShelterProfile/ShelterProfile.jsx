import React, { useEffect, useState } from "react";
import { ShelterHeader } from "../../components/ShelterHeader";
import InputApp from "../../components/input/input";
import axios from 'axios';

export const ShelterProfile = () => {
  const [shelterData, setShelterData] = useState({
    name: '',
    description: '',
    address: '',
    phone: '',
    email: '',
    current_occupancy: '',
    current_occupancy_pets: ''
  });

  useEffect(() => {
    const fetchShelterData = async () => {
      try {
        const response = await axios.get('http://localhost:8080/api/v1/shelters/1');
        setShelterData(response.data);
      } catch (error) {
        console.error("Erro ao buscar dados do abrigo:", error);
      }
    };

    fetchShelterData();
  }, []);

  return (
    <div className="min-h-screen bg-custom-gradient">
      <ShelterHeader page="profile" />
      <div className="flex justify-center items-center h-[680px]">
        <div className="w-5/6 flex-row p-16 space-y-14 bg-[#F5F9E9] rounded-2xl shadow-lg">
          <h2 className="text-3xl font-bold text-center text-black">Minhas Informações</h2>
          <div className="w-full flex">
            <div className="w-1/2 px-6">
              <div className="w-full mb-3 space-y-1">
                <label>Nome:</label>
                <InputApp value={shelterData.name} readOnly />
              </div>
              <div className="w-full mb-3 space-y-1">
                <label>Descrição:</label>
                <InputApp value={shelterData.description} readOnly />
              </div>
              <div className="w-full mb-3 space-y-1">
                <label>Endereço:</label>
                <InputApp value={shelterData.address} readOnly />
              </div>
            </div>

            <div className="w-1/2 px-6">
              <div className="w-full mb-3 space-y-1">
                <label>Telefone:</label>
                <InputApp value={shelterData.phone} readOnly />
              </div>
              <div className="w-full mb-3 space-y-1">
                <label>Email:</label>
                <InputApp value={shelterData.email} readOnly />
              </div>
              <div className="w-full mb-3 space-y-1">
                <label>Ocupação:</label>
                <InputApp value={shelterData.current_occupancy} readOnly />
              </div>
              <div className="w-full mb-3 space-y-1">
                <label>Ocupação de pets:</label>
                <InputApp value={shelterData.current_occupancy_pets} readOnly />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
