import { useState } from "react";
import ButtonComponent from "../components/ButtonComponent";
import InputApp from "../components/input/input";

export default function CompanyRegister() {
    const [compName,SetCompName] = useState('')
    const [cnpj,SetCnpj] = useState('')
    const [address,SetAddress] = useState('')
    const [phone,SetPhone] = useState('')
    const [email,SetEmail] = useState('')
    const [password,SetPassword] = useState('')
    return (
        <div className="flex items-center justify-center min-h-screen bg-custom-gradient">
            <div className="w-5/6 flex-row p-16 space-y-14 bg-[#F5F9E9] rounded-2xl shadow-lg">
                <h2 className="text-3xl  font-bold text-center text-black">Cadastre-se como empresa</h2>
                <div className="w-full flex">

                    <div className="w-1/2 px-6">
                        <div className="w-full mb-3 space-y-1">
                            <label>Nome da empresa:</label>
                            <InputApp onChangeText={SetCompName}/>
                        </div>
                        <div className="w-full mb-3 space-y-1">
                            <label>CNPJ:</label>
                            <InputApp onChangeText={SetCnpj}/>
                        </div>
                        <div className="w-full mb-3 space-y-1">
                            <label>Endereço:</label>
                            <InputApp onChangeText={SetAddress}/>
                        </div>
                    </div>


                    <div className="w-1/2 px-6">
                        <div className="w-full mb-3 space-y-1">
                            <label>Telefone:</label>
                            <InputApp onChangeText={SetPhone}/>
                        </div>
                        <div className="w-full mb-3 space-y-1">
                            <label>Email:</label>
                            <InputApp onChangeText={SetEmail}/>
                        </div>
                        <div className="w-full mb-3 space-y-1">
                            <label>Senha:</label>
                            <InputApp onChangeText={SetPassword}
                            type="password"/>
                        </div>
                    </div>
                </div>
                <div className="flex justify-center">
                    <ButtonComponent label={"Cadastrar"} onClick={()=>{}} />
                </div>
            </div>
        </div>
    )
}