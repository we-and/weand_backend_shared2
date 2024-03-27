package validators

const (
	ValidationFormatKey_Email                   ValidationFormatKey = "Email"
	ValidationFormatKey_Password                ValidationFormatKey = "Password"
	ValidationFormatKey_MaxLength32                                 = "MaxLength32"
	ValidationFormatKey_MaxLength128                                = "MaxLength128"
	ValidationFormatKey_Length3                                     = "Length3"
	ValidationFormatKey_MinLength8                                  = "MinLength8"
	ValidationFormatKey_MinLength3                                  = "MinLength3"
	ValidationFormatKey_Alphabetical                                = "Alphabetical"
	ValidationFormatKey_Alphabetical_Nonempty                       = "Alphabetical_Nonempty"
	ValidationFormatKey_Alphanumerical                              = "Alphanumerical"
	ValidationFormatKey_Number_NonEmpty                             = "Number_NonEmpty"
	ValidationFormatKey_Alphanumerical_Nonempty                     = "Alphanumerical_Nonempty"
	ValidationFormatKey_Nonempty                                    = "Nonempty"
	ValidationFormatKey_NonZero                                     = "NonZero"
)

type ValidationFormat struct {
	KeyForDesc string
	Type       ValidationFormatKey
	Value      string
}

type ValidationFormatInt struct {
	KeyForDesc string
	Type       ValidationFormatKey
	Value      uint32
}
