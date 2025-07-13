-- Sydney Health Database Schema

-- Members table
CREATE TABLE IF NOT EXISTS members (
    member_id VARCHAR(50) PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100),
    date_of_birth DATE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20),
    street1 VARCHAR(255),
    street2 VARCHAR(255),
    city VARCHAR(100),
    state VARCHAR(2),
    zip_code VARCHAR(10),
    country VARCHAR(50) DEFAULT 'USA',
    group_number VARCHAR(50) NOT NULL,
    subscriber_id VARCHAR(50) UNIQUE NOT NULL,
    enrollment_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_email (email),
    INDEX idx_subscriber_id (subscriber_id)
);

-- Member coverages table
CREATE TABLE IF NOT EXISTS member_coverages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    member_id VARCHAR(50) NOT NULL,
    coverage_type ENUM('MEDICAL', 'DENTAL', 'VISION', 'PHARMACY') NOT NULL,
    status ENUM('ACTIVE', 'INACTIVE', 'PENDING') DEFAULT 'ACTIVE',
    effective_date DATE NOT NULL,
    termination_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (member_id) REFERENCES members(member_id),
    UNIQUE KEY unique_member_coverage (member_id, coverage_type),
    INDEX idx_member_coverage (member_id, coverage_type, status)
);

-- Benefits table
CREATE TABLE IF NOT EXISTS benefits (
    benefit_id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    coverage_type ENUM('MEDICAL', 'DENTAL', 'VISION', 'PHARMACY') NOT NULL,
    is_covered BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_coverage_type (coverage_type)
);

-- Coverage levels table
CREATE TABLE IF NOT EXISTS coverage_levels (
    id INT AUTO_INCREMENT PRIMARY KEY,
    benefit_id VARCHAR(50) NOT NULL,
    network_type ENUM('IN_NETWORK', 'OUT_OF_NETWORK') NOT NULL,
    copay_cents INT,
    coinsurance_percentage INT,
    deductible_cents INT,
    annual_limit_cents INT,
    lifetime_limit_cents INT,
    FOREIGN KEY (benefit_id) REFERENCES benefits(benefit_id),
    UNIQUE KEY unique_benefit_network (benefit_id, network_type)
);

-- Providers table
CREATE TABLE IF NOT EXISTS providers (
    provider_id VARCHAR(50) PRIMARY KEY,
    npi VARCHAR(10) UNIQUE,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    practice_name VARCHAR(255),
    phone VARCHAR(20),
    fax VARCHAR(20),
    accepting_new_patients BOOLEAN DEFAULT TRUE,
    rating DECIMAL(2,1),
    review_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_npi (npi)
);

-- Provider specialties table
CREATE TABLE IF NOT EXISTS provider_specialties (
    id INT AUTO_INCREMENT PRIMARY KEY,
    provider_id VARCHAR(50) NOT NULL,
    specialty VARCHAR(100) NOT NULL,
    is_primary BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (provider_id) REFERENCES providers(provider_id),
    INDEX idx_provider_specialty (provider_id, specialty)
);

-- Provider locations table
CREATE TABLE IF NOT EXISTS provider_locations (
    location_id VARCHAR(50) PRIMARY KEY,
    provider_id VARCHAR(50) NOT NULL,
    street1 VARCHAR(255) NOT NULL,
    street2 VARCHAR(255),
    city VARCHAR(100) NOT NULL,
    state VARCHAR(2) NOT NULL,
    zip_code VARCHAR(10) NOT NULL,
    phone VARCHAR(20),
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    FOREIGN KEY (provider_id) REFERENCES providers(provider_id),
    INDEX idx_provider_location (provider_id),
    INDEX idx_location_geo (latitude, longitude)
);

-- Claims table
CREATE TABLE IF NOT EXISTS claims (
    claim_id VARCHAR(50) PRIMARY KEY,
    member_id VARCHAR(50) NOT NULL,
    provider_id VARCHAR(50),
    provider_name VARCHAR(255) NOT NULL,
    service_date DATE NOT NULL,
    processed_date DATE,
    status ENUM('PENDING', 'APPROVED', 'DENIED', 'PROCESSING', 'PAID') DEFAULT 'PENDING',
    coverage_type ENUM('MEDICAL', 'DENTAL', 'VISION', 'PHARMACY') NOT NULL,
    total_charged_cents INT NOT NULL,
    allowed_amount_cents INT,
    deductible_applied_cents INT DEFAULT 0,
    copay_cents INT DEFAULT 0,
    coinsurance_cents INT DEFAULT 0,
    member_responsibility_cents INT,
    plan_paid_cents INT,
    explanation_of_benefits_url VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (member_id) REFERENCES members(member_id),
    INDEX idx_member_claims (member_id, service_date),
    INDEX idx_claim_status (status)
);

-- Claim line items table
CREATE TABLE IF NOT EXISTS claim_line_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    claim_id VARCHAR(50) NOT NULL,
    service_code VARCHAR(20),
    service_description TEXT,
    quantity INT DEFAULT 1,
    charged_amount_cents INT NOT NULL,
    allowed_amount_cents INT,
    paid_amount_cents INT,
    FOREIGN KEY (claim_id) REFERENCES claims(claim_id),
    INDEX idx_claim_lines (claim_id)
);

-- Conversations table
CREATE TABLE IF NOT EXISTS conversations (
    conversation_id VARCHAR(50) PRIMARY KEY,
    member_id VARCHAR(50) NOT NULL,
    subject VARCHAR(255) NOT NULL,
    type ENUM('GENERAL', 'CLAIMS', 'BENEFITS', 'PROVIDER') DEFAULT 'GENERAL',
    status ENUM('OPEN', 'CLOSED', 'PENDING') DEFAULT 'OPEN',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (member_id) REFERENCES members(member_id),
    INDEX idx_member_conversations (member_id, status)
);

-- Messages table
CREATE TABLE IF NOT EXISTS messages (
    message_id VARCHAR(50) PRIMARY KEY,
    conversation_id VARCHAR(50) NOT NULL,
    sender_id VARCHAR(50) NOT NULL,
    sender_name VARCHAR(255) NOT NULL,
    sender_type ENUM('MEMBER', 'SUPPORT_AGENT', 'CARE_COORDINATOR', 'PROVIDER') NOT NULL,
    content TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    read_at TIMESTAMP NULL,
    FOREIGN KEY (conversation_id) REFERENCES conversations(conversation_id),
    INDEX idx_conversation_messages (conversation_id, sent_at)
);

-- Audit log table
CREATE TABLE IF NOT EXISTS audit_logs (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    event_type VARCHAR(50) NOT NULL,
    entity_type VARCHAR(50) NOT NULL,
    entity_id VARCHAR(50) NOT NULL,
    user_id VARCHAR(50),
    ip_address VARCHAR(45),
    user_agent TEXT,
    event_data JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_audit_entity (entity_type, entity_id),
    INDEX idx_audit_user (user_id),
    INDEX idx_audit_created (created_at)
);