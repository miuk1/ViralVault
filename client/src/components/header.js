import React from 'react';
import { Navbar, Nav, Container } from 'react-bootstrap';
import logo from '../logo.png';
import '../style/header.css';

function Header() {
  return (
    <Navbar bg="light" expand="lg">
      <Container>
        <Navbar.Brand href="/">
          <img
            src={logo}
            height="40"
            className="d-inline-block align-top"
            alt="ViralVault logo"
          />
          <span className="ms-3 align-middle">ViralVault</span>
        </Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="ms-auto">
            <Nav.Link href="/">Home</Nav.Link>
            <Nav.Link href="/about">About</Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}

export default Header;