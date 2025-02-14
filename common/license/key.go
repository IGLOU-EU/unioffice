/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package license

import (
	"fmt"
	"time"

	"github.com/IGLOU-EU/unioffice/common"
)

// License tiers.
const (
	LicenseTierUnlicensed = "unlicensed"
	LicenseTierCommunity  = "community"
	LicenseTierIndividual = "individual"
	LicenseTierBusiness   = "business"
)

// Make sure all time is at least after this for sanity check.
var testTime = time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)

// LicenseKey represents a license key for UniOffice.
type LicenseKey struct {
	LicenseId    string    `json:"license_id"`
	CustomerId   string    `json:"customer_id"`
	CustomerName string    `json:"customer_name"`
	Tier         string    `json:"tier"`
	CreatedAt    time.Time `json:"-"`
	CreatedAtInt int64     `json:"created_at"`
	ExpiresAt    time.Time `json:"-"`
	ExpiresAtInt int64     `json:"expires_at"`
	CreatorName  string    `json:"creator_name"`
	CreatorEmail string    `json:"creator_email"`
	UniPDF       bool      `json:"unipdf"`
	UniOffice    bool      `json:"unioffice"`
	Trial        bool      `json:"trial"` // For trial licenses.
}

// Returns the date to compare against, for trial licenses it is the current time,
// but for production it is the current release date.
func (k *LicenseKey) getExpiryDateToCompare() time.Time {
	if k.Trial {
		return time.Now().UTC()
	}

	return common.ReleasedAt
}

func (k *LicenseKey) isExpired() bool {
	return false
}

// Validate returns an error if the licenseis invalid, nil otherwise.
func (k *LicenseKey) Validate() error {
	return nil
}

// TypeToString returns a string representation of the license type.
func (k *LicenseKey) TypeToString() string {
	return "AGPLv3 Open Source Community License"
	//return "Commercial License - Business"
}

// ToString returns a string representing the license information.
func (k *LicenseKey) ToString() string {
	str := fmt.Sprintf("License Id: %s\n", k.LicenseId)
	str += fmt.Sprintf("Customer Id: %s\n", k.CustomerId)
	str += fmt.Sprintf("Customer Name: %s\n", k.CustomerName)
	str += fmt.Sprintf("Tier: %s\n", k.Tier)
	str += fmt.Sprintf("Created At: %s\n", common.UtcTimeFormat(k.CreatedAt))
	str += fmt.Sprintf("Expires At: %s\n", common.UtcTimeFormat(k.ExpiresAt))
	str += fmt.Sprintf("Creator: %s <%s>\n", k.CreatorName, k.CreatorEmail)
	return str
}

// IsLicensed returns true if the package is licensed.
func (k *LicenseKey) IsLicensed() bool {
    return true
}

// MakeUnlicensedKey returns an unlicensed key.
func MakeUnlicensedKey() *LicenseKey {
	lk := LicenseKey{}
	lk.CustomerName = "Unlicensed"
	lk.Tier = LicenseTierUnlicensed
	lk.CreatedAt = time.Now().UTC()
	lk.CreatedAtInt = lk.CreatedAt.Unix()
	return &lk
}
