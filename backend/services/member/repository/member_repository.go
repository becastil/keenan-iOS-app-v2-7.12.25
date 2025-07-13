package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/sydney-health/backend/pkg/database"
	pb "github.com/sydney-health/backend/shared/pb"
)

// MemberRepository handles database operations for members
type MemberRepository struct {
	db *database.DB
}

// NewMemberRepository creates a new member repository
func NewMemberRepository(db *database.DB) *MemberRepository {
	return &MemberRepository{db: db}
}

// GetMember retrieves a member by ID
func (r *MemberRepository) GetMember(ctx context.Context, memberID string) (*pb.Member, error) {
	query := `
		SELECT 
			member_id, email, first_name, last_name, date_of_birth,
			phone, gender, created_at, updated_at
		FROM members
		WHERE member_id = $1
	`

	var member pb.Member
	var dob time.Time
	err := r.db.QueryRowContext(ctx, query, memberID).Scan(
		&member.MemberId,
		&member.Email,
		&member.FirstName,
		&member.LastName,
		&dob,
		&member.Phone,
		&member.Gender,
		&member.CreatedAt,
		&member.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("member not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get member: %w", err)
	}

	member.DateOfBirth = dob.Format("2006-01-02")

	// Get member's coverage
	coverageQuery := `
		SELECT coverage_type, plan_name, group_number, effective_date, termination_date
		FROM member_coverages
		WHERE member_id = $1 AND status = 'active'
	`

	rows, err := r.db.QueryContext(ctx, coverageQuery, memberID)
	if err != nil {
		return nil, fmt.Errorf("failed to get coverage: %w", err)
	}
	defer rows.Close()

	member.Coverages = make([]*pb.MemberCoverage, 0)
	for rows.Next() {
		var coverage pb.MemberCoverage
		var effectiveDate, terminationDate sql.NullTime

		err := rows.Scan(
			&coverage.CoverageType,
			&coverage.PlanName,
			&coverage.GroupNumber,
			&effectiveDate,
			&terminationDate,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan coverage: %w", err)
		}

		if effectiveDate.Valid {
			coverage.EffectiveDate = effectiveDate.Time.Format("2006-01-02")
		}
		if terminationDate.Valid {
			coverage.TerminationDate = terminationDate.Time.Format("2006-01-02")
		}

		member.Coverages = append(member.Coverages, &coverage)
	}

	// Get address
	addressQuery := `
		SELECT street1, street2, city, state, zip_code, country
		FROM member_addresses
		WHERE member_id = $1 AND is_primary = true
		LIMIT 1
	`

	var address pb.Address
	var street2 sql.NullString
	err = r.db.QueryRowContext(ctx, addressQuery, memberID).Scan(
		&address.Street1,
		&street2,
		&address.City,
		&address.State,
		&address.ZipCode,
		&address.Country,
	)

	if err == nil {
		if street2.Valid {
			address.Street2 = street2.String
		}
		member.Address = &address
	}

	return &member, nil
}

// UpdateMember updates member information
func (r *MemberRepository) UpdateMember(ctx context.Context, member *pb.Member) error {
	query := `
		UPDATE members
		SET email = $2, phone = $3, updated_at = CURRENT_TIMESTAMP
		WHERE member_id = $1
	`

	result, err := r.db.ExecContext(ctx, query, member.MemberId, member.Email, member.Phone)
	if err != nil {
		return fmt.Errorf("failed to update member: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("member not found")
	}

	// Update address if provided
	if member.Address != nil {
		addressQuery := `
			UPDATE member_addresses
			SET street1 = $2, street2 = $3, city = $4, state = $5, zip_code = $6, updated_at = CURRENT_TIMESTAMP
			WHERE member_id = $1 AND is_primary = true
		`

		_, err = r.db.ExecContext(ctx, addressQuery,
			member.MemberId,
			member.Address.Street1,
			member.Address.Street2,
			member.Address.City,
			member.Address.State,
			member.Address.ZipCode,
		)
		if err != nil {
			return fmt.Errorf("failed to update address: %w", err)
		}
	}

	return nil
}

// GetMemberCard retrieves member card information
func (r *MemberRepository) GetMemberCard(ctx context.Context, memberID string, coverageType pb.CoverageType) (*pb.MemberCard, error) {
	query := `
		SELECT 
			m.member_id, m.first_name, m.last_name,
			mc.coverage_type, mc.plan_name, mc.group_number, mc.bin_number,
			mc.pcn_number, mc.copay_primary, mc.copay_specialist, mc.copay_er
		FROM members m
		JOIN member_coverages mc ON m.member_id = mc.member_id
		WHERE m.member_id = $1 AND mc.coverage_type = $2 AND mc.status = 'active'
	`

	var card pb.MemberCard
	var copayPrimary, copaySpecialist, copayER sql.NullFloat64

	err := r.db.QueryRowContext(ctx, query, memberID, coverageType.String()).Scan(
		&card.MemberId,
		&card.MemberName,
		&card.MemberName, // Will concatenate first and last
		&card.CoverageType,
		&card.PlanName,
		&card.GroupNumber,
		&card.BinNumber,
		&card.PcnNumber,
		&copayPrimary,
		&copaySpecialist,
		&copayER,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("member card not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get member card: %w", err)
	}

	// Build copay information
	card.CopayInfo = make(map[string]string)
	if copayPrimary.Valid {
		card.CopayInfo["primary"] = fmt.Sprintf("$%.0f", copayPrimary.Float64)
	}
	if copaySpecialist.Valid {
		card.CopayInfo["specialist"] = fmt.Sprintf("$%.0f", copaySpecialist.Float64)
	}
	if copayER.Valid {
		card.CopayInfo["emergency"] = fmt.Sprintf("$%.0f", copayER.Float64)
	}

	return &card, nil
}

// ListDependents retrieves all dependents for a member
func (r *MemberRepository) ListDependents(ctx context.Context, memberID string) ([]*pb.Member, error) {
	query := `
		SELECT 
			d.member_id, d.email, d.first_name, d.last_name, d.date_of_birth,
			d.phone, d.gender, d.created_at, d.updated_at
		FROM members d
		JOIN member_dependents md ON d.member_id = md.dependent_id
		WHERE md.primary_member_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, memberID)
	if err != nil {
		return nil, fmt.Errorf("failed to list dependents: %w", err)
	}
	defer rows.Close()

	var dependents []*pb.Member
	for rows.Next() {
		var member pb.Member
		var dob time.Time

		err := rows.Scan(
			&member.MemberId,
			&member.Email,
			&member.FirstName,
			&member.LastName,
			&dob,
			&member.Phone,
			&member.Gender,
			&member.CreatedAt,
			&member.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan dependent: %w", err)
		}

		member.DateOfBirth = dob.Format("2006-01-02")
		dependents = append(dependents, &member)
	}

	return dependents, nil
}