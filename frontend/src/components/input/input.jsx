import { FaEye, FaEyeSlash } from 'react-icons/fa';
import { useState } from 'react';

export const InputApp = ({ placeHolder, onChangeText, value, type = "text" }) => {
  const [isHidden, setIsHidden] = useState(true);

  const toggleVisibility = () => {
    setIsHidden(!isHidden);
  };

  return (
    <div className="relative w-full">
      <input
        type={isHidden && type === "password" ? "password" : "text"}
        value={value}
        placeholder={placeHolder}
        onChange={(e) => onChangeText(e.target.value)}
        className="w-full drop-shadow-lg rounded-md h-10 px-4 py-2 placeholder-gray-400 border-none shadow-md focus:outline-none focus:ring-2 focus:ring-[#74D480] pr-10"
      />
      {type === 'password' && (
        <div
          className="absolute inset-y-0 right-0 flex items-center pr-3 cursor-pointer"
          onClick={toggleVisibility}
        >
          {isHidden ? (
            <FaEyeSlash className="text-gray-600" size={20} />
          ) : (
            <FaEye className="text-gray-600" size={20} />
          )}
        </div>
      )}
    </div>
  );
};

export default InputApp;
