import React from 'react';
import { CardButtonComponent } from '../components/CardButtonComponent';
// import { RegisterVolunteer } from './RegisterVolunteer';

export const RegisterScreen = () => {

    return (
        <div className="flex flex-col h-screen bg-custom-gradient">
                <h1 className="text-3xl font-semibold mt-10 text-center">Cadastrar</h1>
                    <div className="flex h-full items-center justify-center place-content-center space-x-12">
                        <CardButtonComponent onClick={() => console.log('Voluntário')} title='Voluntário'/>
                        <CardButtonComponent onClick={() => console.log('Empresa')} title='Empresa'/>
                        <CardButtonComponent onClick={() => console.log('Abrigo')} title='Abrigo' image='https://images03.brasildefato.com.br/ba2b3042aea1c53aeb6429a7e63a0719.webp'/>
                </div>
        </div>
    );
}