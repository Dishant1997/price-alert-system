import React, { useState } from "react";
import axios from "axios";

const AlertForm = () => {
  const [crypto, setCrypto] = useState("");
  const [targetPrice, setTargetPrice] = useState("");
  const [email, setEmail] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post("http://localhost:8080/alerts", {
        crypto,
        targetPrice: parseFloat(targetPrice),
        email,
      });
      alert("Alert created successfully!");
    } catch (error) {
      console.error("Error creating alert:", error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        placeholder="Cryptocurrency (e.g., BTC)"
        value={crypto}
        onChange={(e) => setCrypto(e.target.value)}
      />
      <input
        type="number"
        placeholder="Target Price"
        value={targetPrice}
        onChange={(e) => setTargetPrice(e.target.value)}
      />
      <input
        type="email"
        placeholder="Your Email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />
      <button type="submit">Create Alert</button>
    </form>
  );
};

export default AlertForm;