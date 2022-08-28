# Setup

## Set environment variables

- `PORT`: Defines the port the server will be using
- `POSTGRES_USER`: Database username
- `POSTGRES_PASSWORD`: Database password
- `POSTGRES_DB`: Name of the database
- `DATABASE_HOST`: Hostname where the database is hosted
- `DATABASE_PORT`: Port where the database is hosted
- `PASSWORD_PEPPER`: Sets the password pepper (secret) for hashing the passwords
- `SIGNATURE`: Signature secret for JWT signing
- `JWT_COOKIE_NAME (optional)`: Sets the name for the JWT token cookie's name. `(default: token)`

# Progress

## Features checklist

- [ ] Public/private keys
- [ ] Authentication
- [ ] Sessions
- [ ] Database models
  - [ ] Products
  - [ ] Categories
  - [ ] Users
  - [ ] Cart
  - [ ] Discounts
  - [ ] Addresses
  - [ ] Orders
  - [ ] Payment methods
- [ ] Public API functionalities
  - [x] Consistent error handling pattern
  - [x] Get products list
  - [ ] Get categories list
  - [x] Register new user
  - [ ] Get user data
  - [ ] Update user data
  - [ ] Delete account
  - [ ] Add item to cart
  - [ ] Delete item from cart
  - [ ] Update item in cart
  - [ ] Apply discounts to cart
  - [ ] Create new address
  - [ ] Update address
  - [ ] Delete address
  - [ ] Create new order (purchase products)
  - [ ] Cancel order
  - [ ] Add payment method
  - [ ] Update payment method
  - [ ] Delete payment method
