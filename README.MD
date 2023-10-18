# Work in Progress Backend Template Go Gin-Gonic

Please note that this backend template is currently under development, and I actively working on adding more features and improvements to it. Stay tuned for future updates and enhancements that will further enhance its functionality and usability.

A straightforward backend template built with Go Gin-Gonic, designed to kickstart your backend development with ease. It comes packed with essential features, including JSON Web Token (JWT) authentication, user management, and a containerized PostgreSQL database. The entire stack is containerized with Docker, simplifying deployment.

## Key Features

- **Gin-Gonic**: Harness the power of Gin-Gonic, a high-performance web framework for building APIs with Go.

- **GORM ORM**: Utilize GORM as the ORM to interact with the PostgreSQL database, enabling efficient database operations and data modeling. (Next Update)

- **Goroutines**: Leverage goroutines and channels for concurrent operations, making your application more efficient. (Soon)

- **Authentication**: Secure your APIs with JSON Web Tokens (JWT) to ensure robust user authentication and authorization. (Soon)

- **User Management**: Easily manage user accounts, including registration, login, and profile management.

- **Database**: Utilize a containerized PostgreSQL database to store your application's data, ensuring data integrity and scalability.

- **Dockerized**: Simplify deployment with Docker and Docker Compose. Run `docker-compose up`, and your entire application stack is up and running.

## Getting Started

1. Clone this repository.

2. Inside the `app` folder, create the following `.env` files for configuration:

   - Create `auth.env` (For JWT secret key):

     ```plaintext
     JWT_SECRET_KEY=09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7
     ```

   - Create `service.env` (For Service config):

     ```plaintext
     SERVICE_PORT=8100
     SERVICE_HOST=0.0.0.0
     SERVICE_WORKERS=4
     ```

   - Create `config.env` (For Database config):
     ```plaintext
     POSTGRES_USER=postgres
     POSTGRES_DB=rest_template
     POSTGRES_PASSWORD=default15p
     POSTGRES_HOST=postgres_rest_template_go
     POSTGRES_PORT=8103
     ```

   Feel free to change the values in the `.env` files based on your needs.

3. On the root directory, run the following command to start the development server along with the database. Make sure you have docker installed:
   ```bash
   docker-compose up
   ```

This Template is on progress, more updates will come
