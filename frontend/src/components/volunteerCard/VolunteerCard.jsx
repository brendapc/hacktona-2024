export default function VolunteerCard({ name, type, email }) {
    return (
        <div className="w-1/2 bg-white rounded-md h-14 flex justify-around items-center mb-8">
            <p className="font-bold">{name}</p>
            <p>Tipo: {type}</p>
            <p>Email: {email}</p>
        </div>
    )
}