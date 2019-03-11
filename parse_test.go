package gousdl

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	sampleData := `ANSI 636000090002DL00410278ZV03190008DLDAQT64235789L
DCSSAMPLE
DDEN
DACMICHAEL
DDFN
DADJOHN
DDGN
DCUJR
DCAD
DCBK
DCDPH
DBD06062016
DBB06061986
DBA12102024
DBC1
DAU068 IN
DAYBRO
DAG2300 WEST BROAD STREET
DAIRICHMOND
DAJVA
DAK232690000
DCF2424244747474786102204
DCGUSA
DAW240
DCK123456789
DDAF
DDB06062008
DDC06062009
DDD1
ZVZVA01`

	sampleStruct := &USDLData{
		JurisdictionVehicleClass:      "D",
		JurisdictionRestrictionCodes:  "K",
		JurisdictionEndorsementCodes:  "PH",
		DateOfExpiry:                  1733788800,
		LastName:                      "SAMPLE",
		FirstName:                     "MICHAEL",
		MiddleName:                    "JOHN",
		DateOfIssue:                   1465171200,
		DateOfBirth:                   518400000,
		Sex:                           "M",
		EyeColor:                      "BRO",
		HeightIn:                      68,
		HeightCm:                      172,
		AddressStreet:                 "2300 WEST BROAD STREET",
		AddressCity:                   "RICHMOND",
		AddressState:                  "VA",
		AddressPostalCode:             "232690000",
		DocumentNumber:                "",
		DocumentDiscriminator:         "2424244747474786102204",
		Issuer:                        "USA",
		LastNameTruncated:             "N",
		FirstNameTruncated:            "N",
		MiddleNameTruncated:           "N",
		AddressStreet2:                "",
		HairColor:                     "",
		PlaceOfBirth:                  "",
		AuditInformation:              "",
		InventoryControlNumber:        "123456789",
		OtherLastName:                 "",
		OtherFirstName:                "",
		OtherSuffixName:               "",
		NameSuffix:                    "JR",
		WeightRange:                   "",
		Race:                          "",
		StandardVehicleClassification: "",
		StandardEndorsementCode:       "",
		StandardRestrictionCode:       "",
		JurisdictionVehicleClassificationDescription: "",
		JurisdictionEndorsementCodeDescription:       "",
		JurisdictionRestrictionCodeDescription:       "",
		ComplianceType:                               "F",
		DateCardRevised:                              1212710400,
		DateOfExpiryHazmatEndorsement:                1244246400,
		LimitedDurationDocumentIndicator:             true,
		WeightLb:                                     240,
		WeightKg:                                     0,
		DateAge18:                                    0,
		DateAge19:                                    0,
		DateAge21:                                    0,
		OrganDonor:                                   false,
		Veteran:                                      false}

	usdlData, err := ParseString(sampleData)
	assert.Nil(err)
	require.NoError(err)
	assert.Equal(sampleStruct, usdlData)
}
