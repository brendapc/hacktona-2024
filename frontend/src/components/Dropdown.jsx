/* eslint-disable react/prop-types */

const Dropdown = ({ options, selectedOption, setSelectedOption }) => {
  return (
    <select 
      value={selectedOption} 
      onChange={(e) => setSelectedOption(e.target.value)}
      className="border rounded px-4 py-1 ml-1"
    >
      {options.map((option, index) => (
        <option key={index} value={option}>
          {option}
        </option>
      ))}
    </select>
  );
};

export default Dropdown;