import React, { useState } from "react";
import axios from "axios";
import "./App.css";
import Destination from "./components/Destination";
import PlaceList from "./components/PlaceList";


const App = () => {
  const [destinationPlaces, setDestinationPlaces] = useState([]);
 
  const [nearbyPlaces, setNearbyPlaces] = useState([]);

  const fetchPlaces = async (destination) => {
    try {
      const response = await axios.get(
        `${process.env.REACT_APP_BACKEND_URL}/places?destination=${destination}`
      );
      console.log("API Response:", response.data);
      console.log("Destination Data:", response.data.destination);
      console.log("Nearby Data:", response.data.nearby);

      // Check and set destination places
      if (response.data.destination) {
        setDestinationPlaces(response.data.destination);
      } else {
        setDestinationPlaces([]);
      }

      if (response.data.nearby) {
        setNearbyPlaces(response.data.nearby);
      } else {
        setNearbyPlaces([]);
      }

    } catch (error) {
      console.error("Error fetching places:", error);
      setDestinationPlaces([]);
      setNearbyPlaces([]);
    }

  };

  return (
    <div className="app-container">
      <Destination fetchPlaces={fetchPlaces} />
      <h2>Destination Places</h2>
      <PlaceList places={destinationPlaces} />
      
      <h2>Nearby Places</h2>
      <PlaceList places={nearbyPlaces} />
    </div>
  );
};

export default App;
