import ButtonComponent from "../components/ButtonComponent";
import InputApp from "../components/input/input";

export default function CompanyRegister() {
    return (
        <div className="flex items-center justify-center min-h-screen bg-custom-gradient">
            <div className="w-5/6 flex-row p-16 space-y-14 bg-[#F5F9E9] rounded-2xl shadow-lg">
                <h2 className="text-3xl  font-bold text-center text-black">Cadastre-se como empresa</h2>
                <div className="w-full flex">

                    <div className="w-1/2 px-6">
                        <div className="w-full mb-3 space-y-1">
                            <label>Nome da empresa:</label>
                            <InputApp placeHolder="" />
                        </div>
                        <div className="w-full mb-3 space-y-1">
                            <label>CNPJ:</label>
                            <InputApp placeHolder="" />
                        </div>
                        <div className="w-full mb-3 space-y-1">
                            <label>Endereço:</label>
                            <InputApp placeHolder="" />
                        </div>
                    </div>


                    <div className="w-1/2 px-6">
                        <div className="w-full mb-3 space-y-1">
                            <label>Telefone:</label>
                            <InputApp placeHolder="" />
                        </div>
                        <div className="w-full mb-3 space-y-1">
                            <label>Email:</label>
                            <InputApp placeHolder="" />
                        </div>
                        <div className="w-full mb-3 space-y-1">
                            <label>Senha:</label>
                            <InputApp placeHolder="" />
                        </div>
                    </div>
                </div>
                <div className="flex justify-center">
                    <ButtonComponent label={"Cadastrar"} />
                </div>
            </div>
        </div>
    )
}