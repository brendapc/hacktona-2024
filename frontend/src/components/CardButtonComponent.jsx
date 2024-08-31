export const CardButtonComponent = ({ onClick, label, title, image }) => {
    return (
      <button
        onClick={onClick}
        className="w-full max-w-xs h-auto md:w-80 md:h-96 rounded-3xl bg-[#F5FCE8] overflow-hidden shadow-lg border text-center flex flex-col items-center justify-between p-4 transition-transform duration-300 transform hover:scale-105"
      >
        <div className="text-lg font-medium mt-4">{title}</div>
        
        <div className="flex-grow flex items-center justify-center mt-2 mb-4">
          {image && <img src={image} alt={title} className="w-3/4 h-auto object-contain" />}
        </div>
  
        <div className="text-sm mb-4">{label}</div>
      </button>
    );
  };
  