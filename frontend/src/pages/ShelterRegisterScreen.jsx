import InputApp from "../components/input/input"
import ButtonComponent from "../components/ButtonComponent"

export default function ShelterRegister() {
  const handleTextareaChange = (event) => {
    event.target.style.height = 'auto';
    event.target.style.height = event.target.scrollHeight + 'px';
  };
  return (
      <div className="flex items-center justify-center min-h-screen bg-custom-gradient">
          <div className="w-5/6 flex-row p-16 space-y-14 bg-[#F5F9E9] rounded-2xl shadow-lg">
              <h2 className="text-3xl  font-bold text-center text-black">Cadastre-se como Abrigo</h2>
              <div className="w-full flex">

                  <div className="w-1/2 px-6">
                      <div className="w-full mb-3 space-y-1">
                          <label>Nome do abrigo:</label>
                          <InputApp placeHolder="" />
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
                          <InputApp placeHolder="" />
                      </div>
                      <div className="w-full mb-3 space-y-1">
                          <label>Telefone:</label>
                          <InputApp placeHolder="" />
                      </div>
                  </div>


                  <div className="w-1/2 px-6">
                      <div className="w-full mb-3 space-y-1">
                          <label>Email:</label>
                          <InputApp placeHolder="" />
                      </div>
                      <div className="w-full mb-3 space-y-1">
                          <label>Senha:</label>
                          <InputApp placeHolder="" />
                      </div>
                      <div className="w-full mb-3 space-y-1">
                          <label>Capacidade:</label>
                          <InputApp placeHolder="" />
                      </div>
                      <div className="w-full mb-3 space-y-1">
                          <label>Capacidade para Pets:</label>
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