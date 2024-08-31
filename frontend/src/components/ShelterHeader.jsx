export const ShelterHeader = ({ page }) => {
    return (
        <header className="bg-white/70">
            <nav>
                <ul className="flex p-2 space-x-10">
                    <li className="p-2">
                        <a className={`rounded-xl py-1 px-6 hover:bg-[#b0b0b0] focus:bg-[#a0a0a0] font-bold ${page === 'profile' ? 'bg-[#2929FF] text-white' : 'bg-[#c9c9c9]'}`}
                            href="/perfil-abrigo"
                        >Minha conta</a>
                    </li>
                    <li className="p-2">
                        <a className={`rounded-xl py-1 px-6 hover:bg-[#b0b0b0] focus:bg-[#a0a0a0] font-bold ${page === 'sheltered' ? 'bg-[#2929FF] text-white' : 'bg-[#c9c9c9]'}`} href="/cadastro-abrigados">Cadastrar abrigados</a>
                    </li>
                    <li className="p-2">
                        <a className={`rounded-xl py-1 px-6 hover:bg-[#b0b0b0] focus:bg-[#a0a0a0] font-bold ${page === 'volunteer-list' ? 'bg-[#2929FF] text-white' : 'bg-[#c9c9c9]'}`} href="/banco-de-voluntarios">Banco de voluntários</a>
                    </li>
                </ul>
            </nav>
        </header>
    )
}