import { useState } from "react";
import ButtonComponent from "../components/ButtonComponent";
import InputApp from "../components/input/input";
import { FaEye, FaEyeSlash } from "react-icons/fa";

export default function CompanyRegister() {
    const [compName, setCompName] = useState('');
    const [cnpj, setCnpj] = useState('');
    const [address, setAddress] = useState('');
    const [phone, setPhone] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [showPassword, setShowPassword] = useState(false);
    const [errors, setErrors] = useState({});

    const validateForm = () => {
        const newErrors = {};
        if (!email) newErrors.email = "Email é obrigatório.";
        else if (!/\S+@\S+\.\S+/.test(email)) newErrors.email = "Email inválido.";
        if (!password) newErrors.password = "Senha é obrigatória.";
        if (!compName) newErrors.compName = "Nome da empresa é obrigatório.";
        if (!cnpj) newErrors.cnpj = "CNPJ é obrigatório.";
        else if (cnpj.length !== 14) newErrors.cnpj = "CNPJ inválido.";
        if (!address) newErrors.address = "Endereço é obrigatório.";
        if (!phone) newErrors.phone = "Telefone é obrigatório.";

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
                <h2 className="text-3xl font-bold text-center text-black">Cadastre-se como empresa</h2>
                <form className="w-full flex" onSubmit={handleSubmit}>
                    <div className="w-1/2 px-6">
                        <div className="w-full mb-3 space-y-1">
                            <label>Nome da empresa:</label>
                            <InputApp 
                                value={compName}
                                onChangeText={setCompName}
                            />
                            {errors.compName && <p className="text-red-500 text-sm mt-1">{errors.compName}</p>}
                        </div>
                        <div className="w-full mb-3 space-y-1">
                            <label>CNPJ (Digite apenas os números):</label>
                            <InputApp 
                                value={cnpj}
                                onChangeText={setCnpj}
                            />
                            {errors.cnpj && <p className="text-red-500 text-sm mt-1">{errors.cnpj}</p>}
                        </div>
                        <div className="w-full mb-3 space-y-1">
                            <label>Endereço:</label>
                            <InputApp 
                                value={address}
                                onChangeText={setAddress}
                            />
                            {errors.address && <p className="text-red-500 text-sm mt-1">{errors.address}</p>}
                        </div>
                    </div>

                    <div className="w-1/2 px-6">
                        <div className="w-full mb-3 space-y-1">
                            <label>Telefone:</label>
                            <InputApp 
                                value={phone}
                                onChangeText={setPhone}
                            />
                            {errors.phone && <p className="text-red-500 text-sm mt-1">{errors.phone}</p>}
                        </div>
                        <div className="w-full mb-3 space-y-1">
                            <label>Email:</label>
                            <InputApp 
                                value={email}
                                onChangeText={setEmail}
                            />
                            {errors.email && <p className="text-red-500 text-sm mt-1">{errors.email}</p>}
                        </div>
                        <div className="w-full mb-3 space-y-1 relative">
                            <label>Senha:</label>
                            <InputApp 
                                value={password}
                                onChangeText={setPassword}
                                type={showPassword ? "text" : "password"}
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
