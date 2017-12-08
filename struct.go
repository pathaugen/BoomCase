
package main
import (
	"time"

    "golang.org/x/net/context"

	//"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// ========== START: case_struct ========== ========== ========== ========== ========== ========== ========== ========== ==========
type Case struct {
	Name							string
	Overview					string
	Featuring					string

	//FrequencyResponse	string	// Old field: Replacced by low/high
	FrequencyLow			int32
	FrequencyHigh			int32

	Length						int8	// uint8 = 0-255 (int8 = 127)
	Width							int8	// uint8 = 0-255 (int8 = 127)
	Height						int8	// uint8 = 0-255 (int8 = 127)

	Weight						int8	// uint8 = 0-255 (int8 = 127)
	Battery						int8	// uint8 = 0-255 (int8 = 127)
	Notes							string

	Price							int32	// uint16 = 0-65,535
	Watts							int16
	Sold							bool	// bool - Mark as sold

	DriverMultiplier	string	// float32 // 1.12345678 (eight digits from decimal point) -> Switch to string for precision

	BlobKey						string

	DateAdded					time.Time
}
// caseKey returns the key used for all case entries.
func caseKey(ctx context.Context) *datastore.Key { // context.Context vs appengine.Context
	// The string "default_case" here could be varied to have multiple types of cases.
	return datastore.NewKey(ctx, "Case", "default_case", 0, nil)
}
// ========== END: case_struct ========== ========== ========== ========== ========== ========== ========== ========== ==========

// ========== START: image_struct ========== ========== ========== ========== ========== ========== ========== ========== ==========
// Images for each Case, and marking which one can be customized
type CaseImage struct {
	CaseId						int64	// e.g. (4,767,482,418,036,736) uint64 = 0-18,446,744,073,709,551,615 uint32 = 0-4,294,967,295
	BlobKey						string	//appengine.BlobKey // e.g. (ahBkZXZ-cG1kLWJvb21jYXNlcicLEgRDYXNlIgxkZWZhdWx0X2Nhc2UMCxIEQ2FzZRiAgICAgIC8CAw)

	Customizable			bool	// bool - Mark as customizable

	DateAdded					time.Time
}
// imageKey returns the key used for all case entries.
func caseImageKey(ctx context.Context) *datastore.Key { // context.Context vs appengine.Context
	return datastore.NewKey(ctx, "CaseImage", "default_caseimage", 0, nil)
}
// ========== END: image_struct ========== ========== ========== ========== ========== ========== ========== ========== ==========

// ========== START: driver_struct ========== ========== ========== ========== ========== ========== ========== ========== ==========
type Driver struct {
	Name							string
	Type							string	// low/mid/high
	Circle						bool	// bool - Mark as circle

	//FrequencyResponse	string	// Old field: Replacced by low/high
	FrequencyLow			int32
	FrequencyHigh			int32

	Weight						int8	// uint8 = 0-255 (int8 = 127)
	Diameter					int16	// uint8 = 0-255 -> int8 = -128 - 127 vs. int16 = -32,768 - 32,767
	Price							int32	// uint16 = 0-65,535

	BlobKey						string

	DateAdded					time.Time
}
// driverKey returns the key used for all case entries.
func driverKey(ctx context.Context) *datastore.Key { // context.Context vs appengine.Context
	return datastore.NewKey(ctx, "Driver", "default_driver", 0, nil)
}
// ========== END: driver_struct ========== ========== ========== ========== ========== ========== ========== ========== ==========
