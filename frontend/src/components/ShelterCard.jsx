import React from 'react';

export const ShelterCard = ({ name, address, remainingSpots, needs, description, value }) => {
    return (
        <div className="bg-white shadow-xl rounded-lg p-4 max-w-[1200px] border-green-800 mb-10">
            <h2 className="text-green-800 text-2xl font-bold mb-2">{name}</h2>
            <p className="text-gray-700 mb-2">{address}</p>
            <p className="text-gray-700 mb-2">
                Vagas restantes: <span className="font-bold">{remainingSpots}</span>
            </p>
            <div className="flex items-center mb-2">
                <span className="text-gray-700 mr-2">Necessidades:
                    <ul className="list-disc text-gray-700 ml-2 pl-8">
                        {needs.map((need, index) => (
                            <li key={index}
                            className={need.toLowerCase() === value ? 'font-bold' : ''}
                            >{need}</li>
                        
                        ))}
                    </ul>
                </span>

            </div>
            <p className="text-gray-700">{description}</p>
        </div>
    );
};
