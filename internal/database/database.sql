-- ============================================
-- EXTENSIONS
-- ============================================

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS postgis;
CREATE EXTENSION IF NOT EXISTS pgcrypto;


-- ============================================
-- ENUM TYPES
-- ============================================

CREATE TYPE user_role AS ENUM (
'customer',
'driver',
'admin'
);

CREATE TYPE booking_status AS ENUM (
'pending',
'accepted',
'arriving',
'started',
'completed',
'cancelled'
);

CREATE TYPE payment_status AS ENUM (
'pending',
'paid',
'failed'
);

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

CREATE INDEX idx_users_phone ON users(phone);


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

CREATE INDEX idx_drivers_user ON drivers(user_id);


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

CREATE INDEX idx_cars_driver ON cars(driver_id);
CREATE INDEX idx_cars_status ON cars(status);


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

CREATE INDEX idx_car_images_car ON car_images(car_id);


-- ============================================
-- DRIVER LOCATIONS (REAL TIME GPS)
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

updated_at TIMESTAMP DEFAULT now()
);

CREATE INDEX idx_booking_user ON bookings(user_id);
CREATE INDEX idx_booking_car ON bookings(car_id);
CREATE INDEX idx_booking_status ON bookings(status);


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

CREATE INDEX idx_reviews_driver ON driver_reviews(driver_id);


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


CREATE INDEX idx_payment_booking ON payments(booking_id);


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

CREATE INDEX idx_notifications_user ON notifications(user_id);