package utils

type HelperUtils struct {
}

func NewHelperUtils() *HelperUtils {
	return &HelperUtils{}
}

func (hu *HelperUtils) DefaultIfEmpty(value *string, defaultValue string) string {
	if value == nil || *value == "" {
		return defaultValue
	}

	return *value
}
