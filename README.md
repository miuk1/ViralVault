

Viral Vault

Infect your skills, not your network! Overview

Viral Vault is a web application designed to help users practice and improve their penetration testing skills in a safe and controlled environment. The application provides a list of vulnerable machines that users can deploy and access directly from the website, allowing them to practice their skills without putting real networks or systems at risk.

The application consists of a backend API built with the Go programming language, and a frontend web interface built with React. The backend uses a Postgresql database to store information about the vulnerable machines and user accounts, and implements authentication and validation using Go-AWS-Auth and govalidator, respectively. Getting Started Prerequisites

To get started with Viral Vault, you'll need to have the following software installed on your system:

Go (version 1.16 or later)
Node.js (version 14 or later)
Docker (version 19 or later)
Docker Compose (version 1.28 or later)

You can download and install these tools from their respective websites, or by using a package manager such as Homebrew (for macOS and Linux) or Chocolatey (for Windows). Installation

Clone the Viral Vault repository from GitHub:git clone https://github.com/<username> viral-vault.git
    
Replace <username> with your GitHub username.

Navigate to the project directory and build the Docker images:

docker-compose build

Start the application:

docker-compose up

This will start the backend API and the frontend web interface, and make the application available at http://localhost:3000.

Contributing

If you would like to contribute to Viral Vault, you can fork the repository and submit a pull request with your changes. Before submitting a pull request, please make sure that your code follows the coding standards and guidelines used in the project. License

Viral Vault is licensed under the MIT License. See the LICENSE file for more information. Contact

If you have any questions or feedback about Viral Vault, please feel free to contact us at contact@viralvault.com.

