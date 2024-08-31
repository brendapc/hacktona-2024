export const InputApp = ({placeHolder,onChangeText}) => {
    return (
      
    <input placeholder={placeHolder} onChange={(e)=>onChangeText(e.target.value)} className="w-5/6 drop-shadow-lg rounded-md p-2 mx-2"/>
      
    )
  }