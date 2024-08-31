import MyMap from '../components/Map';
import { HomeHeader } from '../components/home/HomeHEader';

const Home = () => {
  return (
    <div className="bg-custom-gradient h-screen">
      <HomeHeader />
      <main>
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