import React from 'react';
import '../style/home.css';
import heroImage from '../images/hero-image.jpg';

const HomePage = () => {
    return (
        <div>
            {/*Hero Section */}
            <div className="hero-section" style={{ backgroundImage: `url(${heroImage})` }}>
                <div className="hero-overlay">
                    <div className="hero-content">
                        <h1>Infect Your Skills, Not Your Network</h1>
                        <p>Practice and improve your penetration testing skills in a safe and controlled environment.</p>
                        <button className="btn btn-primary">Get Started</button>
                    </div>
                </div>
            </div>


            {/* Features Section */}
            <section className="features-section py-5">
            <h1 className="text-center mb-4">Features</h1>
                <div className="container">
                    <div className="row">
                        <div className="col-md-4">
                            <div className="feature-item">
                                <div className="feature-icon">
                                    <i className="fas fa-shield-alt"></i>
                                </div>
                                <h3>Secure Practice Environment</h3>
                                <p>Practice penetration testing in a safe and controlled environment without the risk of harming real networks or systems.</p>
                            </div>
                        </div>
                        <div className="col-md-4">
                            <div className="feature-item">
                                <div className="feature-icon">
                                    <i className="fas fa-code"></i>
                                </div>
                                <h3>Realistic Scenarios</h3>
                                <p>Deploy and access a variety of vulnerable machines directly from the website to practice your skills.</p>
                            </div>
                        </div>
                        <div className="col-md-4">
                            <div className="feature-item">
                                <div className="feature-icon">
                                    <i className="fas fa-certificate"></i>
                                </div>
                                <h3>Certification</h3>
                                <p>Earn a certificate of completion for completing challenges and demonstrating your penetration testing skills.</p>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
        </div>
    );
}

export default HomePage;
