export const HomeHeader = () => {
  return (
    <header className='flex justify-between h-10 items-center'>
      <div>
        <span className="rounded-md px-3 py-1 mr-2 bg-slate-50">icon</span>
        <span className="rounded-md px-3 py-1 bg-slate-50">searchbar placeholder</span>
      </div>
      <div>
        <a href="/login" className="rounded-md px-3 py-1 bg-slate-50">Entrar</a>
      </div>
    </header>
  )
}