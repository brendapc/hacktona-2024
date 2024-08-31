export const ShelterHeader = () => {
    return (
        <header className="bg-white/70">
            <nav>
                <ul className="flex p-2 space-x-10">
                    <li className="p-2">
                        <a className="bg-[#c9c9c9] rounded-xl py-1 px-6 hover:bg-[#b0b0b0] focus:bg-[#a0a0a0] font-bold" href="/minha-conta">Minha conta</a>
                    </li>
                    <li className="p-2">
                        <a className="bg-[#c9c9c9] rounded-xl py-1 px-4 hover:bg-[#b0b0b0]  focus:bg-[#a0a0a0]  font-bold" href="/cadastrar-abrigados">Cadastrar abrigados</a>
                    </li>
                    <li className="p-2">
                        <a className="bg-[#c9c9c9] rounded-xl py-1 px-4 hover:bg-[#b0b0b0] focus:bg-[#a0a0a0 font-bold" href="/banco-de-voluntarios">Banco de voluntários</a>
                    </li>
                </ul>
            </nav>
        </header>
    )
}