export default function VolunteerCard({name,tipe,email}){
    return(
        <div className="w-1/2 bg-white rounded-md h-14 flex justify-around items-center">
            <p className="font-bold">{name}</p>
            <p>Tipo: {tipe}</p>
            <p>Email: {email}</p>
        </div>
    )
}