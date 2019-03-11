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
DAQ485723A48750328947502398745
DAU068 IN
DAYBRO
DAG2300 WEST BROAD STREET
DAIRICHMOND
DAJVA
DAK232690000
DCF2424244747474786102204
DCGUSA
DAW240
DAZBRN
DAHUNIT 202
DCISPAIN
DCJ3456789
DBNHENDRIX
DBGJIMI
DBSJR
DCE3
DCLWHITE
DCMC
DCOC
DCPFOO
DCQFOO
DCRFOO
DCK123456789
DDAF
DAX12
DDH06062002
DDI06062003
DDJ06062007
DDK1
DDL1
DDB06062008
DDC06062009
DDD1
ZVZVA01`

	sampleStruct := &USDLData{
		Hash:                          "baeb91fe5f52fa73936743436cc482e2115a8f8565c76ca5a19fdc828175338c",
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
		DocumentNumber:                "485723A48750328947502398745",
		DocumentDiscriminator:         "2424244747474786102204",
		Issuer:                        "USA",
		LastNameTruncated:             "N",
		FirstNameTruncated:            "N",
		MiddleNameTruncated:           "N",
		AddressStreet2:                "UNIT 202",
		HairColor:                     "BRN",
		PlaceOfBirth:                  "SPAIN",
		AuditInformation:              "3456789",
		InventoryControlNumber:        "123456789",
		OtherLastName:                 "HENDRIX",
		OtherFirstName:                "JIMI",
		OtherSuffixName:               "JR",
		NameSuffix:                    "JR",
		WeightRange:                   "60 - 70 kg (131  160 lbs)",
		Race:                          "WHITE",
		StandardVehicleClassification: "C",
		StandardEndorsementCode:       "",
		StandardRestrictionCode:       "C",
		JurisdictionVehicleClassificationDescription: "FOO",
		JurisdictionEndorsementCodeDescription:       "FOO",
		JurisdictionRestrictionCodeDescription:       "FOO",
		ComplianceType:                               "F",
		DateCardRevised:                              1212710400,
		DateOfExpiryHazmatEndorsement:                1244246400,
		LimitedDurationDocumentIndicator:             true,
		WeightLb:                                     240,
		WeightKg:                                     12,
		DateAge18:                                    1023321600,
		DateAge19:                                    1054857600,
		DateAge21:                                    1181088000,
		OrganDonor:                                   true,
		Veteran:                                      true,
	}

	usdlData, err := ParseString(sampleData)
	assert.Nil(err)
	require.NoError(err)
	assert.Equal(sampleStruct, usdlData)
}
