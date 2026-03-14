-- ============================================
-- EXTENSIONS
-- ============================================

-- UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Cryptography functions (optional, for password hashing)
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- PostGIS for spatial data (GPS locations)
CREATE EXTENSION IF NOT EXISTS postgis;

-- ============================================
-- ENUM TYPES
-- ============================================

-- User roles in the system
CREATE TYPE user_role AS ENUM (
    'customer',
    'driver',
    'admin'
);

-- Status of a booking
CREATE TYPE booking_status AS ENUM (
    'pending',
    'accepted',
    'arriving',
    'started',
    'completed',
    'cancelled'
);

-- Payment status
CREATE TYPE payment_status AS ENUM (
    'pending',
    'paid',
    'failed'
);

-- Status of a car
CREATE TYPE car_status AS ENUM (
    'available',
    'busy',
    'offline',
    'maintenance'
);

-- ============================================
-- USERS
-- ============================================

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    phone TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE,
    password_hash TEXT NOT NULL,
    role user_role DEFAULT 'customer',
    profile_image TEXT,
    version INT DEFAULT 1,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP
);

-- Covering index for fast queries by phone
CREATE INDEX idx_users_phone_cover
ON users(phone, id, name, role);

-- ============================================
-- DRIVERS
-- ============================================

CREATE TABLE drivers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    license_number TEXT UNIQUE NOT NULL,
    rating NUMERIC(2,1) DEFAULT 5.0,
    total_reviews INT DEFAULT 0,
    is_online BOOLEAN DEFAULT false,
    version INT DEFAULT 1,
    created_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP
);

-- Covering index for fast lookup by user_id
CREATE INDEX idx_drivers_user_cover
ON drivers(user_id, id, rating, is_online);

-- ============================================
-- CARS
-- ============================================

CREATE TABLE cars (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    driver_id UUID REFERENCES drivers(id),
    brand TEXT NOT NULL,
    model TEXT NOT NULL,
    year INT,
    plate_number TEXT UNIQUE,
    color TEXT,
    status car_status DEFAULT 'available',
    version INT DEFAULT 1,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP
);

-- Covering index for fast lookup of cars by driver
CREATE INDEX idx_cars_driver_cover
ON cars(driver_id, id, brand, model, status);

-- Index for car status queries
CREATE INDEX idx_cars_status
ON cars(status);

-- Partial index for only available cars (optimizes driver assignment)
CREATE INDEX idx_cars_available
ON cars(status, id, driver_id)
WHERE status='available';

-- ============================================
-- CAR IMAGES (FILE SYSTEM STORAGE)
-- ============================================

CREATE TABLE car_images (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    car_id UUID REFERENCES cars(id) ON DELETE CASCADE,
    image_path TEXT NOT NULL,
    is_primary BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT now()
);

CREATE INDEX idx_car_images_car
ON car_images(car_id);

-- ============================================
-- DRIVER LOCATIONS (REAL-TIME GPS)
-- ============================================

CREATE TABLE driver_locations (
    driver_id UUID PRIMARY KEY REFERENCES drivers(id),
    location GEOGRAPHY(Point,4326),
    updated_at TIMESTAMP DEFAULT now()
);

-- GiST index for spatial queries (nearest driver)
CREATE INDEX idx_driver_location_geo
ON driver_locations USING GIST(location);

-- ============================================
-- BOOKINGS
-- ============================================

CREATE TABLE bookings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    driver_id UUID REFERENCES drivers(id),
    car_id UUID REFERENCES cars(id),
    pickup_location GEOGRAPHY(Point,4326),
    dropoff_location GEOGRAPHY(Point,4326),
    distance_km NUMERIC,
    price NUMERIC,
    status booking_status DEFAULT 'pending',
    payment_status payment_status DEFAULT 'pending',
    requested_at TIMESTAMP DEFAULT now(),
    accepted_at TIMESTAMP,
    started_at TIMESTAMP,
    completed_at TIMESTAMP,
    version INT DEFAULT 1,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- Covering index for fetching bookings per user efficiently
CREATE INDEX idx_booking_user_cover
ON bookings(user_id, created_at DESC, id, status, price);

-- Covering index for driver's active bookings
CREATE INDEX idx_booking_driver_status
ON bookings(driver_id, status, id);

CREATE INDEX idx_booking_car
ON bookings(car_id);

-- ============================================
-- DRIVER REVIEWS
-- ============================================

CREATE TABLE driver_reviews (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    booking_id UUID UNIQUE REFERENCES bookings(id),
    driver_id UUID REFERENCES drivers(id),
    user_id UUID REFERENCES users(id),
    rating INT CHECK (rating BETWEEN 1 AND 5),
    comment TEXT,
    created_at TIMESTAMP DEFAULT now()
);

-- Covering index for fetching reviews fast
CREATE INDEX idx_reviews_driver_cover
ON driver_reviews(driver_id, rating, comment);

-- ============================================
-- PAYMENTS
-- ============================================

CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    booking_id UUID REFERENCES bookings(id),
    amount NUMERIC NOT NULL,
    payment_method TEXT,
    status payment_status DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT now()
);

-- Covering index for payment queries
CREATE INDEX idx_payment_booking_cover
ON payments(booking_id, amount, status);

-- ============================================
-- NOTIFICATIONS (MOBILE APP)
-- ============================================

CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    is_read BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT now()
);

-- Covering index for fast notification fetch per user
CREATE INDEX idx_notifications_user_cover
ON notifications(user_id, created_at DESC, id, title, is_read);