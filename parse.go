package gousdl

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/blake2b"
)

const lineSeparator = "\n"

func ParseString(data string) (*USDLData, error) {
	return Parse([]byte(data))
}

func Parse(data []byte) (*USDLData, error) {
	rawLines := strings.Split(strings.TrimSpace(string(data)), lineSeparator)
	lines := sanitizeData(rawLines)
	started := false

	usdlData := &USDLData{}

	h := blake2b.Sum256(data)
	usdlData.Hash = fmt.Sprintf("%x", h)

	for _, line := range lines {
		if !started {
			if strings.Index(line, "ANSI ") == 0 {
				started = true
			}
			continue
		} else {

			code := getCode(line)
			value := getValue(line)

			if code != "" {
				switch code {
				case "DCA":
					usdlData.JurisdictionVehicleClass = value
				case "DCB":
					usdlData.JurisdictionRestrictionCodes = value
				case "DCD":
					usdlData.JurisdictionEndorsementCodes = value
				case "DBA":
					usdlData.DateOfExpiry = getDate(value)
				case "DCS":
					usdlData.LastName = value
				case "DAC":
					usdlData.FirstName = value
				case "DAD":
					usdlData.MiddleName = value
				case "DBD":
					usdlData.DateOfIssue = getDate(value)
				case "DBB":
					usdlData.DateOfBirth = getDate(value)
				case "DBC":
					usdlData.Sex = getSex(value)
				case "DAY":
					usdlData.EyeColor = value
				case "DAU":
					usdlData.HeightIn = getHeight(value, 0)
					usdlData.HeightCm = getHeight(value, 1)
				case "DAG":
					usdlData.AddressStreet = value
				case "DAI":
					usdlData.AddressCity = value
				case "DAJ":
					usdlData.AddressState = value
				case "DAK":
					usdlData.AddressPostalCode = value
				case "DAQ":
					usdlData.DocumentNumber = value
				case "DCF":
					usdlData.DocumentDiscriminator = value
				case "DCG":
					usdlData.Issuer = value
				case "DDE":
					usdlData.LastNameTruncated = value
				case "DDF":
					usdlData.FirstNameTruncated = value
				case "DDG":
					usdlData.MiddleNameTruncated = value
				case "DAZ":
					usdlData.HairColor = value
				case "DAH":
					usdlData.AddressStreet2 = value
				case "DCI":
					usdlData.PlaceOfBirth = value
				case "DCJ":
					usdlData.AuditInformation = value
				case "DCK":
					usdlData.InventoryControlNumber = value
				case "DBN":
					usdlData.OtherLastName = value
				case "DBG":
					usdlData.OtherFirstName = value
				case "DBS":
					usdlData.OtherSuffixName = value
				case "DCU":
					usdlData.NameSuffix = value
				case "DCE":
					usdlData.WeightRange = getWightRange(value)
				case "DCL":
					usdlData.Race = value
				case "DCM":
					usdlData.StandardVehicleClassification = value
				case "DCN":
					usdlData.StandardEndorsementCode = value
				case "DCO":
					usdlData.StandardRestrictionCode = value
				case "DCP":
					usdlData.JurisdictionVehicleClassificationDescription = value
				case "DCQ":
					usdlData.JurisdictionEndorsementCodeDescription = value
				case "DCR":
					usdlData.JurisdictionRestrictionCodeDescription = value
				case "DDA":
					usdlData.ComplianceType = value
				case "DDB":
					usdlData.DateCardRevised = getDate(value)
				case "DDC":
					usdlData.DateOfExpiryHazmatEndorsement = getDate(value)
				case "DDD":
					usdlData.LimitedDurationDocumentIndicator = getBool(value)
				case "DAW":
					usdlData.WeightLb = getInt(value)
				case "DAX":
					usdlData.WeightKg = getInt(value)
				case "DDH":
					usdlData.DateAge18 = getDate(value)
				case "DDI":
					usdlData.DateAge19 = getDate(value)
				case "DDJ":
					usdlData.DateAge21 = getDate(value)
				case "DDK":
					usdlData.OrganDonor = getBool(value)
				case "DDL":
					usdlData.Veteran = getBool(value)
				}
			}
		}

	}
	return usdlData, nil
}

func sanitizeData(rawLines []string) []string {
	lines := []string{}
	re := regexp.MustCompile(`[\011\012\015\040-\177]*`)
	for _, line := range rawLines {
		if re.MatchString(line) {
			l := strings.TrimSpace(line)
			lines = append(lines, l)
		}
	}

	return lines
}

func getCode(line string) string {
	return line[0:3]
}

func getValue(line string) string {
	return line[3:]
}

func getSex(value string) string {
	switch v := getInt(value); v {
	case 1:
		return "M"
	case 2:
		return "F"
	case 9:
		return "U"
	}
	return ""
}

func getBool(value string) bool {
	return value == "1"
}

func getWightRange(value string) string {
	switch v := getInt(value); v {
	case 0:
		return "up to 31 kg (up to 70 lbs)"
	case 1:
		return "32  45 kg (71  100 lbs)"
	case 2:
		return "46 - 59 kg (101  130 lbs)"
	case 3:
		return "60 - 70 kg (131  160 lbs)"
	case 4:
		return "71 - 86 kg (161  190 lbs)"
	case 5:
		return "87 - 100 kg (191  220 lbs)"
	case 6:
		return "101 - 113 kg (221  250 lbs)"
	case 7:
		return "114 - 127 kg (251  280 lbs)"
	case 8:
		return "128  145 kg (281  320 lbs)"
	case 9:
		return "146+ kg (321+ lbs)"
	}
	return ""
}

func getHeight(value string, units int) int {
	heightAr := strings.Split(value, " ")
	h := getInt(heightAr[0])
	if len(heightAr) == 2 {
		if heightAr[1] == "IN" && units == 0 {
			return h
		} else if heightAr[1] == "CM" && units == 0 {
			return (h * 100) / 254
		} else if heightAr[1] == "IN" && units == 1 {
			return ((h * 100) * 254) / 10000
		} else if heightAr[1] == "CM" && units == 1 {
			return h
		}
	}
	return -1
}

func getDate(value string) int64 {
	t, err := time.Parse("01022006", value)
	if err != nil {
		return -1
	}
	return t.Unix()
}

func getInt(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		return -1
	}
	return i
}
