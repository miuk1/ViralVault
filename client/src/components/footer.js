import React from 'react';
import logo from '../logo.png';

function Footer() {
    return (
      <footer className="bg-light text-center text-lg-start">
        <div className="text-center p-3">
          <img
            src={logo}
            height="30"
            className="d-inline-block align-top"
            alt="ViralVault logo"
          />
          ViralVault © 2023 All Rights Reserved.
        </div>
      </footer>
    );
  }

  export default Footer;