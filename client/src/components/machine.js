import React, { useState, useEffect } from "react";
import { Card, Button, Modal } from "react-bootstrap";

function VulnerableMachines() {
  const [machines, setMachines] = useState([]);
  const [currentMachine, setCurrentMachine] = useState(null);
  const [show, setShow] = useState(false);

  // Fetch vulnerable machines from API on component mount
  useEffect(() => {
    fetch("http://localhost:8000/api/vulnerable-machines", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      }
    })
      .then((response) => response.json())
      .then((data) => setMachines(data))
      .catch((error) => console.log(error));
  }, []);

  // Show details modal when a card is clicked
  const handleCardClick = (machine) => {
    setCurrentMachine(machine);
    setShow(true);
  };

  return (
    <>
      <h1 className="text-center mb-4">Vulnerable Machines</h1>
      <div className="row">
        {machines.map((machine) => (
          <div className="col-md-4 mb-4" key={machine.ID}>
            <Card>
              <Card.Body>
                <Card.Title>{machine.name}</Card.Title>
                <Card.Text>{machine.description}</Card.Text>
                <Button onClick={() => handleCardClick(machine)}>
                  View Details
                </Button>
              </Card.Body>
            </Card>
          </div>
        ))}
      </div>

      <Modal show={show} onHide={() => setShow(false)}>
        <Modal.Header closeButton>
          <Modal.Title>{currentMachine && currentMachine.name}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          {currentMachine && (
            <>
              <p>{currentMachine.description}</p>
              <p>OS: {currentMachine.os}</p>
              <p>Docker Image: {currentMachine.dockerimage}</p>
              <h5>Vulnerabilities:</h5>
              {currentMachine.vulnerabilities.map((vulnerability) => (
                <p key={vulnerability.ID}>
                  {vulnerability.name} ({vulnerability.difficulty})
                </p>
              ))}
            </>
          )}
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={() => setShow(false)}>
            Close
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );
}

export default VulnerableMachines;
