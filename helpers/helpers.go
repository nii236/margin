package helpers

import (
	"strconv"
	"strings"

	"github.com/nii236/margin/models"
)

type Mapping struct {
	Key   string
	Value int
}

var CurrentFields = []*Mapping{
	&Mapping{Key: "TYPE", Value: 0x0},
	&Mapping{Key: "MARKET", Value: 0x0},
	&Mapping{Key: "FROMSYMBOL", Value: 0x0},
	&Mapping{Key: "TOSYMBOL", Value: 0x0},
	&Mapping{Key: "FLAGS", Value: 0x0},
	&Mapping{Key: "PRICE", Value: 0x1},
	&Mapping{Key: "BID", Value: 0x2},
	&Mapping{Key: "OFFER", Value: 0x4},
	&Mapping{Key: "LASTUPDATE", Value: 0x8},
	&Mapping{Key: "AVG", Value: 0x10},
	&Mapping{Key: "LASTVOLUME", Value: 0x20},
	&Mapping{Key: "LASTVOLUMETO", Value: 0x40},
	&Mapping{Key: "LASTTRADEID", Value: 0x80},
	&Mapping{Key: "VOLUMEHOUR", Value: 0x100},
	&Mapping{Key: "VOLUMEHOURTO", Value: 0x200},
	&Mapping{Key: "VOLUME24HOUR", Value: 0x400},
	&Mapping{Key: "VOLUME24HOURTO", Value: 0x800},
	&Mapping{Key: "OPENHOUR", Value: 0x1000},
	&Mapping{Key: "HIGHHOUR", Value: 0x2000},
	&Mapping{Key: "LOWHOUR", Value: 0x4000},
	&Mapping{Key: "OPEN24HOUR", Value: 0x8000},
	&Mapping{Key: "HIGH24HOUR", Value: 0x10000},
	&Mapping{Key: "LOW24HOUR", Value: 0x20000},
	&Mapping{Key: "LASTMARKET", Value: 0x40000},
}

func UnpackCurrent(msg string) (*models.CurrentMessage, error) {
	valuesArray := strings.Split(msg, "~")
	maskHex := valuesArray[len(valuesArray)-1]
	// fmt.Println(maskHex)
	maskInt64, err := strconv.ParseInt(maskHex, 16, 64)
	if err != nil {
		return nil, err
	}
	maskInt := int(maskInt64)
	// maskBin := strconv.FormatInt(maskInt, 2)
	// fmt.Println(maskBin)
	unpackedCurrent := &models.CurrentMessage{}
	currentField := 0
	for _, v := range CurrentFields {
		if v.Value == 0x0 {
			switch v.Key {
			case "TYPE":
				unpackedCurrent.TYPE = valuesArray[currentField]
			case "MARKET":
				unpackedCurrent.MARKET = valuesArray[currentField]
			case "FROMSYMBOL":
				unpackedCurrent.FROMSYMBOL = valuesArray[currentField]
			case "TOSYMBOL":
				unpackedCurrent.TOSYMBOL = valuesArray[currentField]
			case "FLAGS":
				unpackedCurrent.FLAGS = valuesArray[currentField]
			}
			currentField++
		} else if (maskInt & v.Value) != 0 {
			switch v.Key {
			case "PRICE":
				unpackedCurrent.PRICE = valuesArray[currentField]
			case "BID":
				unpackedCurrent.BID = valuesArray[currentField]
			case "OFFER":
				unpackedCurrent.OFFER = valuesArray[currentField]
			case "LASTUPDATE":
				unpackedCurrent.LASTUPDATE = valuesArray[currentField]
			case "AVG":
				unpackedCurrent.AVG = valuesArray[currentField]
			case "LASTVOLUME":
				unpackedCurrent.LASTVOLUME = valuesArray[currentField]
			case "LASTVOLUMETO":
				unpackedCurrent.LASTVOLUMETO = valuesArray[currentField]
			case "LASTTRADEID":
				unpackedCurrent.LASTTRADEID = valuesArray[currentField]
			case "VOLUMEHOUR":
				unpackedCurrent.VOLUMEHOUR = valuesArray[currentField]
			case "VOLUMEHOURTO":
				unpackedCurrent.VOLUMEHOURTO = valuesArray[currentField]
			case "VOLUME24HOUR":
				unpackedCurrent.VOLUME24HOUR = valuesArray[currentField]
			case "VOLUME24HOURTO":
				unpackedCurrent.VOLUME24HOUR = valuesArray[currentField]
			case "OPENHOUR":
				unpackedCurrent.OPENHOUR = valuesArray[currentField]
			case "HIGHHOUR":
				unpackedCurrent.HIGHHOUR = valuesArray[currentField]
			case "LOWHOUR":
				unpackedCurrent.LOWHOUR = valuesArray[currentField]
			case "OPEN24HOUR":
				unpackedCurrent.OPEN24HOUR = valuesArray[currentField]
			case "HIGH24HOUR":
				unpackedCurrent.HIGH24HOUR = valuesArray[currentField]
			case "LOW24HOUR":
				unpackedCurrent.LOW24HOUR = valuesArray[currentField]
			case "LASTMARKET":
				unpackedCurrent.LASTMARKET = valuesArray[currentField]
			}
			currentField++
		}
	}
	return unpackedCurrent, nil
}

// func PackCurrent(msg *models.CurrentMessage) string {
// 	// mask := strconv.FormatInt(0, 16)
// 	return "%s~%s~%s~%s~%s~%s~%s~%s~%s~%s~%s~%s~%s"
// }
