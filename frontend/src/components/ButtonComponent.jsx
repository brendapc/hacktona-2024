import React from 'react';


const ButtonComponent = ({ onClick, disabled = false, children, className = '' }) => {
    return (
        <button
            
            onClick={onClick}
            disabled={disabled}
            className="W-40 h-7 bg-[#007bff] text-black rounded-md text-xs[12px] font-regular p-8 border-none cursor-pointer"  
        >
        </button>
    );
};

export default ButtonComponent;
