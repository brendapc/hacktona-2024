import ButtonComponent from "../../components/ButtonComponent"
import InputApp from "../../components/input/input"
import { ShelterHeader } from "../../components/ShelterHeader"

export const ShelteredForm = () => {
    return (
        <div className="bg-custom-gradient w-screen h-screen overflow-auto" >
            <ShelterHeader page="sheltered" />
            <div className="flex items-center justify-center h-[800px]">
                <div className="w-[880px] h-[600px] p-16 bg-[#F5F9E9] rounded-2xl shadow-lg">
                    <h2 className="text-xl font-bold">Adicionar abrigado</h2>
                    <div className="mt-12 mb-28 space-y-2">
                        <label htmlFor="">Nome completo</label>
                        <div className="flex space-x-6">
                            <InputApp placeHolder={"João da Silva"} />
                            <ButtonComponent label={"Adicionar"} />
                        </div>
                    </div>
                    <h2 className="text-xl font-bold">Remover abrigado</h2>
                    <div className="mt-12 space-y-2">
                        <label htmlFor="">Nome completo</label>
                        <div className="flex space-x-6">
                            <InputApp placeHolder={"João da Silva"} />
                            <ButtonComponent label={"Remover"} />
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}