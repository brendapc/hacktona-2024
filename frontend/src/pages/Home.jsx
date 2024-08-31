import MyMap from '../components/Map';
import { HomeHeader } from '../components/home/HomeHEader';

const Home = () => {
  return (
    <div className='bg-slate-200'>
      <HomeHeader />
      <MyMap />
    </div>
  );
};

export default Home;