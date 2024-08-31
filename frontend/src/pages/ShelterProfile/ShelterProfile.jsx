import React from "react"
import { ShelterHeader } from "../../components/ShelterHeader"
import InputApp from "../../components/input/input"



export const ShelterProfile = () => {
    return (
        <div className="min-h-screen bg-custom-gradient">
            <ShelterHeader page="profile" />
            <div className="flex justify-center items-center h-[680px]">
                <div className="w-5/6 flex-row p-16 space-y-14 bg-[#F5F9E9] rounded-2xl shadow-lg ">
                    <h2 className="text-3xl  font-bold text-center text-black">Minhas Informações</h2>
                    <div className="w-full flex">

                        <div className="w-1/2 px-6">
                            <div className="w-full mb-3 space-y-1">
                                <label>Nome:</label>
                                <InputApp />
                            </div>
                            <div className="w-full mb-3 space-y-1">
                                <label>Descrição:</label>
                                <InputApp/>
                            </div>
                            <div className="w-full mb-3 space-y-1">
                                <label>Endereço:</label>
                                <InputApp />
                            </div>
                        </div>


                        <div className="w-1/2 px-6">
                            <div className="w-full mb-3 space-y-1">
                                <label>Telefone:</label>
                                <InputApp />
                            </div>
                            <div className="w-full mb-3 space-y-1">
                                <label>Email:</label>
                                <InputApp />
                            </div>
                            <div className="w-full mb-3 space-y-1">
                                <label>Capacidade:</label>
                                <InputApp />
                            </div>
                            <div className="w-full mb-3 space-y-1">
                                <label>Capacidade de pets:</label>
                                <InputApp />
                            </div>
                        </div>
                    </div>
                    <div className="flex justify-center">

                    </div>
                </div>
            </div>
        </div>
    )
}