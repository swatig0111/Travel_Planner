import React, { useState } from "react";
import "./Destination.css";

const Destination = ({ fetchPlaces }) => {
  const [destination, setDestination] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    if (destination.trim()) {
      fetchPlaces(destination);
    }
  };

  return (
    <div className="container" >
      <h2>Travel Planner</h2>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Enter destination..."
          value={destination}
          onChange={(e) => setDestination(e.target.value)}
        />
        <button type="submit">Search</button>
      </form>
    </div>
  );
};

export default Destination;
