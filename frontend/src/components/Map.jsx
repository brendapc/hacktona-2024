import { useState, useEffect } from 'react';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import L from 'leaflet';

// Function to create custom icons
const createCustomIcon = (color, myLocation) => {
  return new L.Icon({
    iconUrl: myLocation ? `data:image/svg+xml;utf8,<svg width="24" height="24" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><circle cx="12" cy="12" r="10" stroke="black" stroke-width="1" fill="${color}"/></svg>` :  `data:image/svg+xml;utf8,<svg width="24" height="24" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path d="M12 2L2 12h3v8h6v-6h2v6h6v-8h3z" fill="${color}" stroke="black" stroke-width="1"/></svg>`,
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

  useEffect(() => {
    const fetchAndGeocodeAddresses = async () => {
      // Simulated fetch from backend
      const addresses = [
        { id: 1, name: 'Local 1', address: 'Av. Ipiranga, 7060 - Partenon, Porto Alegre', color: 'yellow' },
        { id: 2, name: 'Local 2', address: 'Av. Ipiranga, 6681 - Partenon, Porto Alegre', color: 'yellow' },
        { id: 3, name: 'Local 3', address: 'Praça da Matriz, Porto Alegre, Brazil', color: 'yellow' }
      ];

      try {
        const locationPromises = addresses.map(async (loc) => {
          const coords = await geocodeAddress(loc.address);
          if (coords) {
            return { id: loc.id, name: loc.name, ...coords, color: loc.color };
          }
          return null;
        });
        const locations = (await Promise.all(locationPromises)).filter(Boolean);
        setLocations(locations);
      } catch (error) {
        console.error('Error fetching and geocoding addresses:', error);
      }
    };
    fetchAndGeocodeAddresses();
  }, []);

  useEffect(() => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition((position) => {
        const { latitude, longitude } = position.coords;
        setUserLocation([latitude, longitude]);
      });
    }
  }, []);

  return (
    <MapContainer center={userLocation || [-30.0586, -51.1756]} zoom={13} style={{ height: '89vh', width: '100%' }}>
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
              <p>Click here for more details...</p>
              <a href='abrigo-X'>go to</a>
              {/* You can include more content here */}
            </div>
          </Popup>
        </Marker>
      ))}
    </MapContainer>
  );
}

export default MyMap;