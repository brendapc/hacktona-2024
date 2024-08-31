import React, { useState, useEffect } from 'react';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import L from 'leaflet';

// Function to create custom icons
const createCustomIcon = (color) => {
  return new L.Icon({
    iconUrl: `data:image/svg+xml;utf8,<svg width="24" height="24" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><circle cx="12" cy="12" r="10" stroke="black" stroke-width="2" fill="${color}" /></svg>`,
    iconSize: [24, 24],
    iconAnchor: [12, 24],
    popupAnchor: [0, -24]
  });
};

// Function to get latitude and longitude from an address using Nominatim
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

  useEffect(() => {
    const fetchAndGeocodeAddresses = async () => {
      // Simulated fetch from backend
      const addresses = [
        { id: 1, name: 'Local 1', address: 'Av. Ipiranga, 7060 - Partenon, Porto Alegre', color: 'red' },
        { id: 2, name: 'Local 2', address: 'Av. Ipiranga, 6681 - Partenon, Porto Alegre', color: 'blue' },
        { id: 3, name: 'Local 3', address: 'Praça da Matriz, Porto Alegre, Brazil', color: 'green' }
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

  return (
    <MapContainer center={[-30.0346, -51.2177]} zoom={13} style={{ height: '100vh', width: '100%' }}>
      <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      />
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
              <a href='abrigo-X' >go to </a>
              {/* You can include more content here */}
            </div>
          </Popup>
        </Marker>
      ))}
    </MapContainer>
  );
}

export default MyMap;
