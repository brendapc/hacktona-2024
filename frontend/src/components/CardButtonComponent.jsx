export const CardButtonComponent = ({ onClick, label, title, image }) => {
    return (
        <button
            onClick={onClick}
            className="w-80 h-96 rounded-3xl bg-[#F5FCE8] overflow-hidden shadow-lg border text-center flex flex-col items-center justify-between p-2"
        >
            <div className="text-lg font-regular">{title}</div>
            <div className="flex-grow flex items-center justify-center">
                {image && <img src={image} alt={title} className="w- h- object-cover" />}
            </div>
            <div className="text-sm">{label}</div>
        </button>
    );
};
