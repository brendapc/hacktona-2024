export const CardButtonComponent = ({ onClick, label, title, image }) => {
    return (
        <button
            onClick={onClick}
            className="w-64 h-80 rounded-lg bg-[#F5FCE8] overflow-hidden shadow-lg border text-center flex flex-col items-center justify-between p-2"
        >
            <div className="text-lg font-semibold">{title}</div>
            <div className="flex-grow flex items-center justify-center">
                {image && <img src={image} alt={title} className="w-16 h-16 object-cover" />}
            </div>
            <div className="text-sm">{label}</div>
        </button>
    );
};
