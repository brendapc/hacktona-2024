import { useState, useEffect } from 'react';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import L from 'leaflet';
import axios from 'axios';

const createCustomIcon = (color, myLocation) => {
  return new L.Icon({
    iconUrl: myLocation ? `data:image/svg+xml;utf8,<svg width="24" height="24" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><circle cx="12" cy="12" r="10" stroke="black" stroke-width="1" fill="${color}"/></svg>` : `data:image/svg+xml;utf8,<svg width="24" height="24" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path d="M12 2L2 12h3v8h6v-6h2v6h6v-8h3z" fill="${color}" stroke="black" stroke-width="1"/></svg>`,
    iconSize: [24, 24],
    iconAnchor: [12, 24],
    popupAnchor: [0, -24]
  });
};

const geocodeAddress = async (address) => {
  try {
    const response = await fetch(`https://nominatim.openstreetmap.org/search?q=${encodeURIComponent(address)}&format=json&addressdetails=1`);
    const data = await response.json();
    if (data.length > 0) {
      const { lat, lon } = data[0];
      return { lat: parseFloat(lat), lng: parseFloat(lon) };
    } else {
      throw new Error('Address not found');
    }
  } catch (error) {
    console.error('Error geocoding address:', error);
    return null;
  }
};

function MyMap() {
  const [locations, setLocations] = useState([]);
  const [userLocation, setUserLocation] = useState(null);
  const [shelters, setShelters] = useState([]);

  useEffect(() => {
    const fetchAndGeocodeAddresses = async () => {
      const response = await axios.get("http://localhost:8080/api/v1/shelters");
      const addresses = response.data.map((shelter) => ({ ...shelter, color: 'green' }));
  
  
      setShelters([...addresses]);
    };
  
    fetchAndGeocodeAddresses();
  }, []);  // Executa apenas uma vez ao montar o componente
  
  useEffect(() => {
    const geocodeShelters = async () => {
      try {
        const locationPromises = shelters.map(async (loc) => {
          const coords = await geocodeAddress(loc.address);
          if (coords) {
            return { id: loc.id, name: loc.name, address: loc.address, ...coords, color: loc.color, capacity: loc.capacity, current_occupancy: loc.current_occupancy};
          }
          return null;
        });
        const locations = (await Promise.all(locationPromises)).filter(Boolean);
        setLocations(locations);
      } catch (error) {
        console.error('Error fetching and geocoding addresses:', error);
      }
    };
  
    if (shelters.length > 0) {
      geocodeShelters();
    }
  }, [shelters]);  // Executa toda vez que `shelters` for atualizado
  

  return (
    <MapContainer center={userLocation || [-30.0586, -51.1756]} zoom={13} style={{ height: '80vh', width: '100%' }}>
      <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      />
      {userLocation && (
        <Marker position={userLocation} icon={createCustomIcon('blue', true)}>
          <Popup>You are here</Popup>
        </Marker>
      )}
      {locations.map((loc) => (
        <Marker
          key={loc.id}
          position={[loc.lat, loc.lng]}
          icon={createCustomIcon(loc.color)}
        >
          <Popup>
            <div style={{ width: '200px', height: 'auto' }}>
              <h3>{loc.name}</h3>
              <p>{loc.address}</p>
              <p>Vagas restantes: {loc.capacity-loc.current_occupancy}</p>
              {/* You can include more content here */}
            </div>
          </Popup>
        </Marker>
      ))}
    </MapContainer>
  );
}

export default MyMap;