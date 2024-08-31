import { useState } from "react";
import InputApp from "../components/input/input";
import ButtonComponent from "../components/ButtonComponent";
import { FaEye, FaEyeSlash } from "react-icons/fa";

export default function ShelterRegister() {
    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [address, setAddress] = useState('');
    const [phone, setPhone] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [capacity, setCapacity] = useState('');
    const [capacityPets, setCapacityPets] = useState('');
    const [showPassword, setShowPassword] = useState(false);
    const [errors, setErrors] = useState({});

    const handleTextareaChange = (event) => {
        event.target.style.height = 'auto';
        event.target.style.height = event.target.scrollHeight + 'px';
        setDescription(event.target.value); 
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
        if (!description) newErrors.description = "Descrição é obrigatória.";
        if (!address) newErrors.address = "Endereço é obrigatório.";
        if (!phone) newErrors.phone = "Telefone é obrigatório.";
        if (!capacity) newErrors.capacity = "Capacidade é obrigatória.";
        if (!capacityPets) newErrors.capacityPets = "Capacidade para Pets é obrigatória.";

        setErrors(newErrors);
        return Object.keys(newErrors).length === 0; 
    };

    const handleSubmit = (event) => {
        event.preventDefault(); 
        if (validateForm()) {
            console.log("Sucesso!");
        }
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-custom-gradient">
            <div className="w-5/6 flex-row p-16 space-y-14 bg-[#F5F9E9] rounded-2xl shadow-lg">
                <h2 className="text-3xl font-bold text-center text-black">Cadastre-se como Abrigo</h2>
                <form className="w-full flex" onSubmit={handleSubmit}>
                    <div className="w-1/2 px-6">
                        <div className="w-full mb-3 space-y-1">
                            <label>Nome do abrigo:</label>
                            <InputApp 
                                value={name}
                                onChangeText={setName}
                                placeHolder="Nome do abrigo"
                            />
                            {errors.name && <p className="text-red-500 text-sm mt-1">{errors.name}</p>}
                        </div>
                        <div className="w-full mb-2 space-y-1">
                            <label>Descrição:</label>
                            <textarea
                                placeholder="Descrição do abrigo"
                                className="w-full px-3 py-2 placeholder-gray-400 bg-white border border-none rounded-md shadow-md drop-shadow-md focus:outline-none focus:ring-2 focus:ring-[#74D480] resize-none"
                                rows="1"
                                value={description}
                                onChange={handleTextareaChange}
                            />
                            {errors.description && <p className="text-red-500 text-sm mt-1">{errors.description}</p>}
                        </div>
                        <div className="w-full mb-3 space-y-1">
                            <label>Endereço:</label>
                            <InputApp 
                                value={address}
                                onChangeText={setAddress}
                                placeHolder="Endereço do abrigo"
                            />
                            {errors.address && <p className="text-red-500 text-sm mt-1">{errors.address}</p>}
                        </div>
                        <div className="w-full mb-3 space-y-1">
                            <label>Telefone:</label>
                            <InputApp 
                                value={phone}
                                onChangeText={setPhone}
                                placeHolder="Telefone do abrigo"
                            />
                            {errors.phone && <p className="text-red-500 text-sm mt-1">{errors.phone}</p>}
                        </div>
                    </div>

                    <div className="w-1/2 px-6">
                        <div className="w-full mb-3 space-y-1">
                            <label>Email:</label>
                            <InputApp 
                                value={email}
                                onChangeText={setEmail}
                                placeHolder="Email do abrigo"
                            />
                            {errors.email && <p className="text-red-500 text-sm mt-1">{errors.email}</p>}
                        </div>
                        <div className="w-full mb-3 space-y-1 relative">
                            <label>Senha:</label>
                            <InputApp 
                                value={password}
                                onChangeText={setPassword}
                                type={showPassword ? "text" : "password"}
                                placeHolder="Senha do abrigo"
                            />
                            <button
                                type="button"
                                className="absolute right-4 top-1/2 transform -translate-y-1/2 flex-row-reverse"
                                onClick={() => setShowPassword(!showPassword)}
                            >
                                {showPassword ? <FaEyeSlash /> : <FaEye />}
                            </button>
                            {errors.password && <p className="text-red-500 text-sm mt-1">{errors.password}</p>}
                        </div>
                        <div className="w-full mb-3 space-y-1">
                            <label>Capacidade:</label>
                            <InputApp 
                                value={capacity}
                                onChangeText={setCapacity}
                                placeHolder="Capacidade do abrigo"
                            />
                            {errors.capacity && <p className="text-red-500 text-sm mt-1">{errors.capacity}</p>}
                        </div>
                        <div className="w-full mb-3 space-y-1">
                            <label>Capacidade para Pets:</label>
                            <InputApp 
                                value={capacityPets}
                                onChangeText={setCapacityPets}
                                placeHolder="Capacidade para pets"
                            />
                            {errors.capacityPets && <p className="text-red-500 text-sm mt-1">{errors.capacityPets}</p>}
                        </div>
                    </div>
                </form>
                <div className="flex justify-center">
                    <ButtonComponent 
                        label={"Cadastrar"} 
                        onClick={handleSubmit} 
                    />
                </div>
            </div>
        </div>
    );
}
