package utils

import (
	"fmt"
	"strconv"
	"stretches-common-api/publicid"
	"strings"

	"github.com/lib/pq"
)

// Auth godoc
// @Summary User
// @Description returns the list of users
// @Router /v1/users/all [post]
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {object} []user.User
// @Failure 400 {object} types.ErrorMsg
// @Failure 500 {object} types.ErrorMsg
func ArrayToString(arr []uint32) string {
	var strBuilder strings.Builder
	strBuilder.WriteString("[")
	for i, num := range arr {
		if i > 0 {
			strBuilder.WriteString(", ")
		}
		strBuilder.WriteString(strconv.FormatUint(uint64(num), 10))
	}
	strBuilder.WriteString("]")
	return strBuilder.String()
}

func Map3ToString(arr map[string]bool) string {
	tmp := []string{}
	for key, _ := range arr {
		i, err := strconv.Atoi(key)
		if err == nil {
			id := publicid.Unobfuscate32bit(uint32(i))
			tmp = append(tmp, fmt.Sprintf("%v", id))
		}

	}
	return strings.Join(tmp, ",")
}
func MapuintToString(arr map[uint32]bool) string {
	tmp := []string{}
	for key, _ := range arr {
//		i, err := strconv.Atoi(key)
//		if err == nil {
//			id := publicid.Unobfuscate32bit(key)
			tmp = append(tmp, fmt.Sprintf("%v", key))
//		}

	}
	return strings.Join(tmp, ",")
}
func IntMap3ToString(arr map[uint32]bool) string {
	tmp := []string{}
	for key, _ := range arr {
		id := publicid.Unobfuscate32bit(key)
		tmp = append(tmp, fmt.Sprintf("%v", id))

	}
	return strings.Join(tmp, ",")
}
func Map4ToString(arr map[string]string) string {
	tmp := []string{}
	for key, _ := range arr {
		i, err := strconv.Atoi(key)
		if err == nil {
			id := publicid.Unobfuscate32bit(uint32(i))
			tmp = append(tmp, fmt.Sprintf("%v", id))
		}

	}
	return strings.Join(tmp, ",")
}
func MapuintstrToString(arr map[uint32]string) string {
	tmp := []string{}
	for key, _ := range arr {
//			id := publicid.Unobfuscate32bit(key)
			tmp = append(tmp, fmt.Sprintf("%v", key))
	
	}
	return strings.Join(tmp, ",")
}
func IntMap4ToString(arr map[uint32]string) string {
	tmp := []string{}
	for key, _ := range arr {
		id := publicid.Unobfuscate32bit(key)
		tmp = append(tmp, fmt.Sprintf("%v", id))

	}
	return strings.Join(tmp, ",")
}
func Array3ToString(arr []string) string {
	tmp := []string{}
	for _, item := range arr {
		if len(item) != 0 {
			i, err := strconv.Atoi(item)
			if err == nil {
				id := publicid.Unobfuscate32bit(uint32(i))
				tmp = append(tmp, fmt.Sprintf("%v", id))
			}
		}
	}
	return strings.Join(tmp, ",")

}
func IntArray3ToString(arr []uint32) string {
	tmp := []string{}
	for _, item := range arr {
		if item != 0 {
			id := publicid.Unobfuscate32bit(item)
			tmp = append(tmp, fmt.Sprintf("%v", id))
		}
	}
	return strings.Join(tmp, ",")

}
func Int64ArrayToString(arr pq.Int64Array) string {
	tmp := []string{}
	for _, item := range arr {
	//	it := int32(item)	 
//		if item != 0 {
//			id := publicid.Unobfuscate32bit(item)
			tmp = append(tmp, fmt.Sprintf("%v", item))
//		}
	}
	return strings.Join(tmp, ",")

}
