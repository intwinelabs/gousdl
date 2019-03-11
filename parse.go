package gousdl

import (
	"regexp"
	"strconv"
	"strings"
	"time"
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

	for _, line := range lines {
		if !started {
			if strings.Index(line, "ANSI ") == 0 {
				started = true
			}
			continue
		} else {

			code := getCode(line)
			value := getValue(line)
			field := getField(code)

			if field != "" {
				switch field {
				case "jurisdictionVehicleClass":
					usdlData.JurisdictionVehicleClass = value
				case "jurisdictionRestrictionCodes":
					usdlData.JurisdictionRestrictionCodes = value
				case "jurisdictionEndorsementCodes":
					usdlData.JurisdictionEndorsementCodes = value
				case "dateOfExpiry":
					usdlData.DateOfExpiry = getDate(value)
				case "lastName":
					usdlData.LastName = value
				case "firstName":
					usdlData.FirstName = value
				case "middleName":
					usdlData.MiddleName = value
				case "dateOfIssue":
					usdlData.DateOfIssue = getDate(value)
				case "dateOfBirth":
					usdlData.DateOfBirth = getDate(value)
				case "sex":
					usdlData.Sex = getSex(value)
				case "eyeColor":
					usdlData.EyeColor = value
				case "height":
					usdlData.HeightIn = getHeight(value, 0)
					usdlData.HeightCm = getHeight(value, 1)
				case "addressStreet":
					usdlData.AddressStreet = value
				case "addressCity":
					usdlData.AddressCity = value
				case "addressState":
					usdlData.AddressState = value
				case "addressPostalCode":
					usdlData.AddressPostalCode = value
				case "documentNumber":
					usdlData.DocumentNumber = value
				case "documentDiscriminator":
					usdlData.DocumentDiscriminator = value
				case "issuer":
					usdlData.Issuer = value
				case "lastNameTruncated":
					usdlData.LastNameTruncated = value
				case "firstNameTruncated":
					usdlData.FirstNameTruncated = value
				case "middleNameTruncated":
					usdlData.MiddleNameTruncated = value
				case "hairColor":
					usdlData.HairColor = value
				case "addressStreet2":
					usdlData.AddressStreet2 = value
				case "placeOfBirth":
					usdlData.PlaceOfBirth = value
				case "auditInformation":
					usdlData.AuditInformation = value
				case "inventoryControlNumber":
					usdlData.InventoryControlNumber = value
				case "otherLastName":
					usdlData.OtherLastName = value
				case "otherFirstName":
					usdlData.OtherFirstName = value
				case "otherSuffixName":
					usdlData.OtherSuffixName = value
				case "nameSuffix":
					usdlData.NameSuffix = value
				case "weightRange":
					usdlData.WeightRange = getWightRange(value)
				case "race":
					usdlData.Race = value
				case "standardVehicleClassification":
					usdlData.StandardVehicleClassification = value
				case "standardEndorsementCode":
					usdlData.StandardEndorsementCode = value
				case "standardRestrictionCode":
					usdlData.StandardRestrictionCode = value
				case "jurisdictionVehicleClassificationDescription":
					usdlData.JurisdictionVehicleClassificationDescription = value
				case "jurisdictionEndorsementCodeDescription":
					usdlData.JurisdictionEndorsementCodeDescription = value
				case "jurisdictionRestrictionCodeDescription":
					usdlData.JurisdictionRestrictionCodeDescription = value
				case "complianceType":
					usdlData.ComplianceType = value
				case "dateCardRevised":
					usdlData.DateCardRevised = getDate(value)
				case "dateOfExpiryHazmatEndorsement":
					usdlData.DateOfExpiryHazmatEndorsement = getDate(value)
				case "limitedDurationDocumentIndicator":
					usdlData.LimitedDurationDocumentIndicator = getBool(value)
				case "weightLb":
					usdlData.WeightLb = getInt(value)
				case "weightKg":
					usdlData.WeightKg = getInt(value)
				case "dateAge18":
					usdlData.DateAge18 = getDate(value)
				case "dateAge19":
					usdlData.DateAge19 = getDate(value)
				case "dateAge21":
					usdlData.DateAge21 = getDate(value)
				case "organDonor":
					usdlData.OrganDonor = getBool(value)
				case "veteran":
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

func getField(code string) string {
	if field, ok := codeToField[code]; ok {
		return field
	}
	return ""
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
