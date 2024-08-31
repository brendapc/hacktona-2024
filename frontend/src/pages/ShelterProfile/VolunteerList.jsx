import VolunteerCard from "../../components/volunteerCard/VolunteerCard";
import { ShelterHeader } from "../../components/ShelterHeader"


export const VolunteerList = () => {

    const mockVolunteers = [
        { name: "João Silva", type: "Voluntário", email: "joao.silva@example.com" },
        { name: "Maria Oliveira", type: "Coordenador", email: "maria.oliveira@example.com" },
        { name: "Carlos Souza", type: "Assistente", email: "carlos.souza@example.com" }
    ];

    return (
        <div className="bg-custom-gradient h-screen w-screen">
            <ShelterHeader page={"volunteer-list"}/>
            <div className="pt-8 ml-5">
                <div>
                    <h1 className="text-3xl font-bold mb-8">Voluntários cadastrados:</h1>
                </div>
                {mockVolunteers.map((volunteer, index) => (
                    <VolunteerCard
                        key={index}
                        name={volunteer.name}
                        type={volunteer.type}
                        email={volunteer.email}
                    />

                ))}
            </div>
        </div>
    )
}