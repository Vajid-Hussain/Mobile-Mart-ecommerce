
# Mobile-Mart

Mobile-Mart is a robust RESTful API built using Go and the Gin framework. It serves as the backbone for mobile e-commerce applications, providing essential functionalities such as user management, product listing, and more.

## Features

- **Architecture**: Clean code architecture with Golang for maintainability and scalability.
- **Security**: Twilio for OTP service, access tokens, and refresh tokens.
- **Payment**: Integrated Razorpay for secure payment processing.
- **Storage**: AWS S3 for scalable file storage.
- **Hosting**: AWS EC2 for high availability.
- **CI/CD**: Jenkins pipeline for automated testing and deployment.
- **Containerization**: Docker for efficient deployment.
- **Orchestration**: Kubernetes for container scalability.
- **Documentation**: Swagger for API documentation.
- **Testing**: Comprehensive unit testing for code reliability.

## User Roles

- **Admin**: Manage users and products.
- **Seller**: List products and manage orders.
- **User**: Browse products, add to cart, and make purchases.

## Technologies Used

- **Backend**: Go, Gin framework
- **Database**: PostgreSQL
- **Security**: Twilio (OTP), Access Tokens, Refresh Tokens
- **Payment Gateway**: Razorpay
- **Storage**: AWS S3
- **Hosting**: AWS EC2
- **CI/CD**: Jenkins
- **Containerization**: Docker, Kubernetes
- **Documentation**: Swagger

## Unit Testing

Mobile-Mart is thoroughly tested to ensure code reliability and maintainability. Unit tests cover all critical components and functionalities, providing confidence in the stability of the application.


To clone the repository, run the following commands in your terminal:

```bash
git clone https://github.com/Vajid-Hussain/Mobile-Mart-ecommerce.git
cd Mobile-Mart-ecommerce
```

To run unit tests locally, execute the following command:

```bash
make test ./...
```
