const ButtonComponent = ({ onClick, label }) => {
    return (
        <button 
            onClick={onClick}
            className="w-1/2 h-12 py-2 bg-[#6060D4] text-white rounded-md shadow-lg border-none cursor-pointer focus:outline-none focus:ring-2 focus:ring-[#6060D4] focus:ring-opacity-50 hover:bg-[#4E4EB8]"  
        >
            {label}
        </button>
    );
};

export default ButtonComponent;
