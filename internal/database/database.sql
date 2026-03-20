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
    deleted_at TIMESTAMP,

    -- Full-Text Search column
    search_vector tsvector
);

-- Trigger to update search_vector on INSERT or UPDATE
CREATE TRIGGER users_search_vector_update
BEFORE INSERT OR UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION tsvector_update_trigger(search_vector, 'pg_catalog.english', name, email, phone);

-- Indexes
CREATE INDEX idx_users_phone_cover
ON users(phone, id, name, role);

CREATE INDEX idx_users_search_vector
ON users USING GIN(search_vector);

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
    deleted_at TIMESTAMP,

    -- Full-Text Search column
    search_vector tsvector
);

-- Trigger for updating search_vector automatically
CREATE TRIGGER cars_search_vector_update
BEFORE INSERT OR UPDATE ON cars
FOR EACH ROW
EXECUTE FUNCTION tsvector_update_trigger(search_vector, 'pg_catalog.english', brand, model, plate_number, color);

-- Indexes
CREATE INDEX idx_cars_driver_cover
ON cars(driver_id, id, brand, model, status);

CREATE INDEX idx_cars_status
ON cars(status);

CREATE INDEX idx_cars_available
ON cars(status, id, driver_id)
WHERE status='available';

CREATE INDEX idx_cars_search_vector
ON cars USING GIN(search_vector);

-- ============================================
-- CAR IMAGES
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
    updated_at TIMESTAMP DEFAULT now(),

    -- Full-Text Search column
    search_vector tsvector
);

-- Trigger to update search_vector automatically
CREATE TRIGGER bookings_search_vector_update
BEFORE INSERT OR UPDATE ON bookings
FOR EACH ROW
EXECUTE FUNCTION tsvector_update_trigger(search_vector, 'pg_catalog.english', status, payment_status);

-- Indexes
CREATE INDEX idx_booking_user_cover
ON bookings(user_id, created_at DESC, id, status, price);

CREATE INDEX idx_booking_driver_status
ON bookings(driver_id, status, id);

CREATE INDEX idx_booking_car
ON bookings(car_id);

CREATE INDEX idx_bookings_search_vector
ON bookings USING GIN(search_vector);

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
    created_at TIMESTAMP DEFAULT now(),

    -- Full-Text Search column
    search_vector tsvector
);

CREATE TRIGGER driver_reviews_search_vector_update
BEFORE INSERT OR UPDATE ON driver_reviews
FOR EACH ROW
EXECUTE FUNCTION tsvector_update_trigger(search_vector, 'pg_catalog.english', comment);

CREATE INDEX idx_reviews_driver_cover
ON driver_reviews(driver_id, rating, comment);

CREATE INDEX idx_driver_reviews_search_vector
ON driver_reviews USING GIN(search_vector);

-- ============================================
-- PAYMENTS
-- ============================================

CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    booking_id UUID REFERENCES bookings(id),
    amount NUMERIC NOT NULL,
    payment_method TEXT,
    status payment_status DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT now(),

    -- Full-Text Search column
    search_vector tsvector
);

CREATE TRIGGER payments_search_vector_update
BEFORE INSERT OR UPDATE ON payments
FOR EACH ROW
EXECUTE FUNCTION tsvector_update_trigger(search_vector, 'pg_catalog.english', payment_method);

CREATE INDEX idx_payment_booking_cover
ON payments(booking_id, amount, status);

CREATE INDEX idx_payments_search_vector
ON payments USING GIN(search_vector);

-- ============================================
-- NOTIFICATIONS
-- ============================================

CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    is_read BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT now(),

    -- Full-Text Search column
    search_vector tsvector
);

CREATE TRIGGER notifications_search_vector_update
BEFORE INSERT OR UPDATE ON notifications
FOR EACH ROW
EXECUTE FUNCTION tsvector_update_trigger(search_vector, 'pg_catalog.english', title, body);

CREATE INDEX idx_notifications_user_cover
ON notifications(user_id, created_at DESC, id, title, is_read);

CREATE INDEX idx_notifications_search_vector
ON notifications USING GIN(search_vector);