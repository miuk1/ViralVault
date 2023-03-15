import React from 'react';
import { Navbar, Nav, Container, Button } from 'react-bootstrap';
import logo from '../logo-nav.png';
import '../style/header.css';

function Header() {
  return (
    <Navbar bg="light" expand="lg" className="sticky-top">
      <Container>
        <div class="d-flex align-items-center">
          <Navbar.Brand href="/" className="d-flex align-items-center">
            <img
              src={logo}
              width="60"
              height="60"
              className="d-inline-block align-top"
              alt="ViralVault logo"
            />{' '}
            <span className="ms-2">ViralVault</span> 
          </Navbar.Brand>
        </div>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav" className="justify-content-end">
            <Nav>
            <Button variant="outline-primary" href="/signin" className="nav-link">Sign In</Button>
              <Nav.Link href="/about">About</Nav.Link>
            </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}

export default Header;