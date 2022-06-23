package convert_pointer

import "time"

func StringPointerToString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func TimePointerToTime(s *time.Time) time.Time {
	if s != nil {
		return *s
	}
	return time.Time{}
}

func BooleanPointerToBoolean(s *bool) bool {
	if s != nil {
		return *s
	}
	return false
}

func IntPointerToInt(s *int) int {
	if s != nil {
		return *s
	}
	return 0
}

func Int32PointerToInt32(s *int32) int32 {
	if s != nil {
		return *s
	}
	return 0
}

func Int64PointerToInt64(s *int64) int64 {
	if s != nil {
		return *s
	}
	return 0
}
