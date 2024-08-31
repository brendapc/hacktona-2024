export const InputApp = ({placeHolder,onChangeText,value}) => {
    return (
      
    <input value={value} placeholder={placeHolder} onChange={(e)=>onChangeText(e.target.value)} className="w-full drop-shadow-lg rounded-md h-10 px-4 py-2 placeholder-gray-400 border-none shadow-md focus:outline-none focus:ring-2 focus:ring-[#74D480]"/>
      
    )
  }
  export default InputApp;