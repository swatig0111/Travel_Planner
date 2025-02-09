import React from "react";

const PlaceList = ({ places }) => {
  console.log("Places received in PlaceList:", places);

  return (
    <div className="places-container">
      {places.length === 0 ? (
        <p>No places found.</p>
      ) : (
        <ul>
          {places.map((place, index) => (
  <li key={index}>
    <h3>{place.name ? place.name : "No Name Available"}</h3>
    <p>{place.formatted_address || place.vicinity || "Address not available"}</p>
  </li>
))}

        </ul>
      )}
    </div>
  );
};

export default PlaceList;
