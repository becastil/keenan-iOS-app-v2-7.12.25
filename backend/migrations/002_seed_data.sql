-- Seed data for development environment
-- This creates test data for the demo user (M123456)

-- Insert test member
INSERT INTO members (member_id, email, first_name, last_name, date_of_birth, phone, gender, created_at, updated_at)
VALUES 
    ('M123456', 'john.doe@example.com', 'John', 'Doe', '1985-03-15', '555-123-4567', 'MALE', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert member address
INSERT INTO member_addresses (member_id, street1, street2, city, state, zip_code, country, is_primary, created_at, updated_at)
VALUES 
    ('M123456', '123 Main Street', 'Apt 4B', 'New York', 'NY', '10001', 'USA', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert coverage levels
INSERT INTO coverage_levels (coverage_type, tier, monthly_premium, annual_deductible, out_of_pocket_max, created_at)
VALUES 
    ('MEDICAL', 'GOLD', 450.00, 1500.00, 6000.00, CURRENT_TIMESTAMP),
    ('DENTAL', 'STANDARD', 45.00, 250.00, 1500.00, CURRENT_TIMESTAMP),
    ('VISION', 'PREMIUM', 20.00, 150.00, 500.00, CURRENT_TIMESTAMP),
    ('PHARMACY', 'ENHANCED', 0.00, 100.00, 2000.00, CURRENT_TIMESTAMP);

-- Insert member coverages
INSERT INTO member_coverages (member_id, coverage_type, plan_name, group_number, effective_date, status, 
    bin_number, pcn_number, copay_primary, copay_specialist, copay_er, created_at)
VALUES 
    ('M123456', 'MEDICAL', 'Gold PPO Plan', 'GRP-001', '2024-01-01', 'active', 
     NULL, NULL, 20.00, 40.00, 150.00, CURRENT_TIMESTAMP),
    ('M123456', 'DENTAL', 'Standard Dental', 'GRP-001', '2024-01-01', 'active', 
     NULL, NULL, NULL, NULL, NULL, CURRENT_TIMESTAMP),
    ('M123456', 'VISION', 'Premium Vision', 'GRP-001', '2024-01-01', 'active', 
     NULL, NULL, NULL, NULL, NULL, CURRENT_TIMESTAMP),
    ('M123456', 'PHARMACY', 'Enhanced Rx', 'GRP-001', '2024-01-01', 'active', 
     '123456', 'PCN001', 10.00, 25.00, 50.00, CURRENT_TIMESTAMP);

-- Insert some benefits
INSERT INTO benefits (coverage_type, category, subcategory, description, in_network_coverage, out_network_coverage, 
    requires_preauth, created_at)
VALUES 
    ('MEDICAL', 'Preventive Care', 'Annual Physical', 'Annual wellness exam', '100%', '70%', false, CURRENT_TIMESTAMP),
    ('MEDICAL', 'Preventive Care', 'Immunizations', 'Routine vaccinations', '100%', '70%', false, CURRENT_TIMESTAMP),
    ('MEDICAL', 'Specialist Care', 'Cardiology', 'Heart specialist visits', '80%', '60%', true, CURRENT_TIMESTAMP),
    ('MEDICAL', 'Emergency', 'ER Visit', 'Emergency room services', '80%', '80%', false, CURRENT_TIMESTAMP),
    ('DENTAL', 'Preventive', 'Cleaning', 'Routine cleaning (2x per year)', '100%', '80%', false, CURRENT_TIMESTAMP),
    ('VISION', 'Routine', 'Eye Exam', 'Annual eye examination', '100%', '80%', false, CURRENT_TIMESTAMP);

-- Insert some providers
INSERT INTO providers (provider_id, name, type, npi, tax_id, created_at, updated_at)
VALUES 
    ('P001', 'New York Medical Center', 'FACILITY', '1234567890', '12-3456789', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('P002', 'Dr. Sarah Johnson, MD', 'INDIVIDUAL', '0987654321', '98-7654321', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('P003', 'Downtown Dental Group', 'FACILITY', '1122334455', '11-2233445', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert provider specialties
INSERT INTO provider_specialties (provider_id, specialty, is_primary, created_at)
VALUES 
    ('P001', 'General Medicine', true, CURRENT_TIMESTAMP),
    ('P001', 'Emergency Medicine', false, CURRENT_TIMESTAMP),
    ('P002', 'Cardiology', true, CURRENT_TIMESTAMP),
    ('P003', 'General Dentistry', true, CURRENT_TIMESTAMP);

-- Insert provider locations
INSERT INTO provider_locations (provider_id, name, street1, city, state, zip_code, phone, is_primary, 
    accepts_new_patients, created_at)
VALUES 
    ('P001', 'Main Campus', '100 Hospital Way', 'New York', 'NY', '10002', '555-100-1000', true, true, CURRENT_TIMESTAMP),
    ('P002', 'Cardiology Associates', '200 Medical Plaza', 'New York', 'NY', '10003', '555-200-2000', true, true, CURRENT_TIMESTAMP),
    ('P003', 'Downtown Office', '300 Dental Drive', 'New York', 'NY', '10004', '555-300-3000', true, true, CURRENT_TIMESTAMP);

-- Insert some claims
INSERT INTO claims (claim_id, member_id, provider_id, service_date, received_date, processed_date, 
    status, total_amount, member_responsibility, paid_amount, created_at, updated_at)
VALUES 
    ('CLM-2024-001', 'M123456', 'P002', '2024-06-15', '2024-06-20', '2024-06-25', 
     'PAID', 250.00, 40.00, 210.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('CLM-2024-002', 'M123456', 'P001', '2024-07-01', '2024-07-05', '2024-07-10', 
     'PAID', 150.00, 20.00, 130.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('CLM-2024-003', 'M123456', 'P003', '2024-07-10', '2024-07-12', NULL, 
     'PROCESSING', 200.00, 0.00, 0.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert claim line items
INSERT INTO claim_line_items (claim_id, service_code, description, quantity, unit_price, total_price, created_at)
VALUES 
    ('CLM-2024-001', '99213', 'Office Visit - Established Patient', 1, 150.00, 150.00, CURRENT_TIMESTAMP),
    ('CLM-2024-001', '93000', 'EKG', 1, 100.00, 100.00, CURRENT_TIMESTAMP),
    ('CLM-2024-002', '99281', 'Emergency Dept Visit', 1, 150.00, 150.00, CURRENT_TIMESTAMP),
    ('CLM-2024-003', 'D0120', 'Periodic Oral Evaluation', 1, 75.00, 75.00, CURRENT_TIMESTAMP),
    ('CLM-2024-003', 'D1110', 'Prophylaxis - Adult', 1, 125.00, 125.00, CURRENT_TIMESTAMP);

-- Insert a conversation for messaging
INSERT INTO conversations (conversation_id, member_id, subject, status, created_at, updated_at)
VALUES 
    (uuid_generate_v4(), 'M123456', 'Question about coverage', 'OPEN', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert a message (using the conversation_id from above - you'll need to update this)
-- This is just an example structure
-- INSERT INTO messages (message_id, conversation_id, sender_id, sender_type, content, created_at)
-- VALUES 
--     (uuid_generate_v4(), '<conversation_id>', 'M123456', 'MEMBER', 'I have a question about my dental coverage.', CURRENT_TIMESTAMP);

-- Add some audit log entries
INSERT INTO audit_logs (entity_type, entity_id, action, actor_id, actor_type, changes, created_at)
VALUES 
    ('member', 'M123456', 'login', 'M123456', 'member', '{"ip": "192.168.1.1", "user_agent": "Mozilla/5.0"}', CURRENT_TIMESTAMP),
    ('claim', 'CLM-2024-001', 'status_change', 'SYSTEM', 'system', '{"old_status": "PROCESSING", "new_status": "PAID"}', CURRENT_TIMESTAMP);