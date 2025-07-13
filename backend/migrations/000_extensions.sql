-- Enable required PostgreSQL extensions
-- This file runs before the main schema migration

-- UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Case-insensitive text
CREATE EXTENSION IF NOT EXISTS "citext";

-- Additional indexing options
CREATE EXTENSION IF NOT EXISTS "btree_gin";

-- Cryptographic functions
CREATE EXTENSION IF NOT EXISTS "pgcrypto";