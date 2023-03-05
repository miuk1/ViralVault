import React from "react";
import Header from './components/header';
import Footer from "./components/footer";
import HomePage from "./components/home";
import VulnerableMachines from "./components/machine";

function App() {
  return (
    <div className="container-fluid">
      <Header />
      <HomePage />
      <VulnerableMachines />
      <Footer />
    </div>
  );
}

export default App;
