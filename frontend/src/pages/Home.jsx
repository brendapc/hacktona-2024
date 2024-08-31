import MyMap from '../components/Map';
import { HomeHeader } from '../components/home/HomeHEader';

const Home = () => {
  return (
    <div className="relative bg-slate-200">
      <HomeHeader />
      <main className="pt-20">
        <div className="container mx-auto px-4">
          <div className="rounded-xl overflow-hidden shadow-lg border border-gray-200">
            <MyMap />
          </div>
        </div>
      </main>
    </div>
  );
};

export default Home;