package err

func GetErrorMessage(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
