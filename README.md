# Event Booking REST API

## Overview
This is a Go-powered REST API designed for managing event bookings. It provides functionalities for users to sign up, log in, manage events, and register for events. The API uses JSON Web Tokens (JWT) for authentication and enforces specific rules for creating, updating, and deleting events.

---

## Features
- User authentication via JWT.
- CRUD operations for events (Create, Read, Update, Delete).
- Event registration and cancellation.
- Access control for event management (only creators can modify or delete their events).

---

## Endpoints

### Public Endpoints
- **GET /events**  
  Fetch a list of all available events.

- **GET /events/{id}**  
  Fetch details of a specific event by its ID.

- **POST /signup**  
  Register a new user.

- **POST /login**  
  Authenticate an existing user and retrieve a JWT token.

---

### Authenticated Endpoints
> **Note**: Authentication is required for the following endpoints using a JWT token.

- **POST /events**  
  Create a new bookable event.  
  **Access**: Authenticated users only.

- **PUT /events/{id}**  
  Update an existing event.  
  **Access**: Only the creator of the event.

- **DELETE /events/{id}**  
  Delete an existing event.  
  **Access**: Only the creator of the event.

- **POST /events/{id}/register**  
  Register a user for an event.  
  **Access**: Authenticated users only.

- **DELETE /events/{id}/register**  
  Cancel a user’s registration for an event.  
  **Access**: Authenticated users only.

---

## Authentication
This API uses JSON Web Tokens (JWT) for user authentication. After logging in, users receive a JWT token that must be included in the `Authorization` header of all authenticated requests.
